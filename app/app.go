package app

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/kataras/golog"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/node"
	"github.com/sonr-io/core/pkg/common"
	"github.com/spf13/viper"
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
	Ctx  context.Context
	Node api.NodeImpl
	Mode node.StubMode
)

// Start starts the Sonr Node
func Start(req *api.InitializeRequest, options ...Option) {
	// Check if Node is already running
	if ok := HasStarted(); ok {
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

	// Set Logging Settings
	golog.SetLevel(opts.logLevel)
	golog.SetPrefix(Mode.Prefix())

	// Create Node
	Ctx = context.Background()
	err := req.Parse()
	if err != nil {
		golog.Default.Child("(app)").Fatalf("%s - Failed to parse Initialize Request", err)
	}
	// Open Listener on Port
	listener, err := net.Listen(opts.network, opts.Address())
	if err != nil {
		golog.Default.Child("(app)").Fatalf("%s - Failed to Create New Listener", err)
	}

	// Set Node Stub
	Node, _, err = node.NewNode(Ctx, listener,
		node.WithMode(Mode),
		node.WithRequest(req))
	if err != nil {
		golog.Default.Child("(app)").Fatalf("%s - Failed to Start new Node", err)
	}

	// Serve Node for GRPC
	Persist(listener)
}

// Persist contains the main loop for the Node
func Persist(l net.Listener) {
	// Check if CLI Mode
	if common.IsMobile() {
		golog.Default.Child("(app)").Info("Skipping Serve, Node is mobile...")
		return
	}

	// Wait for Exit Signal
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		Exit(0)
	}()

	// Hold until Exit Signal
	for {
		// Start GRPC Server
		golog.Default.Child("(app)").Infof("Starting GRPC Server on %s", l.Addr().String())
		select {
		case <-Ctx.Done():
			golog.Default.Child("(app)").Info("Context Done")
			l.Close()
			return
		}
	}
}

// Exit handles cleanup on Sonr Node
func Exit(code int) {
	if Node == nil {
		golog.Default.Child("(app)").Debug("Skipping Exit, instance is nil...")
		return
	}
	golog.Default.Child("(app)").Debug("Cleaning up Node on Exit...")
	Node.Close()

	defer Ctx.Done()

	// Check for Full Desktop Node
	if common.IsDesktop() {
		golog.Default.Child("(app)").Debug("Removing Bitcask DB...")
		ex, err := os.Executable()
		if err != nil {
			golog.Default.Child("(app)").Errorf("%s - Failed to find Executable", err)
			return
		}

		// Delete Executable Path
		exPath := filepath.Dir(ex)
		err = os.RemoveAll(filepath.Join(exPath, "sonr_bitcask"))
		if err != nil {
			golog.Default.Child("(app)").Warn("Failed to remove Bitcask, ", err)
		}
		err = viper.SafeWriteConfig()
		if err == nil {
			golog.Default.Child("(app)").Debug("Wrote new config file to Disk")
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
