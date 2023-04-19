package keeper

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sonrhq/core/x/identity/types"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace

		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
		groupKeeper   types.GroupKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,

	accountKeeper types.AccountKeeper, bankKeeper types.BankKeeper, groupKeeper types.GroupKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}
	k := &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramstore:    ps,
		accountKeeper: accountKeeper, bankKeeper: bankKeeper, groupKeeper: groupKeeper,
	}
	return k
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams()
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                          DIDDocument Keeper Functions                          ||
// ! ||--------------------------------------------------------------------------------||

// CheckAlias checks if an alias is already used
func (k Keeper) CheckAlias(ctx sdk.Context, alias string) error {
	_, found := k.GetPrimaryIdentityByAlias(ctx, alias)
	if found {
		return status.Error(codes.AlreadyExists, "Alias already exists")
	}
	return nil
}

// SetDidDocument set a specific didDocument in the store from its index
func (k Keeper) SetPrimaryIdentity(ctx sdk.Context, didDocument types.DidDocument) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PrimaryIdentityPrefix))

	ptrs := strings.Split(didDocument.Id, ":")
	addr := ptrs[len(ptrs)-1]
	didDocument.Owner = addr

	b := k.cdc.MustMarshal(&didDocument)
	store.Set(types.DidDocumentKey(
		didDocument.Id,
	), b)
}

// GetDidDocument returns a didDocument from its index
func (k Keeper) GetPrimaryIdentity(
	ctx sdk.Context,
	did string,
) (val types.DidDocument, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PrimaryIdentityPrefix))
	b := store.Get(types.DidDocumentKey(
		did,
	))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetPrimaryIdentityByAlias returns a didDocument from its index
func (k Keeper) GetPrimaryIdentityByAlias(
	ctx sdk.Context,
	alias string,
) (val types.DidDocument, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PrimaryIdentityPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var doc types.DidDocument
		k.cdc.MustUnmarshal(iterator.Value(), &doc)
		if doc.AlsoKnownAs[0] == alias {
			val = doc
			found = true
		}
	}
	return val, found
}

// GetPrimaryIdentityByAddress iterates over all didDocuments and returns the first one that matches the address
func (k Keeper) GetPrimaryIdentityByAddress(
	ctx sdk.Context,
	addr string,
) (val types.DidDocument, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PrimaryIdentityPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var doc types.DidDocument
		k.cdc.MustUnmarshal(iterator.Value(), &doc)
		if doc.Owner == addr {
			val = doc
			found = true
		}
	}
	return val, found
}

// RemoveDidDocument removes a didDocument from the store
func (k Keeper) RemovePrimaryIdentity(
	ctx sdk.Context,
	did string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PrimaryIdentityPrefix))
	store.Delete(types.DidDocumentKey(
		did,
	))
}

// GetAllDidDocument returns all didDocument
func (k Keeper) GetAllPrimaryIdentities(ctx sdk.Context) (list []types.DidDocument) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PrimaryIdentityPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DidDocument
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// SetDidDocument set a specific didDocument in the store from its index
func (k Keeper) SetBlockchainIdentity(ctx sdk.Context, didDocument types.DidDocument) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BlockchainIdentityPrefix))
	b := k.cdc.MustMarshal(&didDocument)
	store.Set(types.DidDocumentKey(
		didDocument.Id,
	), b)
}

// SetDidDocument set a specific didDocument in the store from its index
func (k Keeper) SetBlockchainIdentities(ctx sdk.Context, docs ...*types.DidDocument) {
	for _, doc := range docs {
		k.SetBlockchainIdentity(ctx, *doc)
	}
}

// GetDidDocument returns a didDocument from its index
func (k Keeper) GetBlockchainIdentity(
	ctx sdk.Context,
	did string,

) (val types.DidDocument, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BlockchainIdentityPrefix))
	b := store.Get(types.DidDocumentKey(
		did,
	))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetBlockchainIdentityByAddress iterates over all didDocuments and returns the first one that matches the address
func (k Keeper) GetBlockchainIdentityByAddress(
	ctx sdk.Context,
	addr string,
) (val types.DidDocument, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BlockchainIdentityPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var doc types.DidDocument
		k.cdc.MustUnmarshal(iterator.Value(), &doc)
		if doc.MatchesAddress(addr) {
			val = doc
			found = true
		}
	}
	return val, found
}

// RemoveDidDocument removes a didDocument from the store
func (k Keeper) RemoveBlockchainIdentity(
	ctx sdk.Context,
	did string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BlockchainIdentityPrefix))
	store.Delete(types.DidDocumentKey(
		did,
	))
}

// GetAllDidDocument returns all didDocument
func (k Keeper) GetAllBlockchainIdentities(ctx sdk.Context) (list []types.DidDocument) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BlockchainIdentityPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DidDocument
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}
	return
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                         Relationships Keeper Functions                         ||
// ! ||--------------------------------------------------------------------------------||

// HasRelationship checks if the element exists in the store
func (k Keeper) HasRelationship(ctx sdk.Context, reference string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RelationshipKeyPrefix))
	return store.Has(types.RelationshipKey(reference))
}

// SetRelationship set a specific Service in the store from its index
func (k Keeper) SetRelationship(ctx sdk.Context, VerificationRelationship types.VerificationRelationship) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RelationshipKeyPrefix))
	b := k.cdc.MustMarshal(&VerificationRelationship)
	store.Set(types.RelationshipKey(VerificationRelationship.Reference), b)
}

// GetRelationship returns a Service from its index
func (k Keeper) GetRelationship(ctx sdk.Context, reference string) (val types.VerificationRelationship, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RelationshipKeyPrefix))

	b := store.Get(types.RelationshipKey(reference))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllRelationships returns all Relationship
func (k Keeper) GetAllRelationships(ctx sdk.Context) (list []types.VerificationRelationship) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RelationshipKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.VerificationRelationship
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) GetRelationshipsFromList(ctx sdk.Context, addrs ...string) ([]types.VerificationRelationship, error) {
	vrs := make([]types.VerificationRelationship, 0, len(addrs))

	for _, addr := range addrs {
		if vr, found := k.GetRelationship(sdk.UnwrapSDKContext(ctx), addr); found {
			vrs = append(vrs, vr)
		} else {
			return nil, status.Error(codes.NotFound, "not found")
		}
	}

	return vrs, nil
}
