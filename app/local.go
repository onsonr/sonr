package app

import (
	"github.com/ipfs/kubo/client/rpc"
	"github.com/onsonr/sonr/internal/files"
)

var (
	ChainID = "testnet"
	ValAddr = "val1"
)

// Initialize initializes the local configuration values
func init() {
	err := files.Assemble(".data/vaults/0")
	if err != nil {
		panic(err)
	}
}

// SetLocalContextSessionID sets the session ID for the local context
func SetLocalValidatorAddress(address string) {
	ValAddr = address
}

// SetLocalContextChainID sets the chain ID for the local
func SetLocalChainID(id string) {
	ChainID = id
}

// IPFSClient is an interface for interacting with an IPFS node.
type IPFSClient = *rpc.HttpApi

// NewLocalClient creates a new IPFS client that connects to the local IPFS node.
func GetIPFSClient() (IPFSClient, error) {
	rpcClient, err := rpc.NewLocalApi()
	if err != nil {
		return nil, err
	}
	return rpcClient, nil
}
