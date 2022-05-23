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
	ownerWhoIs, ownerFound := k.FindWhoIsByAlias(ctx, msg.GetAlias())
	if !ownerFound {
		return nil, sdkerrors.Wrap(types.ErrAliasUnavailable, "Name does not exist")
	}

	_, alias, err := ownerWhoIs.FindAliasByName(msg.GetAlias())
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrAliasUnavailable, "Name does not exist")
	}

	// Get buyerWhoIs from Owner
	// TODO: Implement Multisig for root level owner #322
	buyerWhoIs, buyerFound := k.GetWhoIsFromOwner(ctx, msg.Creator)
	if !buyerFound {
		return nil, sdkerrors.Wrapf(types.ErrControllerNotFound, "creator %s", msg.Creator)
	}

	// Convert Alias Owner address strings to sdk.AccAddress
	ownerAddr, err := ownerWhoIs.OwnerAccAddress()
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidLengthWhoIs, err.Error())
	}

	// Get Buyer address from WhoIs
	buyerAddr, err := buyerWhoIs.OwnerAccAddress()
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidLengthWhoIs, err.Error())
	}

	//TODO: put this in escrow to mitigate transfer/alias race condition attacks

	// Send Coins to new owner
	err = k.bankKeeper.SendCoins(ctx, buyerAddr, ownerAddr, sdk.NewCoins(sdk.NewCoin("snr", sdk.NewInt(int64(msg.GetAmount())))))
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to send coins to %s", msg.GetRecipient())
	}

	// Update Alias Owner
	newOwnerWhois, err := buyerWhoIs.AddAlsoKnownAs(alias.GetName(), true)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidLengthWhoIs, err.Error())
	}
	k.SetWhoIs(ctx, newOwnerWhois)

	// Remove Alias from old owner
	oldOwnerWhois, err := ownerWhoIs.RemoveAlsoKnownAs(alias.GetName(), true)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidLengthWhoIs, err.Error())
	}
	k.SetWhoIs(ctx, oldOwnerWhois)

	// Update WhoIs in keeper store
	return &types.MsgTransferAliasResponse{
		Success: true,
		WhoIs:   &buyerWhoIs,
	}, nil

}
