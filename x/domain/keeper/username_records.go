package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonrhq/core/x/domain/types"
)

// SetUsernameRecords set a specific UsernameRecord in the store from its index
func (k Keeper) SetUsernameRecords(ctx sdk.Context, UsernameRecord types.UsernameRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UsernameRecordsKeyPrefix))
	b := k.cdc.MustMarshal(&UsernameRecord)
	store.Set(types.UsernameRecordsKey(
		UsernameRecord.Index,
	), b)
}

// GetUsernameRecords returns a UsernameRecord from its index
func (k Keeper) GetUsernameRecords(
	ctx sdk.Context,
	index string,

) (val types.UsernameRecord, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UsernameRecordsKeyPrefix))

	b := store.Get(types.UsernameRecordsKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveUsernameRecords removes a UsernameRecord from the store
func (k Keeper) RemoveUsernameRecords(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UsernameRecordsKeyPrefix))
	store.Delete(types.UsernameRecordsKey(
		index,
	))
}

// GetAllUsernameRecords returns all UsernameRecord
func (k Keeper) GetAllUsernameRecords(ctx sdk.Context) (list []types.UsernameRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UsernameRecordsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.UsernameRecord
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
