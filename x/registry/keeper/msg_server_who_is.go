package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/x/registry/types"
)

// CreateWhoIs creates a whoIs from the store
func (k msgServer) CreateWhoIs(goCtx context.Context, msg *types.MsgCreateWhoIs) (*types.MsgCreateWhoIsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// UnmarshalJSON from DID document
	doc, err := did.NewDocument(msg.GetCreatorDid())
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Copy buffer to the document
	err = doc.CopyFromBytes(msg.DidDocument)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Create the new buffer
	didDocBuf, err := doc.MarshalJSON()
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// TODO: Implement Multisig for root level owner #322
	var whoIs = types.WhoIs{
		Owner:       msg.Creator,
		DidDocument: didDocBuf,
		Type:        msg.WhoisType,
		Controllers: doc.ControllersAsString(),
		IsActive:    true,
		Timestamp:   time.Now().Unix(),
		Alias:       make([]*types.Alias, 0),
	}

	// Add the also known as to the whois
	k.SetWhoIs(ctx, whoIs)
	return &types.MsgCreateWhoIsResponse{
		WhoIs: &whoIs,
	}, nil
}

// UpdateWhoIs updates a whoIs from the store
func (k msgServer) UpdateWhoIs(goCtx context.Context, msg *types.MsgUpdateWhoIs) (*types.MsgUpdateWhoIsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetWhoIsFromOwner(ctx, msg.GetCreator())
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.GetCreator()))
	}

	// Checks if the msg owner is the same as the current owner
	if msg.Creator != val.Owner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	newWhoIs, err := val.UpdateDidBuffer(msg.DidDocument)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	k.SetWhoIs(ctx, newWhoIs)
	return &types.MsgUpdateWhoIsResponse{}, nil
}

// DeleteWhoIs deletes a whoIs from the store
func (k msgServer) DeleteWhoIs(goCtx context.Context, msg *types.MsgDeactivateWhoIs) (*types.MsgDeactivateWhoIsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetWhoIsFromOwner(ctx, msg.GetCreator())
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.GetCreator()))
	}

	// Checks if the msg owner is the same as the current owner
	if msg.Creator != val.Owner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	doc, err := did.NewDocument(msg.GetCreatorDid())
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Deactivates the element
	val.IsActive = false
	val.Timestamp = time.Now().Unix()
	val.Alias = make([]*types.Alias, 0)
	val.DidDocument, err = doc.MarshalJSON()
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	k.SetWhoIs(ctx, val)
	return &types.MsgDeactivateWhoIsResponse{}, nil
}
