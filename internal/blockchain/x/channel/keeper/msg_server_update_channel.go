package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/sonr/internal/blockchain/x/channel/types"
)

func (k msgServer) UpdateChannel(goCtx context.Context, msg *types.MsgUpdateChannel) (*types.MsgUpdateChannelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if Object exists
	howis, found := k.GetHowIs(ctx, msg.GetDid())
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Channel was not found in this Application")
	}

	// Check if Channel is IsActive
	if !howis.IsActive {
		return nil, types.ErrInactiveChannel
	}

	// Replace Channel Properties with new ones
	howis.Channel.Description = msg.GetDescription()
	howis.Channel.Label = msg.GetLabel()
	howis.Channel.RegisteredObject = msg.GetObjectToRegister()

	return &types.MsgUpdateChannelResponse{
		Code:    100,
		Message: fmt.Sprintf("Existing Channel %s has been updated", howis.Did),
	}, nil
}
