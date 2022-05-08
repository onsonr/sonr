package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/x/registry/types"
)

func (k msgServer) BuyAppAlias(goCtx context.Context, msg *types.MsgBuyAppAlias) (*types.MsgBuyAppAliasResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgBuyAppAliasResponse{}, nil
}
