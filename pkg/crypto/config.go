package crypto

import "github.com/taurusgroup/multi-party-sig/pkg/party"

// The default shards that are added to the MPC wallet
var defaultParticipants = party.IDSlice{"vault", "shared"}

// Preset options struct
type walletConfig struct {
	participants party.IDSlice
	threshold    int
	network      *Network
}

// default configuration options
func defaultConfig() *walletConfig {
	return &walletConfig{
		participants: party.IDSlice{"vault", "shared"},
		threshold:    1,
		network:      NewNetwork(party.IDSlice{"vault", "shared"}),
	}
}

// Applies the options and returns a new walletConfig
func (wc *walletConfig) Apply(opts ...WalletOption) *MPCWallet {
	for _, opt := range opts {
		opt(wc)
	}

	return &MPCWallet{
		Threshold: wc.threshold,
		Network:   wc.network,
	}
}

// WalletOption is a function that applies a configuration option to a walletConfig
type WalletOption func(*walletConfig)

// WithParticipants adds a list of participants to the wallet
func WithParticipants(participants ...party.ID) WalletOption {
	return func(c *walletConfig) {
		// Update participants and network.
		c.participants = append(defaultParticipants, participants...)
		c.network = NewNetwork(c.participants)
	}
}

// WithThreshold sets the threshold of the MPC wallet
func WithThreshold(threshold int) WalletOption {
	return func(c *walletConfig) {
		c.threshold = threshold
		if c.threshold == 0 {
			c.threshold = 1
		}
	}
}
