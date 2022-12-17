package mpc

import (
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/protocol"
)

const (
	// MPC_KEYGEN_PROTOCOL is the protocol ID for the MPC keygen protocol that is attached to the node.
	MPC_KEYGEN_PROTOCOL = protocol.ID("/sonr/mpc/keygen/1.0.0")

	// MPC_SIGN_PROTOCOL is the protocol ID for the MPC sign protocol that is attached to the node.
	MPC_SIGN_PROTOCOL = protocol.ID("/sonr/mpc/sign/1.0.0")

	// MPC_REFRESH_PROTOCOL is the protocol ID for the MPC refresh protocol that is attached to the node.
	MPC_REFRESH_PROTOCOL = protocol.ID("/sonr/mpc/refresh/1.0.0")

	// MPC_PRE_SIGN_PROTOCOL is the protocol ID for the MPC pre-sign protocol that is attached to the node.
	MPC_PRE_SIGN_PROTOCOL = protocol.ID("/sonr/mpc/pre-sign/1.0.0")

	// MPC_PRE_SIGN_ONLINE_PROTOCOL is the protocol ID for the MPC pre-sign online protocol that is attached to the node.
	MPC_PRE_SIGN_ONLINE_PROTOCOL = protocol.ID("/sonr/mpc/pre-sign-online/1.0.0")
)

func (p *MpcProtocol) KeygenHandler(stream network.Stream) {
	

}
