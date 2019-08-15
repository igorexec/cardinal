package api

import (
	"context"
	"fmt"
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth_chi"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-pkgz/lgr"
	"github.com/rakyll/statik/fs"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Rest struct {
	Version     string
	WebRoot     string
	CardinalURL string

	lock sync.Mutex

	httpServer *http.Server
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

func (s *Rest) Run(port int) {
	lgr.Printf("[INFO] activate http rest server on port %d", port)

	s.lock.Lock()
	s.httpServer = s.makeHTTPServer(port, nil)
	s.httpServer.ErrorLog = lgr.ToStdLogger(lgr.Default(), "WARN")
	s.lock.Unlock()

	err := s.httpServer.ListenAndServe()
	lgr.Printf("[WARN] http server terminated, %s", err)
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
	router.Use(middleware.Throttle(1000), middleware.RealIP)

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-XSRF-Token", "X-JWT"},
		ExposedHeaders:   []string{"Authorization"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	router.Use(corsMiddleware.Handler)

	router.Route("/api/v1", func(rapi chi.Router) {

	})

	addFileServer(router, "/web", http.Dir(s.WebRoot))
	return router
}

func addFileServer(r chi.Router, path string, root http.FileSystem) {
	var webFS http.Handler

	statikFS, err := fs.New()
	if err != nil {
		lgr.Printf("[DEBUG] no embedded assets loaded, %s", err)
		lgr.Printf("[INFO] run file server for %s, path %s", root, path)
		webFS = http.FileServer(root)
	} else {
		lgr.Printf("[INFO] run file server for %s, embedded", root)
		webFS = http.FileServer(statikFS)
	}

	origPath := path
	webFS = http.StripPrefix(path, webFS)
	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.With(tollbooth_chi.LimitHandler(tollbooth.NewLimiter(20, nil)), middleware.Timeout(10*time.Second)).
		Get(path, func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/") && len(r.URL.Path) > 1 && r.URL.Path != (origPath+"/") {
				http.NotFound(w, r)
				return
			}
			webFS.ServeHTTP(w, r)
		})
}
