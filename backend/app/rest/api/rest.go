package api

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/igorexec/cardinal/app/collect"
	"github.com/igorexec/cardinal/app/store/engine/service"
	"log"
	"net/http"
	"sync"
	"time"
)

type Rest struct {
	Version string

	DataService *service.DataStore
	CardinalURL string

	PageSpeedCollector *collect.PageSpeedCollector

	httpServer *http.Server
	lock       sync.Mutex

	pubRest  public
	privRest private
}

const hardBodyLimit = 1024 * 64 // limit size of body

func (s *Rest) Run(port int) {
	// todo: add switch statement according to SSL config. HTTP or HTTPS
	log.Printf("[info] activate http server on port: %d", port)

	s.lock.Lock()
	s.httpServer = s.makeHTTPServer(port, s.routes())
	s.lock.Unlock()

	err := s.httpServer.ListenAndServe()
	log.Printf("[warn] http server terminated: %s", err)
}

func (s *Rest) Shutdown() {
	log.Print("[warn] shutdown rest server")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	s.lock.Lock()
	if s.httpServer != nil {
		if err := s.httpServer.Shutdown(ctx); err != nil {
			log.Printf("[debug] http shutdown error: %s", err)
		}
		log.Print("[debug] shutdown http server completed")
	}
	s.lock.Unlock()
}

func (s *Rest) makeHTTPServer(port int, router http.Handler) *http.Server {
	return &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       30 * time.Second,
	}
}

func (s *Rest) routes() chi.Router {
	router := chi.NewRouter()

	// todo: add middlewares

	s.pubRest, s.privRest = s.controllerGroups()

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-XSRF-Token", "X-JWT"},
		ExposedHeaders:   []string{"Authorization"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	router.Use(corsMiddleware.Handler)

	router.Route("/api/v1", func(rapi chi.Router) {
		rapi.Group(func(ropen chi.Router) {
			ropen.Get("/config", s.configCtrl)

			ropen.Group(func(rps chi.Router) {
				rps.Get("/pagespeed", s.pubRest.findPageSpeed)
				rps.Post("/pagespeed", s.privRest.collectCtrl)
			})
		})
	})
	return router
}

func (s *Rest) configCtrl(w http.ResponseWriter, r *http.Request) {
	cnf := struct {
		Version string
	}{Version: s.Version}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, cnf)
}

func (s *Rest) controllerGroups() (public, private) {
	pubGrp := public{dataService: s.DataService}
	privGrp := private{dataService: s.DataService, pageSpeedCollector: s.PageSpeedCollector}

	return pubGrp, privGrp
}
