package keeper

import (
	"crypto/sha256"
	"fmt"

	"cosmossdk.io/collections"
	storetypes "cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"cosmossdk.io/orm/model/ormdb"
	nftkeeper "cosmossdk.io/x/nft/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"gopkg.in/macaroon.v2"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	stakkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"

	"github.com/onsonr/crypto/mpc"
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

	authority string
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
	k := Keeper{
		cdc:    cdc,
		logger: logger,
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

func (k Keeper) NewController(ctx sdk.Context) (uint64, types.ControllerI, error) {
	shares, err := mpc.GenerateKeyshares()
	if err != nil {
		return 0, nil, err
	}
	controller, err := types.NewController(shares)
	if err != nil {
		return 0, nil, err
	}
	entry, err := controller.GetTableEntry()
	if err != nil {
		return 0, nil, err
	}
	num, err := k.OrmDB.ControllerTable().InsertReturningNumber(ctx, entry)
	if err != nil {
		return 0, nil, err
	}
	return num, controller, nil
}

// IssueMacaroon creates a macaroon with the specified parameters.
func (k Keeper) IssueMacaroon(ctx sdk.Context, sharedMPCPubKey, location, id string, blockExpiry uint64) (*macaroon.Macaroon, error) {
	// Derive the root key by hashing the shared MPC public key
	rootKey := sha256.Sum256([]byte(sharedMPCPubKey))
	// Create the macaroon
	m, err := macaroon.New(rootKey[:], []byte(id), location, macaroon.LatestVersion)
	if err != nil {
		return nil, err
	}

	// Add the block expiry caveat
	caveat := fmt.Sprintf("block-expiry=%d", blockExpiry)
	err = m.AddFirstPartyCaveat([]byte(caveat))
	if err != nil {
		return nil, err
	}

	return m, nil
}
