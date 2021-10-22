package app

import (
	"context"
	"fmt"
	"net"
	"os"
	"path/filepath"

	"github.com/kataras/golog"
	"github.com/pterm/pterm"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/node"
	"github.com/sonr-io/core/pkg/common"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func init() {
	golog.SetStacktraceLimit(2)
}

// Sonr is the main struct for Sonr Node
type Sonr struct {
	// Properties
	Ctx        context.Context
	Node       api.NodeImpl
	Mode       node.StubMode
	Listener   net.Listener
	GRPCServer *grpc.Server
}

// instance is the global Sonr Instance
var instance Sonr

// Start starts the Sonr Node
func Start(req *api.InitializeRequest, options ...Option) {
	// Check if Node is already running
	if instance.Node != nil {
		golog.Error("Sonr Instance already active")
		return
	}

	// Set Options
	opts := defaultOptions()
	for _, o := range options {
		o(opts)
	}

	// Set Logging Settings
	golog.SetPrefix(opts.mode.Prefix())
	if opts.mode.IsCLI() {
		pterm.SetDefaultOutput(golog.Default.Printer)
	}

	// Initialize Wallet, and FS
	err := req.Parse()
	if err != nil {
		golog.Fatal("Failed to initialize Device", golog.Fields{"error": err})
		os.Exit(1)
	}

	// Apply Options
	err = opts.Apply(req)
	if err != nil {
		golog.Fatal("Failed to initialize Sonr", golog.Fields{"error": err})
		os.Exit(1)
	}

	// Handle Node Events
	if err := instance.GRPCServer.Serve(instance.Listener); err != nil {
		golog.Fatal("Failed to serve gRPC", err)
	}
}

// Exit handles cleanup on Sonr Node
func Exit(code int) {
	if instance.Node == nil {
		golog.Info("Skipping Exit, instance is nil...")
		return
	}
	golog.Info("Cleaning up on Exit...")
	instance.Node.Close()
	instance.GRPCServer.Stop()
	instance.Listener.Close()
	defer instance.Ctx.Done()

	// Check for Full Desktop Node
	if common.IsDesktop() {
		ex, err := os.Executable()
		if err != nil {
			golog.Error("Failed to find Executable, ", err)
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
			golog.Info("Wrote new config file to Disk")
		}
		os.Exit(code)
	}
}

// Option is a function that can be passed to Start
type Option func(*options)

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
	host    string
	network string
	port    int
	mode    node.StubMode
}

// Address returns the address of the node.
func (opts *options) Address() string {
	return fmt.Sprintf("%s%d", opts.host, opts.port)
}

// defaultOptions returns the default options
func defaultOptions() *options {
	return &options{
		host:    ":",
		port:    26225,
		mode:    node.StubMode_LIB,
		network: "tcp",
	}
}

// Apply applies the options to the request
func (o *options) Apply(req *api.InitializeRequest) error {
	// Create Node
	ctx := context.Background()

	// Open Listener on Port
	listener, err := net.Listen(o.network, o.Address())
	if err != nil {
		return err
	}

	// Set Instance
	instance = Sonr{
		Ctx:        ctx,
		Mode:       o.mode,
		Listener:   listener,
		GRPCServer: grpc.NewServer(),
	}

	// Set Node Stub
	n, _, err := node.NewNode(ctx, node.WithGRPC(instance.GRPCServer),
		node.WithMode(o.mode),
		node.WithRequest(req))
	if err != nil {
		return err
	}
	instance.Node = n
	return nil
}
