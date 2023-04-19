package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/sonrhq/core/x/domain/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
	}
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
// ! ||                               Keeper SLD Methods                               ||
// ! ||--------------------------------------------------------------------------------||

// SetSLDRecord set a specific sLDRecord in the store from its index
func (k Keeper) SetSLDRecord(ctx sdk.Context, sLDRecord types.SLDRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SLDRecordKeyPrefix))
	b := k.cdc.MustMarshal(&sLDRecord)
	store.Set(types.SLDRecordKey(
		sLDRecord.Index,
	), b)
}

// GetSLDRecord returns a sLDRecord from its index
func (k Keeper) GetSLDRecord(
	ctx sdk.Context,
	index string,

) (val types.SLDRecord, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SLDRecordKeyPrefix))

	b := store.Get(types.SLDRecordKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSLDRecord removes a sLDRecord from the store
func (k Keeper) RemoveSLDRecord(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SLDRecordKeyPrefix))
	store.Delete(types.SLDRecordKey(
		index,
	))
}

// GetAllSLDRecord returns all sLDRecord
func (k Keeper) GetAllSLDRecord(ctx sdk.Context) (list []types.SLDRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SLDRecordKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SLDRecord
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                               Keeper TLD Methods                               ||
// ! ||--------------------------------------------------------------------------------||

// SetTLDRecord set a specific tLDRecord in the store from its index
func (k Keeper) SetTLDRecord(ctx sdk.Context, tLDRecord types.TLDRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TLDRecordKeyPrefix))
	b := k.cdc.MustMarshal(&tLDRecord)
	store.Set(types.TLDRecordKey(
		tLDRecord.Index,
	), b)
}

// GetTLDRecord returns a tLDRecord from its index
func (k Keeper) GetTLDRecord(
	ctx sdk.Context,
	index string,

) (val types.TLDRecord, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TLDRecordKeyPrefix))

	b := store.Get(types.TLDRecordKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTLDRecord removes a tLDRecord from the store
func (k Keeper) RemoveTLDRecord(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TLDRecordKeyPrefix))
	store.Delete(types.TLDRecordKey(
		index,
	))
}

// GetAllTLDRecord returns all tLDRecord
func (k Keeper) GetAllTLDRecord(ctx sdk.Context) (list []types.TLDRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TLDRecordKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TLDRecord
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                                 HNS DNS Utility                                ||
// ! ||--------------------------------------------------------------------------------||

// ResolveHNSTLD resolves a HNS TLD record
func ResolveHNSTLD(options ...types.DNSOption) ([]*types.DNSRecord, error) {
	opts := types.DefaultDNSOptions()
	res := opts.Apply(options...)
	var records []*types.DNSRecord
	for target := range res.ResMap {
		for _, record := range res.ResMap[target] {
			records = append(records, types.NewDNSRecordFromResultItem(target, record))
		}
	}
	return records, nil
}
