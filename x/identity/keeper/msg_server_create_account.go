package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-hq/sonr/x/identity/types"
)

func (k msgServer) CreateAccount(goCtx context.Context, msg *types.MsgCreateAccount) (*types.MsgCreateAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// Check if the value already exists
	addr := msg.AssertionMethod.Address()
	_, isFound := k.GetDidDocument(
		ctx,
		addr,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}
	doc := types.BlankDocument(addr)
	doc.AddAssertion(msg.AssertionMethod)
	doc.AddAuthentication(msg.AssertionMethod)
	doc.AddCapabilityDelegation(msg.AssertionMethod)
	doc.AddService(msg.Service)

	// Create the account
	acc := k.accountKeeper.NewAccountWithAddress(ctx, sdk.AccAddress(addr))
	k.accountKeeper.SetAccount(ctx, acc)
	k.SetDidDocument(
		ctx,
		*doc,
	)
	return &types.MsgCreateAccountResponse{}, nil
}
