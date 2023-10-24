package domain

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/sonr-io/core/testutil/sample"
	domainsimulation "github.com/sonr-io/core/x/domain/simulation"
	"github.com/sonr-io/core/x/domain/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = domainsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateUsernameRecords = "op_weight_msg_username_records"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateUsernameRecords int = 100

	opWeightMsgUpdateUsernameRecords = "op_weight_msg_username_records"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateUsernameRecords int = 100

	opWeightMsgDeleteUsernameRecords = "op_weight_msg_username_records"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteUsernameRecords int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	domainGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		UsernameRecordsList: []types.UsernameRecord{
			{
				Address: sample.AccAddress(),
				Index:   "0",
			},
			{
				Address: sample.AccAddress(),
				Index:   "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&domainGenesis)
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

	var weightMsgCreateUsernameRecords int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateUsernameRecords, &weightMsgCreateUsernameRecords, nil,
		func(_ *rand.Rand) {
			weightMsgCreateUsernameRecords = defaultWeightMsgCreateUsernameRecords
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateUsernameRecords,
		domainsimulation.SimulateMsgCreateUsernameRecords(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateUsernameRecords int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateUsernameRecords, &weightMsgUpdateUsernameRecords, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateUsernameRecords = defaultWeightMsgUpdateUsernameRecords
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateUsernameRecords,
		domainsimulation.SimulateMsgUpdateUsernameRecords(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteUsernameRecords int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteUsernameRecords, &weightMsgDeleteUsernameRecords, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteUsernameRecords = defaultWeightMsgDeleteUsernameRecords
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteUsernameRecords,
		domainsimulation.SimulateMsgDeleteUsernameRecords(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateUsernameRecords,
			defaultWeightMsgCreateUsernameRecords,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				domainsimulation.SimulateMsgCreateUsernameRecords(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateUsernameRecords,
			defaultWeightMsgUpdateUsernameRecords,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				domainsimulation.SimulateMsgUpdateUsernameRecords(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteUsernameRecords,
			defaultWeightMsgDeleteUsernameRecords,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				domainsimulation.SimulateMsgDeleteUsernameRecords(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
