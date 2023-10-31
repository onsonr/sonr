package service

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/sonrhq/sonr/testutil/sample"
	servicesimulation "github.com/sonrhq/sonr/x/service/simulation"
	"github.com/sonrhq/sonr/x/service/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = servicesimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateServiceRecord = "op_weight_msg_service_record"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateServiceRecord int = 100

	opWeightMsgUpdateServiceRecord = "op_weight_msg_service_record"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateServiceRecord int = 100

	opWeightMsgDeleteServiceRecord = "op_weight_msg_service_record"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteServiceRecord int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	serviceGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		ServiceRecordList: []types.ServiceRecord{
			{
				Id:         "0",
				Controller: sample.AccAddress(),
			},
			{
				Id:         "1",
				Controller: sample.AccAddress(),
			},
		},
		ServiceRecordCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&serviceGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateServiceRecord int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateServiceRecord, &weightMsgCreateServiceRecord, nil,
		func(_ *rand.Rand) {
			weightMsgCreateServiceRecord = defaultWeightMsgCreateServiceRecord
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateServiceRecord,
		servicesimulation.SimulateMsgCreateServiceRecord(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateServiceRecord int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateServiceRecord, &weightMsgUpdateServiceRecord, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateServiceRecord = defaultWeightMsgUpdateServiceRecord
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateServiceRecord,
		servicesimulation.SimulateMsgUpdateServiceRecord(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteServiceRecord int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteServiceRecord, &weightMsgDeleteServiceRecord, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteServiceRecord = defaultWeightMsgDeleteServiceRecord
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteServiceRecord,
		servicesimulation.SimulateMsgDeleteServiceRecord(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateServiceRecord,
			defaultWeightMsgCreateServiceRecord,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				servicesimulation.SimulateMsgCreateServiceRecord(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateServiceRecord,
			defaultWeightMsgUpdateServiceRecord,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				servicesimulation.SimulateMsgUpdateServiceRecord(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteServiceRecord,
			defaultWeightMsgDeleteServiceRecord,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				servicesimulation.SimulateMsgDeleteServiceRecord(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
