package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"cosmossdk.io/errors"
	"github.com/onsonr/hway/x/did/types"
)

type msgServer struct {
	k Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the module MsgServer interface.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{k: keeper}
}

// InitializeController implements types.MsgServer.
func (ms msgServer) InitializeController(goCtx context.Context, msg *types.MsgInitializeController) (*types.MsgInitializeControllerResponse, error) {
	_ = sdk.UnwrapSDKContext(goCtx)
	return &types.MsgInitializeControllerResponse{}, nil
}

// UpdateParams updates the x/did module parameters.
func (ms msgServer) UpdateParams(ctx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if ms.k.authority != msg.Authority {
		return nil, errors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", ms.k.authority, msg.Authority)
	}

	return nil, ms.k.Params.Set(ctx, msg.Params)
}

// AuthenticateController implements types.MsgServer.
func (ms msgServer) AuthenticateController(ctx context.Context, msg *types.MsgAuthenticateController) (*types.MsgAuthenticateControllerResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("AuthenticateController is unimplemented")
	return &types.MsgAuthenticateControllerResponse{}, nil
}
