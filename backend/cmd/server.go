package cmd

import (
	"context"
	"github.com/go-pkgz/lgr"
	"github.com/go-pkgz/mongo"
	"github.com/icheliadinski/cardinal/rest/api"
	"github.com/icheliadinski/cardinal/store/engine"
	"github.com/icheliadinski/cardinal/store/service"
	"github.com/pkg/errors"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type ServerCommand struct {
	Store StoreGroup `group:"store" namespace:"store" env-namespace:"STORE"`
	Mongo MongoGroup `group:"mongo" namespace:"mongo" env-namespace:"MONGO"`

	Port int `long:"port" env:"CARDINAL_PORT" default:"8080" description:"port"`
	CommonOpts
}

type StoreGroup struct {
	Type string `long:"type" env:"TYPE" description:"type of storage" choice:"mongo" default:"mongo"`
}

type MongoGroup struct {
	URL string `long:"url" env:"url" description:"mongo url"`
	DB  string `long:"db" env:"DB" default:"cardinal" description:"mongo database"`
}

type serverApp struct {
	*ServerCommand
	restSrv     *api.Rest
	dataService *service.DataStore
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

	if err := app.run(ctx); err != nil {
		lgr.Printf("[ERROR] cardinal terminated with error %+v", err)
		return err
	}
	lgr.Print("[INFO] cardinal terminated")
	return nil
}

func (s *ServerCommand) newServerApp() (*serverApp, error) {
	if !strings.HasPrefix(s.CardinalURL, "http://") && !strings.HasPrefix(s.CardinalURL, "https://") {
		return nil, errors.Errorf("invalid cardinal url %s", s.CardinalURL)
	}
	lgr.Printf("[INFO] root url=%s:%d", s.CardinalURL, s.Port)

	storeEngine, err := s.makeDataStore()
	if err != nil {
		return nil, errors.Wrap(err, "failed to make data store engine")
	}

	dataService := &service.DataStore{
		Engine: storeEngine,
	}

	srv := &api.Rest{
		DataService: dataService,
		Version:     s.Revision,
		CardinalURL: s.CardinalURL,
	}

	return &serverApp{
		ServerCommand: s,
		restSrv:       srv,
		dataService:   dataService,
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
	return nil
}

func (s *ServerCommand) makeDataStore() (result engine.Interface, err error) {
	lgr.Printf("[INFO] make data store, type=%s", s.Store.Type)

	switch s.Store.Type {
	case "mongo":
		mgServer, err := s.makeMongo()
		if err != nil {
			return result, errors.Wrap(err, "failed to create mongo server")
		}
		conn := mongo.NewConnection(mgServer, s.Mongo.DB, "")
		result, err = engine.NewMongo(conn, 500, 100*time.Millisecond)
	default:
		return nil, errors.Errorf("unsupported store type %s", s.Store.Type)
	}
	return result, errors.Wrapf(err, "can't initialize data store")
}

func (s *ServerCommand) makeMongo() (result *mongo.Server, err error) {
	if s.Mongo.URL == "" {
		return nil, errors.New("no mongo URL provided")
	}
	return mongo.NewServerWithURL(s.Mongo.URL, 10*time.Second)
}
