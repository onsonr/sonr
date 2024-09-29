package keeper

import (
	"context"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"cosmossdk.io/errors"
	"github.com/onsonr/sonr/x/macaroon/types"
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

// AuthorizeService implements types.MsgServer.
func (ms msgServer) AuthorizeService(ctx context.Context, msg *types.MsgIssueMacaroon) (*types.MsgIssueMacaroonResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("AuthorizeService is unimplemented")
	return &types.MsgIssueMacaroonResponse{}, nil
}

// IssueMacaroon implements types.MsgServer.
func (ms msgServer) IssueMacaroon(ctx context.Context, msg *types.MsgIssueMacaroon) (*types.MsgIssueMacaroonResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("IssueMacaroon is unimplemented")
	return &types.MsgIssueMacaroonResponse{}, nil
}
