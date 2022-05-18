package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/x/registry/types"
)

func (k msgServer) BuyAlias(goCtx context.Context, msg *types.MsgBuyAlias) (*types.MsgBuyAliasResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := types.ValidateAlias(msg.GetName()); err != nil {
		return nil, err
	}

	// Check if Alias exists
	_, aliasIsFound := k.GetWhoIsFromAlias(ctx, msg.Name)
	// If a name is found in store
	if aliasIsFound {
		return nil, sdkerrors.Wrap(types.ErrAliasUnavailable, "Name already has an owner")
	}

	// Get whois from Owner
	// TODO: Implement Multisig for root level owner
	whois, isFound := k.GetWhoIsFromOwner(ctx, msg.Creator)
	if !isFound {
		return nil, sdkerrors.Wrapf(types.ErrControllerNotFound, "creator %s", msg.Creator)
	}

	// Unmarshal DIDDoc from whois
	doc := did.Document{}
	err := doc.UnmarshalJSON(whois.GetDidDocument())
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrDidDocumentInvalid, "Failed to unmarshal DID document")
	}

	// Convert owner and buyer address strings to sdk.AccAddress
	buyer, _ := sdk.AccAddressFromBech32(msg.Creator)
	minPrice := sdk.Coins{sdk.NewInt64Coin("snr", 10)}
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, buyer, types.ModuleName, minPrice)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Not enough snr coins to buy name")
	}

	// Create an updated whois record
	doc.AddAlias(types.FormatAppAlias(msg.GetName()))
	whois.CopyFromDidDocument(&doc)
	k.SetWhoIs(ctx, whois)

	// Return response
	return &types.MsgBuyAliasResponse{
		Did:   whois.Owner,
		WhoIs: &whois,
	}, nil
}
