package mpc

import (
	"context"

	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/sonr-hq/sonr/pkg/node"

	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/pkg/pool"
	"github.com/taurusgroup/multi-party-sig/pkg/protocol"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

func CMPKeygen(n *node.Node, ids ...peer.ID) (*cmp.Config, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create a new network.
	net, err := node.NewNetwork(ctx, n, ids)
	if err != nil {
		return nil, err
	}

	// Create a new keygen session.
	pl := pool.NewPool(0)
	handler, err := protocol.NewMultiHandler(cmp.Keygen(curve.Secp256k1{}, party.ID(n.ID()), peerIdListToPartyIdList(ids), 1, pl), nil)
	if err != nil {
		return nil, err
	}
	HandlerLoop(handler, net)

	// Wait for the network to finish.
	r, err := handler.Result()
	if err != nil {
		return nil, err
	}
	return r.(*cmp.Config), nil
}
