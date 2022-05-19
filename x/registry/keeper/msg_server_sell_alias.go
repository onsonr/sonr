package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/sonr/x/registry/types"
)

func (k msgServer) SellAlias(goCtx context.Context, msg *types.MsgSellAlias) (*types.MsgSellAliasResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if Alias exists
	_, aliasIsFound := k.GetWhoIsFromAlias(ctx, msg.GetAlias())
	// If a name is found in store
	if aliasIsFound {
		return nil, sdkerrors.Wrap(types.ErrAliasUnavailable, "Name already has an owner")
	}

	// Get whois from controller
	// TODO: Implement Multisig for root level owner
	whois, isFound := k.GetWhoIsFromOwner(ctx, msg.Creator)
	if !isFound {
		return nil, sdkerrors.Wrapf(types.ErrControllerNotFound, "creator %s", msg.Creator)
	}

	// Find associated Alias in whoIs
	foundAlias := false
	for _, a := range whois.Alias {
		if a.GetName() == msg.GetAlias() {
			a.IsForSale = true
			a.Amount = msg.GetAmount()
			foundAlias = true
		}
	}

	// If alias not found
	if !foundAlias {
		return nil, sdkerrors.Wrapf(types.ErrAliasNotFound, "Alias %s not found", msg.GetAlias())
	}

	// Update WhoIs in Keeper store
	k.SetWhoIs(ctx, whois)
	return &types.MsgSellAliasResponse{
		Success: true,
		WhoIs:   &whois,
	}, nil
}
