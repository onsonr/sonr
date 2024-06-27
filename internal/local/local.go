package local

import (
	"os"

	"github.com/di-dao/sonr/crypto/daed"
	"github.com/ipfs/kubo/client/rpc"
)

var (
	chainID           = "testnet"
	valAddr           = "val1"
	nodeDir           = ".sonr"
	contextSessionKey = contextKey("session-id")

	defaultNodeHome = os.ExpandEnv("$HOME/") + nodeDir

	kh *daed.AESSIV
)

// Initialize initializes the local configuration values
func init() {
}

// SetLocalContextSessionID sets the session ID for the local context
func SetLocalValidatorAddress(address string) {
	valAddr = address
}

// SetLocalContextChainID sets the chain ID for the local
func SetLocalChainID(id string) {
	chainID = id
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
