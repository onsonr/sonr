package keeper

import (
	"context"
	"encoding/json"

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

//	# AllocateVault
//
// AllocateVault implements types.MsgServer.
func (ms msgServer) AllocateVault(
	goCtx context.Context,
	msg *types.MsgAllocateVault,
) (*types.MsgAllocateVaultResponse, error) {
	// 1.Check if the service origin is valid
	ctx := ms.k.UnwrapCtx(goCtx)
	if err := ctx.ValidateOrigin(msg.Origin); err != nil {
		return nil, err
	}

	// 2.Allocate the vault
	cid, expiryBlock, err := ms.k.AssembleVault(ctx, msg.GetSubject(), msg.GetOrigin())
	if err != nil {
		return nil, err
	}

	// 3.Return the response
	return &types.MsgAllocateVaultResponse{
		ExpiryBlock: expiryBlock,
		Cid:         cid,
	}, nil
}

// # RegisterController
//
// RegisterController implements types.MsgServer.
func (ms msgServer) RegisterController(
	goCtx context.Context,
	msg *types.MsgRegisterController,
) (*types.MsgRegisterControllerResponse, error) {
	_ = sdk.UnwrapSDKContext(goCtx)
	return &types.MsgRegisterControllerResponse{}, nil
}

// # RegisterService
//
// RegisterService implements types.MsgServer.
func (ms msgServer) RegisterService(
	goCtx context.Context,
	msg *types.MsgRegisterService,
) (*types.MsgRegisterServiceResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)

	// 1.Check if the service origin is valid
	// if !ms.k.IsValidServiceOrigin(ctx, msg.Service.Origin) {
	// 	return nil, types.ErrInvalidServiceOrigin
	// }
	return nil, errors.Wrapf(types.ErrInvalidServiceOrigin, "invalid service origin")
}

// # AuthorizeService
//
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

// # UpdateParams
//
// UpdateParams updates the x/did module parameters.
func (ms msgServer) UpdateParams(
	ctx context.Context,
	msg *types.MsgUpdateParams,
) (*types.MsgUpdateParamsResponse, error) {
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
