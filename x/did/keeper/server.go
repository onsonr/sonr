package keeper

import (
	"context"

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

// UpdateParams updates the x/did module parameters.
func (ms msgServer) UpdateParams(ctx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if ms.k.authority != msg.Authority {
		return nil, errors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", ms.k.authority, msg.Authority)
	}

	return nil, ms.k.Params.Set(ctx, msg.Params)
}

// RegisterController implements types.MsgServer.
func (ms msgServer) RegisterController(ctx context.Context, msg *types.MsgRegisterController) (*types.MsgRegisterControllerResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("RegisterController is unimplemented")
	return &types.MsgRegisterControllerResponse{}, nil
}

// Authenticate implements types.MsgServer.
func (ms msgServer) Authenticate(ctx context.Context, msg *types.MsgAuthenticate) (*types.MsgAuthenticateResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("Authenticate is unimplemented")
	return &types.MsgAuthenticateResponse{}, nil
}

// RegisterService implements types.MsgServer.
func (ms msgServer) RegisterService(ctx context.Context, msg *types.MsgRegisterService) (*types.MsgRegisterServiceResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("RegisterService is unimplemented")
	return &types.MsgRegisterServiceResponse{}, nil
}
