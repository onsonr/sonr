package keeper

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/patrickmn/go-cache"
	"github.com/sonr-hq/sonr/x/identity/types"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace

		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
		groupKeeper   types.GroupKeeper
		mintKeeper    types.MintKeeper
		web           *webauthn.WebAuthn
		userCache     *cache.Cache
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,

	accountKeeper types.AccountKeeper, bankKeeper types.BankKeeper, groupKeeper types.GroupKeeper, mintKeeper types.MintKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	wan, err := webauthn.New(&webauthn.Config{
		RPDisplayName: "Sonr ID",
		RPID:          "sonr.network",
		RPOrigin:      "https://auth.sonr.network",
	})
	if err != nil {
		panic(err)
	}

	k := &Keeper{
		web:           wan,
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramstore:    ps,
		accountKeeper: accountKeeper, bankKeeper: bankKeeper, groupKeeper: groupKeeper, mintKeeper: mintKeeper,
		userCache: cache.New(5*time.Minute, 10*time.Minute),
	}
	return k
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
