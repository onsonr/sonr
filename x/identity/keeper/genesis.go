package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonrhq/core/x/identity/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams()
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// SetDIDDocument set a specific dIDDocument in the store from its index
func (k Keeper) SetDIDDocument(ctx sdk.Context, dIDDocument types.DIDDocument) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DIDDocumentKeyPrefix))
	b := k.cdc.MustMarshal(&dIDDocument)
	store.Set(types.DIDDocumentKey(
		dIDDocument.Id,
	), b)
}

// GetDIDDocument returns a dIDDocument from its index
func (k Keeper) GetDIDDocument(
	ctx sdk.Context,
	index string,

) (val types.DIDDocument, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DIDDocumentKeyPrefix))

	b := store.Get(types.DIDDocumentKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDIDDocument removes a dIDDocument from the store
func (k Keeper) RemoveDIDDocument(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DIDDocumentKeyPrefix))
	store.Delete(types.DIDDocumentKey(
		index,
	))
}

// GetAllDIDDocument returns all dIDDocument
func (k Keeper) GetAllDIDDocument(ctx sdk.Context) (list []types.DIDDocument) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DIDDocumentKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DIDDocument
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                          DIDDocument Keeper Functions                          ||
// ! ||--------------------------------------------------------------------------------||

// CheckAlsoKnownAs checks if an alias is already used
func (k Keeper) CheckAlsoKnownAs(ctx sdk.Context, alias string) error {
	_, found := k.GetIdentityByPrimaryAlias(ctx, alias)
	if found {
		return status.Error(codes.AlreadyExists, "Alias already exists")
	}
	return nil
}

// GetDidDocumentByAlsoKnownAs returns a didDocument from its index
func (k Keeper) GetIdentityByPrimaryAlias(
	ctx sdk.Context,
	alias string,
) (val types.DIDDocument, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DIDDocumentKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var doc types.DIDDocument
		k.cdc.MustUnmarshal(iterator.Value(), &doc)
		if doc.AlsoKnownAs[0] == alias {
			val = doc
			found = true
		}
	}
	return val, found
}
