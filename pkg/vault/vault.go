package vault

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/sonr-io/sonr/internal/projectpath"
	"github.com/sonr-io/sonr/pkg/did"
)

type VaultClient interface {
	CreateVault(d string, deviceShards [][]byte, dscPub string, encDscShard, pskShard, recShard []byte) (did.Service, error)
	GetVaultShards(d string) (Vault, error)
	PopShard(d string) (Shard, error)
	IssueShard(d, shardSuffix, dscPub, dscShard string) (did.Service, error)
}

type vaultImpl struct {
	vaultEndpoint string
}

func getVaultUri() (string, error) {
	env_path := filepath.Join(projectpath.Root, ".env")

	const uriDefault string = "https://vault.sonr.ws"

	// by default use .env if it exists
	_, err := os.Stat(env_path)
	if errors.Is(err, os.ErrNotExist) {
		return uriDefault, nil
	}

	err = godotenv.Load(env_path)
	if err != nil {
		return uriDefault, err
	}

	return os.Getenv("VAULT_ENDPOINT"), nil
}

func New() VaultClient {
	uri, err := getVaultUri()
	if err != nil {
		fmt.Printf("Error when retrieving vault uri: %s\n", err)
	}

	return &vaultImpl{
		vaultEndpoint: uri,
	}
}
