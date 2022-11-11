package client

import (
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

	// -- Dev Blockchain --
	BLOCKCHAIN_REST_DEV   = "http://143.198.29.209:1317"
	BLOCKCHAIN_FAUCET_DEV = "http://143.198.29.209:8000"
	BLOCKCHAIN_RPC_DEV    = "143.198.29.209:9090"

	// -- Beta Blockchain --
	BLOCKCHAIN_REST_BETA   = "http://137.184.190.146:1317"
	BLOCKCHAIN_FAUCET_BETA = "http://137.184.190.146:8000"
	BLOCKCHAIN_RPC_BETA    = "137.184.190.146:9090"

	// -- Services --
	IPFS_ADDRESS      = "https://ipfs.sonr.ws"
	IPFS_API_ADDRESS  = "https://api.ipfs.sonr.ws"
	VAULT_API_ADDRESS = "http://164.92.99.233"
)

type Endpoints struct {
	FaucetAddress  string
	RPCAddress     string
	APIAddress     string
	IPFSAddress    string
	IPFSApiAddress string
}

func defaultEndpoints() *Endpoints {
	return &Endpoints{
		FaucetAddress:  "0.0.0.0:4500",
		RPCAddress:     "0.0.0.0:9090",
		APIAddress:     "0.0.0.0:26657",
		IPFSAddress:    "0.0.0.0:4001",
		IPFSApiAddress: "0.0.0.0:5001",
	}
}

type Client struct {
	clientMode mt.ClientMode
	Endpoints  *Endpoints
}

func getEnvFile(mode mt.ClientMode) string {
	envFileName := ""
	switch mode {
	case mt.ClientMode_ENDPOINT_DEV:
		envFileName = ".env.beta"
	case mt.ClientMode_ENDPOINT_BETA:
		envFileName = ".env.dev"
	default:
		envFileName = ".env"
	}
	env_path := filepath.Join(projectpath.Root, envFileName)
	return env_path
}

func NewClient(mode mt.ClientMode) *Client {
	env_file := getEnvFile(mode)
	err := godotenv.Load(env_file)
	if err != nil {
		log.Printf("failed to loading variables from %s: %s, using default endpoints", env_file, err.Error())
		return &Client{
			clientMode: mt.ClientMode_ENDPOINT_LOCAL,
			Endpoints:  defaultEndpoints(),
		}
	}
	return &Client{
		clientMode: mode,
		Endpoints: &Endpoints{
			FaucetAddress:  os.Getenv("BLOCKCHAIN_FAUCET"),
			RPCAddress:     os.Getenv("BLOCKCHAIN_RPC"),
			APIAddress:     os.Getenv("BLOCKCHAIN_REST"),
			IPFSAddress:    os.Getenv("IPFS_ADDRESS"),
			IPFSApiAddress: os.Getenv("IPFS_API_ADDRESS"),
		},
	}
}

func (c *Client) GetFaucetAddress() string {
	return c.Endpoints.FaucetAddress
}

func (c *Client) GetRPCAddress() string {
	return c.Endpoints.RPCAddress
}

func (c *Client) GetAPIAddress() string {
	return c.Endpoints.APIAddress
}

func (c *Client) GetIPFSAddress() string {
	return c.Endpoints.IPFSAddress
}

func (c *Client) GetIPFSApiAddress() string {
	return c.Endpoints.IPFSApiAddress
}
