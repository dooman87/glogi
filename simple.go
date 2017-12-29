package glogi

import (
	"log"
	"os"
)

//SimpleLogger is a basic implementation
//of a Logger interface that prints error messages
//to stderr and debug to stdout. Using log.Logger
//as a backend.
type SimpleLogger struct {
	debug *log.Logger
	err *log.Logger
}

func New() *SimpleLogger {
	return &SimpleLogger{
		debug: log.New(os.Stdout, "", log.LstdFlags),
		err: log.New(os.Stderr, "ERROR: ", log.LstdFlags),
	}
}

func (l *SimpleLogger) Printf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

func (l *SimpleLogger) Print(v ...interface{}) {
	l.debug.Print(v...)
}

func (l *SimpleLogger) Errorf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}

func (l *SimpleLogger) Error(v ...interface{}) {
	l.err.Print(v...)
}

