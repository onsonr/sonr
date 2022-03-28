package log

import (
	"os"

	"github.com/op/go-logging"
)

// StdoutLogger outputs logs to standard output.
type StdoutLogger struct {
	logger
}

var _ Logger = (*StdoutLogger)(nil)

func NewStdoutLogger(module, logLevel, format string) (*StdoutLogger, error) {
	baseLogger := newLogger(module)

	levelInt, err := logging.LogLevel(logLevel)
	if err != nil {
		return nil, err
	}

	backend, err := baseLogger.setupFormattedBackend(os.Stdout, format)
	if err != nil {
		return nil, err
	}

	leveledBackend := logging.SetBackend(backend)
	leveledBackend.SetLevel(levelInt, module)
	logger := &StdoutLogger{baseLogger}
	logger.SetBackend(leveledBackend)
	return logger, nil
}
