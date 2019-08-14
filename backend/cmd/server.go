package cmd

import "github.com/go-pkgz/lgr"

type ServerCommand struct {
	Port int `long:"port" env:"CARDINAL_PORT" default:"8080" description:"port"`
	CommonOpts
}

func (s *ServerCommand) Execute(args []string) error {
	lgr.Printf("[INFO] start server on port %d", s.Port)
	return nil
}
