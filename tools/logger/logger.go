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
		log, _ = zap.NewDevelopment()
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
func Error(msg string, fields ...zapcore.Field) {
	log.Error(msg, fields...)
}

// Fatal logs a message at FatalLevel. The message includes any fields passed at the log site, as well as any fields accumulated on the logger.
// The logger then calls os.Exit(1), even if logging at FatalLevel is disabled.
func Fatal(msg string, fields ...zapcore.Field) {
	log.Fatal(msg, fields...)
}

// Info logs a message at InfoLevel. The message includes any fields passed at the log site,
// as well as any fields accumulated on the logger.
func Info(msg string, fields ...zapcore.Field) {
	log.Info(msg, fields...)
}

// Panic logs a message at PanicLevel. The message includes any fields passed at the log site,
// as well as any fields accumulated on the logger.
func Panic(msg string, fields ...zapcore.Field) {
	if isDev {
		log.DPanic(msg, fields...)
	} else {
		log.Panic(msg, fields...)
	}
}

// Warn logs a message at WarnLevel. The message includes any fields passed at the log site,
// as well as any fields accumulated on the logger.
func Warn(msg string, fields ...zapcore.Field) {
	log.Warn(msg, fields...)
}

// Close calls the underlying Core's Sync method, flushing any buffered log entries. Applications
// should take care to call Sync before exiting.
func Close() {
	log.Sync()
}
