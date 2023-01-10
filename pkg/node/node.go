package node

import (
	"context"
	"errors"

	"github.com/sonr-hq/sonr/pkg/common"
	"github.com/sonr-hq/sonr/pkg/node/host"
	"github.com/sonr-hq/sonr/pkg/node/ipfs"
)

type Node interface {
	Host() (*host.P2PHost, error)
	IPFS() (ipfs.IPFS, error)
	Type() common.PeerType
}

func NewMotor(ctx context.Context, opts ...host.Option) (Node, error) {
	h, err := host.New(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return &node{
		host:     h,
		peerType: common.PeerType_MOTOR,
	}, nil
}

func NewHighway(ctx context.Context, isLocal bool) (Node, error) {
	var err error
	var i ipfs.IPFS
	if isLocal {
		i, err = ipfs.NewLocal(ctx)
		if err != nil {
			return nil, err
		}
	} else {
		i, err = ipfs.NewRemote(ctx)
		if err != nil {
			return nil, err
		}
	}
	return &node{
		ipfs:     i,
		peerType: common.PeerType_HIGHWAY,
	}, nil
}

type node struct {
	host     *host.P2PHost
	ipfs     ipfs.IPFS
	peerType common.PeerType
}

func (n *node) Host() (*host.P2PHost, error) {
	if n.peerType != common.PeerType_MOTOR {
		return nil, errors.New("Invalid Node type. Only motors can return a libp2p host.")
	}
	return n.host, nil
}

func (n *node) IPFS() (ipfs.IPFS, error) {
	if n.peerType == common.PeerType_MOTOR {
		return nil, errors.New("Invalid Node type. Non-motors cannot return a ipfs instance.")
	}
	return n.ipfs, nil
}

func (n *node) Type() common.PeerType {
	return n.peerType
}
