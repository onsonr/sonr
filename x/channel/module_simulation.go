package channel

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/sonr-io/sonr/internal/blockchain/testutil/sample"
	channelsimulation "github.com/sonr-io/sonr/internal/blockchain/x/channel/simulation"
	"github.com/sonr-io/sonr/internal/blockchain/x/channel/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = channelsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateChannel = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateChannel int = 100

	opWeightMsgReadChannel = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgReadChannel int = 100

	opWeightMsgDeactivateChannel = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeactivateChannel int = 100

	opWeightMsgListenChannel = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgListenChannel int = 100

	opWeightMsgUpdateChannel = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateChannel int = 100

	opWeightMsgCreateHowIs = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateHowIs int = 100

	opWeightMsgUpdateHowIs = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateHowIs int = 100

	opWeightMsgDeleteHowIs = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteHowIs int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	channelGenesis := types.GenesisState{
		HowIsList: []types.HowIs{
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
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&channelGenesis)
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

	var weightMsgCreateChannel int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateChannel, &weightMsgCreateChannel, nil,
		func(_ *rand.Rand) {
			weightMsgCreateChannel = defaultWeightMsgCreateChannel
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateChannel,
		channelsimulation.SimulateMsgCreateChannel(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeactivateChannel int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeactivateChannel, &weightMsgDeactivateChannel, nil,
		func(_ *rand.Rand) {
			weightMsgDeactivateChannel = defaultWeightMsgDeactivateChannel
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeactivateChannel,
		channelsimulation.SimulateMsgDeactivateChannel(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateChannel int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateChannel, &weightMsgUpdateChannel, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateChannel = defaultWeightMsgUpdateChannel
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateChannel,
		channelsimulation.SimulateMsgUpdateChannel(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateHowIs int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateHowIs, &weightMsgCreateHowIs, nil,
		func(_ *rand.Rand) {
			weightMsgCreateHowIs = defaultWeightMsgCreateHowIs
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateHowIs,
		channelsimulation.SimulateMsgCreateHowIs(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateHowIs int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateHowIs, &weightMsgUpdateHowIs, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateHowIs = defaultWeightMsgUpdateHowIs
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateHowIs,
		channelsimulation.SimulateMsgUpdateHowIs(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteHowIs int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteHowIs, &weightMsgDeleteHowIs, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteHowIs = defaultWeightMsgDeleteHowIs
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteHowIs,
		channelsimulation.SimulateMsgDeleteHowIs(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
