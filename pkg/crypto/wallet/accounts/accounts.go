package accounts

import (
	"fmt"

	"github.com/sonrhq/core/pkg/crypto/wallet"
	"github.com/sonrhq/core/pkg/crypto/wallet/accounts/internal"
	v1 "github.com/sonrhq/core/x/identity/types/vault/v1"
)

// New creates a new account with the given options.
func New(opts ...Option) (wallet.Account, error) {
	c := defaultConfig()
	for _, opt := range opts {
		opt(c)
	}
	return c.Keygen()
}

// Load loads an account from a *crypto.AccountConfig.
func Load(ac *wallet.AccountConfig) (wallet.Account, error) {
	shares, err := v1.DeserializeConfigList(ac.Shares)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize shares: %w", err)
	}
	return internal.BaseAccountFromConfig(ac, shares[0]), nil
}
