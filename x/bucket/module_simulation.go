package bucket

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/sonr-io/sonr/testutil/sample"
	bucketsimulation "github.com/sonr-io/sonr/x/bucket/simulation"
	"github.com/sonr-io/sonr/x/bucket/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = bucketsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateWhereIs = "op_weight_msg_where_is"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateWhereIs int = 100

	opWeightMsgUpdateWhereIs = "op_weight_msg_where_is"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateWhereIs int = 100

	opWeightMsgDeleteWhereIs = "op_weight_msg_where_is"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteWhereIs int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	bucketGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		WhereIsList: []types.WhereIs{
			{
				Did:     "did:sonr:1",
				Creator: sample.AccAddress(),
			},
			{
				Did:     "did:sonr:2",
				Creator: sample.AccAddress(),
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&bucketGenesis)
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

	var weightMsgCreateWhereIs int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateWhereIs, &weightMsgCreateWhereIs, nil,
		func(_ *rand.Rand) {
			weightMsgCreateWhereIs = defaultWeightMsgCreateWhereIs
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateWhereIs,
		bucketsimulation.SimulateMsgCreateWhereIs(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateWhereIs int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateWhereIs, &weightMsgUpdateWhereIs, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateWhereIs = defaultWeightMsgUpdateWhereIs
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateWhereIs,
		bucketsimulation.SimulateMsgUpdateWhereIs(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteWhereIs int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteWhereIs, &weightMsgDeleteWhereIs, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteWhereIs = defaultWeightMsgDeleteWhereIs
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteWhereIs,
		bucketsimulation.SimulateMsgDeleteWhereIs(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
