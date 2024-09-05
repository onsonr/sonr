package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"cosmossdk.io/errors"
	didv1 "github.com/onsonr/sonr/api/did/v1"
	"github.com/onsonr/sonr/internal/files"
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

// UpdateParams updates the x/did module parameters.
func (ms msgServer) UpdateParams(ctx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if ms.k.authority != msg.Authority {
		return nil, errors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", ms.k.authority, msg.Authority)
	}

	return nil, ms.k.Params.Set(ctx, msg.Params)
}

// Authorize implements types.MsgServer.
func (ms msgServer) Authorize(ctx context.Context, msg *types.MsgAuthorize) (*types.MsgAuthorizeResponse, error) {
	if ms.k.authority != msg.Authority {
		return nil, errors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", ms.k.authority, msg.Authority)
	}
	// ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.MsgAuthorizeResponse{}, nil
}

// AllocateVault implements types.MsgServer.
func (ms msgServer) AllocateVault(goCtx context.Context, msg *types.MsgAllocateVault) (*types.MsgAllocateVaultResponse, error) {
	//	ctx := sdk.UnwrapSDKContext(goCtx)
	err := files.Assemble("/tmp/sonr-testnet-1/vaults/0")
	if err != nil {
		return nil, err
	}
	return &types.MsgAllocateVaultResponse{}, nil
}

// RegisterController implements types.MsgServer.
func (ms msgServer) RegisterController(goCtx context.Context, msg *types.MsgRegisterController) (*types.MsgRegisterControllerResponse, error) {
	if ms.k.authority != msg.Authority {
		return nil, errors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", ms.k.authority, msg.Authority)
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	svc := didv1.ServiceRecord{
		Controller: msg.Authority,
	}
	ms.k.OrmDB.ServiceRecordTable().Insert(ctx, &svc)
	return &types.MsgRegisterControllerResponse{}, nil
}

// RegisterService implements types.MsgServer.
func (ms msgServer) RegisterService(ctx context.Context, msg *types.MsgRegisterService) (*types.MsgRegisterServiceResponse, error) {
	if ms.k.authority != msg.Controller {
		return nil, errors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", ms.k.authority, msg.Controller)
	}

	// ctx := sdk.UnwrapSDKContext(goCtx)
	svc := didv1.ServiceRecord{
		Controller: msg.Controller,
	}
	ms.k.OrmDB.ServiceRecordTable().Insert(ctx, &svc)
	return &types.MsgRegisterServiceResponse{}, nil
}

// SyncVault implements types.MsgServer.
func (ms msgServer) SyncVault(ctx context.Context, msg *types.MsgSyncVault) (*types.MsgSyncVaultResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.MsgSyncVaultResponse{}, nil
}
