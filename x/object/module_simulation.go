package object

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/sonr-io/sonr/internal/blockchain/testutil/sample"
	objectsimulation "github.com/sonr-io/sonr/internal/blockchain/x/object/simulation"
	"github.com/sonr-io/sonr/internal/blockchain/x/object/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = objectsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateObject = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateObject int = 100

	opWeightMsgReadObject = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgReadObject int = 100

	opWeightMsgUpdateObject = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateObject int = 100

	opWeightMsgDeactivateObject = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeactivateObject int = 100

	opWeightMsgCreateWhatIs = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateWhatIs int = 100

	opWeightMsgUpdateWhatIs = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateWhatIs int = 100

	opWeightMsgDeleteWhatIs = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteWhatIs int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	objectGenesis := types.GenesisState{
		WhatIsList: []types.WhatIs{
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
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&objectGenesis)
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

	var weightMsgCreateObject int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateObject, &weightMsgCreateObject, nil,
		func(_ *rand.Rand) {
			weightMsgCreateObject = defaultWeightMsgCreateObject
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateObject,
		objectsimulation.SimulateMsgCreateObject(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateObject int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateObject, &weightMsgUpdateObject, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateObject = defaultWeightMsgUpdateObject
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateObject,
		objectsimulation.SimulateMsgUpdateObject(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeactivateObject int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeactivateObject, &weightMsgDeactivateObject, nil,
		func(_ *rand.Rand) {
			weightMsgDeactivateObject = defaultWeightMsgDeactivateObject
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeactivateObject,
		objectsimulation.SimulateMsgDeactivateObject(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateWhatIs int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateWhatIs, &weightMsgCreateWhatIs, nil,
		func(_ *rand.Rand) {
			weightMsgCreateWhatIs = defaultWeightMsgCreateWhatIs
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateWhatIs,
		objectsimulation.SimulateMsgCreateWhatIs(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateWhatIs int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateWhatIs, &weightMsgUpdateWhatIs, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateWhatIs = defaultWeightMsgUpdateWhatIs
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateWhatIs,
		objectsimulation.SimulateMsgUpdateWhatIs(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteWhatIs int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteWhatIs, &weightMsgDeleteWhatIs, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteWhatIs = defaultWeightMsgDeleteWhatIs
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteWhatIs,
		objectsimulation.SimulateMsgDeleteWhatIs(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
