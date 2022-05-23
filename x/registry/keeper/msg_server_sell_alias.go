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
	whoIs, found := k.GetWhoIsFromOwner(ctx, msg.GetCreator())
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrAliasNotFound, "Alias %s not found", msg.GetAlias())
	}

	// Find associated Alias in whoIs
	newWhoIs, err := whoIs.UpdateAlias(msg.GetAlias(), int(msg.GetAmount()), true)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "Failed to update alias")
	}

	// Update WhoIs in Keeper store
	k.SetWhoIs(ctx, newWhoIs)
	return &types.MsgSellAliasResponse{
		Success: true,
		WhoIs:   &whoIs,
	}, nil
}
