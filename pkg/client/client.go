package client

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/sonr-io/sonr/internal/projectpath"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
)

const (
	// -- Local Blockchain --
	BLOCKCHAIN_REST_LOCAL   = "http://0.0.0.0:26657"
	BLOCKCHAIN_FAUCET_LOCAL = "http://0.0.0.0:4500"
	BLOCKCHAIN_RPC_LOCAL    = "127.0.0.1:9090"
	IPFS_ADDRESS_LOCAL      = ""
	IPFS_API_ADDRESS_LOCAL  = ""
	VAULT_API_ADDRESS_LOCAL = ""

	// -- Dev Blockchain --
	BLOCKCHAIN_REST_DEV   = "https://v1.dev.sonr.ws"
	BLOCKCHAIN_FAUCET_DEV = "http://v1.dev.sonr.ws:8000"
	BLOCKCHAIN_RPC_DEV    = "v1.dev.sonr.ws:9090"
	IPFS_ADDRESS_DEV      = "https://ipfs.dev.sonr.ws"
	IPFS_API_ADDRESS_DEV  = "https://api.ipfs.dev.sonr.ws"
	VAULT_API_ADDRESS_DEV = "https://vault.dev.sonr.ws"

	// -- Beta Blockchain --
	BLOCKCHAIN_REST_BETA   = "http://v1.beta.sonr.ws:1317"
	BLOCKCHAIN_FAUCET_BETA = "http://v1.beta.sonr.ws:8000"
	BLOCKCHAIN_RPC_BETA    = "v1.beta.sonr.ws:9090"
	IPFS_ADDRESS_BETA      = "https://ipfs.beta.sonr.ws"
	IPFS_API_ADDRESS_BETA  = "https://api.ipfs.beta.sonr.ws"
	VAULT_API_ADDRESS_BETA = "https://vault.beta.sonr.ws"

	// -- Production Blockchain --
	BLOCKCHAIN_REST_PROD   = "http://v1.sonr.ws"
	BLOCKCHAIN_FAUCET_PROD = "http://v1.sonr.ws:8000"
	BLOCKCHAIN_RPC_PROD    = "v1.sonr.ws:9090"
	IPFS_ADDRESS_PROD      = "https://ipfs.sonr.ws"
	IPFS_API_ADDRESS_PROD  = "https://api.ipfs.sonr.ws"
	VAULT_API_ADDRESS_PROD = "http://vault.sonr.ws"
)

type Client struct {
	clientMode mt.ClientMode
}

func NewClient(mode mt.ClientMode) *Client {
	return &Client{
		clientMode: mode,
	}
}

func (c *Client) GetFaucetAddress() string {
	env_path := filepath.Join(projectpath.Root, ".env")

	// by default use .env if it exists
	_, err := os.Stat(env_path)
	if errors.Is(err, os.ErrNotExist) {
		// .env does not exist, use preset client mode
		switch c.clientMode {
		case mt.ClientMode_ENDPOINT_PROD:
			return BLOCKCHAIN_FAUCET_PROD
		case mt.ClientMode_ENDPOINT_BETA:
			return BLOCKCHAIN_FAUCET_BETA
		case mt.ClientMode_ENDPOINT_DEV:
			return BLOCKCHAIN_FAUCET_DEV
		case mt.ClientMode_ENDPOINT_LOCAL:
			return BLOCKCHAIN_FAUCET_LOCAL
		}
	}

	err = godotenv.Load(env_path)
	if err != nil {
		log.Fatal(err)
	}

	return os.Getenv("BLOCKCHAIN_FAUCET")
}

func (c *Client) GetRPCAddress() string {
	env_path := filepath.Join(projectpath.Root, ".env")

	// by default use .env if it exists
	_, err := os.Stat(env_path)
	if errors.Is(err, os.ErrNotExist) {
		// .env does not exist, use preset client mode
		switch c.clientMode {
		case mt.ClientMode_ENDPOINT_PROD:
			return BLOCKCHAIN_RPC_PROD
		case mt.ClientMode_ENDPOINT_BETA:
			return BLOCKCHAIN_RPC_BETA
		case mt.ClientMode_ENDPOINT_DEV:
			return BLOCKCHAIN_RPC_DEV
		case mt.ClientMode_ENDPOINT_LOCAL:
			return BLOCKCHAIN_RPC_LOCAL
		}
	}

	err = godotenv.Load(env_path)
	if err != nil {
		log.Fatal(err)
	}

	return os.Getenv("BLOCKCHAIN_RPC")
}

func (c *Client) GetAPIAddress() string {
	env_path := filepath.Join(projectpath.Root, ".env")

	// by default use .env if it exists
	_, err := os.Stat(env_path)
	if errors.Is(err, os.ErrNotExist) {
		// .env does not exist, use preset client mode
		switch c.clientMode {
		case mt.ClientMode_ENDPOINT_PROD:
			return BLOCKCHAIN_REST_PROD
		case mt.ClientMode_ENDPOINT_BETA:
			return BLOCKCHAIN_REST_BETA
		case mt.ClientMode_ENDPOINT_DEV:
			return BLOCKCHAIN_REST_DEV
		case mt.ClientMode_ENDPOINT_LOCAL:
			return BLOCKCHAIN_REST_LOCAL
		}
	}

	err = godotenv.Load(env_path)
	if err != nil {
		log.Fatal(err)
	}

	return os.Getenv("BLOCKCHAIN_REST")
}

func (c *Client) GetIPFSAddress() string {
	env_path := filepath.Join(projectpath.Root, ".env")

	// by default use .env if it exists
	_, err := os.Stat(env_path)
	if errors.Is(err, os.ErrNotExist) {
		switch c.clientMode {
		case mt.ClientMode_ENDPOINT_PROD:
			return IPFS_ADDRESS_PROD
		case mt.ClientMode_ENDPOINT_BETA:
			return IPFS_ADDRESS_BETA
		case mt.ClientMode_ENDPOINT_DEV:
			return IPFS_ADDRESS_DEV
		case mt.ClientMode_ENDPOINT_LOCAL:
			return IPFS_ADDRESS_LOCAL
		}
	}

	err = godotenv.Load(env_path)
	if err != nil {
		log.Fatal(err)
	}

	return os.Getenv("IPFS_ADDRESS")
}

func (c *Client) GetIPFSApiAddress() string {
	env_path := filepath.Join(projectpath.Root, ".env")

	// by default use .env if it exists
	_, err := os.Stat(env_path)
	if errors.Is(err, os.ErrNotExist) {
		switch c.clientMode {
		case mt.ClientMode_ENDPOINT_PROD:
			return IPFS_ADDRESS_PROD
		case mt.ClientMode_ENDPOINT_BETA:
			return IPFS_ADDRESS_BETA
		case mt.ClientMode_ENDPOINT_DEV:
			return IPFS_ADDRESS_DEV
		case mt.ClientMode_ENDPOINT_LOCAL:
			return IPFS_ADDRESS_LOCAL
		}
	}

	err = godotenv.Load(env_path)
	if err != nil {
		log.Fatal(err)
	}

	return os.Getenv("IPFS_API_ADDRESS")
}
