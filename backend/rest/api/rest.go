package api

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-pkgz/lgr"
	"net/http"
	"sync"
	"time"
)

type Rest struct {
	Version     string
	CardinalURL string

	httpServer *http.Server
	lock       sync.Mutex
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

	router.Route("/api/v1", func(rapi chi.Router) {
		rapi.Group(func(ropen chi.Router) {

		})
	})

	return router
}
