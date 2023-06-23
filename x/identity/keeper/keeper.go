package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sonrhq/core/internal/gateway"
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
		vaultKeeper   types.VaultKeeper
		authenticator gateway.Authenticator
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,

	accountKeeper types.AccountKeeper, bankKeeper types.BankKeeper, groupKeeper types.GroupKeeper,
	vaultKeeper types.VaultKeeper,
	authenticator gateway.Authenticator,
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
		vaultKeeper:   vaultKeeper,
		authenticator: authenticator,
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
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IdentityKeyPrefix))
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

// SetIdentity checks the validity of the identity and set it in the store based off its did method
func (k Keeper) SetIdentity(ctx sdk.Context, identity types.DIDDocument) error {
	// ptrs := strings.Split(identity.Id, ":")
	// addr := ptrs[len(ptrs)-1]
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IdentityKeyPrefix))
	b := k.cdc.MustMarshal(&identity)
	store.Set(types.IdentificationKey(
		identity.Id,
	), b)
	// owner, err := sdk.AccAddressFromBech32(addr)
	// if err != nil {
	// 	return status.Error(codes.InvalidArgument, "Invalid address")
	// }
	// k.accountKeeper.SetAccount(ctx, k.accountKeeper.NewAccountWithAddress(ctx, owner))
	return nil
}

// HasIdentity checks if an identity exists in the store across all did methods
func (k Keeper) HasIdentity(ctx sdk.Context, did string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IdentityKeyPrefix))
	b := store.Get(types.IdentificationKey(did))
	return b != nil
}

// GetIdentity returns the identity from the store
func (k Keeper) GetIdentity(ctx sdk.Context, did string) (val types.DIDDocument, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IdentityKeyPrefix))
	b := store.Get(types.IdentificationKey(did))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllDidDocument returns all didDocument
func (k Keeper) GetAllIdentities(ctx sdk.Context) (list []types.DIDDocument) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IdentityKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DIDDocument
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}
	return
}
