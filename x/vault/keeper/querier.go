package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onsonr/crypto/mpc"
	didtypes "github.com/onsonr/sonr/x/did/types"

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
func (k Querier) Allocate(goCtx context.Context, req *types.QueryAllocateRequest) (*types.QueryAllocateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 1. Get current schema
	sch, err := k.currentSchema(ctx)
	if err != nil {
		ctx.Logger().Error(err.Error())
		return nil, types.ErrInvalidSchema.Wrap(err.Error())
	}
	shares, err := mpc.GenerateKeyshares()
	if err != nil {
		ctx.Logger().Error(err.Error())
		return nil, err
	}
	con, err := didtypes.NewController(ctx, shares)
	if err != nil {
		ctx.Logger().Error(err.Error())
		return nil, err
	}

	usrKs, err := con.ExportUserKs()
	if err != nil {
		ctx.Logger().Error(err.Error())
		return nil, types.ErrInvalidSchema.Wrap(err.Error())
	}
	v, err := types.NewVault(usrKs, con.SonrAddress(), con.ChainID(), sch)
	if err != nil {
		ctx.Logger().Error(err.Error())
		return nil, types.ErrInvalidSchema.Wrap(err.Error())
	}
	cid, err := k.ipfsClient.Unixfs().Add(context.Background(), v.FS)
	if err != nil {
		ctx.Logger().Error(err.Error())
		return nil, types.ErrVaultAssembly.Wrap(err.Error())
	}

	return &types.QueryAllocateResponse{
		Success:     true,
		Cid:         cid.String(),
		ExpiryBlock: calculateBlockExpiry(ctx, time.Second*30),
	}, nil
}

// ╭───────────────────────────────────────────────────────────╮
// │                  Authenticated Endpoints                  │
// ╰───────────────────────────────────────────────────────────╯

// Sync implements types.QueryServer.
func (k Querier) Sync(goCtx context.Context, req *types.QuerySyncRequest) (*types.QuerySyncResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	p, err := k.Keeper.Params.Get(ctx)
	if err != nil {
		return nil, err
	}
	c, _ := k.DIDKeeper.ResolveController(ctx, req.Did)
	if c == nil {
		return &types.QuerySyncResponse{
			Success: false,
			Schema:  p.Schema,
			ChainID: ctx.ChainID(),
		}, nil
	}
	return &types.QuerySyncResponse{
		Success: true,
		Schema:  p.Schema,
		ChainID: ctx.ChainID(),
		Address: c.SonrAddress(),
	}, nil
}
