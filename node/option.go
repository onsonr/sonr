package node

import (
	"fmt"

	common "github.com/sonr-io/core/common"
)

// LogLevel is the type for the log level
type LogLevel string

const (
	// DebugLevel is the debug log level
	DebugLevel LogLevel = "debug"
	// InfoLevel is the info log level
	InfoLevel LogLevel = "info"
	// WarnLevel is the warn log level
	WarnLevel LogLevel = "warn"
	// ErrorLevel is the error log level
	ErrorLevel LogLevel = "error"
	// FatalLevel is the fatal log level
	FatalLevel LogLevel = "fatal"
)

// Option is a function that modifies the node options.
type Option func(*options)

// WithMode starts the Client RPC server as a highway node.
func WithMode(m Role) Option {
	return func(o *options) {
		o.mode = m
	}
}

// WithLogLevel sets the log level for Logger
func WithLogLevel(level LogLevel) Option {
	return func(o *options) {
		o.logLevel = string(level)
	}
}

// WithHost sets the host for the Node Stub Client Host
func WithHost(host string) Option {
	return func(o *options) {
		o.host = host
	}
}

// WithPort sets the port for the Node Stub Client
func WithPort(port int) Option {
	return func(o *options) {
		o.port = port
	}
}

// WithSocketsDir sets the directory for the Node Sockets
func WithSocketsDir(dir string) Option {
	return func(o *options) {
		o.socketsDir = dir
	}
}

// options is a collection of options for the node.
type options struct {
	connection    common.Connection
	location      *common.Location
	mode          Role
	profile       *common.Profile
	configuration *Configuration
	host          string
	network       string
	port          int
	logLevel      string
	socketsDir    string
}

// defaultOptions returns the default options
func defaultOptions() *options {
	return &options{
		connection: common.Connection_WIFI,
		profile:    common.NewDefaultProfile(),
		host:       ":",
		port:       26225,
		mode:       Role_MOTOR,
		network:    "tcp",
		logLevel:   string(InfoLevel),
	}
}

// Address returns the address of the node.
func (opts *options) Address() string {
	return fmt.Sprintf("%s%d", opts.host, opts.port)
}
