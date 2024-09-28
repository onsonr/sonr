package keeper

import (
	"context"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/onsonr/sonr/x/did/types"
)

type msgServer struct {
	k Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the module MsgServer interface.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{k: keeper}
}

// RegisterController implements types.MsgServer.
func (ms msgServer) RegisterController(goCtx context.Context, msg *types.MsgRegisterController) (*types.MsgRegisterControllerResponse, error) {
	_ = sdk.UnwrapSDKContext(goCtx)
	return &types.MsgRegisterControllerResponse{}, nil
}

// AuthorizeService implements types.MsgServer.
func (ms msgServer) AuthorizeService(goCtx context.Context, msg *types.MsgAuthorizeService) (*types.MsgAuthorizeServiceResponse, error) {
	if ms.k.authority != msg.Controller {
		return nil, errors.Wrapf(
			govtypes.ErrInvalidSigner,
			"invalid authority; expected %s, got %s",
			ms.k.authority,
			msg.Controller,
		)
	}
	return &types.MsgAuthorizeServiceResponse{}, nil
}

// UpdateParams updates the x/did module parameters.
func (ms msgServer) UpdateParams(ctx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if ms.k.authority != msg.Authority {
		return nil, errors.Wrapf(
			govtypes.ErrInvalidSigner,
			"invalid authority; expected %s, got %s",
			ms.k.authority,
			msg.Authority,
		)
	}
	return nil, ms.k.Params.Set(ctx, msg.Params)
}

// ExecuteTx implements types.MsgServer.
func (ms msgServer) ExecuteTx(ctx context.Context, msg *types.MsgExecuteTx) (*types.MsgExecuteTxResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.MsgExecuteTxResponse{}, nil
}
