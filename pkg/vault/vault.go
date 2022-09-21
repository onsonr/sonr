package vault

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/sonr-io/sonr/internal/projectpath"
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
	env_path := filepath.Join(projectpath.Root, ".env")
	err := godotenv.Load(env_path)
  if err != nil {
    log.Fatal(err)
  }

	return &vaultImpl{
		vaultEndpoint: os.Getenv("VAULT_ENDPOINT"),
	}
}
