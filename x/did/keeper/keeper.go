package keeper

import (
	"cosmossdk.io/collections"
	storetypes "cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"cosmossdk.io/orm/model/ormdb"
	nftkeeper "cosmossdk.io/x/nft/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	stakkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"github.com/ipfs/kubo/client/rpc"

	apiv1 "github.com/onsonr/sonr/api/did/v1"
	"github.com/onsonr/sonr/x/did/types"
)

// Keeper defines the middleware keeper.
type Keeper struct {
	cdc codec.BinaryCodec

	logger log.Logger

	// state management
	OrmDB  apiv1.StateStore
	Params collections.Item[types.Params]
	Schema collections.Schema

	AccountKeeper authkeeper.AccountKeeper
	NftKeeper     nftkeeper.Keeper
	StakingKeeper *stakkeeper.Keeper

	authority  string
	ipfsClient *rpc.HttpApi
}

// NewKeeper creates a new poa Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec,
	storeService storetypes.KVStoreService,
	accKeeper authkeeper.AccountKeeper,
	nftKeeper nftkeeper.Keeper,
	stkKeeper *stakkeeper.Keeper,
	logger log.Logger,
	authority string,
) Keeper {
	logger = logger.With(log.ModuleKey, "x/"+types.ModuleName)
	sb := collections.NewSchemaBuilder(storeService)
	if authority == "" {
		authority = authtypes.NewModuleAddress(govtypes.ModuleName).String()
	}
	db, err := ormdb.NewModuleDB(
		&types.ORMModuleSchema,
		ormdb.ModuleDBOptions{KVStoreService: storeService},
	)
	if err != nil {
		panic(err)
	}
	store, err := apiv1.NewStateStore(db)
	if err != nil {
		panic(err)
	}

	// Initialize IPFS client
	ipfsClient, _ := rpc.NewLocalApi()
	k := Keeper{
		ipfsClient: ipfsClient,
		cdc:        cdc,
		logger:     logger,
		Params: collections.NewItem(
			sb,
			types.ParamsKey,
			"params",
			codec.CollValue[types.Params](cdc),
		),
		authority:     authority,
		OrmDB:         store,
		AccountKeeper: accKeeper,
		NftKeeper:     nftKeeper,
		StakingKeeper: stkKeeper,
	}
	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}

	k.Schema = schema
	return k
}

// VerifyServicePermissions checks if a service has permission
func (k Keeper) VerifyServicePermissions(
	ctx sdk.Context,
	addr string,
	service string,
	permissions string,
) bool {
	return false
}
