package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/x/registry/types"
)

func (k msgServer) TransferNameAlias(goCtx context.Context, msg *types.MsgTransferNameAlias) (*types.MsgTransferNameAliasResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgTransferNameAliasResponse{}, nil
}
