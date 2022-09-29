package keeper

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/sonr/pkg/did"
	rtv1 "github.com/sonr-io/sonr/x/registry/types"
	"github.com/sonr-io/sonr/x/schema/types"
)

const (
	serviceId = "sonr:what_is"
)

func (k msgServer) CreateSchema(goCtx context.Context, msg *types.MsgCreateSchema) (*types.MsgCreateSchemaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := msg.ValidateBasic()
	if err != nil {
		return nil, err
	}

	k.Logger(ctx).Info("msg validation successful")

	accts := msg.GetSigners()
	if len(accts) < 1 {
		return nil, sdkerrors.ErrNotFound
	}

	creator_did := msg.GetCreatorDid()
	k.Logger(ctx).Info(fmt.Sprintf("Creating schema for creator did %s", creator_did))

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, sdkerrors.Wrapf(err, "Error while pinning schema definition to storage")
	}

	if err != nil {
		return nil, sdkerrors.Wrapf(err, "Error while persisting schema fields")
	}

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
		Creator: creator_did,
		Did:     what_is_did.String(),
		Schema: &types.SchemaDefinition{
			Did:    what_is_did.String(),
			Label:  msg.Label,
			Fields: msg.Fields,
		},
		Timestamp: time.Now().Unix(),
		IsActive:  true,
		Metadata:  metadata,
	}

	k.SetWhatIs(ctx, whatIs)

	who_is, found := k.registryKeeper.GetWhoIs(ctx, creator_did)

	if !found {
		return nil, errors.New(fmt.Sprintf("error while querying who is for creator %s", creator_did))
	}

	who_is.DidDocument.Service = append(who_is.DidDocument.Service, &rtv1.Service{
		Id:   fmt.Sprintf("%s#%s", who_is.DidDocument.Id, whatIs.Schema.Label),
		Type: serviceId,
		ServiceEndpoint: []*rtv1.KeyValuePair{
			{
				Key:   "did",
				Value: who_is.DidDocument.Id,
			},
		},
	})

	k.registryKeeper.SetWhoIs(ctx, who_is)
	resp := types.MsgCreateSchemaResponse{
		Code:    http.StatusAccepted,
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

	who_is, found := k.registryKeeper.GetWhoIs(ctx, what_is.Creator)

	if !found {
		return nil, errors.New(fmt.Sprintf("error while querying who is for creator %s", what_is.Creator))
	}

	//If already deactivated, do nothing.
	//Responsibility of caller to check if isActive beforehand
	if what_is.GetIsActive() {
		what_is.IsActive = false
		k.SetWhatIs(ctx, *what_is)
	}

	var serviceIndex int = -1
	for i, s := range who_is.DidDocument.Service {
		if s.Id == fmt.Sprintf("%s#%s", who_is.DidDocument.Id, what_is.Schema.Label) {
			serviceIndex = i
		}
	}

	if serviceIndex < -1 {
		who_is.DidDocument.Service = append(who_is.DidDocument.Service[:serviceIndex], who_is.DidDocument.Service[serviceIndex+1:]...)
		k.registryKeeper.SetWhoIs(ctx, who_is)
	}

	return &types.MsgDeprecateSchemaResponse{
		Code:    200,
		Message: "Schema deprecated successfully.",
	}, nil
}
