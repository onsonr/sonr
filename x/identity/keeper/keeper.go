package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	storetypes "cosmossdk.io/core/store"
	"cosmossdk.io/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/codec"

	modulev1 "github.com/sonrhq/sonr/api/identity/module/v1"
	"github.com/sonrhq/sonr/x/identity"
)

type Keeper struct {
	cdc          codec.BinaryCodec
	addressCodec address.Codec
	// authority is the address capable of executing a MsgUpdateParams and other authority-gated message.
	// typically, this should be the x/gov module account.
	authority string
	db        modulev1.StateStore

	// referenced keepers
	bankKeeper identity.BankKeeper

	// state management
	Schema  collections.Schema
	Params  collections.Item[identity.Params]
	Counter collections.Map[string, uint64]
}

// NewKeeper creates a new Keeper instance
func NewKeeper(cdc codec.BinaryCodec, addressCodec address.Codec, storeService storetypes.KVStoreService,
	bk identity.BankKeeper,
	authority string,
) Keeper {
	if _, err := addressCodec.StringToBytes(authority); err != nil {
		panic(fmt.Errorf("invalid authority address: %w", err))
	}
	db, err := ormdb.NewModuleDB(identitySchema, ormdb.ModuleDBOptions{KVStoreService: storeService})
	if err != nil {
		panic(err)
	}

	store, err := modulev1.NewStateStore(db)
	if err != nil {
		panic(err)
	}
	sb := collections.NewSchemaBuilder(storeService)
	k := Keeper{
		cdc:          cdc,
		addressCodec: addressCodec,
		authority:    authority,
		Params:       collections.NewItem(sb, identity.ParamsKey, "params", codec.CollValue[identity.Params](cdc)),
		Counter:      collections.NewMap(sb, identity.CounterKey, "counter", collections.StringKey, collections.Uint64Value),
		db:           store,
		bankKeeper:   bk,
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}

	k.Schema = schema

	return k
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// DeriveAccount uses MPC to generate a new Account for an Identity.
func (k Keeper) DeriveAccount(ctx context.Context, id string) error {
	return nil
}

// GenerateIdentity generates a new Identity.
func (k Keeper) GenerateIdentity(ctx context.Context) error {
	return nil
}

// LinkCredential links a Credential to a Persona.
func (k Keeper) LinkCredential(ctx context.Context, identityID string) error {
	return nil
}

// LinkPersona links a Persona to an Account and Identity.
func (k Keeper) LinkPersona(ctx context.Context, identityID string) error {
	return nil
}

// RevokeAccount revokes an Account.
func (k Keeper) RevokeAccount(ctx context.Context, identityID string) error {
	return nil
}

// RevokeIdentity revokes an Identity.
func (k Keeper) RevokeIdentity(ctx context.Context, identityID string) error {
	return nil
}

// SignWithAccount signs a message with an Account.
func (k Keeper) SignWithAccount(ctx context.Context, identityID string) error {
	return nil
}

// UnlinkCredential unlinks a Credential from a Persona.
func (k Keeper) UnlinkCredential(ctx context.Context, identityID string) error {
	return nil
}

// UnlinkPersona unlinks a Persona from an Account and Identity.
func (k Keeper) UnlinkPersona(ctx context.Context, identityID string) error {
	return nil
}

// VerifyAccountSignature verifies a signature with an Account.
func (k Keeper) VerifyAccountSignature(ctx context.Context, identityID string) error {
	return nil
}
