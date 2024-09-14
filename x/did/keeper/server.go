package keeper

import (
	"context"
	"encoding/json"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/onsonr/sonr/x/did/builder"
	"github.com/onsonr/sonr/x/did/types"

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

//	# AllocateVault
//
// AllocateVault implements types.MsgServer.
func (ms msgServer) AllocateVault(
	goCtx context.Context,
	msg *types.MsgAllocateVault,
) (*types.MsgAllocateVaultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// 1.Check if the service origin is valid
	if ms.k.IsValidServiceOrigin(ctx, msg.Origin) {
		return nil, types.ErrInvalidServiceOrigin
	}

	cid, expiryBlock, err := ms.k.assembleInitialVault(ctx)
	if err != nil {
		return nil, err
	}

	regOpts, err := builder.GetPublicKeyCredentialCreationOptions(msg.Origin, msg.Subject, cid, ms.k.GetParams(ctx))
	if err != nil {
		return nil, err
	}

	// Convert to string
	regOptsJSON, err := json.Marshal(regOpts)
	if err != nil {
		return nil, err
	}

	return &types.MsgAllocateVaultResponse{
		ExpiryBlock:         expiryBlock,
		Cid:                 cid,
		RegistrationOptions: string(regOptsJSON),
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
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 1.Check if the service origin is valid
	if !ms.k.IsValidServiceOrigin(ctx, msg.Service.Origin) {
		return nil, types.ErrInvalidServiceOrigin
	}
	return ms.k.insertService(ctx, msg.Service)
}

// # SyncController
//
// SyncController implements types.MsgServer.
func (ms msgServer) SyncController(ctx context.Context, msg *types.MsgSyncController) (*types.MsgSyncControllerResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.MsgSyncControllerResponse{}, nil
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
	return &types.MsgRegisterControllerResponse{}, nil
}

// RegisterService implements types.MsgServer.
func (ms msgServer) RegisterService(goCtx context.Context, msg *types.MsgRegisterService) (*types.MsgRegisterServiceResponse, error) {
	if ms.k.authority != msg.Authority {
		return nil, errors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", ms.k.authority, msg.Authority)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	svc := didv1.Service{
		ControllerDid: msg.Authority,
	}
	err := ms.k.OrmDB.ServiceTable().Insert(ctx, &svc)
	if err != nil {
		return nil, err
	}
	return &types.MsgRegisterServiceResponse{}, nil
}

// ProveWitness implements types.MsgServer.
func (ms msgServer) ProveWitness(ctx context.Context, msg *types.MsgProveWitness) (*types.MsgProveWitnessResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.MsgProveWitnessResponse{}, nil
}

// SyncVault implements types.MsgServer.
func (ms msgServer) SyncVault(ctx context.Context, msg *types.MsgSyncVault) (*types.MsgSyncVaultResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.MsgSyncVaultResponse{}, nil
}
