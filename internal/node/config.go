package node

import (
	"context"
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"

	"github.com/kataras/golog"
	"github.com/pterm/pterm"
	api "github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/pkg/common"
)

// Error Definitions
var (
	logger             = golog.Child("internal/node")
	progressPrinter    *pterm.ProgressbarPrinter
	ErrEmptyQueue      = errors.New("No items in Transfer Queue.")
	ErrInvalidQuery    = errors.New("No SName or PeerID provided.")
	ErrNBClientMissing = errors.New("No Namebase API Client Key provided.")
	ErrNBSecretMissing = errors.New("No Namebase API Secret Key provided.")
	ErrMissingParam    = errors.New("Paramater is missing.")
	ErrProtocolsNotSet = errors.New("Node Protocol has not been initialized.")
)

// NodeStubMode is the type of the node (Client, Highway)
type NodeStubMode int

const (
	// StubMode_CLIENT is the Node utilized by Desktop, Mobile and Web Clients
	StubMode_CLIENT NodeStubMode = iota

	// StubMode_HIGHWAY is the Node utilized by long running Server processes
	StubMode_HIGHWAY
)

// IsClient returns true if the node is a client node.
func (m NodeStubMode) IsClient() bool {
	return m == StubMode_CLIENT
}

// IsHighway returns true if the node is a highway node.
func (m NodeStubMode) IsHighway() bool {
	return m == StubMode_HIGHWAY
}

// Option is a function that modifies the node options.
type Option func(*options)

// WithRequest sets the initialize request.
func WithRequest(req *api.InitializeRequest) Option {
	return func(o *options) {
		o.location = req.GetLocation()
		o.profile = req.GetProfile()
		o.connection = req.GetConnection()
	}
}

// SetHighway starts the Client RPC server as a highway node.
func SetHighway() Option {
	return func(o *options) {
		o.mode = StubMode_HIGHWAY
	}
}

// WithTerminal sets the node as a terminal node.
func SetTerminalMode(val bool) Option {
	return func(o *options) {
		o.isTerminal = val
	}
}

// options is a collection of options for the node.
type options struct {
	address    string
	connection common.Connection
	location   *common.Location
	mode       NodeStubMode
	network    string
	profile    *common.Profile
	isTerminal bool
}

// defaultNodeOptions returns the default node options.
func defaultNodeOptions() *options {
	return &options{
		mode:       StubMode_CLIENT,
		location:   common.DefaultLocation(),
		connection: common.Connection_WIFI,
		network:    "tcp",
		address:    fmt.Sprintf(":%d", common.RPC_SERVER_PORT),
		profile:    common.NewDefaultProfile(),
		isTerminal: false,
	}
}

// Apply applies Options to node
func (opts *options) Apply(ctx context.Context, node *Node) error {
	node.isTerminal = opts.isTerminal
	node.mode = opts.mode

	// Handle by Node Mode
	if opts.mode == StubMode_CLIENT {
		logger.Debug("Starting Client stub...")
		// Client Node Type
		stub, err := node.startClientService(ctx, opts)
		if err != nil {
			logger.Error("Failed to start Client Service", err)
			return err
		}

		// Set Stub to node
		node.clientStub = stub

	} else {
		logger.Debug("Starting Highway stub...")
		// Highway Node Type
		stub, err := node.startHighwayService(ctx, opts)
		if err != nil {
			logger.Error("Failed to start Highway Service", err)
			return err
		}

		// Set Stub to node
		node.highwayStub = stub
	}
	return nil
}

// printTerminal is a helper function that prints to the terminal.
func (n *Node) printTerminal(title string, msg string) {
	if n.isTerminal {
		// Print a section with level one.
		pterm.DefaultSection.Println(title)
		// Print placeholder.
		pterm.Info.Println(msg)
	}
}

// promptTerminal is a helper function that prompts the user for input.
func (n *Node) promptTerminal(title string, onResult func(result bool)) error {
	if n.isTerminal {
		prompt := promptui.Prompt{
			Label:     title,
			IsConfirm: true,
		}

		result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}
		onResult(result == "y")
	}
	return nil
}

// progressTerminal is a helper function that prints a progress bar to the terminal.
func (n *Node) progressTerminal(ev *api.ProgressEvent) {
	if n.isTerminal {
		// Create a new progress bar if it doesn't exist.
		if progressPrinter == nil {
			var err error
			progressPrinter, err = pterm.DefaultProgressbar.Start()
			if err != nil {
				logger.Error("Failed to start progressbar", err)
				return
			}
		}

		// While progress is not complete, update the progress bar.
		for ev.GetProgress() < 100 {
			progressPrinter.Add(int(ev.Progress))
		}
	}
}
