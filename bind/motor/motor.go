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

package motor

import (
	"context"

	"github.com/sonr-hq/sonr/pkg/common"
	"github.com/sonr-hq/sonr/pkg/node"
	"github.com/sonr-hq/sonr/pkg/node/config"
)

type MotorInstance struct {
	// Node is the libp2p host
	Node node.Node
	//
}

// It creates a new node and wallet, and returns a MotorNode struct containing them
func NewMotorInstance(ctx context.Context) (*MotorInstance, error) {
	n, err := node.New(ctx, config.WithPeerType(common.PeerType_MOTOR))
	if err != nil {
		return nil, err
	}

	return &MotorInstance{
		Node: n,
	}, nil
}
