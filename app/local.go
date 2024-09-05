package app

import (
	"github.com/ipfs/kubo/client/rpc"
)

var (
	ChainID = "testnet"
	ValAddr = "val1"
)

// Initialize initializes the local configuration values
func init() {

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
