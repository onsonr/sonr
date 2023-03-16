package keeper

import (
	"context"

	"github.com/bufbuild/connect-go"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonrhq/core/x/identity/types"
)

func (s *msgServerConnectWrapper) RegisterService(
	ctx context.Context,
	req *connect.Request[types.MsgRegisterService],
) (*connect.Response[types.MsgRegisterServiceResponse], error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Check if the value already exists
	_, isFound := s.GetDidDocument(
		sdkCtx,
		req.Msg.Creator,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	res := connect.NewResponse(&types.MsgRegisterServiceResponse{})
	return res, nil
}

func (s *msgServerConnectWrapper) UpdateService(
	ctx context.Context,
	req *connect.Request[types.MsgUpdateService],
) (*connect.Response[types.MsgUpdateServiceResponse], error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// T-18 Handling the message
	_ = sdkCtx

	res := connect.NewResponse(&types.MsgUpdateServiceResponse{})
	return res, nil
}

func (s *msgServerConnectWrapper) DeactivateService(
	ctx context.Context,
	req *connect.Request[types.MsgDeactivateService],
) (*connect.Response[types.MsgDeactivateServiceResponse], error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// T-19 Handling the message
	_ = sdkCtx

	res := connect.NewResponse(&types.MsgDeactivateServiceResponse{})
	return res, nil
}
