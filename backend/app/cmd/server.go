package cmd

import (
	"github.com/igorexec/cardinal/app/rest/api"
	"log"
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
	return nil, nil
}
