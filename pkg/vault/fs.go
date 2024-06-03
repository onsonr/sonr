package vault

import (
	"github.com/di-dao/sonr/crypto/kss"
	"github.com/di-dao/sonr/pkg/vault/auth"
	"github.com/di-dao/sonr/pkg/vault/props"
	"github.com/di-dao/sonr/pkg/vault/wallet"
)

type vaultFS struct {
	Wallet     *wallet.Wallet
	Creds      auth.Credentials
	Properties props.Properties
}

func createVaultFS(set kss.Set) (*vaultFS, error) {
	wallet, err := wallet.New(set)
	if err != nil {
		return nil, err
	}

	return &vaultFS{
		Wallet:     wallet,
		Creds:      auth.NewCredentials(),
		Properties: props.NewProperties(),
	}, nil
}
