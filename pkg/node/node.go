package node

import (
	"context"

	"github.com/sonr-hq/sonr/pkg/common"
	"github.com/sonr-hq/sonr/pkg/node/config"
	"github.com/sonr-hq/sonr/pkg/node/internal/host"
	"github.com/sonr-hq/sonr/pkg/node/internal/ipfs"
)

// `Node` is an interface that has three methods: `Host`, `IPFS`, and `Type`.
//
// The `Host` method returns a `Motor` interface and an error. The `IPFS` method returns a `Highway`
// interface and an error. The `Type` method returns a `Type` type.
//
// The `Motor` interface has two methods: `Start` and `Stop`. The `Start` method returns an error. The
// `Stop` method returns an error.
//
// The `Highway` interface has two methods: `Start` and
// @property Host - The motor that is hosting the node.
// @property IPFS - The IPFS node that the motor is connected to.
// @property {Type} Type - The type of node. This can be either a Motor or a Highway.
type Node interface {
	// Returning a Motor interface and an error.
	Host() config.P2PNode
	IPFS() config.IPFSNode
	Type() common.PeerType
}

// It creates a new host, and then creates a new node with that host
func New(ctx context.Context, opts ...config.Option) (Node, error) {
	config := config.DefaultConfig()
	err := config.Apply(opts...)
	if err != nil {
		return nil, err
	}
	if config.IsMotor() {
		h, err := host.Initialize(ctx, config)
		if err != nil {
			return nil, err
		}
		return &node{
			host:   h,
			config: config,
		}, nil

	}
	i, err := ipfs.Initialize(ctx, config)
	if err != nil {
		return nil, err
	}
	return &node{
		ipfs:   i,
		config: config,
	}, nil
}

// NewIPFS creates a new IPFS node
func NewIPFS(ctx context.Context, opts ...config.Option) (config.IPFSNode, error) {
	config := config.DefaultConfig()
	err := config.Apply(opts...)
	if err != nil {
		return nil, err
	}
	i, err := ipfs.Initialize(ctx, config)
	if err != nil {
		return nil, err
	}
	return i, nil
}

type node struct {
	host   config.P2PNode
	ipfs   config.IPFSNode
	config *config.Config
}

func (n *node) Host() config.P2PNode {
	return n.host
}

func (n *node) IPFS() config.IPFSNode {
	return n.ipfs
}

func (n *node) Type() common.PeerType {
	return n.config.PeerType
}
