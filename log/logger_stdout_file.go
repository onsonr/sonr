package log

import (
	"os"

	"github.com/op/go-logging"
)

// StdoutFileLogger outputs logs both to standard output as well as to a given file.
type StdoutFileLogger struct {
	logger
}

var _ Logger = (*StdoutFileLogger)(nil)

func NewStdoutFileLogger(module, logFilePath, logLevel, formatStdout,
	formatFile string) (*StdoutFileLogger, error) {
	baseLogger := newLogger(module)

	levelInt, err := logging.LogLevel(logLevel)
	if err != nil {
		return nil, err
	}

	formattedStdoutBackend, err := baseLogger.setupFormattedBackend(os.Stdout, formatStdout)
	if err != nil {
		return nil, err
	}
	backends := []logging.Backend{formattedStdoutBackend}

	logFile, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	formattedFileBackend, err := baseLogger.setupFormattedBackend(logFile, formatFile)
	if err != nil {
		return nil, err
	}
	backends = append(backends, formattedFileBackend)

	leveledBackend := logging.SetBackend(backends...)
	leveledBackend.SetLevel(levelInt, module)
	logger := &StdoutFileLogger{baseLogger}
	logger.SetBackend(leveledBackend)

	return logger, nil
}
