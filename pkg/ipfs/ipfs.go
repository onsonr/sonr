package ipfs

import (
	"context"

	"github.com/ipfs/boxo/files"
	"github.com/ipfs/boxo/path"
	"github.com/ipfs/kubo/client/rpc"
	coreiface "github.com/ipfs/kubo/core/coreiface"
	"github.com/ipfs/kubo/core/coreiface/options"
)

type IPNSKey = coreiface.Key

// IPFSClient is an interface for interacting with an IPFS node.
type IPFSClient struct {
	*rpc.HttpApi
	reachable bool
}

// NewKey creates a new IPFS key.
func NewKey(ctx context.Context, addr string) (IPNSKey, error) {
	// Call the IPFS client to create a new key
	c, err := getClient()
	if err != nil {
		return nil, err
	}
	key, err := c.Key().Generate(ctx, addr)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// // GetVFS returns the VFS interface from the client UnixFS API.
func Get(ctx context.Context, path string) (files.Node, error) {
	c, err := getClient()
	if err != nil {
		return nil, err
	}
	cid, err := parsePath(path)
	if err != nil {
		return nil, err
	}

	// Call the IPFS client to get the UnixFS API.
	node, err := c.Unixfs().Get(context.Background(), cid)
	if err != nil {
		return nil, err
	}
	return node, nil
}

// SaveVFS saves the VFS interface to the client UnixFS API.
func Save(ctx context.Context, fs files.Node, keyName string) (string, error) {
	// Call the IPFS client to get the UnixFS API.
	c, err := getClient()
	if err != nil {
		return "", err
	}
	api, err := c.Unixfs().Add(ctx, fs)
	if err != nil {
		return "", err
	}
	name, err := c.Name().Publish(ctx, api, options.Name.Key(keyName))
	if err != nil {
		return "", err
	}
	return name.AsPath().String(), nil
}

// NewLocalClient creates a new IPFS client that connects to the local IPFS node.
func getClient() (*IPFSClient, error) {
	rpcClient, err := rpc.NewLocalApi()
	if err != nil {
		return &IPFSClient{HttpApi: nil, reachable: false}, err
	}
	return &IPFSClient{HttpApi: rpcClient, reachable: true}, nil
}

func parsePath(p string) (path.Path, error) {
	return path.NewPath(p)
}
