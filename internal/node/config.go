package node

import (
	"context"
	"errors"
	"fmt"
	"os"

	olc "github.com/google/open-location-code/go"
	"github.com/kataras/golog"
	api "github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/common"

	"google.golang.org/protobuf/proto"
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
	// Set Profile buffer
	profile := common.NewDefaultProfile(common.WithCheckerProfile(req.GetProfile()), common.WithPicture())
	proBuf, err := proto.Marshal(profile)
	if err != nil {
		logger.Child("Config").Error("Failed to marshal Profile", err)
	}
	code := olc.Encode(req.GetLocation().GetLatitude(), req.GetLocation().GetLongitude(), 8)
	if code == "" {
		logger.Child("Config").Error("Failed to Determine OLC Code, set to Global")
		code = "global"
	}

	return func(o *nodeOptions) {
		// Set Connection
		o.connection = req.Connection

		// Set Env Variables
		if req.Variables != nil {
			for k, v := range req.Variables {
				os.Setenv(k, v)
			}

			if len(req.Variables) > 0 {
				logger.Info("Added Enviornment Variable(s)", golog.Fields{
					"Total": len(req.Variables),
				})
			}
		}

		// Set Properties
		o.olc = code
		o.profileBuf = proBuf
	}
}

// WithMode starts the Client RPC server and sets the node as a client node.
func WithMode(m NodeStubMode) NodeOption {
	return func(o *nodeOptions) {
		o.mode = m
	}
}

// nodeOptions is a collection of options for the node.
type nodeOptions struct {
	mode       NodeStubMode
	network    string
	address    string
	profileBuf []byte
	connection common.Connection
	olc        string
}

// defaultNodeOptions returns the default node options.
func defaultNodeOptions() *nodeOptions {
	return &nodeOptions{
		mode:       StubMode_CLIENT,
		olc:        "global",
		connection: common.Connection_WIFI,
		network:    "tcp",
		address:    fmt.Sprintf(":%d", common.RPC_SERVER_PORT),
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
