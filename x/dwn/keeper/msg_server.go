package keeper

import (
	"context"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"cosmossdk.io/errors"
	"github.com/onsonr/sonr/x/dwn/types"
)

type msgServer struct {
	k Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the module MsgServer interface.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{k: keeper}
}

func (ms msgServer) UpdateParams(ctx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if ms.k.authority != msg.Authority {
		return nil, errors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", ms.k.authority, msg.Authority)
	}

	return nil, ms.k.Params.Set(ctx, msg.Params)
}

// Initialize implements types.MsgServer.
func (ms msgServer) Initialize(ctx context.Context, msg *types.MsgInitialize) (*types.MsgInitializeResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("Initialize is unimplemented")
	return &types.MsgInitializeResponse{}, nil
}

// Spawn implements types.MsgServer.
func (ms msgServer) Spawn(ctx context.Context, msg *types.MsgSpawn) (*types.MsgSpawnResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("Spawn is unimplemented")
	return &types.MsgSpawnResponse{}, nil
}
