package node

import (
	"context"

	"github.com/sonrhq/core/pkg/common"
	identityprotocol "github.com/sonrhq/core/pkg/common"
	"github.com/sonrhq/core/pkg/node/config"
	"github.com/sonrhq/core/pkg/node/internal/host"
	"github.com/sonrhq/core/pkg/node/internal/ipfs"
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
	Host() common.P2PNode
	IPFS() common.IPFSNode
	Type() common.PeerType
}

// It creates a new host, and then creates a new node with that host
func New(ctx context.Context, opts ...config.Option) (Node, error) {
	pctx, err := identityprotocol.NewContext(ctx)
	if err != nil {
		return nil, err
	}
	config := config.DefaultConfig(pctx)
	err = config.Apply(opts...)
	if err != nil {
		return nil, err
	}
	if config.IsMotor() {
		h, err := host.Initialize(config)
		if err != nil {
			return nil, err
		}
		return &node{
			host:   h,
			config: config,
		}, nil
	}
	i, err := ipfs.Initialize(config)
	if err != nil {
		return nil, err
	}
	return &node{
		ipfs:   i,
		config: config,
	}, nil
}

// NewIPFS creates a new IPFS node
func NewIPFS(ctx context.Context, opts ...config.Option) (common.IPFSNode, error) {
	// Start IPFS Node
	pctx, err := identityprotocol.NewContext(ctx)
	if err != nil {
		return nil, err
	}
	config := config.DefaultConfig(pctx)
	err = config.Apply(opts...)
	if err != nil {
		return nil, err
	}
	i, err := ipfs.Initialize(config)
	if err != nil {
		return nil, err
	}
	return i, nil
}

type node struct {
	host   common.P2PNode
	ipfs   common.IPFSNode
	config *config.Config
}

func (n *node) Host() common.P2PNode {
	return n.host
}

func (n *node) IPFS() common.IPFSNode {
	return n.ipfs
}

func (n *node) Type() common.PeerType {
	return n.config.PeerType
}
