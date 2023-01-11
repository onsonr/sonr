package keeper

import (
	"context"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-hq/sonr/x/identity/types"
)

func (k msgServer) CreateDidDocument(goCtx context.Context, msg *types.MsgCreateDidDocument) (*types.MsgCreateDidDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// Check if the value already exists
	_, isFound := k.GetDidDocument(
		ctx,
		msg.Document.ID,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}
	ptrs := strings.Split(msg.Document.ID, ":")
	addr := fmt.Sprintf("%s%s", ptrs[len(ptrs)-2], ptrs[len(ptrs)-1])
	acc := k.accountKeeper.NewAccountWithAddress(ctx, sdk.AccAddress(addr))
	k.accountKeeper.SetAccount(ctx, acc)
	k.SetDidDocument(
		ctx,
		*msg.Document,
	)
	return &types.MsgCreateDidDocumentResponse{}, nil
}

func (k msgServer) UpdateDidDocument(goCtx context.Context, msg *types.MsgUpdateDidDocument) (*types.MsgUpdateDidDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetDidDocument(
		ctx,
		msg.Document.ID,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if valFound.CheckAccAddress(msg.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	k.SetDidDocument(ctx, *msg.Document)
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

	// Checks if the the msg creator is the same as the current owner
	if valFound.CheckAccAddress(msg.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveDidDocument(
		ctx,
		msg.Did,
	)

	return &types.MsgDeleteDidDocumentResponse{}, nil
}
