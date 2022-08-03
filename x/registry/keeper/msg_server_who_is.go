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

	// Create Sonr DID Doc to store in WhoIs
	sonrDidDoc, err := types.NewDIDDocumentFromBytes(didDocBuf)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// TODO: Implement Multisig for root level owner #322
	var whoIs = types.WhoIs{
		Owner:       msg.Creator,
		DidDocument: sonrDidDoc,
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
	val, found := k.GetWhoIs(ctx, msg.GetCreator())
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.GetCreator()))
	}

	// Checks if the msg owner is the same as the current owner
	if msg.Creator != val.Owner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	// unmarshall DID Document from request using did package
	doc := did.BlankDocument()
	err := doc.UnmarshalJSON(msg.DidDocument)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	val.DidDocument, err = types.NewDIDDocumentFromPkg(doc)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, err.Error())
	}

	// Add all the aliases to the whois
	for _, a := range doc.GetAlsoKnownAs() {
		if !val.ContainsAlias(a) {
			val.AddAlsoKnownAs(a, false)
		}
	}

	// Add all the metadata to the whois
	for k, v := range msg.GetMetadata() {
		val.Metadata[k] = v
	}

	d := val.DidDocument.DID()
	if d == nil {
		return nil, fmt.Errorf("error getting did from did document for creator '%s'", msg.Creator)
	}

	val.Timestamp = time.Now().Unix()
	val.IsActive = true
	k.SetWhoIs(ctx, val)

	return &types.MsgUpdateWhoIsResponse{
		Success: true,
		WhoIs:   &val,
	}, nil
}

// DeactivateWhoIs deletes a whoIs from the store
func (k msgServer) DeactivateWhoIs(goCtx context.Context, msg *types.MsgDeactivateWhoIs) (*types.MsgDeactivateWhoIsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetWhoIs(ctx, msg.GetCreator())
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

	sonrDidDoc, err := types.NewDIDDocumentFromPkg(doc)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Deactivates the element
	val.IsActive = false
	val.Timestamp = time.Now().Unix()
	val.Alias = make([]*types.Alias, 0)
	val.DidDocument = sonrDidDoc
	k.SetWhoIs(ctx, val)
	return &types.MsgDeactivateWhoIsResponse{}, nil
}
