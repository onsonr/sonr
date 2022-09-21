package client

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/sonr-io/sonr/internal/projectpath"
)

type Client struct {}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) GetFaucetAddress() string {
	env_path := filepath.Join(projectpath.Root, ".env")
	err := godotenv.Load(env_path)
  if err != nil {
    log.Fatal(err)
  }
	
	return os.Getenv("BLOCKCHAIN_FAUCET")
}

func (c *Client) GetRPCAddress() string {
	env_path := filepath.Join(projectpath.Root, ".env")
	err := godotenv.Load(env_path)
  if err != nil {
    log.Fatal(err)
  }
	
	return os.Getenv("BLOCKCHAIN_RPC")
}

func (c *Client) GetAPIAddress() string {
	env_path := filepath.Join(projectpath.Root, ".env")
	err := godotenv.Load(env_path)
  if err != nil {
    log.Fatal(err)
  }
	
	return os.Getenv("BLOCKCHAIN_REST")
}

func (c *Client) GetIPFSAddress() string {
	env_path := filepath.Join(projectpath.Root, ".env")
	err := godotenv.Load(env_path)
  if err != nil {
    log.Fatal(err)
  }

	return os.Getenv("IPFS_ADDRESS")
}

func (c *Client) GetIPFSApiAddress() string {
	env_path := filepath.Join(projectpath.Root, ".env")
	err := godotenv.Load(env_path)
  if err != nil {
    log.Fatal(err)
  }

	return os.Getenv("IPFS_API_ADDRESS")
}

func (c *Client) PrintConnectionEndpoints() {
	log.Println("Connection Endpoints:")
	log.Printf("\tREST: %s\n", c.GetAPIAddress())
	log.Printf("\tRPC: %s\n", c.GetRPCAddress())
	log.Printf("\tFaucet: %s\n", c.GetFaucetAddress())
	log.Printf("\tIPFS: %s\n", c.GetIPFSAddress())
	log.Printf("\tIPFS API: %s\n", c.GetIPFSApiAddress())
}
