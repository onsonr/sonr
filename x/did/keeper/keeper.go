package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"cosmossdk.io/collections"
	storetypes "cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"cosmossdk.io/orm/model/ormdb"
	sdk "github.com/cosmos/cosmos-sdk/types"

	apiv1 "github.com/di-dao/core/api/did/v1"
	"github.com/di-dao/core/crypto/core/curves"
	"github.com/di-dao/core/x/did/types"
)

var (
	// vault is the global vault instance
	vault vaultStore

	// defaultCurve is the default curve used for key generation
	defaultCurve = curves.P256()
)

// Keeper defines the middleware keeper.
type Keeper struct {
	cdc codec.BinaryCodec

	logger log.Logger

	// state management
	OrmDB  apiv1.StateStore
	Params collections.Item[types.Params]
	Schema collections.Schema

	authority string
}

// NewKeeper creates a new poa Keeper instance
func NewKeeper(cdc codec.BinaryCodec, storeService storetypes.KVStoreService, logger log.Logger, authority string) Keeper {
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
		cdc:       cdc,
		logger:    logger,
		Params:    collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		authority: authority,
		OrmDB:     store,
	}
	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}

	k.Schema = schema
	return k
}

// GenerateKSS generates a new keyshare set
func (k Keeper) GenerateKSS(ctx sdk.Context) (*ValidatorKeyshare, *UserKeyshare, error) {
	return generateKSS()
}
