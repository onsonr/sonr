package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/x/channel/types"
)

// SetHowIs set a specific howIs in the store from its did
func (k Keeper) SetHowIs(ctx sdk.Context, howIs types.HowIs) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HowIsKeyPrefix))
	b := k.cdc.MustMarshal(&howIs)
	store.Set(types.HowIsKey(
		howIs.Did,
	), b)
}

// GetHowIs returns a howIs from its did
func (k Keeper) GetHowIs(
	ctx sdk.Context,
	did string,
) (val types.HowIs, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HowIsKeyPrefix))

	b := store.Get(types.HowIsKey(
		did,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveHowIs removes a howIs from the store
func (k Keeper) RemoveHowIs(
	ctx sdk.Context,
	did string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HowIsKeyPrefix))
	store.Delete(types.HowIsKey(
		did,
	))
}

// GetAllHowIs returns all howIs
func (k Keeper) GetAllHowIs(ctx sdk.Context) (list []types.HowIs) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HowIsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.HowIs
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
