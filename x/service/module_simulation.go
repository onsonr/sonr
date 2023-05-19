package service

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/sonrhq/core/testutil/sample"
	servicesimulation "github.com/sonrhq/core/x/service/simulation"
	"github.com/sonrhq/core/x/service/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = servicesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
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

	opWeightMsgCreateServiceRelationships = "op_weight_msg_service_relationships"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateServiceRelationships int = 100

	opWeightMsgUpdateServiceRelationships = "op_weight_msg_service_relationships"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateServiceRelationships int = 100

	opWeightMsgDeleteServiceRelationships = "op_weight_msg_service_relationships"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteServiceRelationships int = 100

	opWeightMsgRegisterUserEntity = "op_weight_msg_register_user_entity"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterUserEntity int = 100

	opWeightMsgAuthenticateUserEntity = "op_weight_msg_authenticate_user_entity"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAuthenticateUserEntity int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	serviceGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		ServiceRecordList: []types.ServiceRecord{
			{
				Controller: sample.AccAddress(),
				Id:         "0",
			},
			{
				Controller: sample.AccAddress(),
				Id:         "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&serviceGenesis)
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
	var weightMsgRegisterUserEntity int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRegisterUserEntity, &weightMsgRegisterUserEntity, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterUserEntity = defaultWeightMsgRegisterUserEntity
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterUserEntity,
		servicesimulation.SimulateMsgRegisterUserEntity(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAuthenticateUserEntity int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAuthenticateUserEntity, &weightMsgAuthenticateUserEntity, nil,
		func(_ *rand.Rand) {
			weightMsgAuthenticateUserEntity = defaultWeightMsgAuthenticateUserEntity
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAuthenticateUserEntity,
		servicesimulation.SimulateMsgAuthenticateUserEntity(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
