package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/sonr/x/registry/types"
)

func (k msgServer) TransferAlias(goCtx context.Context, msg *types.MsgTransferAlias) (*types.MsgTransferAliasResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if Alias exists
	_, aliasIsFound := k.GetWhoIsFromAlias(ctx, msg.GetAlias())
	// If a name is found in store
	if aliasIsFound {
		return nil, sdkerrors.Wrap(types.ErrAliasUnavailable, "Name already has an owner")
	}

	// Get whois from controller
	whois, isFound := k.GetWhoIsFromController(ctx, msg.Creator)
	if !isFound {
		return nil, sdkerrors.Wrapf(types.ErrControllerNotFound, "creator %s", msg.Creator)
	}

	// Check if alias is sold
	idx, alias, err := whois.FindAliasByName(msg.GetAlias())
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrAliasNotFound, "alias %s", msg.GetAlias())
	}

	// Transfer alias from whois to new owner
	msg.Recipient = msg.Creator
	whois.Alias[idx] = alias

	// Update WhoIs in Keeper store
	k.SetWhoIs(ctx, whois)
	return &types.MsgTransferAliasResponse{
		Success: true,
		WhoIs:   &whois,
	}, nil

}
