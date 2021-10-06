package logger

import (
	"errors"
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger
var isDev bool

// Init sets up Zap Logger for the application, based on isDevelopment
func Init(isDevelopment bool) {
	isDev = isDevelopment

	// Create the logger
	if isDev {
		log, _ = zap.NewDevelopment(zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	} else {
		log, _ = zap.NewProduction()
	}
}

// Debug logs a message at DebugLevel. The message includes any fields passed at the log site,
// as well as any fields accumulated on the logger.
func Debug(msg string, fields ...zapcore.Field) {
	log.Debug(msg, fields...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed at the log site,
// as well as any fields accumulated on the logger.
func Error(msg string, errs ...error) error {
	// Initialize the errors
	err := errors.New(msg)
	fields := []zapcore.Field{}

	// Add the errors
	for _, e := range errs {
		fields = append(fields, zap.Error(e))
		err = fmt.Errorf("%w; Second error", e)
	}

	// Log the message and return the error
	log.Error(msg, fields...)
	return err
}

// Fatal logs a message at FatalLevel. The message includes any fields passed at the log site, as well as any fields accumulated on the logger.
// The logger then calls os.Exit(1), even if logging at FatalLevel is disabled.
func Fatal(msg string, errs ...error) error {
	// Initialize the errors
	err := errors.New(msg)
	fields := []zapcore.Field{}

	// Add the errors
	for _, e := range errs {
		fields = append(fields, zap.Error(e))
		err = fmt.Errorf("%w; Second error", e)
	}

	// Log the message and return the error
	log.Fatal(msg, fields...)
	return err
}

// Info logs a message at InfoLevel. The message includes any fields passed at the log site,
// as well as any fields accumulated on the logger.
func Info(msg string, fields ...zapcore.Field) {
	log.Info(msg, fields...)
}

// Panic logs a message at PanicLevel. The message includes any fields passed at the log site,
// as well as any fields accumulated on the logger.
func Panic(msg string, errs ...error) error {
	// Initialize the errors
	err := errors.New(msg)
	fields := []zapcore.Field{}

	// Add the errors
	for _, e := range errs {
		fields = append(fields, zap.Error(e))
		err = fmt.Errorf("%w; Second error", e)
	}

	if isDev {
		log.DPanic(msg, fields...)
	} else {
		log.Panic(msg, fields...)
	}
	return err
}

// Warn logs a message at WarnLevel. The message includes any fields passed at the log site,
// as well as any fields accumulated on the logger.
func Warn(msg string, errs ...error) error {
	// Initialize the errors
	err := errors.New(msg)
	fields := []zapcore.Field{}

	// Add the errors
	for _, e := range errs {
		fields = append(fields, zap.Error(e))
		err = fmt.Errorf("%w; Second error", e)
	}

	// Log the message and return the error
	log.Warn(msg, fields...)
	return err
}

// Close calls the underlying Core's Sync method, flushing any buffered log entries. Applications
// should take care to call Sync before exiting.
func Close() {
	log.Sync()
}
