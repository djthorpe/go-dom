package main

import (
	"io"
	"log"
	"os"
)

// Logger handles logging at different levels based on verbose flag
type Logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

// NewLogger creates a new logger instance
func NewLogger(verbose bool) *Logger {
	// Use LstdFlags to include date and time (2009/01/23 01:23:23)
	flags := log.LstdFlags

	l := &Logger{
		errorLogger: log.New(os.Stderr, "", flags),
	}

	if verbose {
		l.infoLogger = log.New(os.Stdout, "", flags)
	} else {
		// Discard info logs when not verbose
		l.infoLogger = log.New(io.Discard, "", 0)
	}

	return l
}

// Info logs informational messages (only when verbose)
func (l *Logger) Info(v ...interface{}) {
	l.infoLogger.Print(v...)
}

// Infof logs formatted informational messages (only when verbose)
func (l *Logger) Infof(format string, v ...interface{}) {
	l.infoLogger.Printf(format, v...)
}

// Println is an alias for Info
func (l *Logger) Println(v ...interface{}) {
	l.Info(v...)
}

// Printf is an alias for Infof
func (l *Logger) Printf(format string, v ...interface{}) {
	l.Infof(format, v...)
}

// Error logs error messages (always shown)
func (l *Logger) Error(v ...interface{}) {
	l.errorLogger.Print(v...)
}

// Errorf logs formatted error messages (always shown)
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.errorLogger.Printf(format, v...)
}

// Global logger instance
var logger *Logger

// InitLogger initializes the global logger
func InitLogger(verbose bool) {
	logger = NewLogger(verbose)
}
