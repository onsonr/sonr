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
	ErrEmptyQueue      = errors.New("No items in Transfer Queue.")
	ErrInvalidQuery    = errors.New("No SName or PeerID provided.")
	ErrMissingParam    = errors.New("Paramater is missing.")
	ErrProtocolsNotSet = errors.New("Node Protocol has not been initialized.")
)

// StubMode is the type of the node (Client, Highway)
type StubMode int

const (
	// StubMode_LIB is the Node utilized by Mobile and Web Clients
	StubMode_LIB StubMode = iota

	// StubMode_CLI is the Node utilized by CLI Clients
	StubMode_CLI

	// StubMode_BIN is the Node utilized for Desktop background process
	StubMode_BIN

	// StubMode_HIGHWAY is the Custodian Node that manages Network
	StubMode_HIGHWAY
)

// IsLib returns true if the node is a client node.
func (m StubMode) IsLib() bool {
	return m == StubMode_LIB
}

// IsBin returns true if the node is a bin node.
func (m StubMode) IsBin() bool {
	return m == StubMode_BIN
}

// IsCLI returns true if the node is a CLI node.
func (m StubMode) IsCLI() bool {
	return m == StubMode_CLI
}

// IsHighway returns true if the node is a highway node.
func (m StubMode) IsHighway() bool {
	return m == StubMode_HIGHWAY
}

// Prefix returns golog prefix for the node.
func (m StubMode) Prefix() string {
	var name string
	switch m {
	case StubMode_LIB:
		name = "lib"
	case StubMode_CLI:
		name = "cli"
	case StubMode_BIN:
		name = "bin"
	case StubMode_HIGHWAY:
		name = "highway"
	default:
		name = "unknown"
	}
	return fmt.Sprintf("[SONR.%s] ", name)
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

// WithMode starts the Client RPC server as a highway node.
func WithMode(m StubMode) Option {
	return func(o *options) {
		o.mode = m
	}
}

// options is a collection of options for the node.
type options struct {
	address    string
	connection common.Connection
	location   *common.Location
	mode       StubMode
	network    string
	profile    *common.Profile
}

// defaultNodeOptions returns the default node options.
func defaultNodeOptions() *options {
	return &options{
		mode:       StubMode_LIB,
		location:   common.DefaultLocation(),
		connection: common.Connection_WIFI,
		network:    "tcp",
		address:    fmt.Sprintf(":%d", common.RPC_SERVER_PORT),
		profile:    common.NewDefaultProfile(),
	}
}

// Apply applies Options to node
func (opts *options) Apply(ctx context.Context, node *Node) error {
	node.mode = opts.mode

	// Handle by Node Mode
	if opts.mode == StubMode_LIB {
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
	if n.mode.IsCLI() {
		// Print a section with level one.
		pterm.DefaultSection.Println(title)
		// Print placeholder.
		pterm.Info.Println(msg)
	}
}

// promptTerminal is a helper function that prompts the user for input.
func (n *Node) promptTerminal(title string, onResult func(result bool)) error {
	if n.mode.IsCLI() {
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
