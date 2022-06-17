package crypto

import (
	"crypto/elliptic"
	"math/big"

	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/pkg/pool"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

// The default shards that are added to the MPC wallet
var defaultParticipants = party.IDSlice{"me", "vault", "shared"}

// p256Order returns the curve order for the secp256r1 curve
// NOTE: this is specific to the secp256r1/P256 curve,
// and not taken from the domain params for the key itself
// (which would be a more generic approach for all EC).
var p256Order = elliptic.P256().Params().N

// p256HalfOrder returns half the curve order
// a bit shift of 1 to the right (Rsh) is equivalent
// to division by 2, only faster.
var p256HalfOrder = new(big.Int).Rsh(p256Order, 1)

// Preset options struct
type walletConfig struct {
	participants party.IDSlice
	threshold    int
	network      *Network
}

// default configuration options
func defaultConfig() *walletConfig {
	return &walletConfig{
		participants: defaultParticipants,
		threshold:    1,
		network:      NewNetwork(defaultParticipants),
	}
}

// Applies the options and returns a new walletConfig
func (wc *walletConfig) Apply(opts ...WalletOption) *MPCWallet {
	for _, opt := range opts {
		opt(wc)
	}

	return &MPCWallet{
		pool: pool.NewPool(0),

		Configs:   make(map[party.ID]*cmp.Config),
		ID:        wc.participants[0],
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
