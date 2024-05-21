package vault

import (
	"github.com/di-dao/core/pkg/creds"
	"github.com/di-dao/core/pkg/ipfs"
	"github.com/di-dao/core/pkg/kss"
)

type Vault interface {
}

type vault struct {
	vfs         ipfs.VFS
	wallet      *WalletData
	credentials *creds.CredentialData
}

func New(keyshares kss.SetI) (Vault, error) {
	// Get sonr address and bitcoin address from keyshares
	wallet, err := NewWallet(keyshares)
	if err != nil {
		return nil, err
	}
	return &vault{
		credentials: creds.NewCredentialData(),
		vfs:         ipfs.NewVFS(wallet.Address.String()),
		wallet:      wallet,
	}, nil
}
