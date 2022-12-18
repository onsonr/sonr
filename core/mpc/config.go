package mpc

import (
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	p2p_protocol "github.com/libp2p/go-libp2p/core/protocol"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/pkg/protocol"
)

const (
	// MPC_KEYGEN_FEED_PROTOCOL
	kCmpKeygenFeed = p2p_protocol.ID("/mpc-cmp/keygen/0.1.0")

	// MPC_SIGN_PROTOCOL is the protocol ID for the MPC sign protocol that is attached to the node.
	kCmpSign = p2p_protocol.ID("/mpc-cmp/sign/0.1.0")

	// MPC_REFRESH_PROTOCOL is the protocol ID for the MPC refresh protocol that is attached to the node.
	kCmpRefresh = p2p_protocol.ID("/mpc-cmp/refresh/0.1.0")

	// MPC_PRE_SIGN_PROTOCOL is the protocol ID for the MPC pre-sign protocol that is attached to the node.
	kCmpPreSign = p2p_protocol.ID("/mpc-cmp/pre-sign/0.1.0")

	// MPC_PRE_SIGN_ONLINE_PROTOCOL is the protocol ID for the MPC pre-sign online protocol that is attached to the node.
	kCmpPreSignOnline = p2p_protocol.ID("/mpc-cmp/pre-sign-online/0.1.0")
)

// Option is a function that can be applied to ExchangeProtocol config
type Option func(*options)

// options for ExchangeProtocol config
type options struct {
	// location        *types.Location
	interval        time.Duration
	olcCode         string
	autoPushEnabled bool
}

// defaultOptions for ExchangeProtocol config
func defaultOptions() *options {
	return &options{
		//location:        api.DefaultLocation(),
		interval:        time.Second * 5,
		autoPushEnabled: true,
	}
}

func getTotalRoundsFromCreate(c protocol.StartFunc) int {
	r, err := c([]byte(""))
	if err != nil {
		return 0
	}
	return int(r.FinalRoundNumber())
}

func convertToPartyIDs(ids []peer.ID) []party.ID {
	partyIDs := make([]party.ID, len(ids))
	for i, id := range ids {
		partyIDs[i] = party.ID(id)
	}
	return partyIDs
}
