package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonrhq/core/x/identity/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// ! ||--------------------------------------------------------------------------------||
// ! ||                    DIDDocument Message Server Implementation                   ||
// ! ||--------------------------------------------------------------------------------||

func (k msgServer) CreateDidDocument(goCtx context.Context, msg *types.MsgCreateDidDocument) (*types.MsgCreateDidDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// Check if the value already exists
	_, ok := k.GetPrimaryIdentity(ctx, msg.Primary.Id)
	if ok {
		return nil,  sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}
	_, found := k.GetPrimaryIdentityByAlias(ctx, msg.Primary.AlsoKnownAs[0])
	if found {
		return nil,  sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}
	// Set the value
	k.SetPrimaryIdentity(
		ctx,
		*msg.Primary,
	)

	// Set the blockchain identities
	k.SetBlockchainIdentities(ctx, msg.Blockchains...)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent("NewTx", sdk.NewAttribute("tx-name", "create-did-document"), sdk.NewAttribute("did", msg.Primary.Id), sdk.NewAttribute("creator", msg.Creator), sdk.NewAttribute("alias", msg.Alias)),
	)
	return &types.MsgCreateDidDocumentResponse{}, nil
}

func (k msgServer) UpdateDidDocument(goCtx context.Context, msg *types.MsgUpdateDidDocument) (*types.MsgUpdateDidDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetPrimaryIdentity(
		ctx,
		msg.Primary.Id,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Check if the msg creator is the same as the current owner
	if !valFound.CheckAccAddress(msg.Creator) {
		return nil, types.ErrUnauthorized
	}

	k.SetPrimaryIdentity(ctx, *msg.Primary)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent("NewTx", sdk.NewAttribute("tx-name", "update-did-document"), sdk.NewAttribute("did", msg.Primary.Id), sdk.NewAttribute("creator", msg.Creator)),
	)
	return &types.MsgUpdateDidDocumentResponse{}, nil
}

func (k msgServer) DeleteDidDocument(goCtx context.Context, msg *types.MsgDeleteDidDocument) (*types.MsgDeleteDidDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetPrimaryIdentity(
		ctx,
		msg.Did,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Check if the msg creator is the same as the current owner
	if !valFound.CheckAccAddress(msg.Creator) {
		return nil, types.ErrUnauthorized
	}

	k.RemovePrimaryIdentity(
		ctx,
		msg.Did,
	)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent("NewTx", sdk.NewAttribute("tx-name", "delete-did-document"), sdk.NewAttribute("did", msg.Did), sdk.NewAttribute("creator", msg.Creator)),
	)
	return &types.MsgDeleteDidDocumentResponse{}, nil
}
