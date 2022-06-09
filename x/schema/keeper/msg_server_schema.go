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
	k.Logger(ctx).Info("Querying for signers")
	err := msg.ValidateBasic()
	if err != nil {
		return nil, err
	}

	accts := msg.GetSigners()
	if len(accts) < 1 {
		return nil, sdkerrors.ErrNotFound
	}

	did, err := did.ParseDID(msg.Creator)
	types.

	if err != nil {
		return nil, err
	}

	var schema = types.Schema{
		Label: msg.Label,
		did.Did: did.ID,
		Fields: msg.Fields,
	}

	k.SetSchema(ctx, schema)
	return schema, nil
}

func (k msgServer) DeprecateSchema(goCtx context.Context, msg *types.MsgDeprecateSchema) (*types.MsgDeprecateSchemaResponse, error) {
	// TODO: implement
	return nil, errors.New("unimplemented")
}
