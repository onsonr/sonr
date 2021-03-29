package node

import (
	"fmt"

	// Imported
	// tor "berty.tech/go-libp2p-tor-transport"
	// tconf "berty.tech/go-libp2p-tor-transport/config"
	brprot "berty.tech/berty/v2/go/pkg/bertyprotocol"
	brtypes "berty.tech/berty/v2/go/pkg/protocoltypes"
	// Local
	// mplex "github.com/libp2p/go-libp2p-mplex"
	// direct "github.com/libp2p/go-libp2p-webrtc-direct"
	// "github.com/pion/webrtc/v3"
)

func (n *Node) StartProtocol() error {
	client, err := brprot.New(n.ctx, brprot.Opts{})
	if err != nil {
		return err
	}

	ret, err := client.InstanceGetConfiguration(n.ctx, &brtypes.InstanceGetConfiguration_Request{})
	if err != nil {
		return err
	}

	for _, listener := range ret.Listeners {
		if listener == "/p2p-circuit" {
			fmt.Println(listener)
		}
	}
	return nil
}
