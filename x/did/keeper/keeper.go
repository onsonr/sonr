package keeper

import (
	"cosmossdk.io/collections"
	storetypes "cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"cosmossdk.io/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	apiv1 "github.com/di-dao/core/api/did/v1"
	"github.com/di-dao/core/x/did/controller"
	"github.com/di-dao/core/x/did/types"
)

// defaultCurve is the default curve used for key generation

// Keeper defines the middleware keeper.
type Keeper struct {
	cdc codec.BinaryCodec

	logger log.Logger

	// state management
	OrmDB  apiv1.StateStore
	Params collections.Item[types.Params]
	Schema collections.Schema

	AccountKeeper authkeeper.AccountKeeper

	authority string
}

// NewKeeper creates a new poa Keeper instance
func NewKeeper(cdc codec.BinaryCodec, storeService storetypes.KVStoreService, accKeeper authkeeper.AccountKeeper, logger log.Logger, authority string) Keeper {
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
		cdc:           cdc,
		logger:        logger,
		Params:        collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		authority:     authority,
		OrmDB:         store,
		AccountKeeper: accKeeper,
	}
	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}

	k.Schema = schema
	return k
}

// GenerateKeyshares generates a new keyshare set. First step
func (k Keeper) GenerateKeyshares(ctx sdk.Context) (types.KssI, error) {
	return controller.GenKSS()
}

// LinkController links a user identifier to a kss pair creating a controller. Second step
func (k Keeper) LinkController(ctx sdk.Context, kss types.KssI, identifier string) ([]byte, error) {
	c, err := controller.Create(kss)
	if err != nil {
		return nil, err
	}
	return c.Set("email", identifier)
}

// AssignVault assigns a vault to a controller. Third step
func (k Keeper) AssignVault(ctx sdk.Context, c types.ControllerI) error {
	return nil
}
