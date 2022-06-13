package crypto

import "github.com/taurusgroup/multi-party-sig/pkg/party"

var defaultParticipants = party.IDSlice{"vault", "shared"}

type walletConfig struct {
	participants party.IDSlice
	threshold    int
	network      *Network
}

func defaultConfig() *walletConfig {
	return &walletConfig{
		participants: party.IDSlice{"vault", "shared"},
		threshold:    1,
		network:      NewNetwork(party.IDSlice{"vault", "shared"}),
	}
}

func (wc *walletConfig) Apply(opts ...WalletOption) *MPCWallet {
	for _, opt := range opts {
		opt(wc)
	}

	return &MPCWallet{
		Participants: wc.participants,
		Threshold:    wc.threshold,
		Network:      wc.network,
	}

}

type WalletOption func(*walletConfig)

func WithParticipants(participants ...party.ID) WalletOption {
	return func(c *walletConfig) {
		c.participants = defaultParticipants
		// Update participants and network.
		c.participants = append(c.participants, participants...)
		c.network = NewNetwork(c.participants)
	}
}

func WithThreshold(threshold int) WalletOption {
	return func(c *walletConfig) {
		c.threshold = threshold
		if c.threshold == 0 {
			c.threshold = 1
		}
	}
}
