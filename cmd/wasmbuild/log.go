package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	// Packages
	"github.com/fatih/color"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// Logger handles logging at different levels based on verbose flag
type Logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
	verbose     bool
	infoColor   *color.Color
	errorColor  *color.Color
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// NewLogger creates a new logger instance
func NewLogger(verbose bool) *Logger {
	flags := 0

	l := &Logger{
		errorLogger: log.New(os.Stderr, "", flags),
		verbose:     verbose,
		infoColor:   color.New(color.Bold),
		errorColor:  color.New(color.FgRed),
	}

	if verbose {
		l.infoLogger = log.New(os.Stderr, "", flags)
	} else {
		// Discard info logs when not verbose
		l.infoLogger = log.New(io.Discard, "", 0)
	}

	return l
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// Info logs informational messages (only when verbose)
func (l *Logger) Info(v ...interface{}) {
	if l.verbose {
		l.infoColor.Fprint(os.Stderr, v...)
		fmt.Fprintln(os.Stderr)
	}
}

// Infof logs formatted informational messages (only when verbose)
func (l *Logger) Infof(format string, v ...interface{}) {
	if l.verbose {
		l.infoColor.Fprintf(os.Stderr, format, v...)
		if format[len(format)-1] != '\n' {
			fmt.Fprintln(os.Stderr)
		}
	}
}

// Error logs error messages (always shown)
func (l *Logger) Error(v ...interface{}) {
	l.errorColor.Fprint(os.Stderr, v...)
	fmt.Fprintln(os.Stderr)
}

// Errorf logs formatted error messages (always shown)
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.errorColor.Fprintf(os.Stderr, format, v...)
	if format[len(format)-1] != '\n' {
		fmt.Fprintln(os.Stderr)
	}
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

// logging middleware
func logging(next http.Handler, logger *Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("%s %s %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
