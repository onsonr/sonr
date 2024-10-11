package module

import (
	"os"

	"cosmossdk.io/core/address"
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/core/store"
	"cosmossdk.io/depinject"
	"cosmossdk.io/log"
	nftkeeper "cosmossdk.io/x/nft/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	groupkeeper "github.com/cosmos/cosmos-sdk/x/group/keeper"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"

	modulev1 "github.com/onsonr/sonr/api/service/module/v1"
	didkeeper "github.com/onsonr/sonr/x/did/keeper"
	macaroonkeeper "github.com/onsonr/sonr/x/macaroon/keeper"
	"github.com/onsonr/sonr/x/service/keeper"
	vaultkeeper "github.com/onsonr/sonr/x/vault/keeper"
)

var _ appmodule.AppModule = AppModule{}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (am AppModule) IsOnePerModuleType() {}

// IsAppModule implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

func init() {
	appmodule.Register(
		&modulev1.Module{},
		appmodule.Provide(ProvideModule),
	)
}

type ModuleInputs struct {
	depinject.In

	Cdc          codec.Codec
	StoreService store.KVStoreService
	AddressCodec address.Codec

	DidKeeper      didkeeper.Keeper
	GroupKeeper    groupkeeper.Keeper
	MacaroonKeeper macaroonkeeper.Keeper
	NFTKeeper      nftkeeper.Keeper
	StakingKeeper  stakingkeeper.Keeper
	SlashingKeeper slashingkeeper.Keeper
	VaultKeeper    vaultkeeper.Keeper
}

type ModuleOutputs struct {
	depinject.Out

	Module appmodule.AppModule
	Keeper keeper.Keeper
}

func ProvideModule(in ModuleInputs) ModuleOutputs {
	govAddr := authtypes.NewModuleAddress(govtypes.ModuleName).String()

	k := keeper.NewKeeper(in.Cdc, in.StoreService, log.NewLogger(os.Stderr), govAddr, in.DidKeeper, in.GroupKeeper, in.MacaroonKeeper, in.NFTKeeper, in.VaultKeeper)
	m := NewAppModule(in.Cdc, k, in.DidKeeper, in.MacaroonKeeper)

	return ModuleOutputs{Module: m, Keeper: k, Out: depinject.Out{}}
}
