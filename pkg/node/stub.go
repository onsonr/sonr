package node

import (
	"errors"

	"github.com/sonr-hq/sonr/pkg/common"
	"github.com/sonr-hq/sonr/pkg/node/config"
)

type node struct {
	host   config.P2PNode
	ipfs   config.IPFSNode
	config *config.Config
}

func makeNode(host config.P2PNode, ipfs config.IPFSNode, config *config.Config) (*node, error) {
	if host == nil && ipfs == nil {
		return nil, errors.New("Invalid node. Must have either a host or an ipfs instance.")
	}
	return &node{
		host:   host,
		ipfs:   ipfs,
		config: config,
	}, nil
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
