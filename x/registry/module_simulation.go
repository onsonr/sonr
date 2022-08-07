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
	opWeightMsgCreateWhoIs          = "op_weight_msg_who_is"
	defaultWeightMsgCreateWhoIs int = 100

	opWeightMsgUpdateWhoIs          = "op_weight_msg_who_is"
	defaultWeightMsgUpdateWhoIs int = 60

	opWeightMsgDeactivateWhoIs          = "op_weight_msg_who_is"
	defaultWeightMsgDeactivateWhoIs int = 50

	opWeightMsgBuyAlias = "op_weight_msg_buy_alias"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBuyAlias int = 90

	opWeightMsgSellAlias = "op_weight_msg_sell_alias"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSellAlias int = 80

	opWeightMsgTransferAlias = "op_weight_msg_transfer_alias"
	// TODO: Determine the simulation weight value
	defaultWeightMsgTransferAlias int = 70

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	// Define initial state for each module account
	accs := make([]string, len(simState.Accounts))
	whoIsList := make([]types.WhoIs, len(simState.Accounts))

	// Iterate over all accounts
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
		whoIsList[i], _ = registrysimulation.CreateMockWhoIs(acc)
	}
	registryGenesis := types.GenesisState{
		Params:     types.DefaultParams(),
		PortId:     types.PortID,
		WhoIsList:  whoIsList,
		WhoIsCount: uint64(len(whoIsList)),
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

	var weightMsgBuyAlias int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSellAlias, &weightMsgBuyAlias, nil,
		func(_ *rand.Rand) {
			weightMsgBuyAlias = defaultWeightMsgSellAlias
		},
	)

	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBuyAlias,
		registrysimulation.SimulateMsgBuyAlias(am.accountKeeper, am.bankKeeper, am.keeper),
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

	var weightMsgSellAlias int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSellAlias, &weightMsgSellAlias, nil,
		func(_ *rand.Rand) {
			weightMsgSellAlias = defaultWeightMsgSellAlias
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSellAlias,
		registrysimulation.SimulateMsgSellAlias(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgTransferAlias int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgTransferAlias, &weightMsgTransferAlias, nil,
		func(_ *rand.Rand) {
			weightMsgTransferAlias = defaultWeightMsgTransferAlias
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgTransferAlias,
		registrysimulation.SimulateMsgTransferAlias(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeactivateWhoIs int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeactivateWhoIs, &weightMsgDeactivateWhoIs, nil,
		func(_ *rand.Rand) {
			weightMsgDeactivateWhoIs = defaultWeightMsgDeactivateWhoIs
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeactivateWhoIs,
		registrysimulation.SimulateMsgDeactivateWhoIs(am.accountKeeper, am.bankKeeper, am.keeper),
	))
	// this line is used by starport scaffolding # simapp/module/operation

	return make([]simtypes.WeightedOperation, 0)
}
