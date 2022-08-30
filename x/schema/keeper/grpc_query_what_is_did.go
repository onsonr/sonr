package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	st "github.com/sonr-io/sonr/x/schema/types"
)

func (k Keeper) WhatIsByDid(goCtx context.Context, req *st.QueryWhatIsByDidRequest) (*st.QueryWhatIsByDidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	ref, found := k.GetWhatIsFromDid(ctx, req.Did)
	if !found || ref == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Schema was not found")
	}

	return &st.QueryWhatIsByDidResponse{
		WhatIs: ref,
	}, nil
}
