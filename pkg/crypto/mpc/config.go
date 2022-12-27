package mpc

import (
	peer "github.com/libp2p/go-libp2p/core/peer"
	"github.com/sonr-hq/sonr/pkg/common"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/pkg/protocol"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

// `config` is a struct that contains a public key, a list of participants, a threshold, and a
// map of configurations for each participant.
// @property {[]byte} pubKey - the public key of the protocol
// @property participants - A list of parties that are participating in the protocol.
// @property {int} threshold - The minimum number of parties that must sign the message for it to be
// valid.
// @property configs - a map of party IDs to the configuration of the party.
type config struct {
	pubKey    []byte
	threshold int
	configs   map[party.ID]*cmp.Config
}

// default configuration options
func defaultConfig() *config {
	return &config{
		pubKey:    make([]byte, 0),
		threshold: 1,
		configs:   make(map[party.ID]*cmp.Config),
	}
}

// Applies the options and returns a new walletConfig
func (wc *config) Apply(opts ...Option) *MPCProtocol {
	for _, opt := range opts {
		opt(wc)
	}

	return &MPCProtocol{
		pubKey:    wc.pubKey,
		configs:   wc.configs,
		threshold: wc.threshold,
	}
}

// Option is a function that applies a configuration option to a walletConfig
type Option func(*config)

// WithThreshold sets the threshold of the MPC wallet
func WithThreshold(threshold int) Option {
	return func(c *config) {
		c.threshold = threshold
		if c.threshold == 0 {
			c.threshold = 1
		}
	}
}

// WithWalletShares sets the configs used for the MPC wallet
func WithWalletShares(cnfs ...common.WalletShare) Option {
	return func(c *config) {
		c.configs = make(map[party.ID]*cmp.Config)
		for _, cnf := range cnfs {
			c.configs[cnf.SelfID()] = cnf.CMPConfig()
		}
	}
}

//
// Public Interfaces
//

// A Network is a channel that sends messages to parties and receives messages from parties.
type Network interface {
	// Ls returns a list of peers that are connected to the network.
	Ls() []party.ID

	// A function that takes in a party ID and returns a channel of protocol messages.
	Next(id party.ID) <-chan *protocol.Message

	// Sending a message to the network.
	Send(msg *protocol.Message)

	// A channel that is closed when the party is done with the protocol.
	Done(id party.ID) chan struct{}

	// A function that is called when a party is done with the protocol.
	Quit(id party.ID)

	// IsOnlineNetwork returns true if the network is an online network.
	IsOnlineNetwork() bool
}

//
// Helper Functions
//

// It converts a peer.ID to a party.ID
func peerIdToPartyId(id peer.ID) party.ID {
	return party.ID(id)
}

// It converts a party ID to a peer ID
func partyIdToPeerId(id party.ID) peer.ID {
	return peer.ID(id)
}

// It converts a list of peer IDs to a list of party IDs
func peerIdListToPartyIdList(ids []peer.ID) []party.ID {
	partyIds := make([]party.ID, len(ids))
	for i, id := range ids {
		partyIds[i] = peerIdToPartyId(id)
	}
	return partyIds
}

// It converts a list of party IDs to a list of peer IDs
func partyIdListToPeerIdList(ids []party.ID) []peer.ID {
	peerIds := make([]peer.ID, len(ids))
	for i, id := range ids {
		peerIds[i] = partyIdToPeerId(id)
	}
	return peerIds
}
