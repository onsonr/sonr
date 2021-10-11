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
	logger                = golog.Child("internal/node")
	ErrEmptyQueue         = errors.New("No items in Transfer Queue.")
	ErrInvalidQuery       = errors.New("No SName or PeerID provided.")
	ErrNBClientMissing    = errors.New("No Namebase API Client Key provided.")
	ErrNBSecretMissing    = errors.New("No Namebase API Secret Key provided.")
	ErrRecentsNotCreated  = errors.New("Recents has not been created yet.")
	ErrProfileNotCreated  = errors.New("Profile has not been created yet.")
	ErrProfileNotProvided = errors.New("Profile has not been provided to Store.")
	ErrProfileIsOlder     = errors.New("Profile is older than the oldest one on disk.")
	ErrProfileNoTimestamp = errors.New("Profile has no timestamp.")
	ErrStoreNotCreated    = errors.New("Node Store has not been opened/created.")
	ErrLobbyNotCreated    = errors.New("LobbyProtocol has not been created")
	ErrExchangeNotCreated = errors.New("ExchangeProtocol has not been created")
	ErrTransferNotCreated = errors.New("TransferProtocol has not been created")
)

// NodeStub is the interface for the node based on mode: (client, highway)
type NodeStub interface {
	Close() error
}

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

// WithStubMode starts the Client RPC server and sets the node as a client node.
func WithStubMode(m NodeStubMode) NodeOption {
	return func(o *nodeOptions) {
		o.mode = m
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
}

// defaultNodeOptions returns the default node options.
func defaultNodeOptions() *nodeOptions {
	return &nodeOptions{
		mode:       StubMode_CLIENT,
		location:   &common.Location{},
		connection: common.Connection_WIFI,
		network:    "tcp",
		address:    fmt.Sprintf(":%d", common.RPC_SERVER_PORT),
		profile:    common.NewDefaultProfile(),
	}
}

func (opts *nodeOptions) Apply(ctx context.Context, node *Node) error {
	// Handle by Node Mode
	if opts.mode == StubMode_CLIENT {
		// Client Node Type
		stub, err := node.startClientService(ctx, opts)
		if err != nil {
			logger.Error("Failed to start Client Service", err)
			return err
		}

		// Set Stub to node
		node.stub = stub

	} else {
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
