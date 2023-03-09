package keeper

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonrhq/core/x/identity/types"
)

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

// RemoveDomainRecord removes a DomainRecord from the store
func (k Keeper) RemoveDomainRecord(
	ctx sdk.Context,
	origin string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ServiceKeyPrefix))
	store.Delete(types.ServiceKey(
		cleanServiceDomain(origin),
	))
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
