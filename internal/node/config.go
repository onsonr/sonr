package node

import (
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/tools/logger"
)

// NodeOption is a function that modifies the node options.
type NodeOption func(nodeOptions) nodeOptions

// nodeOptions is a collection of options for the node.
type nodeOptions struct {
	isClient  bool
	isHighway bool
	request   *InitializeRequest
}

// GetNodeType returns the node type from Config
func (no nodeOptions) GetNodeType() NodeType {
	if no.isHighway {
		return NodeType_HIGHWAY
	}
	return NodeType_CLIENT
}

// GetConnection returns the connection type.
func (no nodeOptions) GetConnection() common.Connection {
	return no.request.GetConnection()
}

// GetLocation returns the location of the node.
func (no nodeOptions) GetLocation() *common.Location {
	// Check if the request has a location
	if no.request.Location != nil {
		logger.Warn("No Location was set.")
		// Return Default
		return &common.Location{
			Latitude:  0,
			Longitude: 0,
		}
	}
	return no.request.GetLocation()
}

// WithRequest sets the initialize request.
func WithRequest(req *InitializeRequest) NodeOption {
	return func(o nodeOptions) nodeOptions {
		o.request = req
		return o
	}
}

// WithClient starts the Client RPC server and sets the node as a client node.
func WithClient() NodeOption {
	return func(o nodeOptions) nodeOptions {
		o.isClient = true
		o.isHighway = false
		return o
	}
}

// WithHighway starts the Highway RPC server and sets the node as a highway node.
func WithHighway() NodeOption {
	return func(o nodeOptions) nodeOptions {
		o.isHighway = true
		o.isClient = false
		return o
	}
}

// defaultNodeOptions returns the default node options.
func defaultNodeOptions() nodeOptions {
	return nodeOptions{
		isClient:  true,
		isHighway: false,
		request: &InitializeRequest{
			Connection: common.Connection_MOBILE,
		},
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
