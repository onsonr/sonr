package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/x/registry/types"
)

func (k msgServer) SellAlias(goCtx context.Context, msg *types.MsgSellAlias) (*types.MsgSellAliasResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSellAliasResponse{}, nil
}
