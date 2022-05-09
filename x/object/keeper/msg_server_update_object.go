package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/x/object/types"
)

func (k msgServer) UpdateObject(goCtx context.Context, msg *types.MsgUpdateObject) (*types.MsgUpdateObjectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get Properties
	appName, err := msg.GetSession().GetWhois().ApplicationName()
	if err != nil {
		return nil, err
	}

	// Generate a new Object Did
	did, err := msg.GetSession().GenerateDID(did.WithPathSegments(appName, "object"), did.WithFragment(msg.GetLabel()))
	if err != nil {
		return nil, err
	}

	// Check if Object exists
	whatis, found := k.GetWhatIs(ctx, did.ID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Object was not found in this Application")
	}

	// Check if Object is IsActive
	if !whatis.IsActive {
		return nil, types.ErrInactiveObject
	}

	// Create New Field Map
	whatis.ObjectDoc.AddFields(msg.GetAddedFields()...)
	whatis.ObjectDoc.RemoveFields(msg.GetRemovedFields()...)
	whatis.Timestamp = time.Now().Unix()
	whatis.IsActive = true
	k.SetWhatIs(ctx, whatis)

	// Return Response
	return &types.MsgUpdateObjectResponse{
		WhatIs:  &whatis,
		Code:    100,
		Message: fmt.Sprintf("Existing Object %s has been updated", whatis.Did),
	}, nil
}
