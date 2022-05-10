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
	doc := did.Document{}
	err := doc.UnmarshalJSON(msg.DidDocument)
	if err != nil {
		return nil, err
	}

	var whoIs = types.WhoIs{
		Owner:       msg.Creator,
		DidDocument: msg.DidDocument,
		Type:        msg.WhoisType,
		Alias:       doc.AlsoKnownAs,
		Controllers: doc.ControllersAsString(),
		IsActive:    true,
		Timestamp:   time.Now().Unix(),
	}

	k.SetWhoIs(ctx, whoIs)

	return &types.MsgCreateWhoIsResponse{
		WhoIs: &whoIs,
	}, nil
}

// UpdateWhoIs updates a whoIs from the store
func (k msgServer) UpdateWhoIs(goCtx context.Context, msg *types.MsgUpdateWhoIs) (*types.MsgUpdateWhoIsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// UnmarshalJSON from DID document
	doc := did.Document{}
	err := doc.UnmarshalJSON(msg.DidDocument)
	if err != nil {
		return nil, err
	}
	// Checks that the element exists
	val, found := k.GetWhoIs(ctx, msg.Did)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Did))
	}

	// Checks if the msg owner is the same as the current owner
	if msg.Creator != val.Owner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	err = val.CopyFromDidDocument(&doc)
	if err != nil {
		return nil, err
	}
	k.SetWhoIs(ctx, val)
	return &types.MsgUpdateWhoIsResponse{}, nil
}

// DeleteWhoIs deletes a whoIs from the store
func (k msgServer) DeleteWhoIs(goCtx context.Context, msg *types.MsgDeactivateWhoIs) (*types.MsgDeactivateWhoIsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetWhoIs(ctx, msg.Did)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Did))
	}

	// Checks if the msg owner is the same as the current owner
	if msg.Creator != val.Owner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	// Deactivates the element
	val.IsActive = false
	k.SetWhoIs(ctx, val)
	return &types.MsgDeactivateWhoIsResponse{}, nil
}
