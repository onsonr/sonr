package config

import "github.com/tendermint/starport/starport/pkg/cosmosaccount"

// WithHome sets the data dir of your chain. This option is used to access your chain's
// file based keyring which is only needed when you deal with creating and signing transactions.
// when it is not provided, your data dir will be assumed as `$HOME/.your-chain-id`.
func WithHome(path string) Option {
	return func(c *Config) {
		c.CosmosHomePath = path
	}
}

// WithKeyringServiceName used as the keyring's name when you are using OS keyring backend.
// by default it is `cosmos`.
func WithKeyringServiceName(name string) Option {
	return func(c *Config) {
		c.CosmosKeyringServiceName = name
	}
}

// WithKeyringBackend sets your keyring backend. By default, it is `test`.
func WithKeyringBackend(backend cosmosaccount.KeyringBackend) Option {
	return func(c *Config) {
		c.CosmosKeyringBackend = backend
	}
}

// WithNodeAddress sets the node address of your chain. When this option is not provided
// `http://localhost:26657` is used as default.
func WithNodeAddress(addr string) Option {
	return func(c *Config) {
		c.CosmosNodeAddress = addr
	}
}

func WithAddressPrefix(prefix string) Option {
	return func(c *Config) {
		c.CosmosAddressPrefix = prefix
	}
}

func WithUseFaucet(faucetAddress, denom string, minAmount uint64) Option {
	return func(c *Config) {
		c.CosmosUseFaucet = true
		c.CosmosFaucetAddress = faucetAddress
		if denom != "" {
			c.CosmosFaucetDenom = denom
		}
		if minAmount != 0 {
			c.CosmosFaucetMinAmount = minAmount
		}
	}
}
