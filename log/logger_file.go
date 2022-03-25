package log

import (
	"os"

	"github.com/op/go-logging"
)

// FileLogger outputs logs to a given file.
type FileLogger struct {
	logger
}

var _ Logger = (*FileLogger)(nil)

func NewFileLogger(module, logFilePath, logLevel, format string) (*FileLogger, error) {
	baseLogger := newLogger(module)

	levelInt, err := logging.LogLevel(logLevel)
	if err != nil {
		return nil, err
	}
	logFile, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	backend, err := baseLogger.setupFormattedBackend(logFile, format)
	if err != nil {
		return nil, err
	}

	leveledBackend := logging.SetBackend(backend)
	leveledBackend.SetLevel(levelInt, module)
	logger := &FileLogger{baseLogger}
	logger.SetBackend(leveledBackend)

	return logger, nil
}
