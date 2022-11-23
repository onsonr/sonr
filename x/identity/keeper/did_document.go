package keeper

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/x/identity/types"
)

// SetDidDocument set a specific didDocument in the store from its index
func (k Keeper) SetDidDocument(ctx sdk.Context, didDocument types.DidDocument) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidDocumentKeyPrefix))
	b := k.cdc.MustMarshal(&didDocument)
	store.Set(types.DidDocumentKey(
		didDocument.ID,
	), b)
}

// GetDidDocument returns a didDocument from its index
func (k Keeper) GetDidDocument(
	ctx sdk.Context,
	did string,

) (val types.DidDocument, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidDocumentKeyPrefix))

	b := store.Get(types.DidDocumentKey(
		did,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDidDocument removes a didDocument from the store
func (k Keeper) RemoveDidDocument(
	ctx sdk.Context,
	did string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidDocumentKeyPrefix))
	store.Delete(types.DidDocumentKey(
		did,
	))
}

// GetAllDidDocument returns all didDocument
func (k Keeper) GetAllDidDocument(ctx sdk.Context) (list []types.DidDocument) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidDocumentKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DidDocument
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) GetDidDocumentByAKA(ctx sdk.Context, aka string) (types.DidDocument, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidDocumentKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()
	aka = strings.TrimSpace(aka)

	var val types.DidDocument
	for ; iterator.Valid(); iterator.Next() {
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		for _, s := range val.AlsoKnownAs {
			if aka == s {
				return val, true
			}
		}
	}

	return val, false
}
