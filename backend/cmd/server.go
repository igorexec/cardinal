package cmd

import (
	"context"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/go-pkgz/lgr"
	"github.com/icheliadinski/cardinal/rest/api"
	"github.com/pkg/errors"
)

type ServerCommand struct {
	Port    int    `long:"port" env:"CARDINAL_PORT" default:"8080" description:"port"`
	WebRoot string `long:"web-root" env:"CARDINAL_WEB_ROOT" default:"./web" description:"web root directory"`
	CommonOpts
}

type serverApp struct {
	*ServerCommand
	restSrv *api.Rest

	terminated chan struct{}
}

func (s *ServerCommand) Execute(args []string) error {
	lgr.Printf("[INFO] start server on port %d", s.Port)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
		<-stop
		lgr.Printf("[WARN] interrupt signal")
		cancel()
	}()

	app, err := s.newServerApp()
	if err != nil {
		lgr.Printf("[PANIC] failed to setup application, %+v", err)
		return err
	}

	lgr.Printf("[INFO] cardinal terminated")
	return nil
}

func (s *ServerCommand) newServerApp() (*serverApp, error) {
	if !strings.HasPrefix(s.CardinalURL, "http://") && !strings.HasPrefix(s.CardinalURL, "https://") {
		return nil, errors.Errorf("invalid cardinal url %s", s.CardinalURL)
	}
	lgr.Printf("[INFO] root url=%s", s.CardinalURL)

	srv := &api.Rest{
		Version:     s.Revision,
		WebRoot:     s.WebRoot,
		CardinalURL: s.CardinalURL,
	}
	return &serverApp{
		ServerCommand: s,
		restSrv:       srv,
		terminated:    make(chan struct{}),
	}, nil
}

func (a *serverApp) run(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		lgr.Print("[INFO] shutdown initiated")
		a.restSrv.Shutdown()
		lgr.Print("[INFO] shutdown completed")
	}()

	a.restSrv.Run(a.Port)
	close(a.terminated)
	return nil
}
