package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/x/object/types"
)

// SetWhatIs set a specific whatIs in the store from its did
func (k Keeper) SetWhatIs(ctx sdk.Context, whatIs types.WhatIs) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhatIsKeyPrefix))
	b := k.cdc.MustMarshal(&whatIs)
	store.Set(types.WhatIsKey(
		whatIs.Did,
	), b)
}

// GetWhatIs returns a whatIs from its did
func (k Keeper) GetWhatIs(
	ctx sdk.Context,
	did string,
) (val types.WhatIs, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhatIsKeyPrefix))

	b := store.Get(types.WhatIsKey(
		did,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveWhatIs removes a whatIs from the store
func (k Keeper) RemoveWhatIs(
	ctx sdk.Context,
	did string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhatIsKeyPrefix))
	store.Delete(types.WhatIsKey(did))
}

// GetAllWhatIs returns all whatIs
func (k Keeper) GetAllWhatIs(ctx sdk.Context) (list []types.WhatIs) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhatIsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.WhatIs
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
