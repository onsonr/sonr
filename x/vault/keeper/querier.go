package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/onsonr/sonr/pkg/crypto/mpc"
	"github.com/onsonr/sonr/x/did/controller"
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
		ctx.Logger().Error(fmt.Sprintf("Error getting current schema: %s", err.Error()))
		return nil, types.ErrInvalidSchema.Wrap(err.Error())
	}

	// 2. Generate MPC Keyshares for new Account
	shares, err := mpc.NewKeyshareSource()
	if err != nil {
		ctx.Logger().Error(fmt.Sprintf("Error generating keyshares: %s", err.Error()))
		return nil, types.ErrInvalidSchema.Wrap(err.Error())
	}

	// 3. Create Controller from Keyshares
	con, err := controller.New(shares)
	if err != nil {
		ctx.Logger().Error(fmt.Sprintf("Error creating controller: %s", err.Error()))
		return nil, types.ErrControllerCreation.Wrap(err.Error())
	}

	// 4. Create a new vault PWA for service-worker
	v, err := types.SpawnVault("", con.SonrAddress(), con.ChainID(), sch)
	if err != nil {
		ctx.Logger().Error(fmt.Sprintf("Error creating vault: %s", err.Error()))
		return nil, types.ErrInvalidSchema.Wrap(err.Error())
	}

	// 5. Add to IPFS and Return CID for User Claims in Gateway
	cid, err := k.ipfsClient.Unixfs().Add(context.Background(), v)
	if err != nil {
		ctx.Logger().Error(fmt.Sprintf("Error adding to IPFS: %s", err.Error()))
		return nil, types.ErrVaultAssembly.Wrap(err.Error())
	}

	// 6. Return final response
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
