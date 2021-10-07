package node

import (
	"context"
	"errors"
	"os"

	olc "github.com/google/open-location-code/go"
	"github.com/kataras/golog"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/pkg/transfer"

	"google.golang.org/protobuf/proto"
)

// Error Definitions
var (
	logger                = golog.Child("Node")
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

// Initialize initializes the node by Type.
func (nt NodeType) Initialize(ctx context.Context, n *Node, olc string) {
	switch nt {
	case NodeType_CLIENT:
		n.startClientService(ctx, olc)
	case NodeType_HIGHWAY:
		n.startHighwayService(ctx)
	}
}

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
			logger.Error("Failed to Determine OLC Code, set to Global")
			o.olc = "global"
		} else {
			o.olc = code
		}

		// Set Profile buffer
		profile := common.NewDefaultProfile(common.WithCheckerProfile(req.GetProfile()), common.WithPicture())
		proBuf, err := proto.Marshal(profile)
		if err != nil {
			logger.Error("Failed to marshal Profile", err)
		}
		o.profileBuf = proBuf
	}
}

// WithClient starts the Client RPC server and sets the node as a client node.
func WithClient() NodeOption {
	return func(o nodeOptions) {
		o.kind = NodeType_CLIENT
		o.enableTCPPort = true
	}
}

// WithHighway starts the Highway RPC server and sets the node as a highway node.
func WithHighway() NodeOption {
	return func(o nodeOptions) {
		o.kind = NodeType_HIGHWAY
		o.enableTCPPort = false
	}
}

// nodeOptions is a collection of options for the node.
type nodeOptions struct {
	enableTCPPort bool
	kind          NodeType
	profileBuf    []byte
	connection    common.Connection
	olc           string
}

// defaultNodeOptions returns the default node options.
func defaultNodeOptions() nodeOptions {
	return nodeOptions{
		enableTCPPort: true,
		kind:          NodeType_CLIENT,
		olc:           "global",
		connection:    common.Connection_WIFI,
	}
}

// Apply applies the node options to the node.
func (no nodeOptions) Apply(ctx context.Context, n *Node) {
	no.kind.Initialize(ctx, n, no.olc)
}

// Share a peer to have a transfer
func (n *Node) NewRequest(to *common.Peer) (peer.ID, *transfer.InviteRequest, error) {
	// Fetch Element from Queue
	elem := n.queue.Front()
	if elem != nil {
		// Get Payload
		payload := n.queue.Remove(elem).(*common.Payload)

		// Create New ID for Invite
		id, err := device.KeyChain.CreateUUID()
		if err != nil {
			logger.Error("Failed to create new id for Shared Invite", err)
			return "", nil, err
		}

		// Create new Metadata
		meta, err := device.KeyChain.CreateMetadata(n.host.ID())
		if err != nil {
			logger.Error("Failed to create new metadata for Shared Invite", err)
			return "", nil, err
		}

		// Fetch User Peer
		from, err := n.Peer()
		if err != nil {
			logger.Error("Failed to get Node Peer Object", err)
			return "", nil, err
		}

		// Create Invite Request
		req := &transfer.InviteRequest{
			Payload:  payload,
			Metadata: common.SignedMetadataToProto(meta),
			To:       to,
			From:     from,
			Uuid:     common.SignedUUIDToProto(id),
		}

		// Fetch Peer ID from Public Key
		toId, err := to.PeerID()
		if err != nil {
			logger.Error("Failed to fetch peer id from public key", err)
			return "", nil, err
		}
		return toId, req, nil
	}
	return "", nil, errors.New("No items in Transfer Queue.")
}

// Respond to an invite request
func (n *Node) NewResponse(decs bool, to *common.Peer) (peer.ID, *transfer.InviteResponse, error) {
	// Create new Metadata
	meta, err := device.KeyChain.CreateMetadata(n.host.ID())
	if err != nil {
		logger.Error("Failed to create new metadata for Shared Invite", err)
		return "", nil, err
	}

	// Fetch User Peer
	from, err := n.Peer()
	if err != nil {
		return "", nil, err
	}

	// Create Invite Response
	resp := &transfer.InviteResponse{
		Decision: decs,
		Metadata: common.SignedMetadataToProto(meta),
		From:     from,
		To:       to,
	}

	// Fetch Peer ID from Public Key
	toId, err := to.PeerID()
	if err != nil {
		logger.Error("Failed to fetch peer id from public key", err)
		return "", nil, err
	}
	return toId, resp, nil
}
