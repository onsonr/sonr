package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ipfs/boxo/files"
	"github.com/ipfs/boxo/path"
	"github.com/ipfs/kubo/core/coreiface/options"
)

// AddToLocalIPFS adds a file to the local IPFS node
func (k Keeper) AddToLocalIPFS(ctx sdk.Context, data files.Node) (string, error) {
	cid, err := k.ipfsClient.Unixfs().Add(ctx, data)
	if err != nil {
		return "", err
	}
	return cid.String(), nil
}

// GetFromLocalIPFS gets a file from the local IPFS node
func (k Keeper) GetFromLocalIPFS(ctx sdk.Context, cid string) (files.Node, error) {
	path, err := path.NewPath(cid)
	if err != nil {
		return nil, err
	}
	return k.ipfsClient.Unixfs().Get(ctx, path)
}

// HasPathInLocalIPFS checks if a file is in the local IPFS node
func (k Keeper) HasPathInLocalIPFS(ctx sdk.Context, cid string) (bool, error) {
	path, err := path.NewPath(cid)
	if err != nil {
		return false, err
	}
	v, err := k.ipfsClient.Unixfs().Get(ctx, path)
	if err != nil {
		return false, err
	}

	if v == nil {
		return false, nil
	}
	return true, nil
}

// PinToLocalIPFS pins a file to the local IPFS node
func (k Keeper) PinToLocalIPFS(ctx sdk.Context, cid string, name string) error {
	path, err := path.NewPath(cid)
	if err != nil {
		return err
	}
	err = k.ipfsClient.Pin().Add(ctx, path, options.Pin.Name(name))
	if err != nil {
		return err
	}
	return nil
}
