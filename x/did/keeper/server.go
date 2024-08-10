package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"cosmossdk.io/errors"
	didv1 "github.com/onsonr/hway/api/did/v1"
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

// Authenticate implements types.MsgServer.
func (ms msgServer) Authenticate(ctx context.Context, msg *types.MsgAuthenticate) (*types.MsgAuthenticateResponse, error) {
	if ms.k.authority != msg.Authority {
		return nil, errors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", ms.k.authority, msg.Authority)
	}
	// ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.MsgAuthenticateResponse{}, nil
}

// RegisterController implements types.MsgServer.
func (ms msgServer) RegisterController(goCtx context.Context, msg *types.MsgRegisterController) (*types.MsgRegisterControllerResponse, error) {
	if ms.k.authority != msg.Authority {
		return nil, errors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", ms.k.authority, msg.Authority)
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	svc := didv1.Service{
		ControllerDid: msg.Authority,
	}
	ms.k.OrmDB.ServiceTable().Insert(ctx, &svc)
	return &types.MsgRegisterControllerResponse{}, nil
}

// RegisterService implements types.MsgServer.
func (ms msgServer) RegisterService(ctx context.Context, msg *types.MsgRegisterService) (*types.MsgRegisterServiceResponse, error) {
	if ms.k.authority != msg.Authority {
		return nil, errors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", ms.k.authority, msg.Authority)
	}
	// ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.MsgRegisterServiceResponse{}, nil
}
