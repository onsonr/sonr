package vault

import (
	"context"

	"github.com/di-dao/sonr/crypto/mpc"
	"github.com/di-dao/sonr/pkg/ipfs"
	"github.com/di-dao/sonr/pkg/vault/auth"
	"github.com/di-dao/sonr/pkg/vault/props"
	"github.com/di-dao/sonr/pkg/vault/wallet"
)

// Vault is an interface that defines the methods for a vault.
type Vault interface{}

// vault is a struct that contains the information of a vault to be stored in the vault
type vault struct {
	credentials auth.Credentials
	properties  props.Properties
	wallet      *wallet.Wallet
	vfs         ipfs.VFS
}

// New creates a new vault from a set of keyshares.
func Generate(ctx context.Context) (Vault, error) {
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
		wallet:      wallet,
		credentials: auth.NewCredentials(),
		properties:  props.NewProperties(),
		vfs:         ipfs.NewFileSystem(wallet.Accounts[wallet.Coin][0].Address),
	}, nil
}
