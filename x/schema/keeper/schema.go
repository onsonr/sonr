package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/x/schema/types"
)

func (k Keeper) GetSchemaCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SchemaCountKey)
	bz := store.Get(byteKey)

	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}

// SetWhoIsCount set the total number of whoIs
func (k Keeper) SetSchemaCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SchemaCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// GetSchemaFromCreator returns a WhoIs whos DIDDocument contains the given controller
func (k Keeper) GetSchemasFromID(ctx sdk.Context, creator string) (val []types.Schema, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SchemaKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	var vals []types.Schema = make([]types.Schema, 0)
	for ; iterator.Valid(); iterator.Next() {
		var instance types.Schema
		k.cdc.MustUnmarshal(iterator.Value(), &instance)
		if instance.Did == creator {
			vals = append(vals, instance)
		}
	}
	return vals, true
}

// GetSchemaFromCreator returns a WhoIs whos DIDDocument contains the given controller
func (k Keeper) GetSchemasFromLabel(ctx sdk.Context, label string) (val []types.Schema, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SchemaKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	var vals []types.Schema = make([]types.Schema, 0)
	for ; iterator.Valid(); iterator.Next() {
		var instance types.Schema
		k.cdc.MustUnmarshal(iterator.Value(), &instance)
		if instance.Label == label {
			vals = append(vals, instance)
		}
	}
	return vals, true
}

// SetSchema set a specific schema in the store from its did
func (k Keeper) SetSchema(ctx sdk.Context, schema types.Schema) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SchemaKeyPrefix))
	b := k.cdc.MustMarshal(&schema)
	store.Set(types.SchemaKey(
		schema.Did,
	), b)
}

// GetSchema returns an instance of a schema from its id
func (k Keeper) GetSchema(ctx sdk.Context, id string) (val types.Schema, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SchemaKeyPrefix))

	b := store.Get(types.SchemaKey(
		id,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}
