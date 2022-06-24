package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/sonr/x/schema/types"
)

func (k Keeper) QuerySchema(goCtx context.Context, req *types.QuerySchemaRequest) (*types.QuerySchemaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	ref, found := k.GetWhatIsFromCreator(ctx, req.Creator)
	if !found || len(ref) < 1 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Schema was not found")
	}

	for _, item := range ref {
		if item.Did == req.Did {
			var resp = types.QuerySchemaResponse{
				Schema: &item,
			}

			return &resp, nil
		}
	}

	return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Schema was not found")
}
