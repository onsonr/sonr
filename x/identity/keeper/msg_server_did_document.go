package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonrhq/core/x/identity/types"
)

func (k msgServer) CreateDidDocument(goCtx context.Context, msg *types.MsgCreateDidDocument) (*types.MsgCreateDidDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// Check if the value already exists
	_, isFound := k.GetDidDocument(
		ctx,
		msg.Document.Id,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	k.SetDidDocument(
		ctx,
		*msg.Document,
	)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent("create-did-document", sdk.NewAttribute("did", msg.Document.Id), sdk.NewAttribute("creator", msg.Creator), sdk.NewAttribute("address", msg.Document.Address())),
	)

	return &types.MsgCreateDidDocumentResponse{}, nil
}

func (k msgServer) UpdateDidDocument(goCtx context.Context, msg *types.MsgUpdateDidDocument) (*types.MsgUpdateDidDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetDidDocument(
		ctx,
		msg.Document.Id,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Check if the msg creator is the same as the current owner
	if !valFound.CheckAccAddress(msg.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetDidDocument(ctx, *msg.Document)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent("update-did-document", sdk.NewAttribute("did", msg.Document.Id), sdk.NewAttribute("creator", msg.Creator), sdk.NewAttribute("address", msg.Document.Address())),
	)
	return &types.MsgUpdateDidDocumentResponse{}, nil
}

func (k msgServer) DeleteDidDocument(goCtx context.Context, msg *types.MsgDeleteDidDocument) (*types.MsgDeleteDidDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetDidDocument(
		ctx,
		msg.Did,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Check if the msg creator is the same as the current owner
	if !valFound.CheckAccAddress(msg.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveDidDocument(
		ctx,
		msg.Did,
	)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent("delete-did-document", sdk.NewAttribute("document-did", msg.Did), sdk.NewAttribute("creator", msg.Creator)),
	)
	return &types.MsgDeleteDidDocumentResponse{}, nil
}
