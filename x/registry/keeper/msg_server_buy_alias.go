package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/sonr/x/registry/types"
)

// TODO(https://github.com/sonr-io/sonr/issues/333): Make default price configurable here and elsewhere
func (k msgServer) BuyAlias(goCtx context.Context, msg *types.MsgBuyAlias) (*types.MsgBuyAliasResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := msg.ValidateAlias(); err != nil {
		return nil, err
	}

	// Check if Alias exists
	_, isAliasOwned := k.FindWhoIsByAlias(ctx, msg.Name)
	if isAliasOwned {
		return nil, sdkerrors.Wrap(types.ErrAliasUnavailable, "Name already has an owner")
	}

	// Get whois from Owner
	// TODO: Implement Multisig for root level owner #322
	whois, isFound := k.GetWhoIsFromOwner(ctx, msg.Creator)
	if !isFound {
		return nil, sdkerrors.Wrapf(types.ErrControllerNotFound, "creator %s", msg.Creator)
	}

	// Convert owner and buyer address strings to sdk.AccAddress
	ownerAddr, err := whois.OwnerAccAddress()
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidLengthWhoIs, err.Error())
	}

	// Send Coins to the Registry Module
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, ownerAddr, types.ModuleName, types.MIN_BUY_ALIAS)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "Not enough snr coins to buy name")
	}

	// Create an updated whois record
	newWhoIs, err := whois.AddAlsoKnownAs(msg.GetName(), true)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "Failed to add alias")
	}

	k.SetWhoIs(ctx, newWhoIs)

	// Return response
	return &types.MsgBuyAliasResponse{
		Success: true,
		WhoIs:   &whois,
	}, nil
}
