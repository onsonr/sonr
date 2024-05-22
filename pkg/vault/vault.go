package vault

import (
	"github.com/di-dao/sonr/crypto/kss"
	"github.com/di-dao/sonr/pkg/vault/chain"
	"github.com/di-dao/sonr/pkg/vault/controller"
	"github.com/di-dao/sonr/pkg/vault/props"
	"github.com/di-dao/sonr/pkg/vault/wallet"
	"github.com/di-dao/sonr/pkg/vfs"
	"github.com/ipfs/boxo/path"
)

// Vault is an interface that defines the methods for a vault.
type Vault interface {
}

// vault is a struct that contains the information of a vault to be stored in the vault
type vault struct {
	controller controller.Controller
	path       path.Path
	properties props.Properties
	wallet     *wallet.Wallet
	vfs        vfs.VFS
}

// New creates a new vault from a set of keyshares.
func New(keyshares kss.Set) (Vault, error) {
	// Get sonr address and bitcoin address from keyshares
	wallet, err := wallet.New(keyshares)
	if err != nil {
		return nil, err
	}
	return &vault{
		wallet:     wallet,
		properties: props.NewProperties(),
		controller: controller.New(keyshares),
		vfs:        vfs.New(wallet.Accounts[chain.CoinSNRType][0].Address),
	}, nil
}
