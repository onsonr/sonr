// ---
// Motor Implementation
//
// Motor is an embedded light node which interacts with the Sonr network. Motors utilize
// the Sonr protocol to leverage account assets in a secure and efficient manner.
//
// Modules: DIDComm, MPC Wallet, Blockchain Client
// Interface: libp2p host
// Transports: TCP, UDP, QUIC, WebTransport, WebSockets
// ---

package internal

import (
	"context"

	"github.com/sonrhq/core/pkg/node"
	"github.com/sonrhq/core/pkg/node/config"
	types "github.com/sonrhq/core/types/common"
)

type MotorInstance struct {
	// Node is the libp2p host
	Node node.Node
	//
}

// It creates a new node and wallet, and returns a MotorNode struct containing them
func NewMotorInstance(ctx context.Context) (*MotorInstance, error) {
	n, err := node.New(ctx, config.WithPeerType(types.PeerType_MOTOR))
	if err != nil {
		return nil, err
	}

	return &MotorInstance{
		Node: n,
	}, nil
}
