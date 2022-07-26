package keeper

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/x/schema/types"
)

func (k msgServer) CreateSchema(goCtx context.Context, msg *types.MsgCreateSchema) (*types.MsgCreateSchemaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := msg.ValidateBasic()
	if err != nil {
		return nil, err
	}

	accts := msg.GetSigners()
	if len(accts) < 1 {
		return nil, sdkerrors.ErrNotFound
	}

	creator_did := msg.GetCreatorDid()

	guid := k.GenerateKeyForDID()

	if err != nil {
		return nil, err
	}

	addr := string(guid)
	what_is_did, err := did.ParseDID(fmt.Sprintf("did:snr:%s", addr))

	if err != nil {
		return nil, err
	}
	schemaDef := make(map[string]types.SchemaKind)
	for _, c := range msg.Definition.GetFields() {
		schemaDef[c.Name] = c.Field
	}
	b, err := json.Marshal(schemaDef)
	if err != nil {
		return nil, err
	}

	cidStr, err := k.ipfs.PutData(ctx.Context(), b)
	if err != nil {
		return nil, err
	}

	if err = k.ipfs.PinFile(ctx.Context(), cidStr); err != nil {
		return nil, err
	}

	if err != nil {
		return nil, sdkerrors.Wrapf(err, "Error while persisting schema fields")
	}

	var schema = types.SchemaReference{
		Label: msg.Definition.Label,
		Did:   what_is_did.String(),
		Cid:   cidStr,
	}

	var whatIs = types.WhatIs{
		Creator:   creator_did,
		Did:       what_is_did.String(),
		Schema:    &schema,
		Timestamp: time.Now().Unix(),
		IsActive:  true,
	}

	k.SetWhatIs(ctx, whatIs)

	resp := types.MsgCreateSchemaResponse{
		Code:    200,
		Message: "Schema Registered Sucessfully",
		WhatIs:  &whatIs,
	}

	return &resp, nil
}

func (k msgServer) DeprecateSchema(goCtx context.Context, msg *types.MsgDeprecateSchema) (*types.MsgDeprecateSchemaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := msg.ValidateBasic()
	if err != nil {
		return nil, err
	}

	schemas, found := k.GetWhatIsFromCreator(ctx, msg.GetCreator())
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "No Schemas found under same creator as message creator.")
	}

	var what_is types.WhatIs
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
		k.SetWhatIs(ctx, what_is)
	}

	return &types.MsgDeprecateSchemaResponse{
		Code:    200,
		Message: "Schema deprecated successfully.",
	}, nil
}
