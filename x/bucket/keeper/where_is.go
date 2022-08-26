package keeper

import (
	"encoding/binary"
	"strings"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/uuid"
	"github.com/sonr-io/sonr/x/bucket/types"
)

// GetWhereIsCount get the total number of whereIs
func (k Keeper) GetWhereIsCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.WhereIsCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetWhereIsCount set the total number of whereIs
func (k Keeper) SetWhereIsCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.WhereIsCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendWhereIs appends a whereIs in the store with a new id and update the count
func (k Keeper) AppendWhereIs(
	ctx sdk.Context,
	whereIs types.WhereIs,
) string {
	// Create the whereIs
	count := k.GetWhereIsCount(ctx)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhereIsKeyPrefix))
	appendedValue := k.cdc.MustMarshal(&whereIs)
	store.Set(types.WhereIsKey(whereIs.Did), appendedValue)

	// Update whereIs count
	k.SetWhereIsCount(ctx, count+1)

	return whereIs.Did
}

// SetWhereIs set a specific whereIs in the store
func (k Keeper) SetWhereIs(ctx sdk.Context, whereIs types.WhereIs) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhereIsKeyPrefix))
	b := k.cdc.MustMarshal(&whereIs)
	store.Set(types.WhereIsKey(whereIs.Did), b)
}

// GetWhereIs returns a whereIs from its id
func (k Keeper) GetWhereIs(ctx sdk.Context, creator, id string) (val types.WhereIs, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhereIsKeyPrefix))
	b := store.Get(
		types.WhereIsKey(
			id,
		),
	)

	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	if val.Visibility == types.BucketVisibility_PUBLIC || val.Creator == creator {
		return val, true
	}

	return val, false
}

// GetWhereIs returns a whereIs from its id
func (k Keeper) GetWhereIsByCreator(ctx sdk.Context, creator string) (list []types.WhereIs) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhereIsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.WhereIs
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.Creator == creator {
			list = append(list, val)
		}
	}

	return
}

// RemoveWhereIs removes a whereIs from the store
func (k Keeper) RemoveWhereIs(ctx sdk.Context, id string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhereIsKeyPrefix))
	store.Delete(types.WhereIsKey(id))
}

// GetAllWhereIs returns all whereIs
func (k Keeper) GetAllWhereIs(ctx sdk.Context) (list []types.WhereIs) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhereIsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.WhereIs
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.Visibility == types.BucketVisibility_PUBLIC || val.Visibility == types.BucketVisibility_UNSPECIFIED {
			list = append(list, val)
		}
	}

	return
}

func (k Keeper) GenerateKeyForDID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
