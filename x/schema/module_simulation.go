package schema

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/sonr-io/sonr/testutil/sample"
	schemasimulation "github.com/sonr-io/sonr/x/schema/simulation"
	"github.com/sonr-io/sonr/x/schema/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = schemasimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	// this line is used by starport scaffolding # simapp/module/const
	opWeightMsgCreateSchema          = "op_weight_msg_create_schema"
	defaultWeightMsgCreateSchema int = 100
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	schemaGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&schemaGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	// this line is used by starport scaffolding # simapp/module/operation
	var weightMsgCreateSchema int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateSchema, &weightMsgCreateSchema, nil,
		func(_ *rand.Rand) {
			weightMsgCreateSchema = defaultWeightMsgCreateSchema
		})

	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateSchema,
		schemasimulation.SimulateMsgCreateScehma(am.accountKeeper, am.bankKeeper, am.keeper)))

	return operations
}
