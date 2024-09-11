package keeper

import (
	"context"
	"encoding/json"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/onsonr/sonr/x/did/builder"
	snrctx "github.com/onsonr/sonr/x/did/context"
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

// AllocateVault implements types.MsgServer.
func (ms msgServer) AllocateVault(
	goCtx context.Context,
	msg *types.MsgAllocateVault,
) (*types.MsgAllocateVaultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	clientInfo, err := snrctx.ExtractClientInfo(goCtx)
	if err != nil {
		return nil, err
	}

	// 1.Check if the service origin is valid
	if ms.k.IsValidServiceOrigin(ctx, msg.Origin, clientInfo) {
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

// RegisterController implements types.MsgServer.
func (ms msgServer) RegisterController(
	goCtx context.Context,
	msg *types.MsgRegisterController,
) (*types.MsgRegisterControllerResponse, error) {
	_ = sdk.UnwrapSDKContext(goCtx)
	return &types.MsgRegisterControllerResponse{}, nil
}

// RegisterService implements types.MsgServer.
func (ms msgServer) RegisterService(
	goCtx context.Context,
	msg *types.MsgRegisterService,
) (*types.MsgRegisterServiceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	clientInfo, err := snrctx.ExtractClientInfo(goCtx)
	if err != nil {
		return nil, err
	}

	// 1.Check if the service origin is valid
	if !ms.k.IsValidServiceOrigin(ctx, msg.Service.Origin, clientInfo) {
		return nil, types.ErrInvalidServiceOrigin
	}
	return ms.k.insertService(ctx, msg.Service)
}

// SyncController implements types.MsgServer.
func (ms msgServer) SyncController(ctx context.Context, msg *types.MsgSyncController) (*types.MsgSyncControllerResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.MsgSyncControllerResponse{}, nil
}

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
