package keeper

import (
	"context"
	"time"

	"cosmossdk.io/log"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ipfs/boxo/path"
	"github.com/ipfs/kubo/client/rpc"

	"github.com/onsonr/sonr/x/vault/types"
)

func (k Keeper) Logger() log.Logger {
	return k.logger
}

// InitGenesis initializes the module's state from a genesis state.
func (k *Keeper) InitGenesis(ctx context.Context, data *types.GenesisState) error {
	// this line is used by starport scaffolding # genesis/module/init
	if err := data.Params.Validate(); err != nil {
		return err
	}

	return k.Params.Set(ctx, data.Params)
}

// ExportGenesis exports the module's state to a genesis state.
func (k *Keeper) ExportGenesis(ctx context.Context) *types.GenesisState {
	params, err := k.Params.Get(ctx)
	if err != nil {
		panic(err)
	}

	// this line is used by starport scaffolding # genesis/module/export

	return &types.GenesisState{
		Params: params,
	}
}

// IPFSConnected returns true if the IPFS client is initialized
func (c Keeper) IPFSConnected() bool {
	if c.ipfsClient == nil {
		ipfsClient, err := rpc.NewLocalApi()
		if err != nil {
			return false
		}
		c.ipfsClient = ipfsClient
	}
	return c.ipfsClient != nil
}

// CalculateExpiration calculates the expiration time for a vault
func (k Keeper) CalculateExpiration(c sdk.Context, duration time.Duration) int64 {
	blockTime := c.BlockTime()
	avgBlockTime := float64(blockTime.Sub(blockTime).Seconds())
	return int64(duration.Seconds() / avgBlockTime)
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
