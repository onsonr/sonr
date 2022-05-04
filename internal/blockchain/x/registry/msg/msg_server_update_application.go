package msg

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/internal/blockchain/x/registry/types"
)

func (k msgServer) UpdateApplication(goCtx context.Context, msg *types.MsgUpdateApplication) (*types.MsgUpdateApplicationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUpdateApplicationResponse{}, nil
}
