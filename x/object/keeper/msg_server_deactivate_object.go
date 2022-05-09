package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/sonr/internal/blockchain/x/object/types"
)

func (k msgServer) DeactivateObject(goCtx context.Context, msg *types.MsgDeactivateObject) (*types.MsgDeactivateObjectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if Object exists
	whatis, found := k.GetWhatIs(ctx, msg.GetDid())
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Object was not found in this Application")
	}
	whatis.IsActive = false
	k.SetWhatIs(ctx, whatis)
	return &types.MsgDeactivateObjectResponse{
		Code:    100,
		Message: fmt.Sprintf("Object %s has been deactivated", whatis.Did),
	}, nil
}
