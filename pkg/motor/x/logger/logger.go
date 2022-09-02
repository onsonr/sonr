package logger

import (
	"log"
	"os"
)

const (
	INFO  = "[INFO]"
	DEBUG = "[DEBUG]"
	WARN  = "[WARN]"
	ERROR = "[ERROR]"
	FATAL = "[FATAL]"
)

type Logger struct {
	logger   *log.Logger
	level    string
	category string
}

func New(level, category string) *Logger {
	return &Logger{
		logger: log.New(os.Stdout, category, 0),
	}
}

func (l *Logger) Debug(v ...string) {
	l.logger.Print(DEBUG, v)
}

func (l *Logger) Info(v ...string) {
	l.logger.Print(INFO, v)
}

func (l *Logger) Warn(v ...string) {
	l.logger.Print(WARN, v)
}

func (l *Logger) Error(v ...string) {
	l.logger.Print(ERROR, v)
}

func (c *Client) PrintConnectionEndpoints() {
	log.Println("Connection Endpoints:")
	log.Printf("\tREST: %s\n", c.GetAPIAddress())
	log.Printf("\tRPC: %s\n", c.GetRPCAddress())
	log.Printf("\tFaucet: %s\n", c.GetFaucetAddress())
	log.Printf("\tIPFS: %s\n", c.GetIPFSAddress())
	log.Printf("\tIPFS API: %s\n", c.GetIPFSApiAddress())
}
