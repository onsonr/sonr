package node

import (
	"context"
	"errors"
	"os"

	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/pkg/exchange"
	"github.com/sonr-io/core/tools/logger"
	"go.uber.org/zap"
)

// Error Definitions
var (
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

// nodeOptions is a collection of options for the node.
type nodeOptions struct {
	isClient  bool
	isHighway bool
	request   *InitializeRequest
}

// Apply applies the node options to the node.
func (no nodeOptions) Apply(ctx context.Context, n *Node) {
	n.profile = no.request.GetProfile()
	no.GetNodeType().Initialize(ctx, n, no.GetLocalOLC())
}

// GetConnection returns the node internet connection type
func (no nodeOptions) GetConnection() common.Connection {
	return no.request.GetConnection()
}

// GetLocalOLC returns the local OLC for LobbyProtocol
func (no nodeOptions) GetLocalOLC() string {
	return no.request.GetLocation().OLC()
}

// GetNodeType returns the node type from Config
func (no nodeOptions) GetNodeType() NodeType {
	if no.isHighway {
		return NodeType_HIGHWAY
	}
	return NodeType_CLIENT
}

// GetIPAddresses returns host.HostListenAddr from hostOpts
func (no nodeOptions) GetIPAddresses() []host.HostListenAddr {
	// Define Listen Addresses
	providedAddrs := no.request.GetHostOptions().GetListenAddrs()
	addrs := make([]host.HostListenAddr, len(providedAddrs))

	// Iterate over provided addresses
	for i, addr := range providedAddrs {
		addrs[i] = host.HostListenAddr{
			Addr:   addr.GetAddress(),
			Family: addr.GetFamily().String(),
		}
	}
	return addrs
}

// WithRequest sets the initialize request.
func WithRequest(req *InitializeRequest) NodeOption {
	return func(o nodeOptions) {
		o.request = req
	}
}

// WithClient starts the Client RPC server and sets the node as a client node.
func WithClient() NodeOption {
	return func(o nodeOptions) {
		o.isClient = true
		o.isHighway = false
	}
}

// WithClient starts the Client RPC server and sets the node as a client node.
func WithEnvMap(vars map[string]string) NodeOption {
	return func(o nodeOptions) {
		for k, v := range vars {
			os.Setenv(k, v)
		}

		if len(vars) > 0 {
			logger.Info("Added Enviornment Variable(s)", zap.Int("Total", len(vars)))
		}
	}
}

// WithHighway starts the Highway RPC server and sets the node as a highway node.
func WithHighway() NodeOption {
	return func(o nodeOptions) {
		o.isHighway = true
		o.isClient = false
	}
}

// defaultNodeOptions returns the default node options.
func defaultNodeOptions() nodeOptions {
	return nodeOptions{
		isClient:  true,
		isHighway: false,
	}
}

// newInitResponse creates a response for the initialize request.
func (n *Node) newInitResponse(err error) *InitializeResponse {
	// Check for provided error
	if err != nil {
		return &InitializeResponse{
			Success: false,
			Error:   err.Error(),
		}
	}

	// Return Response
	return &InitializeResponse{
		Success: true,
		Profile: n.profile,
		//	Recents: r,
	}
}

// ToExchangeQueryRequest converts a query request to an exchange query request.
func (f *SearchRequest) ToExchangeQueryRequest() (*exchange.QueryRequest, error) {
	if f.GetSName() != "" {
		return &exchange.QueryRequest{
			SName: f.GetSName(),
		}, nil
	}

	if f.GetPeerId() != "" {
		return &exchange.QueryRequest{
			PeerId: f.GetPeerId(),
		}, nil
	}
	return nil, logger.Error("Failed to convert FindRequest", ErrInvalidQuery)
}

// ToFindResponse converts PeerInfo to a FindResponse.
func ToFindResponse(p *common.PeerInfo) *SearchResponse {
	return &SearchResponse{
		Success: true,
		Peer:    p.Peer,
		PeerId:  p.PeerID.String(),
		SName:   p.SName,
	}
}
