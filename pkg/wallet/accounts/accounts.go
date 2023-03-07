package accounts

import (
	"fmt"

	"github.com/sonrhq/core/pkg/client/rosetta"
	"github.com/sonrhq/core/pkg/wallet"
	"github.com/sonrhq/core/pkg/wallet/accounts/internal"
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

// GetCosmosAccount returns a cosmos account.
func GetCosmosAccount(root wallet.Account, cosmos wallet.Account, client rosetta.Client) wallet.CosmosAccount {
	return internal.LoadCosmosAccount(root, cosmos, client)
}

// Load loads an account from a *crypto.AccountConfig.
func Load(ac *wallet.AccountConfig) (wallet.Account, error) {
	shares, err := v1.DeserializeConfigList(ac.Shares)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize shares: %w", err)
	}
	return internal.BaseAccountFromConfig(ac, shares[0]), nil
}

// LoadFromBytes loads an account from a byte slice.
func LoadFromBytes(b []byte) (wallet.Account, error) {
	accCfg := &v1.AccountConfig{}
	if err := accCfg.Unmarshal(b); err != nil {
		return nil, fmt.Errorf("failed to unmarshal account config: %w", err)
	}
	return Load(accCfg)
}
