package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonrhq/core/x/identity/types"
)

func (k msgServer) RegisterService(goCtx context.Context, msg *types.MsgRegisterService) (*types.MsgRegisterServiceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetDidDocument(
		ctx,
		msg.Creator,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	return &types.MsgRegisterServiceResponse{}, nil
}

func (k msgServer) UpdateService(goCtx context.Context, msg *types.MsgUpdateService) (*types.MsgUpdateServiceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// T-18 Handling the message
	_ = ctx

	return &types.MsgUpdateServiceResponse{}, nil
}

func (k msgServer) DeactivateService(goCtx context.Context, msg *types.MsgDeactivateService) (*types.MsgDeactivateServiceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// T-19 Handling the message
	_ = ctx

	return &types.MsgDeactivateServiceResponse{}, nil
}
