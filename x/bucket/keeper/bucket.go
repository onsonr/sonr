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
	byteKey := types.KeyPrefix(types.BucketCountKey)
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
	byteKey := types.KeyPrefix(types.BucketCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendWhereIs appends a whereIs in the store with a new id and update the count
func (k Keeper) AppendWhereIs(
	ctx sdk.Context,
	whereIs types.Bucket,
) string {
	// Create the whereIs
	count := k.GetWhereIsCount(ctx)

	k.SetWhereIs(ctx, whereIs)

	// Update whereIs count
	k.SetWhereIsCount(ctx, count+1)

	return whereIs.Uuid
}

// SetWhereIs set a specific whereIs in the store
func (k Keeper) SetWhereIs(ctx sdk.Context, whereIs types.Bucket) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BucketKeyPrefix))
	b := k.cdc.MustMarshal(&whereIs)
	store.Set(types.BucketKey(whereIs.Uuid), b)
}

// GetBucket returns a whereIs from its id
func (k Keeper) GetBucket(ctx sdk.Context, id string) (val types.Bucket, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BucketKeyPrefix))
	b := store.Get(
		types.BucketKey(
			id,
		),
	)

	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	if val.Visibility == types.BucketVisibility_PUBLIC {
		return val, true
	}

	return val, false
}

// GetWhereIs returns a whereIs from its id
func (k Keeper) GetWhereIsByCreator(ctx sdk.Context, creator string) (list []types.Bucket) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BucketKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Bucket
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.Creator == creator {
			list = append(list, val)
		}
	}

	return
}

// RemoveWhereIs removes a whereIs from the store
func (k Keeper) RemoveWhereIs(ctx sdk.Context, id string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BucketKeyPrefix))
	store.Delete(types.BucketKey(id))
}

// GetAllWhereIs returns all whereIs
func (k Keeper) GetAllWhereIs(ctx sdk.Context) (list []types.Bucket) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BucketKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Bucket
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.Visibility == types.BucketVisibility_PUBLIC {
			list = append(list, val)
		}
	}
	return
}

func (k Keeper) GenerateKeyForDID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
