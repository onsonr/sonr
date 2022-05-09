package registry

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/sonr-io/sonr/testutil/sample"
	registrysimulation "github.com/sonr-io/sonr/x/registry/simulation"
	"github.com/sonr-io/sonr/x/registry/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = registrysimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgRegisterApplication = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterApplication int = 100

	opWeightMsgRegisterName = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterName int = 100

	opWeightMsgAccessName = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAccessName int = 100

	opWeightMsgUpdateName = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateName int = 100

	opWeightMsgAccessApplication = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAccessApplication int = 100

	opWeightMsgUpdateApplication = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateApplication int = 100

	opWeightMsgCreateWhoIs = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateWhoIs int = 100

	opWeightMsgUpdateWhoIs = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateWhoIs int = 100

	opWeightMsgDeleteWhoIs = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteWhoIs int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	registryGenesis := types.GenesisState{
		WhoIsList: []types.WhoIs{
			{
				Creator: sample.AccAddress(),
				Name:    "0",
			},
			{
				Creator: sample.AccAddress(),
				Name:    "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&registryGenesis)
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

	var weightMsgRegisterApplication int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRegisterApplication, &weightMsgRegisterApplication, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterApplication = defaultWeightMsgRegisterApplication
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterApplication,
		registrysimulation.SimulateMsgRegisterApplication(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRegisterName int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRegisterName, &weightMsgRegisterName, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterName = defaultWeightMsgRegisterName
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterName,
		registrysimulation.SimulateMsgRegisterName(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAccessName int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAccessName, &weightMsgAccessName, nil,
		func(_ *rand.Rand) {
			weightMsgAccessName = defaultWeightMsgAccessName
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAccessName,
		registrysimulation.SimulateMsgAccessName(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateName int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateName, &weightMsgUpdateName, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateName = defaultWeightMsgUpdateName
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateName,
		registrysimulation.SimulateMsgUpdateName(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAccessApplication int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAccessApplication, &weightMsgAccessApplication, nil,
		func(_ *rand.Rand) {
			weightMsgAccessApplication = defaultWeightMsgAccessApplication
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAccessApplication,
		registrysimulation.SimulateMsgAccessApplication(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateApplication int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateApplication, &weightMsgUpdateApplication, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateApplication = defaultWeightMsgUpdateApplication
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateApplication,
		registrysimulation.SimulateMsgUpdateApplication(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateWhoIs int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateWhoIs, &weightMsgCreateWhoIs, nil,
		func(_ *rand.Rand) {
			weightMsgCreateWhoIs = defaultWeightMsgCreateWhoIs
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateWhoIs,
		registrysimulation.SimulateMsgCreateWhoIs(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateWhoIs int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateWhoIs, &weightMsgUpdateWhoIs, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateWhoIs = defaultWeightMsgUpdateWhoIs
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateWhoIs,
		registrysimulation.SimulateMsgUpdateWhoIs(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteWhoIs int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteWhoIs, &weightMsgDeleteWhoIs, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteWhoIs = defaultWeightMsgDeleteWhoIs
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteWhoIs,
		registrysimulation.SimulateMsgDeleteWhoIs(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
