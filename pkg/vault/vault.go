package vault

import (
	"context"

	"github.com/di-dao/sonr/crypto/mpc"
	"github.com/di-dao/sonr/internal/local"
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
	snrCtx := local.UnwrapContext(ctx)
	// Generate keyshare
	keyshares, err := mpc.GenerateKss()
	if err != nil {
		return nil, err
	}

	// Create a new wallet
	wal, err := wallet.New(keyshares)
	if err != nil {
		return nil, err
	}

	// Update the context with the wallet address
	snrCtx.UserAddress = wal.SonrAddress()
	local.WrapContext(snrCtx)

	// Create a new vault
	return &vault{
		wallet:      wal,
		credentials: auth.NewCredentials(),
		properties:  props.NewProperties(),
		vfs:         ipfs.NewFileSystem(wal.SonrAddress()),
	}, nil
}

// Assign creates a new IPNS key named for the Sonr Address and initializes the VFS
func (v *vault) Assign(ctx context.Context, credential *auth.Credential) error {
	snrCtx := local.UnwrapContext(ctx)
	err := v.credentials.LinkCredential(snrCtx.ServiceOrigin, credential)
	if err != nil {
		return err
	}
	key, err := ipfs.NewKey(context.Background(), v.wallet.SonrAddress())
	if err != nil {
		return err
	}
	snrCtx.PeerID = key.ID().String()
	local.WrapContext(snrCtx)
	return nil
}
