package vault

import (
	"context"

	"github.com/di-dao/sonr/crypto/mpc"
	"github.com/di-dao/sonr/pkg/vault/chain"
	"github.com/di-dao/sonr/pkg/vault/props"
	"github.com/di-dao/sonr/pkg/vault/wallet"
	"github.com/di-dao/sonr/pkg/vfs"
	"github.com/ipfs/boxo/path"
)

// Vault is an interface that defines the methods for a vault.
type Vault interface{}

// vault is a struct that contains the information of a vault to be stored in the vault
type vault struct {
	path       path.Path
	properties props.Properties
	wallet     *wallet.Wallet
	vfs        vfs.FileSystem
}

// New creates a new vault from a set of keyshares.
func New(ctx context.Context) (Vault, error) {
	// Generate keyshares
	keyshares, err := mpc.GenerateKss()
	if err != nil {
		return nil, err
	}

	// Create a new wallet
	wallet, err := wallet.New(keyshares)
	if err != nil {
		return nil, err
	}

	// Create a new vault
	return &vault{
		wallet:     wallet,
		properties: props.NewProperties(),
		vfs:        vfs.New(wallet.Accounts[chain.CoinSNRType][0].Address),
	}, nil
}
