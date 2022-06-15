package keeper

import (
	"context"
	"errors"
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

	creator_did, err := did.NewDocument(msg.GetCreatorDid())

	if err != nil {
		return nil, err
	}

	guid := k.GenerateKeyForDID()

	if err != nil {
		return nil, err
	}

	addr := string(guid)
	what_is_did, err := did.ParseDID(fmt.Sprintf("did:snr:%s", addr))

	if err != nil {
		return nil, err
	}

	var schema = types.SchemaReference{
		Label: msg.Label,
		Did:   what_is_did.String(),
		Cid:   msg.Cid,
	}

	var whatIs = types.WhatIs{
		Creator:   creator_did.GetID().String(),
		Did:       what_is_did.String(),
		Schema:    &schema,
		Timestamp: time.Now().Unix(),
		IsActive:  true,
	}

	k.SetWhatIs(ctx, whatIs)

	resp := types.MsgCreateSchemaResponse{
		Code:    200,
		Message: "Schema Created Successfully",
		WhatIs:  &whatIs,
	}

	return &resp, nil
}

func (k msgServer) DeprecateSchema(goCtx context.Context, msg *types.MsgDeprecateSchema) (*types.MsgDeprecateSchemaResponse, error) {
	// TODO: implement
	return nil, errors.New("unimplemented")
}
