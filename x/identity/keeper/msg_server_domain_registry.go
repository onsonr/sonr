package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-hq/sonr/x/identity/types"
)

func (k msgServer) CreateDomainRecord(goCtx context.Context, msg *types.MsgCreateDomainRecord) (*types.MsgCreateDomainRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetDomainRecord(
		ctx,
		msg.Index,
		msg.Domain,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var DomainRecord = types.DomainRecord{
		Creator: msg.Creator,
		Index:   msg.Index,
	}

	k.SetDomainRecord(
		ctx,
		DomainRecord,
	)
	return &types.MsgCreateDomainRecordResponse{}, nil
}

func (k msgServer) UpdateDomainRecord(goCtx context.Context, msg *types.MsgUpdateDomainRecord) (*types.MsgUpdateDomainRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetDomainRecord(
		ctx,
		msg.Index,
		msg.Domain,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var DomainRecord = types.DomainRecord{
		Creator: msg.Creator,
		Index:   msg.Index,
	}

	k.SetDomainRecord(ctx, DomainRecord)

	return &types.MsgUpdateDomainRecordResponse{}, nil
}

func (k msgServer) DeleteDomainRecord(goCtx context.Context, msg *types.MsgDeleteDomainRecord) (*types.MsgDeleteDomainRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetDomainRecord(
		ctx,
		msg.Index,
		msg.Domain,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveDomainRecord(
		ctx,
		msg.Index,
		msg.Domain,
	)

	return &types.MsgDeleteDomainRecordResponse{}, nil
}
