package keeper

import (
	"time"

	"cosmossdk.io/collections"
	storetypes "cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"cosmossdk.io/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/ipfs/kubo/client/rpc"

	apiv1 "github.com/onsonr/sonr/api/vault/v1"
	dwngen "github.com/onsonr/sonr/pkg/motr/config"
	didkeeper "github.com/onsonr/sonr/x/did/keeper"
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

	ipfsClient  *rpc.HttpApi
	hasIpfsConn bool

	AccountKeeper authkeeper.AccountKeeper
	DIDKeeper     didkeeper.Keeper
}

// NewKeeper creates a new Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec,
	storeService storetypes.KVStoreService,
	logger log.Logger,
	authority string,
	authKeeper authkeeper.AccountKeeper,
	didKeeper didkeeper.Keeper,
) Keeper {
	var hasIpfs bool
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

	ipfsClient, err := rpc.NewLocalApi()
	if err != nil {
		hasIpfs = false
	}

	if ipfsClient != nil {
		hasIpfs = true
	}

	k := Keeper{
		cdc:           cdc,
		logger:        logger,
		DIDKeeper:     didKeeper,
		AccountKeeper: authKeeper,
		Params:        collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		OrmDB:         store,

		ipfsClient:  ipfsClient,
		hasIpfsConn: hasIpfs,
		authority:   authority,
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}

	k.Schema = schema

	return k
}

// currentSchema returns the current schema
func (k Keeper) currentSchema(ctx sdk.Context) (*dwngen.Schema, error) {
	p, err := k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}
	schema := p.Schema
	return &dwngen.Schema{
		Version:    int(schema.Version),
		Account:    schema.Account,
		Asset:      schema.Asset,
		Chain:      schema.Chain,
		Credential: schema.Credential,
		Jwk:        schema.Jwk,
		Grant:      schema.Grant,
		Keyshare:   schema.Keyshare,
		Profile:    schema.Profile,
	}, nil
}

func calculateBlockExpiry(sdkctx sdk.Context, duration time.Duration) int64 {
	blockTime := sdkctx.BlockTime()
	avgBlockTime := float64(blockTime.Sub(blockTime).Seconds())
	return int64(duration.Seconds() / avgBlockTime)
}
