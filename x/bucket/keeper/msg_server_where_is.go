package keeper

import (
	"context"
	"fmt"
	"net/http"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/sonr/x/bucket/types"
)

func (k msgServer) CreateWhereIs(goCtx context.Context, msg *types.MsgCreateWhereIs) (*types.MsgCreateWhereIsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := msg.ValidateBasic()
	k.Logger(ctx).Info("basic request validation finished")

	if err != nil {
		return nil, err
	}

	accts := msg.GetSigners()
	if len(accts) < 1 {
		k.Logger(ctx).Error("Error while querying account: not found")
		return nil, sdkerrors.ErrNotFound
	}

	uuid := k.GenerateKeyForDID()

	did := fmt.Sprintf("did:snr:%s", uuid)

	var whereIs = types.WhereIs{
		Creator:    msg.Creator,
		Label:      msg.Label,
		Did:        did,
		Visibility: msg.Visibility,
		Role:       msg.Role,
		IsActive:   true,
		Content:    msg.Content,
		Timestamp:  time.Now().Unix(),
	}
	fmt.Printf("label: %s, vi: %d, role: %d \n", whereIs.Label, whereIs.Visibility, whereIs.Role)
	k.SetWhereIs(
		ctx,
		whereIs,
	)

	return &types.MsgCreateWhereIsResponse{
		Status:  http.StatusAccepted,
		WhereIs: &whereIs,
	}, nil
}

func (k msgServer) UpdateWhereIs(goCtx context.Context, msg *types.MsgUpdateWhereIs) (*types.MsgUpdateWhereIsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := msg.ValidateBasic()
	k.Logger(ctx).Info("basic request validation finished")

	if err != nil {
		return nil, err
	}

	accts := msg.GetSigners()
	if len(accts) < 1 {
		k.Logger(ctx).Error("Error while querying account: not found")
		return nil, sdkerrors.ErrNotFound
	}

	var whereIs = types.WhereIs{
		Label:      msg.Label,
		Creator:    msg.Creator,
		Did:        msg.Did,
		Visibility: msg.Visibility,
		Role:       msg.Role,
		IsActive:   true,
		Content:    msg.Content,
		Timestamp:  time.Now().Unix(),
	}

	// Checks that the element exists
	val, found := k.GetWhereIs(ctx, msg.Creator, msg.Did)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Did))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetWhereIs(ctx, whereIs)

	return &types.MsgUpdateWhereIsResponse{
		Status:  http.StatusAccepted,
		WhereIs: &whereIs,
	}, nil
}

func (k msgServer) DeleteWhereIs(goCtx context.Context, msg *types.MsgDeleteWhereIs) (*types.MsgDeleteWhereIsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetWhereIs(ctx, msg.Creator, msg.Did)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Did))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	val.IsActive = false
	k.SetWhereIs(ctx, val)

	return &types.MsgDeleteWhereIsResponse{}, nil
}
