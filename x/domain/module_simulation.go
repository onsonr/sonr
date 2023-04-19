package domain

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/sonrhq/core/testutil/sample"
	domainsimulation "github.com/sonrhq/core/x/domain/simulation"
	"github.com/sonrhq/core/x/domain/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = domainsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateTLDRecord = "op_weight_msg_tld_record"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateTLDRecord int = 100

	opWeightMsgUpdateTLDRecord = "op_weight_msg_tld_record"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateTLDRecord int = 100

	opWeightMsgDeleteTLDRecord = "op_weight_msg_tld_record"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteTLDRecord int = 100

	opWeightMsgCreateSLDRecord = "op_weight_msg_sld_record"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateSLDRecord int = 100

	opWeightMsgUpdateSLDRecord = "op_weight_msg_sld_record"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateSLDRecord int = 100

	opWeightMsgDeleteSLDRecord = "op_weight_msg_sld_record"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteSLDRecord int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	domainGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		TLDRecordList: []types.TLDRecord{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		SLDRecordList: []types.SLDRecord{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&domainGenesis)
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

	var weightMsgCreateTLDRecord int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateTLDRecord, &weightMsgCreateTLDRecord, nil,
		func(_ *rand.Rand) {
			weightMsgCreateTLDRecord = defaultWeightMsgCreateTLDRecord
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateTLDRecord,
		domainsimulation.SimulateMsgCreateTLDRecord(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateTLDRecord int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateTLDRecord, &weightMsgUpdateTLDRecord, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateTLDRecord = defaultWeightMsgUpdateTLDRecord
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateTLDRecord,
		domainsimulation.SimulateMsgUpdateTLDRecord(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteTLDRecord int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteTLDRecord, &weightMsgDeleteTLDRecord, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteTLDRecord = defaultWeightMsgDeleteTLDRecord
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteTLDRecord,
		domainsimulation.SimulateMsgDeleteTLDRecord(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateSLDRecord int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateSLDRecord, &weightMsgCreateSLDRecord, nil,
		func(_ *rand.Rand) {
			weightMsgCreateSLDRecord = defaultWeightMsgCreateSLDRecord
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateSLDRecord,
		domainsimulation.SimulateMsgCreateSLDRecord(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateSLDRecord int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateSLDRecord, &weightMsgUpdateSLDRecord, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateSLDRecord = defaultWeightMsgUpdateSLDRecord
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateSLDRecord,
		domainsimulation.SimulateMsgUpdateSLDRecord(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteSLDRecord int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteSLDRecord, &weightMsgDeleteSLDRecord, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteSLDRecord = defaultWeightMsgDeleteSLDRecord
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteSLDRecord,
		domainsimulation.SimulateMsgDeleteSLDRecord(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
