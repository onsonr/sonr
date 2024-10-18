package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/onsonr/sonr/x/vault/types"
)

var _ types.QueryServer = Querier{}

type Querier struct{ Keeper }

func NewQuerier(keeper Keeper) Querier {
	return Querier{Keeper: keeper}
}

// ╭───────────────────────────────────────────────────────────╮
// │                  Fixed Query Methods                      │
// ╰───────────────────────────────────────────────────────────╯

// Params implements types.QueryServer.
func (k Querier) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	p, err := k.Keeper.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	return &types.QueryParamsResponse{Params: &p}, nil
}

// Schema implements types.QueryServer.
func (k Querier) Schema(goCtx context.Context, req *types.QuerySchemaRequest) (*types.QuerySchemaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	p, err := k.Keeper.Params.Get(ctx)
	if err != nil {
		return nil, err
	}
	return &types.QuerySchemaResponse{
		Schema: p.Schema,
	}, nil
}

// ╭───────────────────────────────────────────────────────────╮
// │                  Pre-Authenticated Queries                │
// ╰───────────────────────────────────────────────────────────╯

// Allocate implements types.QueryServer.
func (k Querier) Allocate(goCtx context.Context, req *types.AllocateRequest) (*types.AllocateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 1. Get current schema
	sch, err := k.currentSchema(ctx)
	if err != nil {
		return nil, types.ErrInvalidSchema.Wrap(err.Error())
	}

	// 2.Allocate the vault msg.GetSubject(), msg.GetOrigin()
	cid, expiryBlock, err := k.assembleVault(ctx, sch)
	if err != nil {
		return nil, types.ErrVaultAssembly.Wrap(err.Error())
	}

	return &types.AllocateResponse{
		Success:     true,
		Cid:         cid,
		ExpiryBlock: expiryBlock,
	}, nil
}

// ╭───────────────────────────────────────────────────────────╮
// │                  Authenticated Endpoints                  │
// ╰───────────────────────────────────────────────────────────╯

// Sync implements types.QueryServer.
func (k Querier) Sync(goCtx context.Context, req *types.SyncRequest) (*types.SyncResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	p, err := k.Keeper.Params.Get(ctx)
	if err != nil {
		return nil, err
	}
	c, _ := k.DIDKeeper.ResolveController(ctx, req.Did)
	if c == nil {
		return &types.SyncResponse{
			Success: false,
			Schema:  p.Schema,
			ChainID: ctx.ChainID(),
		}, nil
	}
	return &types.SyncResponse{
		Success: true,
		Schema:  p.Schema,
		ChainID: ctx.ChainID(),
		Address: c.SonrAddress(),
	}, nil
}
