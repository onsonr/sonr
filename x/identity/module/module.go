package module

import (
	"context"
	"encoding/json"
	"fmt"

	"cosmossdk.io/core/appmodule"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"

	"github.com/sonrhq/sonr/x/identity"
	"github.com/sonrhq/sonr/x/identity/keeper"
)

var (
	_ module.AppModuleBasic = AppModule{}
	_ module.HasGenesis     = AppModule{}
	_ appmodule.AppModule   = AppModule{}
)

// ConsensusVersion defines the current module consensus version.
const ConsensusVersion = 1

type AppModule struct {
	cdc        codec.Codec
	keeper     keeper.Keeper
	bankKeeper identity.BankKeeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(cdc codec.Codec, keeper keeper.Keeper, bk identity.BankKeeper) AppModule {
	return AppModule{
		cdc:        cdc,
		keeper:     keeper,
		bankKeeper: bk,
	}
}

func NewAppModuleBasic(m AppModule) module.AppModuleBasic {
	return module.CoreAppModuleBasicAdaptor(m.Name(), m)
}

// Name returns the identity module's name.
func (AppModule) Name() string { return identity.ModuleName }

// RegisterLegacyAminoCodec registers the identity module's types on the LegacyAmino codec.
// New modules do not need to support Amino.
func (AppModule) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the identity module.
func (AppModule) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *gwruntime.ServeMux) {
	if err := identity.RegisterQueryHandlerClient(context.Background(), mux, identity.NewQueryClient(clientCtx)); err != nil {
		panic(err)
	}
}

// RegisterInterfaces registers interfaces and implementations of the identity module.
func (AppModule) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	identity.RegisterInterfaces(registry)
}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 { return ConsensusVersion }

// RegisterServices registers a gRPC query service to respond to the module-specific gRPC queries.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	identity.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
	identity.RegisterQueryServer(cfg.QueryServer(), keeper.NewQueryServerImpl(am.keeper))

	// Register in place module state migration migrations
	// m := keeper.NewMigrator(am.keeper)
	// if err := cfg.RegisterMigration(identity.ModuleName, 1, m.Migrate1to2); err != nil {
	// 	panic(fmt.Sprintf("failed to migrate x/%s from version 1 to 2: %v", identity.ModuleName, err))
	// }
}

// DefaultGenesis returns default genesis state as raw bytes for the module.
func (AppModule) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(identity.NewGenesisState())
}

// ValidateGenesis performs genesis state validation for the circuit module.
func (AppModule) ValidateGenesis(cdc codec.JSONCodec, _ client.TxEncodingConfig, bz json.RawMessage) error {
	var data identity.GenesisState
	if err := cdc.UnmarshalJSON(bz, &data); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", identity.ModuleName, err)
	}

	return data.Validate()
}

// InitGenesis performs genesis initialization for the identity module.
// It returns no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) {
	var genesisState identity.GenesisState
	cdc.MustUnmarshalJSON(data, &genesisState)

	if err := am.keeper.InitGenesis(ctx, &genesisState); err != nil {
		panic(fmt.Sprintf("failed to initialize %s genesis state: %v", identity.ModuleName, err))
	}
}

// ExportGenesis returns the exported genesis state as raw bytes for the circuit
// module.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	gs, err := am.keeper.ExportGenesis(ctx)
	if err != nil {
		panic(fmt.Sprintf("failed to export %s genesis state: %v", identity.ModuleName, err))
	}

	return cdc.MustMarshalJSON(gs)
}
