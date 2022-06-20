package vault

import "github.com/sonr-io/sonr/pkg/did"

type VaultClient interface {
	CreateVault(d string, deviceShards []string, dscPub, encDscShard, pskShard, recShard string) (did.Service, error)
	PopShard() (string, error)
	IssueShard(shardPrefix, dscPub, dscShard string) (did.Service, error)
}

type vaultImpl struct {
	vaultEndpoint string
}

func New() VaultClient {
	return &vaultImpl{
		vaultEndpoint: "https://vault.sonr.ws",
	}
}
