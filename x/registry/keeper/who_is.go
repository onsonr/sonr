package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/x/registry/types"
)

// GetWhoIsCount get the total number of whoIs
func (k Keeper) GetWhoIsCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.WhoIsCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetWhoIsCount set the total number of whoIs
func (k Keeper) SetWhoIsCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.WhoIsCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// SetWhoIs set a specific whoIs in the store
func (k Keeper) SetWhoIs(ctx sdk.Context, whoIs types.WhoIs) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhoIsKeyPrefix))
	b := k.cdc.MustMarshal(&whoIs)
	store.Set(types.WhoIsKey(
		whoIs.Owner,
	), b)
}

// GetWhoIs returns a whoIs from its id
func (k Keeper) GetWhoIs(ctx sdk.Context, id string) (val types.WhoIs, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhoIsKeyPrefix))

	b := store.Get(types.WhoIsKey(
		id,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllWhoIs returns all whoIs
func (k Keeper) GetAllWhoIs(ctx sdk.Context) (list []types.WhoIs) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhoIsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.WhoIs
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}
	return
}

// GetWhoIsIDBytes returns the byte representation of the ID
func GetWhoIsIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetWhoIsIDFromBytes returns ID in uint64 format from a byte array
func GetWhoIsIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
