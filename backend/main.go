package main

import (
	"fmt"
	"github.com/go-pkgz/lgr"
	"github.com/icheliadinski/cardinal/cmd"
	"github.com/jessevdk/go-flags"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

type Opts struct {
	CardinalURL string `long:"url" env:"CARDINAL_URL" required:"true" description:"url to cardinal"`

	Dbg bool `long:"dbg" env:"DEBUG" description:"debug mode"`
}

var revision = "unknown"

func main() {
	fmt.Printf("cardinal %s\n", revision)

	var opts Opts
	p := flags.NewParser(&opts, flags.Default)
	p.CommandHandler = func(command flags.Commander, args []string) error {
		setupLog(opts.Dbg)
		c := command.(cmd.CommonOptionsCommander)
		c.SetCommon(cmd.CommonOpts{
			CardinalURL: opts.CardinalURL,
			Revision:    revision,
		})
		err := c.Execute(args)
		if err != nil {
			lgr.Printf("[ERROR] failed with %+v", err)
		}
		return err
	}

	if _, err := p.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
}

func setupLog(dbg bool) {
	if dbg {
		lgr.Setup(lgr.Debug, lgr.CallerFile, lgr.CallerFunc, lgr.Msec, lgr.LevelBraces)
		return
	}
	lgr.Setup(lgr.Msec, lgr.LevelBraces)
}

func init() {
	// catch SIGQUIT and print stack traces
	sigChan := make(chan os.Signal)
	go func() {
		for range sigChan {
			lgr.Printf("[INFO] SIGQUIT detected, dump:\n%s", getDump())
		}
	}()
	signal.Notify(sigChan, syscall.SIGQUIT)
}

// getDump reads runtime stack and returns as a string
func getDump() string {
	maxSize := 5 * 1024 * 1024
	stacktrace := make([]byte, maxSize)
	length := runtime.Stack(stacktrace, true)
	if length > maxSize {
		length = maxSize
	}
	return string(stacktrace[:length])
}
