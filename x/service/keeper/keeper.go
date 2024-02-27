package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	storetypes "cosmossdk.io/core/store"
	"cosmossdk.io/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/codec"

	modulev1 "github.com/sonrhq/sonr/api/sonr/service/module/v1"
	"github.com/sonrhq/sonr/x/service"
)

// Keeper defines the module's keeper.
type Keeper struct {
	cdc          codec.BinaryCodec
	addressCodec address.Codec
	db           modulev1.StateStore

	// authority is the address capable of executing a MsgUpdateParams and other authority-gated message.
	// typically, this should be the x/gov module account.
	authority string

	// referenced keepers
	identityKeeper service.IdentityKeeper
	bankKeeper     service.BankKeeper
	groupKeeper    service.GroupKeeper

	// state management
	CollSchema collections.Schema
	Params     collections.Item[service.Params]
}

// NewKeeper creates a new Keeper instance
func NewKeeper(cdc codec.BinaryCodec, addressCodec address.Codec, storeService storetypes.KVStoreService, identityKeeper service.IdentityKeeper,
	bankKeeper service.BankKeeper, groupKeeper service.GroupKeeper,
	authority string,
) Keeper {
	if _, err := addressCodec.StringToBytes(authority); err != nil {
		panic(fmt.Errorf("invalid authority address: %w", err))
	}
	db, err := ormdb.NewModuleDB(serviceSchema, ormdb.ModuleDBOptions{KVStoreService: storeService})
	if err != nil {
		panic(err)
	}

	store, err := modulev1.NewStateStore(db)
	if err != nil {
		panic(err)
	}

	sb := collections.NewSchemaBuilder(storeService)
	k := Keeper{
		cdc:            cdc,
		addressCodec:   addressCodec,
		authority:      authority,
		Params:         collections.NewItem(sb, service.ParamsKey, "params", codec.CollValue[service.Params](cdc)),
		db:             store,
		identityKeeper: identityKeeper,
		bankKeeper:     bankKeeper,
		groupKeeper:    groupKeeper,
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}

	k.CollSchema = schema
	return k
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// setupInitialRecords sets up the initial records.
func (k Keeper) setupInitialRecords(ctx context.Context) error {
	err := k.db.ServiceTable().Save(ctx, &modulev1.Service{
		Origin:      "localhost",
		Name:        "Sonr LocalAuth",
		Description: "Sonr authentication service",
		Permissions: modulev1.ServicePermissions_SERVICE_PERMISSIONS_OWN,
	})
	if err != nil {
		return err
	}

	// Set default permissions for the base, read, write and own modules
	if err := k.db.BaseParamsTable().Save(ctx, &modulev1.BaseParams{
		Permissions:              modulev1.ServicePermissions_SERVICE_PERMISSIONS_BASE,
		Algorithm:                -7,
		AuthenticationAttachment: "platform",
	}); err != nil {
		return err
	}

	if err := k.db.ReadParamsTable().Save(ctx, &modulev1.ReadParams{
		Permissions:              modulev1.ServicePermissions_SERVICE_PERMISSIONS_READ,
		Algorithm:                -7,
		AuthenticationAttachment: "platform",
	}); err != nil {
		return err
	}
	if err := k.db.WriteParamsTable().Save(ctx, &modulev1.WriteParams{
		Permissions:              modulev1.ServicePermissions_SERVICE_PERMISSIONS_WRITE,
		Algorithm:                -8,
		ResidentKey:              "preferred",
		AuthenticationAttachment: "cross-platform",
	}); err != nil {
		return err
	}
	if err := k.db.OwnParamsTable().Save(ctx, &modulev1.OwnParams{
		Permissions:              modulev1.ServicePermissions_SERVICE_PERMISSIONS_OWN,
		Algorithm:                -8,
		ResidentKey:              "preferred",
		AuthenticationAttachment: "cross-platform",
	}); err != nil {
		return err
	}
	return nil
}
