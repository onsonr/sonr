package keeper

import (
	"context"
	"encoding/json"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/onsonr/sonr/x/did/builder"
	"github.com/onsonr/sonr/x/did/middleware"
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

// Authorize implements types.MsgServer.
func (ms msgServer) Authorize(
	ctx context.Context,
	msg *types.MsgAuthorize,
) (*types.MsgAuthorizeResponse, error) {
	if ms.k.authority != msg.Authority {
		return nil, errors.Wrapf(
			govtypes.ErrInvalidSigner,
			"invalid authority; expected %s, got %s",
			ms.k.authority,
			msg.Authority,
		)
	}
	// ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.MsgAuthorizeResponse{}, nil
}

// AllocateVault implements types.MsgServer.
func (ms msgServer) AllocateVault(
	goCtx context.Context,
	msg *types.MsgAllocateVault,
) (*types.MsgAllocateVaultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	clientInfo, err := middleware.ExtractClientInfo(goCtx)
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

	regOpts, err := builder.NewRegistrationOptions(msg.Origin, msg.Subject, cid, ms.k.GetParams(ctx))
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

	clientInfo, err := middleware.ExtractClientInfo(goCtx)
	if err != nil {
		return nil, err
	}

	// 1.Check if the service origin is valid
	if !ms.k.IsValidServiceOrigin(ctx, msg.OriginUri, clientInfo) {
		return nil, types.ErrInvalidServiceOrigin
	}
	return ms.k.insertService(ctx, msg)
}

// SyncVault implements types.MsgServer.
func (ms msgServer) SyncVault(
	ctx context.Context,
	msg *types.MsgSyncVault,
) (*types.MsgSyncVaultResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.MsgSyncVaultResponse{}, nil
}
