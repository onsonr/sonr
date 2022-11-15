package client

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/sonr-io/sonr/internal/projectpath"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
)

//Endpoints has address for the services used by the client.
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
		envFileName = ".env.dev"
	case mt.ClientMode_ENDPOINT_BETA:
		envFileName = ".env.beta"
	default:
		envFileName = ".env"
	}
	envPath := filepath.Join(projectpath.Root, envFileName)
	return envPath
}

func NewClient(mode mt.ClientMode) *Client {
	envFile := getEnvFile(mode)
	err := godotenv.Load(envFile)
	if err != nil {
		log.Printf("failed to loading variables from %s: %s, using default endpoints", envFile, err.Error())
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
