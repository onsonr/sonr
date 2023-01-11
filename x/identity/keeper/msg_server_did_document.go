package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-hq/sonr/x/identity/types"
)

func (k msgServer) CreateDidDocument(goCtx context.Context, msg *types.MsgCreateDidDocument) (*types.MsgCreateDidDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// Check if the value already exists
	_, isFound := k.GetDidDocument(
		ctx,
		msg.Did,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var didDocument = types.DidDocument{
		ID:                   types.ConvertAccAddressToDid(msg.Creator),
		Context:              []string{msg.Context},
		Controller:           []string{msg.Controller},
		VerificationMethod:   new(types.VerificationMethods),
		Authentication:       new(types.VerificationRelationships),
		AssertionMethod:      new(types.VerificationRelationships),
		CapabilityInvocation: new(types.VerificationRelationships),
		CapabilityDelegation: new(types.VerificationRelationships),
		KeyAgreement:         new(types.VerificationRelationships),
		Service:              new(types.Services),
		AlsoKnownAs:          []string{msg.AlsoKnownAs},
	}
	acc := k.accountKeeper.NewAccountWithAddress(ctx, sdk.AccAddress(msg.Creator))
	k.accountKeeper.SetAccount(ctx, acc)
	k.SetDidDocument(
		ctx,
		didDocument,
	)
	return &types.MsgCreateDidDocumentResponse{}, nil
}

func (k msgServer) UpdateDidDocument(goCtx context.Context, msg *types.MsgUpdateDidDocument) (*types.MsgUpdateDidDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetDidDocument(
		ctx,
		msg.Did,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if valFound.CheckAccAddress(msg.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var didDocument = types.DidDocument{
		ID:                   types.ConvertAccAddressToDid(msg.Creator),
		Context:              []string{msg.Context},
		Controller:           []string{msg.Controller},
		VerificationMethod:   new(types.VerificationMethods),
		Authentication:       new(types.VerificationRelationships),
		AssertionMethod:      new(types.VerificationRelationships),
		CapabilityInvocation: new(types.VerificationRelationships),
		CapabilityDelegation: new(types.VerificationRelationships),
		KeyAgreement:         new(types.VerificationRelationships),
		Service:              new(types.Services),
		AlsoKnownAs:          []string{msg.AlsoKnownAs},
	}
	k.SetDidDocument(ctx, didDocument)

	return &types.MsgUpdateDidDocumentResponse{}, nil
}

func (k msgServer) DeleteDidDocument(goCtx context.Context, msg *types.MsgDeleteDidDocument) (*types.MsgDeleteDidDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetDidDocument(
		ctx,
		msg.Did,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if valFound.CheckAccAddress(msg.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveDidDocument(
		ctx,
		msg.Did,
	)

	return &types.MsgDeleteDidDocumentResponse{}, nil
}
