package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/x/registry/types"
)

func (k msgServer) TransferAlias(goCtx context.Context, msg *types.MsgTransferAlias) (*types.MsgTransferAliasResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgTransferAliasResponse{}, nil
}
