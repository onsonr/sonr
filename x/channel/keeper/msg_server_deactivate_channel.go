package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/sonr/internal/blockchain/x/channel/types"
)

func (k msgServer) DeactivateChannel(goCtx context.Context, msg *types.MsgDeactivateChannel) (*types.MsgDeactivateChannelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if Channel exists
	howis, found := k.GetHowIs(ctx, msg.GetDid())
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Channel was not found in this Application")
	}
	howis.IsActive = false
	k.SetHowIs(ctx, howis)
	return &types.MsgDeactivateChannelResponse{
		Code:    100,
		Message: fmt.Sprintf("Channel %s has been deactivated", howis.Did),
	}, nil
}
