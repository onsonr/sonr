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

// ! ||--------------------------------------------------------------------------------||
// ! ||                          DIDDocument Keeper Functions                          ||
// ! ||--------------------------------------------------------------------------------||

// SetDidDocument set a specific didDocument in the store from its index
func (k Keeper) SetDidDocument(ctx sdk.Context, didDocument types.DidDocument) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidDocumentKeyPrefix))
	b := k.cdc.MustMarshal(&didDocument)
	store.Set(types.DidDocumentKey(
		didDocument.Id,
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

// ! ||--------------------------------------------------------------------------------||
// ! ||                         Relationships Keeper Functions                         ||
// ! ||--------------------------------------------------------------------------------||

// HasRelationship checks if the element exists in the store
func (k Keeper) HasRelationship(ctx sdk.Context, reference string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RelationshipKeyPrefix))
	return store.Has(types.RelationshipKey(reference))
}

// Set Resolved Document sets all the relationships in the document
func (k Keeper) SetResolvedDocument(ctx sdk.Context, doc types.ResolvedDidDocument) {
	// Set AssertionMethod
	for _, v := range doc.AssertionMethod {
		k.SetRelationship(ctx, *v)
	}

	// Set Authentication
	for _, v := range doc.Authentication {
		k.SetRelationship(ctx, *v)
	}

	// Set CapabilityDelegation
	for _, v := range doc.CapabilityDelegation {
		k.SetRelationship(ctx, *v)
	}

	// Set CapabilityInvocation
	for _, v := range doc.CapabilityInvocation {
		k.SetRelationship(ctx, *v)
	}

	// Set KeyAgreement
	for _, v := range doc.KeyAgreement {
		k.SetRelationship(ctx, *v)
	}
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

func (k Keeper) ResolveDidDocument(ctx sdk.Context, doc types.DidDocument) (types.ResolvedDidDocument, error) {
	resolvedDidDocument := doc.ToResolved()

	vrs := []types.VerificationRelationship{}
	for _, relationship := range doc.VerificationMethod {
		vr, ok := k.GetRelationship(ctx, relationship.Id)
		if !ok {
			return types.ResolvedDidDocument{}, status.Error(codes.NotFound, fmt.Sprintf("verification relationship %s not found", relationship.Id))
		}
		vrs = append(vrs, vr)
	}

	resolvedDidDocument.AddVerificationRelationship(vrs)
	return *resolvedDidDocument, nil
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

// ! ||--------------------------------------------------------------------------------||
// ! ||                         Service Record Keeper Functions                        ||
// ! ||--------------------------------------------------------------------------------||

// SetService set a specific Service in the store from its index
func (k Keeper) SetService(ctx sdk.Context, Service types.Service) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ServiceKeyPrefix))
	b := k.cdc.MustMarshal(&Service)
	store.Set(types.ServiceKey(
		cleanServiceDomain(Service.Origin),
	), b)
}

// GetDomainRecord returns a DomainRecord from its index
func (k Keeper) GetService(
	ctx sdk.Context,
	origin string,
) (val types.Service, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ServiceKeyPrefix))

	b := store.Get(types.ServiceKey(
		cleanServiceDomain(origin),
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllServices returns all Services
func (k Keeper) GetAllServices(ctx sdk.Context) (list []types.Service) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ServiceKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Service
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}
	return
}

// cleanServiceDomain removes the url scheme and path from a service origin
func cleanServiceDomain(origin string) string {
	// Remove url scheme
	r := strings.NewReplacer("https://", "", "http://", "")
	origin = r.Replace(origin)

	if strings.Contains(origin, "/") {
		return strings.Split(origin, "/")[0]
	}
	return origin
}
