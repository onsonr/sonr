package node

import (
	"context"
	"errors"
	"fmt"

	"github.com/kataras/golog"
	api "github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/common"
)

// Error Definitions
var (
	logger             = golog.Child("internal/node")
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

// NodeOption is a function that modifies the node options.
type NodeOption func(*nodeOptions)

// WithRequest sets the initialize request.
func WithRequest(req *api.InitializeRequest) NodeOption {
	return func(o *nodeOptions) {
		o.location = req.GetLocation()
		o.profile = req.GetProfile()
		o.connection = req.GetConnection()
	}
}

// WitHighway starts the Client RPC server as a highway node.
func WitHighway() NodeOption {
	return func(o *nodeOptions) {
		o.mode = StubMode_HIGHWAY
	}
}

// WithTerminal sets the node as a terminal node.
func WithTerminal() NodeOption {
	return func(o *nodeOptions) {
		o.isTerminal = true
	}
}

// nodeOptions is a collection of options for the node.
type nodeOptions struct {
	address    string
	connection common.Connection
	location   *common.Location
	mode       NodeStubMode
	network    string
	profile    *common.Profile
	isTerminal bool
}

// defaultNodeOptions returns the default node options.
func defaultNodeOptions() *nodeOptions {
	return &nodeOptions{
		mode:       StubMode_CLIENT,
		location:   common.DefaultLocation(),
		connection: common.Connection_WIFI,
		network:    "tcp",
		address:    fmt.Sprintf(":%d", common.RPC_SERVER_PORT),
		profile:    common.NewDefaultProfile(),
		isTerminal: false,
	}
}

func (opts *nodeOptions) Apply(ctx context.Context, node *Node) error {
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
		node.stub = stub

	} else {
		logger.Debug("Starting Highway stub...")
		// Highway Node Type
		stub, err := node.startHighwayService(ctx, opts)
		if err != nil {
			logger.Error("Failed to start Highway Service", err)
			return err
		}

		// Set Stub to node
		node.stub = stub
	}
	return nil
}
