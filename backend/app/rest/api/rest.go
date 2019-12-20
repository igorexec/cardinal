package api

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Rest struct {
	Version string

	CardinalURL string

	httpServer *http.Server
	lock       sync.Mutex
}

func (s *Rest) Run(port int) {
	// todo: add switch statement according to SSL config. HTTP or HTTPS
	log.Printf("[info] activate http server on port: %d", port)

	s.lock.Lock()
	s.httpServer = s.makeHTTPServer(port, nil)
	s.lock.Unlock()

	err := s.httpServer.ListenAndServe()
	log.Printf("[warn] http server terminated: %s", err)
}

func (s *Rest) makeHTTPServer(port int, router http.Handler) *http.Server {
	return &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       30 * time.Second,
	}
}
