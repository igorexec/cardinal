package main

import (
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

type Opts struct {
}

func main() {

}

func init() {
	// catch SIGQUIT
	sigChan := make(chan os.Signal)

	go func() {
		for range sigChan {
			log.Printf("[INFO] SIGQUIT detected, dump:\n%s", getDump())
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
