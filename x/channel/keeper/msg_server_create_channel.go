package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/x/channel/types"
)

func (k msgServer) CreateChannel(goCtx context.Context, msg *types.MsgCreateChannel) (*types.MsgCreateChannelResponse, error) {
	// Get Properties
	ctx := sdk.UnwrapSDKContext(goCtx)
	// Generate a new channel Did
	did, err := did.ParseDID(msg.Creator)
	if err != nil {
		return nil, err
	}

	// Check if Channel exists
	_, found := k.GetHowIs(ctx, did.ID)
	if found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Channel already registered to this Application")
	}

	// Create new Document for Channel
	doc := &types.ChannelDoc{
		Label:       msg.GetLabel(),
		Did:         did.ID,
		Description: msg.GetDescription(),
		// RegisteredObject: msg.GetObjectToRegister(),
	}

	// Create a new channel record
	newHowis := types.HowIs{
		Did:       did.ID,
		Creator:   msg.GetCreator(),
		Channel:   doc,
		Timestamp: time.Now().Unix(),
		IsActive:  true,
	}

	// Store the channel record
	k.SetHowIs(ctx, newHowis)
	return &types.MsgCreateChannelResponse{
		Code:    100,
		Message: fmt.Sprintf("New Channel %s has been created", newHowis.Did),
		HowIs:   &newHowis,
	}, nil
}
