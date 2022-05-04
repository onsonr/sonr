package msg

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/sonr/internal/blockchain/x/object/types"
	"github.com/sonr-io/sonr/pkg/did"
)

func (k msgServer) CreateObject(goCtx context.Context, msg *types.MsgCreateObject) (*types.MsgCreateObjectResponse, error) {
	// Get Properties
	ctx := sdk.UnwrapSDKContext(goCtx)
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
	_, found := k.GetWhatIs(ctx, did.ID)
	if found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Object already registered to this Application")
	}

	// Create Document for Object
	doc := &types.ObjectDoc{
		Label:  msg.GetLabel(),
		Did:    did.ID,
		Fields: msg.GetInitialFields(),
	}

	// Create a new Object record
	newWhatIs := types.WhatIs{
		Did:       did.ID,
		Creator:   msg.GetSession().Creator(),
		ObjectDoc: doc,
		Timestamp: time.Now().Unix(),
		IsActive:  true,
	}

	// Store the Object record
	k.SetWhatIs(ctx, newWhatIs)
	return &types.MsgCreateObjectResponse{
		WhatIs:  &newWhatIs,
		Code:    100,
		Message: fmt.Sprintf("New Object %s has been created", newWhatIs.Did),
	}, nil
}
