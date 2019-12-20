package main

import (
	"fmt"
	"github.com/igorexec/cardinal/app/cmd"
	"github.com/jessevdk/go-flags"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

type Opts struct {
	ServerCmd   struct{} `command:"server"`
	CardinalURL string   `long:"url" env:"CARDINAL_URL" required:"true" description:"url to cardinal"`
}

func main() {
	fmt.Printf("cardinal version 1")

	var opts Opts
	p := flags.NewParser(&opts, flags.Default)
	p.CommandHandler = func(command flags.Commander, args []string) error {
		c := command.(cmd.CommonOptionsCommander)
		c.SetCommon(cmd.CommonOpts{CardinalURL: opts.CardinalURL})

		err := c.Execute(args)
		if err != nil {
			log.Printf("[error] failed with %+v", err)
		}
		return err
	}

	if _, err := p.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.HelpFlag {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
}

func init() {
	// catch SIGQUIT
	sigChan := make(chan os.Signal)

	go func() {
		for range sigChan {
			log.Printf("[info] SIGQUIT detected, dump:\n%s", getDump())
		}
	}()

	signal.Notify(sigChan, syscall.SIGQUIT)
}

func getDump() string {
	maxSize := 5 * 1024 * 1024
	stacktrace := make([]byte, maxSize)
	length := runtime.Stack(stacktrace, true)
	if length > maxSize {
		length = maxSize
	}
	return string(stacktrace[:length])
}
