package node

import (
	"errors"

	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/pkg/exchange"
	"github.com/sonr-io/core/tools/logger"
)

// Error Definitions
var (
	ErrEmptyQueue      = errors.New("No items in Transfer Queue.")
	ErrInvalidQuery    = errors.New("No SName or PeerID provided.")
	ErrNBClientMissing = errors.New("No Namebase API Client Key provided.")
	ErrNBSecretMissing = errors.New("No Namebase API Secret Key provided.")
)

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

// nodeOptions is a collection of options for the node.
type nodeOptions struct {
	isClient          bool
	isHighway         bool
	connection        common.Connection
	hostOpts          *InitializeRequest_HostOptions
	location          *common.Location
	profile           *common.Profile
	namebaseApiClient string
	namebaseApiSecret string
}

// GetNBKeys returns the namebase client,secret key
func (no nodeOptions) GetNBKeys() (string, string, error) {
	if no.namebaseApiClient == "" {
		return "", "", ErrNBClientMissing
	}

	if no.namebaseApiSecret == "" {
		return "", "", ErrNBSecretMissing
	}
	return no.namebaseApiClient, no.namebaseApiSecret, nil
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
	providedAddrs := no.hostOpts.GetListenAddrs()
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
		o.hostOpts = req.GetHostOptions()
		o.connection = req.GetConnection()
		o.location = req.GetLocation()
		o.profile = req.GetProfile()
	}
}

// WithClient starts the Client RPC server and sets the node as a client node.
func WithClient() NodeOption {
	return func(o nodeOptions) {
		o.isClient = true
		o.isHighway = false
	}
}

// WithHighway starts the Highway RPC server and sets the node as a highway node.
func WithHighway() NodeOption {
	return func(o nodeOptions) {
		o.isHighway = true
		o.isClient = false
	}
}

// WithNamebase sets the namebase client and secret api for DomainProtocol
func WithNamebase(client, secret string) NodeOption {
	return func(o nodeOptions) {
		o.namebaseApiClient = client
		o.namebaseApiSecret = secret
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

	// Fetch Profile
	p, err := n.Profile()
	if err != nil {
		return &InitializeResponse{
			Success: true,
			Error:   err.Error(),
		}
	}

	// Fetch Recents
	r, err := n.Recents()
	if err != nil {
		return &InitializeResponse{
			Success: true,
			Error:   err.Error(),
			Profile: p,
		}
	}

	// Return Response
	return &InitializeResponse{
		Success: true,
		Profile: p,
		Recents: r,
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
