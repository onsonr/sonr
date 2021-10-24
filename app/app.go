package app

import (
	"context"
	"fmt"
	"net"
	"os"
	"path/filepath"

	"github.com/kataras/golog"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/node"
	"github.com/sonr-io/core/pkg/common"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
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

var (
	Ctx        context.Context
	Node       api.NodeImpl
	Mode       node.StubMode
	GRPCServer *grpc.Server
)

// Start starts the Sonr Node
func Start(req *api.InitializeRequest, options ...Option) {
	// Check if Node is already running
	if Node != nil {
		golog.Error("Sonr Instance already active")
		return
	}

	// Set Options
	opts := defaultOptions()
	for _, o := range options {
		o(opts)
	}

	// Apply Options
	Mode = opts.mode
	GRPCServer = grpc.NewServer()

	// Set Logging Settings
	golog.SetLevel(opts.logLevel)
	golog.SetPrefix(opts.mode.Prefix())

	// Create Node
	Ctx = context.Background()
	err := req.Parse()
	if err != nil {
		golog.Fatalf("%s - Failed to parse Initialize Request", err)
	}

	// Set Node Stub
	Node, _, err = node.NewNode(Ctx, node.WithGRPC(GRPCServer),
		node.WithMode(opts.mode),
		node.WithRequest(req))
	if err != nil {
		golog.Fatalf("%s - Failed to Start new Node", err)
	}

	// Open Listener on Port
	listener, err := net.Listen(opts.network, opts.Address())
	if err != nil {
		golog.Fatalf("%s - Failed to Create New Listener", err)
	}

	// Serve Node for GRPC
	if opts.mode.IsBin() {
		Serve(listener)
	} else {
		go Serve(listener)
	}
}

// Serve starts the GRPC Server
func Serve(l net.Listener) {
	// Start GRPC Server
	golog.Infof("Starting GRPC Server on %s", l.Addr().String())
	err := GRPCServer.Serve(l)
	if err != nil {
		golog.Fatalf("%s - Failed to start GRPC Server", err)
		Exit(1)
	}
}

// Exit handles cleanup on Sonr Node
func Exit(code int) {
	if Node == nil {
		golog.Debug("Skipping Exit, instance is nil...")
		return
	}
	golog.Debug("Cleaning up on Exit...")
	Node.Close()
	GRPCServer.Stop()
	defer Ctx.Done()

	// Check for Full Desktop Node
	if common.IsDesktop() {
		ex, err := os.Executable()
		if err != nil {
			golog.Errorf("%s - Failed to find Executable", err)
			return
		}

		// Delete Executable Path
		exPath := filepath.Dir(ex)
		err = os.RemoveAll(filepath.Join(exPath, "sonr_bitcask"))
		if err != nil {
			golog.Warn("Failed to remove Bitcask, ", err)
		}
		err = viper.SafeWriteConfig()
		if err == nil {
			golog.Debug("Wrote new config file to Disk")
		}
		os.Exit(code)
	}
}

// Option is a function that can be passed to Start
type Option func(*options)

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

// WithMode sets the mode for the Node
func WithMode(mode node.StubMode) Option {
	return func(o *options) {
		o.mode = mode
	}
}

// options is the struct for the options
type options struct {
	host     string
	network  string
	port     int
	mode     node.StubMode
	logLevel string
}

// Address returns the address of the node.
func (opts *options) Address() string {
	return fmt.Sprintf("%s%d", opts.host, opts.port)
}

// defaultOptions returns the default options
func defaultOptions() *options {
	return &options{
		host:     ":",
		port:     26225,
		mode:     node.StubMode_LIB,
		network:  "tcp",
		logLevel: string(InfoLevel),
	}
}
