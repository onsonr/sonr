package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/sonr/x/registry/types"
)

func (k msgServer) TransferAlias(goCtx context.Context, msg *types.MsgTransferAlias) (*types.MsgTransferAliasResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if Alias exists if not return error
	_, aliasIsFound := k.GetWhoIsFromAlias(ctx, msg.GetAlias())
	// If a name is not found in store return error
	if !aliasIsFound {
		return nil, sdkerrors.Wrap(types.ErrAliasUnavailable, "Name does not exist")
	}

	// Get whois from Owner
	// TODO: Implement Multisig for root level owner #322
	whois, isFound := k.GetWhoIsFromOwner(ctx, msg.Creator)
	if !isFound {
		return nil, sdkerrors.Wrapf(types.ErrControllerNotFound, "creator %s", msg.Creator)
	}

	// Find associated Alias in whoIs
	idx, alias, err := whois.FindAliasByName(msg.GetAlias())
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrAliasNotFound, "alias %s", msg.GetAlias())
	}

	// Convert Alias Owner address strings to sdk.AccAddress
	aliasOwner, err := sdk.AccAddressFromBech32(msg.GetRecipient())
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "alias owner %s", msg.GetRecipient())
	}

	//TODO: put this in escrow to mitigate transfer/alias race condition attacks

	// Send Coins to new owner
	err = k.bankKeeper.SendCoins(ctx, aliasOwner, aliasOwner, sdk.NewCoins(sdk.NewCoin("snr", sdk.NewInt(int64(msg.GetAmount())))))
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to send coins to %s", msg.GetRecipient())
	}

	// Transfer Alias to new owner if transaction is successful
	msg.Recipient = msg.Creator
	whois.Alias[idx] = alias

	// Update WhoIs in keeper store
	return &types.MsgTransferAliasResponse{
		Success: true,
		WhoIs:   &whois,
	}, nil

}
