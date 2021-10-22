package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/kataras/golog"
	"github.com/pterm/pterm"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/node"
	"github.com/sonr-io/core/pkg/common"
	"github.com/spf13/viper"
)

func init() {
	golog.SetStacktraceLimit(2)
}

// Sonr is the main struct for Sonr Node
type Sonr struct {
	// Properties
	Ctx  context.Context
	Node api.NodeImpl
	Mode node.StubMode
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

	// Create Node
	ctx := context.Background()
	n, _, err := node.NewNode(ctx, opts.Apply(req)...)
	if err != nil {
		golog.Fatal("Failed to update Profile for Node", golog.Fields{"error": err})
		os.Exit(1)
	}

	// Set Lib
	instance = Sonr{
		Ctx:  ctx,
		Mode: opts.mode,
		Node: n,
	}
	instance.Serve()
}

// AppHeader prints Node Info onto Terminal
func AppHeader() {
	// Get Peer Info
	p, err := instance.Node.Peer()
	if err != nil {
		golog.Error("Failed to get Peer", golog.Fields{"error": err})
		Exit(1)
	}

	// Print Header on Terminal CLI Mode
	if instance.Mode.IsCLI() {
		pterm.DefaultSection.Println(fmt.Sprintf("Sonr Node Online: %s", p.PeerID))
		pterm.Info.Println(fmt.Sprintf("SName: %s \nOS: %s \nArch: %s", p.GetSName(), p.OS(), p.Arch()))
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

// Serve waits for Exit Signal from Terminal
func (sh Sonr) Serve() {
	// Check if CLI Mode
	if common.IsMobile() {
		golog.Info("Skipping Serve, Node is mobile...")
		return
	}

	// Wait for Exit Signal
	AppHeader()
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		Exit(0)
	}()

	// Hold until Exit Signal
	for {
		select {
		case <-sh.Ctx.Done():
			golog.Info("Context Done")
			return
		}
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
	host string
	port int
	mode node.StubMode
}

// defaultOptions returns the default options
func defaultOptions() *options {
	return &options{
		host: ":",
		port: 26225,
		mode: node.StubMode_LIB,
	}
}

// Apply applies the options to the request
func (o *options) Apply(r *api.InitializeRequest) []node.Option {
	return []node.Option{
		node.WithHost(o.host),
		node.WithPort(o.port),
		node.WithMode(o.mode),
		node.WithRequest(r),
	}
}
