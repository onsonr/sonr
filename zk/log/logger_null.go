package log

import (
	"io/ioutil"

	"github.com/op/go-logging"
)

// NullLogger suppresses logging output
type NullLogger struct {
	logger
}

var _ Logger = (*NullLogger)(nil)

func NewNullLogger() *NullLogger {
	backend := logging.NewLogBackend(ioutil.Discard, "", 0)
	leveledBackend := logging.SetBackend(backend)
	logger := &NullLogger{newLogger("")}
	logger.SetBackend(leveledBackend)
	return logger
}
