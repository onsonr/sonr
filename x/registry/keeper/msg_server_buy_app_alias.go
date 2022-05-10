package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/x/registry/types"
)

func (k msgServer) BuyAppAlias(goCtx context.Context, msg *types.MsgBuyAppAlias) (*types.MsgBuyAppAliasResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if Alias exists
	_, aliasIsFound := k.GetWhoIsFromAlias(ctx, msg.Name)
	// If a name is found in store
	if aliasIsFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Name already has an owner")
	}

	// Get whois from controller
	whois, isFound := k.GetWhoIsFromController(ctx, msg.Creator)
	if !isFound {
		return nil, sdkerrors.Wrapf(types.ErrControllerNotFound, "creator %s", msg.Creator)
	}

	// Unmarshal DIDDoc from whois
	doc := did.Document{}
	err := doc.UnmarshalJSON(whois.GetDidDocument())
	if err != nil {
		return nil, sdkerrors.Wrap(err, "Failed to unmarshal DID document")
	}

	// Convert owner and buyer address strings to sdk.AccAddress
	buyer, _ := sdk.AccAddressFromBech32(msg.Creator)
	minPrice := sdk.Coins{sdk.NewInt64Coin("snr", 10)}
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, buyer, types.ModuleName, minPrice)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Not enough snr coins to buy name")
	}

	// Create an updated whois record
	doc.AddAlias(msg.Name)
	whois.CopyFromDidDocument(&doc)
	k.SetWhoIs(ctx, whois)

	// Return response
	return &types.MsgBuyAppAliasResponse{
		Did:   whois.Owner,
		WhoIs: &whois,
	}, nil
}
