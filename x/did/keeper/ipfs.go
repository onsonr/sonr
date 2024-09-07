package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ipfs/boxo/files"
	"github.com/ipfs/boxo/path"
	"github.com/ipfs/kubo/client/rpc"
	"github.com/ipfs/kubo/core/coreiface/options"
	"github.com/onsonr/sonr/internal/vfs"
)

// assembleInitialVault assembles the initial vault
func (k Keeper) assembleInitialVault(ctx sdk.Context) (string, int64, error) {
	cid, err := k.ipfsClient.Unixfs().Add(context.Background(), vfs.AssembleDirectory())
	if err != nil {
		return "", 0, err
	}
	return cid.String(), k.GetExpirationBlockHeight(ctx, time.Second*15), nil
}

// pinInitialVault pins the initial vault to the local IPFS node
func (k Keeper) pinInitialVault(_ sdk.Context, cid string, address string) error {
	// Resolve the path
	path, err := path.NewPath(cid)
	if err != nil {
		return err
	}

	// 1. Initialize vault.db sqlite database in local IPFS with Mount

	// 2. Insert the InitialWalletAccounts

	// 3. Publish the path to the IPNS
	_, err = k.ipfsClient.Name().Publish(context.Background(), path, options.Name.Key(address))
	if err != nil {
		return err
	}

	// 4. Insert the accounts into x/auth

	// 5. Insert the controller into state
	return nil
}

// GetFromIPFS gets a file from the local IPFS node
func (k Keeper) GetFromIPFS(ctx sdk.Context, cid string) (files.Directory, error) {
	path, err := path.NewPath(cid)
	if err != nil {
		return nil, err
	}
	node, err := k.ipfsClient.Unixfs().Get(ctx, path)
	if err != nil {
		return nil, err
	}
	dir, ok := node.(files.Directory)
	if !ok {
		return nil, fmt.Errorf("retrieved node is not a directory")
	}
	return dir, nil
}

// HasIPFSConnection returns true if the IPFS client is initialized
func (k *Keeper) HasIPFSConnection() bool {
	if k.ipfsClient == nil {
		ipfsClient, err := rpc.NewLocalApi()
		if err != nil {
			return false
		}
		k.ipfsClient = ipfsClient
	}
	return k.ipfsClient != nil
}

// HasPathInIPFS checks if a file is in the local IPFS node
func (k Keeper) HasPathInIPFS(ctx sdk.Context, cid string) (bool, error) {
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

// PinToIPFS pins a file to the local IPFS node
func (k Keeper) PinToIPFS(ctx sdk.Context, cid string, name string) error {
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
