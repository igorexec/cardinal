package api

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/go-pkgz/lgr"
	"github.com/icheliadinski/cardinal/collector"
	"github.com/icheliadinski/cardinal/store/service"
	"net/http"
	"sync"
	"time"
)

const hardBodyLimit = 1024 * 64 // limit size of body

type Rest struct {
	Version string

	DataService      *service.DataStore
	CollectorService *collector.Collector
	CardinalURL      string

	httpServer *http.Server
	lock       sync.Mutex

	pubRest public
}

func (s *Rest) Run(port int) {
	lgr.Printf("[INFO] activate http rest server on port %d", port)

	s.lock.Lock()
	s.httpServer = s.makeHTTPServer(port, s.routes())
	s.httpServer.ErrorLog = lgr.ToStdLogger(lgr.Default(), "WARN")
	s.lock.Unlock()

	err := s.httpServer.ListenAndServe()
	lgr.Printf("[WARN] http server terminated, %s", err)
}

func (s *Rest) Shutdown() {
	lgr.Print("[WARN] shutdown rest server")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	s.lock.Lock()
	if s.httpServer != nil {
		if err := s.httpServer.Shutdown(ctx); err != nil {
			lgr.Printf("[DEBUG] http shutdown error, %s", err)
		}
		lgr.Print("[DEBUG] shutdown http server completed")
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

	// TODO: Add middlewares

	s.pubRest = s.controllerGroups()

	router.Route("/api/v1", func(rapi chi.Router) {
		rapi.Group(func(ropen chi.Router) {

			ropen.Get("/config", s.configCtrl)

			ropen.Group(func(rps chi.Router) {
				rps.Post("/pagespeed", s.pubRest.collectPageSpeedCtrl)
			})
		})
	})
	return router
}

func (s *Rest) configCtrl(w http.ResponseWriter, r *http.Request) {
	cnf := struct {
		Version string `json:"version"`
	}{
		Version: s.Version,
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, cnf)
}

func (s *Rest) controllerGroups() public {

	pubGrp := public{
		dataService: s.DataService,
		collector:   s.CollectorService,
	}
	return pubGrp
}
