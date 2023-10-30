package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/x/identity/types"
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

// ! ||--------------------------------------------------------------------------------||
// ! ||                            Controller Account Store                            ||
// ! ||--------------------------------------------------------------------------------||

// GetControllerAccountCount get the total number of controllerAccount
func (k Keeper) GetControllerAccountCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ControllerAccountCountKeyPrefix)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetControllerAccountCount set the total number of controllerAccount
func (k Keeper) SetControllerAccountCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ControllerAccountCountKeyPrefix)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}


// SetControllerAccount set a specific controllerAccount in the store
func (k Keeper) SetControllerAccount(ctx sdk.Context, controllerAccount types.ControllerAccount) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ControllerAccountKeyPrefix))
	b := k.cdc.MustMarshal(&controllerAccount)
	count := k.GetControllerAccountCount(ctx)
	store.Set(types.ControllerAccountKey(controllerAccount.Address), b)
	k.SetControllerAccountCount(ctx, count+1)
	return count
}

// GetControllerAccount returns a controllerAccount from its id
func (k Keeper) GetControllerAccount(ctx sdk.Context, address string) (val types.ControllerAccount, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ControllerAccountKeyPrefix))
	b := store.Get(types.ControllerAccountKey(address))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveControllerAccount removes a controllerAccount from the store
func (k Keeper) RemoveControllerAccount(ctx sdk.Context, address string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ControllerAccountKeyPrefix))
	store.Delete(types.ControllerAccountKey(address))
}

// GetAllControllerAccount returns all controllerAccount
func (k Keeper) GetAllControllerAccount(ctx sdk.Context) (list []types.ControllerAccount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ControllerAccountKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ControllerAccount
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                              Escrow Account Store                              ||
// ! ||--------------------------------------------------------------------------------||

// GetEscrowAccountCount get the total number of escrowAccount
func (k Keeper) GetEscrowAccountCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.EscrowAccountCountKeyPrefix)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetEscrowAccountCount set the total number of escrowAccount
func (k Keeper) SetEscrowAccountCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.EscrowAccountCountKeyPrefix)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}
