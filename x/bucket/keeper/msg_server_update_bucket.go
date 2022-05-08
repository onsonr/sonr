package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/x/bucket/types"
)

func (k msgServer) UpdateBucket(goCtx context.Context, msg *types.MsgUpdateBucket) (*types.MsgUpdateBucketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// Generate a new channel Did
	did, err := did.ParseDID(msg.Creator)
	if err != nil {
		return nil, err
	}

	// Check if Object exists
	whichis, found := k.GetWhichIs(ctx, did.ID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Object was not found in this Application")
	}

	// Check if Bucket is IsActive
	if !whichis.IsActive {
		return nil, types.ErrInactiveBucket
	}

	// Create New Field Map
	whichis.Bucket.AddObjects(msg.GetAddedObjectDids()...)
	whichis.Bucket.RemoveObjects(msg.GetRemovedObjectDids()...)
	whichis.Timestamp = time.Now().Unix()
	whichis.IsActive = true
	k.SetWhichIs(ctx, whichis)

	return &types.MsgUpdateBucketResponse{
		Code:    100,
		Message: fmt.Sprintf("Existing Bucket %s has been updated", whichis.Did),
		WhichIs: &whichis,
	}, nil
}
