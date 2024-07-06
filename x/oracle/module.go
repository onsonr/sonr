package oracle

import (
	"encoding/json"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/onsonr/hway/x/oracle/keeper"
	"github.com/onsonr/hway/x/oracle/types"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	abci "github.com/cometbft/cometbft/abci/types"
)

var (
	_ module.AppModuleBasic      = AppModuleBasic{}
	_ module.AppModule           = AppModule{}
	_ module.AppModuleSimulation = AppModule{}
)

// AppModuleBasic is the middleware AppModuleBasic.
type AppModuleBasic struct{}

// Name implements AppModuleBasic interface.
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterLegacyAminoCodec implements AppModuleBasic interface.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {}

// RegisterInterfaces registers module concrete types into protobuf Any.
func (AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {}

// DefaultGenesis returns default genesis state as raw bytes for the swap module.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return nil
}

// ValidateGenesis performs genesis state validation for the swap module.
func (AppModuleBasic) ValidateGenesis(
	cdc codec.JSONCodec,
	config client.TxEncodingConfig,
	bz json.RawMessage,
) error {
	return nil
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the swap module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {}

// GetTxCmd implements AppModuleBasic interface.
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return nil
}

// GetQueryCmd implements AppModuleBasic interface.
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return nil
}

// AppModule is the middleware AppModule.
type AppModule struct {
	AppModuleBasic
	keeper keeper.Keeper
}

// IsAppModule implements module.AppModule.
func (AppModule) IsAppModule() {
}

// IsOnePerModuleType implements module.AppModule.
func (AppModule) IsOnePerModuleType() {
}

// NewAppModule initializes a new AppModule for the middleware.
func NewAppModule(keeper keeper.Keeper) *AppModule {
	return &AppModule{
		keeper: keeper,
	}
}

// RegisterInvariants implements the AppModule interface.
func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {}

// RegisterServices registers module services.
func (am AppModule) RegisterServices(cfg module.Configurator) {}

// InitGenesis performs genesis initialization for the ibc-router module. It returns
// no validator updates.
func (am AppModule) InitGenesis(
	ctx sdk.Context,
	cdc codec.JSONCodec,
	data json.RawMessage,
) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the exported genesis state as raw bytes for the swap module.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	return nil
}

// ConsensusVersion returns the consensus state breaking version for the swap module.
func (am AppModule) ConsensusVersion() uint64 { return 1 }

// GenerateGenesisState implements the AppModuleSimulation interface.
func (am AppModule) GenerateGenesisState(simState *module.SimulationState) {}

// ProposalContents implements the AppModuleSimulation interface.
func (am AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RegisterStoreDecoder implements the AppModuleSimulation interface.
func (am AppModule) RegisterStoreDecoder(sdr simtypes.StoreDecoderRegistry) {}

// WeightedOperations implements the AppModuleSimulation interface.
func (am AppModule) WeightedOperations(_ module.SimulationState) []simtypes.WeightedOperation {
	return nil
}
