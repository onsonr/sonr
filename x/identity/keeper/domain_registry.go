package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-hq/sonr/x/identity/types"
)

// SetDomainRecord set a specific DomainRecord in the store from its index
func (k Keeper) SetDomainRecord(ctx sdk.Context, DomainRecord types.DomainRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DomainRecordKeyPrefix))
	b := k.cdc.MustMarshal(&DomainRecord)
	store.Set(types.DomainRecordKey(
		DomainRecord.Domain,
		DomainRecord.Index,
	), b)
}

// GetDomainRecord returns a DomainRecord from its index
func (k Keeper) GetDomainRecord(
	ctx sdk.Context,
	domain string,
	tld string,

) (val types.DomainRecord, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DomainRecordKeyPrefix))

	b := store.Get(types.DomainRecordKey(
		domain,
		tld,
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
	index string,
	domain string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DomainRecordKeyPrefix))
	store.Delete(types.DomainRecordKey(
		domain,
		index,
	))
}

// GetAllDomainRecord returns all DomainRecord
func (k Keeper) GetAllDomainRecord(ctx sdk.Context) (list []types.DomainRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DomainRecordKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DomainRecord
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
