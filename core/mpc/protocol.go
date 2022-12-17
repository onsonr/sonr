package mpc

import (
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/sonr-hq/sonr/internal/node"
	"github.com/sonr-io/multi-party-sig/pkg/protocol"
)

type MpcProtocol struct {
	selfNode  *node.Node
	peers     []peer.ID
	threshold int
	messageCh chan *protocol.Message
}
