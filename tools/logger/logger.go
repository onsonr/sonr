package logger

import (
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
func Error(msg string, err error) error {
	log.Error(msg, zap.Error(err))
	return err
}

// Fatal logs a message at FatalLevel. The message includes any fields passed at the log site, as well as any fields accumulated on the logger.
// The logger then calls os.Exit(1), even if logging at FatalLevel is disabled.
func Fatal(msg string, err error) error {
	log.Fatal(msg, zap.Error(err))
	return err
}

// Info logs a message at InfoLevel. The message includes any fields passed at the log site,
// as well as any fields accumulated on the logger.
func Info(msg string, fields ...zapcore.Field) {
	log.Info(msg, fields...)
}

// Panic logs a message at PanicLevel. The message includes any fields passed at the log site,
// as well as any fields accumulated on the logger.
func Panic(msg string, err error) error {
	if isDev {
		log.DPanic(msg, zap.Error(err))
	} else {
		log.Panic(msg, zap.Error(err))
	}
	return err
}

// Warn logs a message at WarnLevel. The message includes any fields passed at the log site,
// as well as any fields accumulated on the logger.
func Warn(msg string, err error) error {
	log.Warn(msg, zap.Error(err))
	return err
}

// Close calls the underlying Core's Sync method, flushing any buffered log entries. Applications
// should take care to call Sync before exiting.
func Close() {
	log.Sync()
}
