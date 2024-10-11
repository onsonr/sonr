package keeper

import (
	"cosmossdk.io/collections"
	storetypes "cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"cosmossdk.io/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/codec"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/ipfs/kubo/client/rpc"

	apiv1 "github.com/onsonr/sonr/api/vault/v1"
	didkeeper "github.com/onsonr/sonr/x/did/keeper"
	macaroonkeeper "github.com/onsonr/sonr/x/macaroon/keeper"
	"github.com/onsonr/sonr/x/vault/types"
)

type Keeper struct {
	cdc codec.BinaryCodec

	logger log.Logger

	// state management
	Schema collections.Schema
	Params collections.Item[types.Params]
	OrmDB  apiv1.StateStore

	authority string

	ipfsClient *rpc.HttpApi

	AccountKeeper  authkeeper.AccountKeeper
	DIDKeeper      didkeeper.Keeper
	MacaroonKeeper macaroonkeeper.Keeper
}

// NewKeeper creates a new Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec,
	storeService storetypes.KVStoreService,
	logger log.Logger,
	authority string,
	authKeeper authkeeper.AccountKeeper,
	didk didkeeper.Keeper,
	macaroonKeeper macaroonkeeper.Keeper,
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

	ipfsClient, _ := rpc.NewLocalApi()
	k := Keeper{
		cdc:            cdc,
		logger:         logger,
		DIDKeeper:      didk,
		MacaroonKeeper: macaroonKeeper,
		AccountKeeper:  authKeeper,
		Params:         collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		OrmDB:          store,

		ipfsClient: ipfsClient,
		authority:  authority,
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}

	k.Schema = schema

	return k
}
