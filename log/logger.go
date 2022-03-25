package log

import (
	"io"

	"github.com/op/go-logging"
)

// Logger is a convenience interface that makes use of all functionality from go-logging's
// Logger struct. In addition, it defines Disable() and SetLevel(level) functions
type Logger interface {
	// This function is our own
	SetLevel(level string) error

	// These are functions hooked on go-logging's Logger struct
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Notice(args ...interface{})
	Noticef(format string, args ...interface{})
	Warning(args ...interface{})
	Warningf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Critical(args ...interface{})
	Criticalf(format string, args ...interface{})
}

// logger embeds *logging.Logger in order to gain access to its implementations of
// logging functions like Debug(...), Debugf(...) and others - see the Logger interface above.
type logger struct {
	*logging.Logger
}

func newLogger(module string) logger {
	return logger{
		logging.MustGetLogger(module),
	}
}

// SetLevel sets the log level for the given logger. In case of an invalid log level argument,
// it propagates the error detected by go-logging's package logging
func (logger *logger) SetLevel(levelStr string) error {
	// obtain logging.Level type from a string argument representing log level
	levelInt, err := logging.LogLevel(levelStr)
	if err != nil {
		return err
	}
	logging.SetLevel(levelInt, logger.Module)
	return nil
}

// setupFormattedBackend accepts io.Writer and a format string. It constructs a logging backend
// that uses io.Writer and outputs logs in a format specified by the format string.
func (logger *logger) setupFormattedBackend(writer io.Writer, format string) (logging.Backend, error) {
	backend := logging.NewLogBackend(writer, "", 0)
	formatter, err := logging.NewStringFormatter(format)
	if err != nil {
		return nil, err
	}
	return logging.NewBackendFormatter(backend, formatter), nil
}
