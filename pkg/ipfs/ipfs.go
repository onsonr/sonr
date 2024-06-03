package ipfs

import (
	"context"

	"github.com/di-dao/sonr/internal/local"
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
}

// NewKey creates a new IPFS key.
func NewKey(ctx context.Context, addr string) (IPNSKey, error) {
	// Call the IPFS client to create a new key
	c, err := local.GetIPFSClient()
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
	c, err := local.GetIPFSClient()
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

// GetFileSystem returns the VFS interface from the client UnixFS API.
func GetFileSystem(ctx context.Context, path string) (VFS, error) {
	c, err := local.GetIPFSClient()
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
	return Load(path, node)
}

// SaveVFS saves the VFS interface to the client UnixFS API.
func Save(ctx context.Context, fs files.Node, keyName string) (string, error) {
	// Call the IPFS client to get the UnixFS API.
	c, err := local.GetIPFSClient()
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

// SaveFileSystem saves the VFS interface to the client UnixFS API.
func SaveFileSystem(ctx context.Context, fs VFS) (string, error) {
	// Call the IPFS client to get the UnixFS API.
	c, err := local.GetIPFSClient()
	if err != nil {
		return "", err
	}
	api, err := c.Unixfs().Add(ctx, fs.Node())
	if err != nil {
		return "", err
	}
	name, err := c.Name().Publish(ctx, api, options.Name.Key(fs.Name()))
	if err != nil {
		return "", err
	}
	return name.AsPath().String(), nil
}

func parsePath(p string) (path.Path, error) {
	return path.NewPath(p)
}
