package main

import (
	"fmt"
	"time"
)

const (
	info  = "INFO"
	error = "ERROR"
	debug = "DEBUG"
	warn  = "WARN"
)

type logEntry struct {
}

type Logger interface {
	Info(msg string) string
}

var logCh = make(chan string, 50)
var doneCh = make(chan struct{}) //signal only channel

func main() {
	var log Logger = &logEntry{}
	/*	defer func() {
		fmt.Println("closing log channel")
		close(logCh)
	}() */
	go logging()
	logCh <- log.Info("hi there")
	logCh <- log.Info("hi there 2")
	logCh <- log.Info("hi there 3")
	logCh <- log.Info("hi there 4")
	logCh <- log.Info("hi there 5")
	logCh <- log.Info("hi there 6")

	time.Sleep(100 * time.Millisecond)

	doneCh <- struct{}{}
}

func (l *logEntry) Info(msg string) string {

	return info + ":" + msg
}

func logging() {
	for {
		select {
		case msg := <-logCh:
			fmt.Println(msg)
		case <-doneCh:
			fmt.Println("Exiting go routine")
			break
		}

	}
}
