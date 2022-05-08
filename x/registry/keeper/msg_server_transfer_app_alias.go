package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/x/registry/types"
)

func (k msgServer) TransferAppAlias(goCtx context.Context, msg *types.MsgTransferAppAlias) (*types.MsgTransferAppAliasResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgTransferAppAliasResponse{}, nil
}
