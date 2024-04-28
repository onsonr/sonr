package ipfs

import (
	"context"

	"github.com/cosmos/btcutil/bech32"
	"github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/ipfs/kubo/client/rpc"
	coreiface "github.com/ipfs/kubo/core/coreiface"
)

type IPNSKey = coreiface.Key

// IPFSClient is an interface for interacting with an IPFS node.
type IPFSClient struct {
	*rpc.HttpApi
	reachable bool
}

// NewLocalClient creates a new IPFS client that connects to the local IPFS node.
func NewLocalClient() (*IPFSClient, error) {
	rpcClient, err := rpc.NewLocalApi()
	if err != nil {
		return &IPFSClient{HttpApi: nil, reachable: false}, err
	}
	return &IPFSClient{HttpApi: rpcClient, reachable: true}, nil
}

// NewKey creates a new IPFS key.
func (c *IPFSClient) NewKey(address types.Address) (IPNSKey, error) {
	// Use bech32 encoding to convert the address to a string.
	addr, err := bech32.Encode("idx", address.Bytes())
	if err != nil {
		return nil, err
	}

	// Call the IPFS client to create a new key.
	ctx := context.Background()
	key, err := c.Key().Generate(ctx, addr)
	if err != nil {
		return nil, err
	}
	return key, nil
}
