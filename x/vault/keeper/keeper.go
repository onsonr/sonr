package keeper

import (
	"context"
	"time"

	"cosmossdk.io/collections"
	storetypes "cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"cosmossdk.io/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/ipfs/kubo/client/rpc"

	apiv1 "github.com/onsonr/sonr/api/vault/v1"
	"github.com/onsonr/sonr/pkg/dwn"
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

	ipfsClient *rpc.HttpApi

	DIDKeeper didkeeper.Keeper
}

// NewKeeper creates a new Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec,
	storeService storetypes.KVStoreService,
	logger log.Logger,
	authority string,
	didk didkeeper.Keeper,
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
		cdc:       cdc,
		logger:    logger,
		DIDKeeper: didk,
		Params:    collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		OrmDB:     store,

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

// currentSchema returns the current schema
func (k Keeper) CurrentSchema(ctx sdk.Context) (*dwn.Schema, error) {
	p, err := k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}
	schema := p.Schema
	return &dwn.Schema{
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

// assembleVault assembles the initial vault
func (k Keeper) AssembleVault(ctx sdk.Context) (string, int64, error) {
	_, con, err := k.DIDKeeper.NewController(ctx)
	if err != nil {
		return "", 0, err
	}
	usrKs, err := con.ExportUserKs()
	if err != nil {
		return "", 0, err
	}
	sch, err := k.CurrentSchema(ctx)
	if err != nil {
		return "", 0, err
	}
	v, err := types.NewVault(usrKs, con.SonrAddress(), con.ChainID(), sch)
	if err != nil {
		return "", 0, err
	}
	cid, err := k.ipfsClient.Unixfs().Add(context.Background(), v.FS)
	if err != nil {
		return "", 0, err
	}
	return cid.String(), k.CalculateExpiration(ctx, time.Second*15), nil
}
