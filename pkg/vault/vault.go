package vault

import (
	"github.com/sonr-io/sonr/pkg/did"
)

type VaultClient interface {
	CreateVault(d string, deviceShards [][]byte, dscPub string, encDscShard, pskShard, recShard []byte) (did.Service, error)
	GetVaultShards(d string) (Vault, error)
	PopShard() (string, error)
	IssueShard(shardPrefix, dscPub, dscShard string) (did.Service, error)
}

type vaultImpl struct {
	vaultEndpoint string
}

func New() VaultClient {
	return &vaultImpl{
		// vaultEndpoint: "https://vault.sonr.ws",
		vaultEndpoint: "http://127.0.0.1:1234",
	}
}
