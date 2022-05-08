package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/x/registry/types"
)

func (k msgServer) BuyNameAlias(goCtx context.Context, msg *types.MsgBuyNameAlias) (*types.MsgBuyNameAliasResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgBuyNameAliasResponse{}, nil
}
