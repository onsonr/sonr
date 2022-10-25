package mpc

import (
	"github.com/mr-tron/base58/base58"
	"github.com/sonr-io/multi-party-sig/pkg/party"
	"github.com/sonr-io/multi-party-sig/pkg/pool"
	"github.com/sonr-io/multi-party-sig/protocols/cmp"
)

// The default shards that are added to the MPC wallet
var defaultParticipants = party.IDSlice{"dsc", "recovery", "psk", "bank0", "bank1", "bank2"}

// Preset options struct
type walletConfig struct {
	pubKey       []byte
	participants party.IDSlice
	threshold    int
	network      *Network
	configs      map[party.ID]*cmp.Config
}

// default configuration options
func defaultConfig() *walletConfig {
	return &walletConfig{
		pubKey:       make([]byte, 0),
		participants: defaultParticipants,
		threshold:    1,
		network:      NewNetwork(defaultParticipants),
		configs:      make(map[party.ID]*cmp.Config),
	}
}

// Applies the options and returns a new walletConfig
func (wc *walletConfig) Apply(opts ...WalletOption) *Wallet {
	for _, opt := range opts {
		opt(wc)
	}

	return &Wallet{
		pool: pool.NewPool(0),

		PubKey:    wc.pubKey,
		Configs:   wc.configs,
		ID:        wc.participants[0],
		Threshold: wc.threshold,
		Network:   wc.network,
	}
}

// WalletOption is a function that applies a configuration option to a walletConfig
type WalletOption func(*walletConfig)

func WithBase58PubKey(key string) WalletOption {
	return func(c *walletConfig) {
		decoded, err := base58.Decode(key)
		if err != nil {
			panic(err)
		}
		c.pubKey = decoded
	}
}

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

// WithConfigs sets the configs used for the MPC wallet
func WithConfigs(cnfs map[party.ID]*cmp.Config) WalletOption {
	return func(c *walletConfig) {
		c.configs = cnfs
	}
}
