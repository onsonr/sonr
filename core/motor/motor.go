// ---
// Motor Implementation
//
// Motor is an embedded light node which interacts with the Sonr network. Motors utilize
// the Sonr protocol to leverage account assets in a secure and efficient manner.
//
// Modules: DIDComm, MPC Wallet
// Interface: libp2p host
// Transports: TCP, UDP, QUIC, WebTransport, WebSockets
// ---

package motor

import (
	"context"

	"github.com/sonr-hq/sonr/pkg/common"
	"github.com/sonr-hq/sonr/pkg/host"
	"github.com/sonr-hq/sonr/pkg/network"
	mt "github.com/sonr-hq/sonr/third_party/types/motor/bind/v1"
)

type MotorNode struct {
	// Node is the libp2p host
	Node   *host.P2PHost
	Wallet common.Wallet
	//
}

// It creates a new node and wallet, and returns a MotorNode struct containing them
func NewMotorInstance(ctx context.Context, req *mt.InitializeRequest) (*MotorNode, error) {
	n, err := host.New(ctx)
	if err != nil {
		return nil, err
	}
	w, err := network.NewWallet("snr")
	if err != nil {
		return nil, err
	}

	return &MotorNode{
		Node:   n,
		Wallet: w,
	}, nil
}
