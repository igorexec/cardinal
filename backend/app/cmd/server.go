package cmd

import (
	"context"
	"github.com/igorexec/cardinal/app/collect"
	"github.com/igorexec/cardinal/app/rest/api"
	"github.com/igorexec/cardinal/app/store"
	"github.com/igorexec/cardinal/app/store/engine"
	"github.com/igorexec/cardinal/app/store/engine/service"
	"github.com/pkg/errors"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type ServerCommand struct {
	Store StoreGroup `group:"store" namespace:"store" env-namespace:"STORE"`

	Port           int    `long:"port" env:"CARDINAL_PORT" default:"8080" description:"cardinal server port"`
	BackupLocation string `long:"backup" env:"BACKUP_PATH" default:"./var/backup" description:"backups location"`

	CommonOpts
}

type StoreGroup struct {
	Type string `long:"type" env:"TYPE" description:"type of storage" choice:"mongo" default:"mongo"`
}

type serverApp struct {
	*ServerCommand

	restSrv     *api.Rest
	dataService *service.DataStore

	terminated chan struct{}
}

func (s *ServerCommand) Execute(args []string) error {
	log.Printf("[info] start server on port %d", s.Port)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
		<-stop
		log.Printf("[warn] interrupt signal")
		cancel()
	}()

	app, err := s.newServerApp()
	if err != nil {
		log.Printf("[panic] failed to setup application: %+v", err)
		return err
	}

	if err := app.run(ctx); err != nil {
		log.Printf("[error] cardinal terminated with error: %+v", err)
		return err
	}
	log.Printf("[info] cardinal terminated")
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

	// todo: add configuration for all services

	storeEngine, err := s.makeDataStore()
	if err != nil {
		return nil, err
	}

	dataService := &service.DataStore{Engine: storeEngine}

	psc := collect.NewPageSpeedCollector("")

	srv := &api.Rest{
		DataService:        dataService,
		Version:            s.Revision,
		CardinalURL:        s.CardinalURL,
		PageSpeedCollector: psc,
	}

	return &serverApp{
		ServerCommand: s,
		restSrv:       srv,
		dataService:   dataService,
		terminated:    make(chan struct{}),
	}, nil
}

func (s *ServerCommand) makeDataStore() (eng engine.Interface, err error) {
	log.Printf("[info] make data store, type=")

	switch s.Store.Type {
	case store.MONGO:
		eng, err = engine.NewMongo("mongodb://localhost:27017")
		if err != nil {
			log.Fatalf("[error] failed to connect to database: %v", err)
			return nil, err
		}
		return eng, nil
	default:
		return nil, errors.Errorf("[error] unsupported store type %s", s.Store.Type)
	}
}

func (a *serverApp) run(ctx context.Context) error {

	go func() {
		<-ctx.Done()
		log.Printf("[info] shutdown initiated")
		a.restSrv.Shutdown()
		// kill all the rest of services
		log.Print("[info] shutdown completed")
	}()

	a.activateBackup(ctx)

	a.restSrv.Run(a.Port)
	close(a.terminated)
	return nil
}

func (a *serverApp) activateBackup(ctx context.Context) {
	// todo: do backup
}
