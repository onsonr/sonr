package keeper

import (
	"context"

	"cosmossdk.io/log"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ipfs/boxo/path"
	"google.golang.org/grpc/peer"

	"github.com/onsonr/sonr/x/did/types"
)

// Logger returns the logger
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

// CheckValidatorExists checks if a validator exists
func (k Keeper) CheckValidatorExists(ctx sdk.Context, addr string) bool {
	address, err := sdk.ValAddressFromBech32(addr)
	if err != nil {
		return false
	}
	ok, err := k.StakingKeeper.Validator(ctx, address)
	if err != nil {
		return false
	}
	if ok != nil {
		return true
	}
	return false
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

func (k Keeper) UnwrapCtx(goCtx context.Context) Context {
	ctx := sdk.UnwrapSDKContext(goCtx)
	peer, _ := peer.FromContext(goCtx)
	return Context{SDKCtx: ctx, Peer: peer, Keeper: k}
}
