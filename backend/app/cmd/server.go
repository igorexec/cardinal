package cmd

import "log"

type ServerCommand struct {
	Port int `long:"port" env:"CARDINAL_PORT" default:"8080" description:"cardinal server port"`

	CommonOpts
}

type serverApp struct {
	*ServerCommand
}

func (s *ServerCommand) Execute(args []string) error {
	log.Printf("[info] start server on port %d", s.Port)
	return nil
}

func (s *ServerCommand) newServerApp() {}
