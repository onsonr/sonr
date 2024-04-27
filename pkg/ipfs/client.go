package ipfs

import (
	"context"

	"github.com/cosmos/btcutil/bech32"
	"github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/ipfs/kubo/client/rpc"
)

// IPFSClient is an interface for interacting with an IPFS node.
type IPFSClient struct {
	*rpc.HttpApi
}

// NewLocalClient creates a new IPFS client that connects to the local IPFS node.
func NewLocalClient() (*IPFSClient, error) {
	rpcClient, err := rpc.NewLocalApi()
	if err != nil {
		return nil, err
	}
	return &IPFSClient{rpcClient}, nil
}

// NewKey creates a new IPFS key.
func (c *IPFSClient) NewKey(ctx context.Context, address types.Address) (IPFSKey, error) {
	// Use bech32 encoding to convert the address to a string.
	addr, err := bech32.Encode("idx", address.Bytes())
	if err != nil {
		return nil, err
	}

	// Call the IPFS client to create a new key.
	key, err := c.Key().Generate(ctx, addr)
	if err != nil {
		return nil, err
	}
	return key, nil
}
