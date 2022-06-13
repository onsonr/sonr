package keeper

import (
	"context"
	"errors"

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

	did, err := did.ParseDID(msg.GetCreatorDid())

	if err != nil {
		return nil, err
	}

	var schema = types.Schema{
		Label:  msg.Label,
		Did:    did.ID,
		Fields: msg.Fields,
	}

	k.SetSchema(ctx, schema)

	resp := types.MsgCreateSchemaResponse{
		Code:    200,
		Message: "Schema Created Successfully",
		Schema:  &schema,
	}

	return &resp, nil
}

func (k msgServer) DeprecateSchema(goCtx context.Context, msg *types.MsgDeprecateSchema) (*types.MsgDeprecateSchemaResponse, error) {
	// TODO: implement
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := msg.ValidateBasic()
	if err != nil {
		return nil, err
	}

	/*oldSchema*/
	oldSchema, found := k.GetSchema(ctx, msg.GetDid()) //TODO: Get ID of schema

	if found {

		//Creator guard
		creatorDid, creatorErr := did.ParseDID(msg.GetDid())
		oldSchemaDid, oldErr := did.ParseDID(oldSchema.GetDid())
		if creatorErr != nil {
			return nil, creatorErr //TODO: Create response
		} else if oldErr != nil {
			return nil, oldErr //TODO: Create response
		} else if creatorDid != oldSchemaDid {
			return nil, errors.New("Permission Denied: Only the owner may deprecate a schema") //TODO: Create response, error
		}

		//If already deactivated, do nothing.
		//Responsibility of caller to check if isActive beforehand
		if oldSchema.IsActive {
			oldSchema.IsActive = false
			k.SetSchema(ctx, oldSchema)
		}

		return nil, nil //TODO: Create response
	} else {
		return nil, errors.New("Schema not found") //TODO: Create response, error
	}
}
