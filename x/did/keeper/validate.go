package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

func (k Keeper) ValidServiceOrigin(ctx sdk.Context, origin string) bool {
	rec, err := k.OrmDB.ServiceRecordTable().GetByOriginUri(ctx, origin)
	if err != nil {
		return false
	}
	if rec == nil {
		return false
	}
	return true
}
