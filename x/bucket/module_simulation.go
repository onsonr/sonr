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
	opWeightMsgCreateBucket = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateBucket int = 100

	opWeightMsgReadBucket = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgReadBucket int = 100

	opWeightMsgUpdateBucket = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateBucket int = 100

	opWeightMsgDeactivateBucket = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeactivateBucket int = 100

	opWeightMsgListenBucket = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgListenBucket int = 100

	opWeightMsgCreateWhichIs = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateWhichIs int = 100

	opWeightMsgUpdateWhichIs = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateWhichIs int = 100

	opWeightMsgDeleteWhichIs = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteWhichIs int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	bucketGenesis := types.GenesisState{
		WhichIsList: []types.WhichIs{
			{
				Creator: sample.AccAddress(),
				Did:     "0",
			},
			{
				Creator: sample.AccAddress(),
				Did:     "1",
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

	var weightMsgCreateBucket int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateBucket, &weightMsgCreateBucket, nil,
		func(_ *rand.Rand) {
			weightMsgCreateBucket = defaultWeightMsgCreateBucket
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateBucket,
		bucketsimulation.SimulateMsgCreateBucket(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateBucket int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateBucket, &weightMsgUpdateBucket, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateBucket = defaultWeightMsgUpdateBucket
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateBucket,
		bucketsimulation.SimulateMsgUpdateBucket(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeactivateBucket int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeactivateBucket, &weightMsgDeactivateBucket, nil,
		func(_ *rand.Rand) {
			weightMsgDeactivateBucket = defaultWeightMsgDeactivateBucket
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeactivateBucket,
		bucketsimulation.SimulateMsgDeactivateBucket(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateWhichIs int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateWhichIs, &weightMsgCreateWhichIs, nil,
		func(_ *rand.Rand) {
			weightMsgCreateWhichIs = defaultWeightMsgCreateWhichIs
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateWhichIs,
		bucketsimulation.SimulateMsgCreateWhichIs(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateWhichIs int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateWhichIs, &weightMsgUpdateWhichIs, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateWhichIs = defaultWeightMsgUpdateWhichIs
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateWhichIs,
		bucketsimulation.SimulateMsgUpdateWhichIs(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteWhichIs int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteWhichIs, &weightMsgDeleteWhichIs, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteWhichIs = defaultWeightMsgDeleteWhichIs
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteWhichIs,
		bucketsimulation.SimulateMsgDeleteWhichIs(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
