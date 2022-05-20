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
	whoIs, aliasIsFound := k.GetWhoIsFromAlias(ctx, msg.GetAlias())
	// If a name is found in store
	if aliasIsFound {
		return nil, sdkerrors.Wrap(types.ErrAliasUnavailable, "Name already has an owner")
	}

	// If alias not found
	if !aliasIsFound {
		return nil, sdkerrors.Wrapf(types.ErrAliasNotFound, "Alias %s not found", msg.GetAlias())
	}

	// Find associated Alias in whoIs
	for _, a := range whoIs.Alias {
		if a.GetName() == msg.GetAlias() {
			a.IsForSale = true
			a.Amount = msg.GetAmount()
		}
	}

	// Update WhoIs in Keeper store
	k.SetWhoIs(ctx, whoIs)
	return &types.MsgSellAliasResponse{
		Success: true,
		WhoIs:   &whoIs,
	}, nil
}
