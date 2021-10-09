package node

import (
	"context"
	"errors"
	"net"
	"os"

	olc "github.com/google/open-location-code/go"
	"github.com/kataras/golog"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/tools/state"

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

func init() {

}

// NodeType is the type of the node (Client, Highway)
type NodeType int

const (
	// NodeType_CLIENT is the Node utilized by Desktop, Mobile and Web Clients
	NodeType_CLIENT NodeType = iota

	// NodeType_HIGHWAY is the Node utilized by long running Server processes
	NodeType_HIGHWAY
)

// NodeOption is a function that modifies the node options.
type NodeOption func(nodeOptions)

// WithRequest sets the initialize request.
func WithRequest(req *InitializeRequest) NodeOption {
	return func(o nodeOptions) {
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

		// Set OLC code
		code := olc.Encode(req.GetLocation().GetLatitude(), req.GetLocation().GetLongitude(), 8)
		if code == "" {
			logger.Child("Config").Error("Failed to Determine OLC Code, set to Global")
			o.olc = "global"
		} else {
			o.olc = code
		}

		// Set Profile buffer
		profile := common.NewDefaultProfile(common.WithCheckerProfile(req.GetProfile()), common.WithPicture())
		proBuf, err := proto.Marshal(profile)
		if err != nil {
			logger.Child("Config").Error("Failed to marshal Profile", err)
		}
		o.profileBuf = proBuf
	}
}

// WithClient starts the Client RPC server and sets the node as a client node.
func WithClient() NodeOption {
	return func(o nodeOptions) {
		o.kind = NodeType_CLIENT
	}
}

// WithEmitter sets the emitter for the node.
func WithEmitter(e *state.Emitter) NodeOption {
	return func(o nodeOptions) {
		o.emitter = e
	}
}

// WithHighway starts the Highway RPC server and sets the node as a highway node.
func WithHighway() NodeOption {
	return func(o nodeOptions) {
		o.kind = NodeType_HIGHWAY
	}
}

// WithListener sets the TCP Listener for Client stub
func WithListener(l net.Listener) NodeOption {
	return func(o nodeOptions) {
		o.listener = l
	}
}

// nodeOptions is a collection of options for the node.
type nodeOptions struct {
	emitter    *state.Emitter
	kind       NodeType
	listener   net.Listener
	profileBuf []byte
	connection common.Connection
	olc        string
}

// defaultNodeOptions returns the default node options.
func defaultNodeOptions() nodeOptions {
	return nodeOptions{
		emitter:    state.NewEmitter(2048),
		kind:       NodeType_CLIENT,
		olc:        "global",
		connection: common.Connection_WIFI,
	}
}

// Apply applies the node options to the node.
func (no nodeOptions) Apply(ctx context.Context, n *Node) error {
	if no.kind == NodeType_CLIENT {
		// Client Node Type
		_, err := n.startClientService(ctx, no.olc)
		if err != nil {
			logger.Error("Failed to start Client Service", err)
			return err
		}
	} else {
		// Highway Node Type
		_, err := n.startHighwayService(ctx)
		if err != nil {
			logger.Error("Failed to start Highway Service", err)
			return err
		}
	}
	return nil
}
