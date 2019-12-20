package cmd

import (
	"context"
	"github.com/igorexec/cardinal/app/rest/api"
	"github.com/pkg/errors"
	"log"
	"strings"
)

type ServerCommand struct {
	Port           int    `long:"port" env:"CARDINAL_PORT" default:"8080" description:"cardinal server port"`
	BackupLocation string `long:"backup" env:"BACKUP_PATH" default:"./var/backup" description:"backups location"`

	CommonOpts
}

type serverApp struct {
	*ServerCommand

	restSrv *api.Rest
}

func (s *ServerCommand) Execute(args []string) error {
	log.Printf("[info] start server on port %d", s.Port)

	_, err := s.newServerApp()
	if err != nil {
		log.Printf("[panic] failed to setup application: %+v", err)
		return err
	}
	return nil
}

func (s *ServerCommand) newServerApp() (*serverApp, error) {
	if err := makeDirs(s.BackupLocation); err != nil {
		return nil, err
	}

	if !strings.HasPrefix(s.CardinalURL, "http://") && !strings.HasPrefix(s.CardinalURL, "https://") {
		return nil, errors.Errorf("invalid cardinal url: %s", s.CardinalURL)
	}

	log.Printf("[info] root url=%s", s.CardinalURL)

	// configuration for all services

	srv := &api.Rest{
		Version:     s.Revision,
		CardinalURL: s.CardinalURL,
	}

	return &serverApp{
		ServerCommand: s,
		restSrv:       srv,
	}, nil
}

func (a *serverApp) run(ctx context.Context) error {

}
