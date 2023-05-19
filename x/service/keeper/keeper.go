package keeper

import (
	"context"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sonrhq/core/x/service/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace

		groupKeeper    types.GroupKeeper
		identityKeeper types.IdentityKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,

	groupKeeper types.GroupKeeper,
	identityKeeper types.IdentityKeeper,
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

		groupKeeper:    groupKeeper,
		identityKeeper: identityKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                         Service Record Keeper Functions                        ||
// ! ||--------------------------------------------------------------------------------||

// SetServiceRecord set a specific serviceRecord in the store from its Id
func (k Keeper) SetServiceRecord(ctx sdk.Context, serviceRecord types.ServiceRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ServiceRecordKeyPrefix))
	b := k.cdc.MustMarshal(&serviceRecord)
	store.Set(types.ServiceRecordKey(
		cleanServiceDomain(serviceRecord.Origin),
	), b)
}

// GetServiceRecord returns a serviceRecord from its Id
func (k Keeper) GetServiceRecord(
	ctx sdk.Context,
	origin string,
) (val types.ServiceRecord, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ServiceRecordKeyPrefix))

	b := store.Get(types.ServiceRecordKey(
		origin,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveServiceRecord removes a serviceRecord from the store
func (k Keeper) RemoveServiceRecord(
	ctx sdk.Context,
	Id string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ServiceRecordKeyPrefix))
	store.Delete(types.ServiceRecordKey(
		Id,
	))
}

// GetAllServiceRecord returns all serviceRecord
func (k Keeper) GetAllServiceRecord(ctx sdk.Context) (list []types.ServiceRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ServiceRecordKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ServiceRecord
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams()
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

func (k Keeper) Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
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

// ! ||--------------------------------------------------------------------------------||
// ! ||                              Service Relationships                             ||
// ! ||--------------------------------------------------------------------------------||

// SetServiceRelationships set a specific serviceRelationships in the store
func (k Keeper) SetServiceRelationship(ctx sdk.Context, serviceRelationships types.ServiceRelationship) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ServiceRelationshipsKey))
	b := k.cdc.MustMarshal(&serviceRelationships)
	store.Set(types.KeyPrefix(serviceRelationships.Did), b)
}

// GetServiceRelationships returns a serviceRelationships from its id
func (k Keeper) GetServiceRelationship(ctx sdk.Context, id string) (val types.ServiceRelationship, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ServiceRelationshipsKey))
	b := store.Get(types.KeyPrefix(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveServiceRelationships removes a serviceRelationships from the store
func (k Keeper) RemoveServiceRelationship(ctx sdk.Context, id string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ServiceRelationshipsKey))
	store.Delete(types.KeyPrefix(id))
}

// GetAllServiceRelationships returns all serviceRelationships
func (k Keeper) GetAllServiceRelationships(ctx sdk.Context) (list []types.ServiceRelationship) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ServiceRelationshipsKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ServiceRelationship
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
