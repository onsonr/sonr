package vault

import (
	"errors"
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

	var uri string
	// by default use .env if it exists
	_, err := os.Stat(env_path)
	if errors.Is(err, os.ErrNotExist) {
		uri = "https://vault.sonr.ws"
	} else {
		err = godotenv.Load(env_path)
		if err != nil {
			log.Fatal(err)
		}

		uri = os.Getenv("VAULT_ENDPOINT")
	}
	
	return &vaultImpl{
		vaultEndpoint: uri,
	}
}
