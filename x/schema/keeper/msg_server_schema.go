package keeper

import (
	"context"
	"fmt"
	"net/http"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/x/schema/types"
)

func (k msgServer) CreateSchema(goCtx context.Context, msg *types.MsgCreateSchema) (*types.MsgCreateSchemaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	what_is_did, err := did.ParseDID(fmt.Sprintf("did:snr:%s", k.GenerateKeyForDID()))
	if err != nil {
		return nil, sdkerrors.Wrap(err, "error while creating did from cid")
	}
	k.Logger(ctx).Info(fmt.Sprintf("Creating schema with did %s", what_is_did))

	metadata := make(map[string]string)

	for _, m := range msg.Metadata {
		metadata[m.Key] = m.Value
	}

	var whatIs = types.WhatIs{
		Creator: msg.Creator,
		Did:     what_is_did.String(),
		Schema: &types.SchemaDefinition{
			Creator: msg.Creator,
			Did:     what_is_did.String(),
			Label:   msg.Label,
			Fields:  msg.Fields,
		},
		Timestamp: time.Now().Unix(),
		IsActive:  true,
		Metadata:  metadata,
	}

	k.SetWhatIs(ctx, whatIs)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.AttributeKeyCreator, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyDID, what_is_did.String()),
			sdk.NewAttribute(types.AttributeKeyLabel, msg.Label),
			sdk.NewAttribute(types.AttributeKeyTxType, types.EventTypeCreateSchema),
		),
	)

	resp := types.MsgCreateSchemaResponse{
		Code:    http.StatusAccepted,
		Message: "Schema Registered Sucessfully",
		WhatIs:  &whatIs,
	}

	return &resp, nil
}

func (k msgServer) DeprecateSchema(goCtx context.Context, msg *types.MsgDeprecateSchema) (*types.MsgDeprecateSchemaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	schemas, found := k.GetWhatIsFromCreator(ctx, msg.GetCreator())
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "No Schemas found under same creator as message creator.")
	}

	var what_is *types.WhatIs
	var foundSchemaWI bool
	for _, a := range schemas {
		if a.GetDid() == msg.GetDid() {
			what_is = a
			foundSchemaWI = true
			break
		}
	}

	if !foundSchemaWI {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "No Schema with same creator as message creator found.")
	}

	//If already deactivated, do nothing.
	//Responsibility of caller to check if isActive beforehand
	if what_is.GetIsActive() {
		what_is.IsActive = false
		k.SetWhatIs(ctx, *what_is)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.AttributeKeyCreator, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyDID, msg.GetDid()),
			sdk.NewAttribute(types.AttributeKeyLabel, what_is.Schema.Label),
			sdk.NewAttribute(types.AttributeKeyTxType, types.EventTypeDeprecateSchema),
		),
	)

	return &types.MsgDeprecateSchemaResponse{
		Code:    200,
		Message: "Schema deprecated successfully.",
	}, nil
}
