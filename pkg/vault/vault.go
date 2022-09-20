package vault

import (
	"log"
	"os"

	"github.com/joho/godotenv"
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
	err := godotenv.Load("../../.env")
  if err != nil {
    log.Fatal(err)
  }

	return &vaultImpl{
		vaultEndpoint: os.Getenv("VAULT_ENDPOINT"),
	}
}
