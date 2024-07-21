package logger

import (
	"flag"
	"log"
	"os"
	"sync"
)

type logger interface {
	Printf(format string, v ...interface{})
}

var (
	verbose       bool
	defaultLogger logger = noopLogger{}
	once          sync.Once
)

func RegisterFlags() {
	flag.BoolVar(&verbose, "verbose", false, "print debug output")
}

func initLogger() {
	if verbose {
		defaultLogger = log.New(os.Stdout, "", log.LstdFlags)
		defaultLogger.Printf("verbose mode enabled")
	}
}

func Debugf(format string, v ...interface{}) {
	once.Do(initLogger)
	defaultLogger.Printf(format, v...)
}

// default logger that does nothing
type noopLogger struct{}

func (n noopLogger) Printf(format string, v ...interface{}) {}
