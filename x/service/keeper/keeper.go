package keeper

import (
	"cosmossdk.io/collections"
	storetypes "cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"cosmossdk.io/orm/model/ormdb"
	nftkeeper "cosmossdk.io/x/nft/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	groupkeeper "github.com/cosmos/cosmos-sdk/x/group/keeper"

	apiv1 "github.com/onsonr/sonr/api/service/v1"
	didkeeper "github.com/onsonr/sonr/x/did/keeper"
	macaroonkeeper "github.com/onsonr/sonr/x/macaroon/keeper"
	"github.com/onsonr/sonr/x/service/types"
	vaultkeeper "github.com/onsonr/sonr/x/vault/keeper"
)

type Keeper struct {
	cdc codec.BinaryCodec

	logger log.Logger

	// state management
	Schema collections.Schema
	Params collections.Item[types.Params]
	OrmDB  apiv1.StateStore

	authority string

	DidKeeper      didkeeper.Keeper
	GroupKeeper    groupkeeper.Keeper
	MacaroonKeeper macaroonkeeper.Keeper
	NFTKeeper      nftkeeper.Keeper
	VaultKeeper    vaultkeeper.Keeper
}

// NewKeeper creates a new Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec,
	storeService storetypes.KVStoreService,
	logger log.Logger,
	authority string,
	didKeeper didkeeper.Keeper,
	groupKeeper groupkeeper.Keeper,
	macaroonKeeper macaroonkeeper.Keeper,
	nftKeeper nftkeeper.Keeper,
	vaultKeeper vaultkeeper.Keeper,
) Keeper {
	logger = logger.With(log.ModuleKey, "x/"+types.ModuleName)

	sb := collections.NewSchemaBuilder(storeService)

	if authority == "" {
		authority = authtypes.NewModuleAddress(govtypes.ModuleName).String()
	}

	db, err := ormdb.NewModuleDB(&types.ORMModuleSchema, ormdb.ModuleDBOptions{KVStoreService: storeService})
	if err != nil {
		panic(err)
	}

	store, err := apiv1.NewStateStore(db)
	if err != nil {
		panic(err)
	}

	k := Keeper{
		cdc:    cdc,
		logger: logger,

		Params: collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		OrmDB:  store,

		authority: authority,

		DidKeeper:      didKeeper,
		GroupKeeper:    groupKeeper,
		MacaroonKeeper: macaroonKeeper,
		NFTKeeper:      nftKeeper,
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}

	k.Schema = schema

	return k
}
