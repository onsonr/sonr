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
	opWeightMsgDefineBucket = "op_weight_msg_where_is"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDefineBucket int = 100

	opWeightMsgUpdateBucket = "op_weight_msg_where_is"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateBucket int = 100

	opWeightMsgDeleteBucket = "op_weight_msg_where_is"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteBucket int = 100

	opWeightMsgGenerateBucket = "op_weight_msg_generate_bucket"
	// TODO: Determine the simulation weight value
	defaultWeightMsgGenerateBucket int = 100

	opWeightMsgDeactivateBucket = "op_weight_msg_deactivate_bucket"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeactivateBucket int = 100

	opWeightMsgBurnBucket = "op_weight_msg_burn_bucket"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBurnBucket int = 100

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
		BucketList: []types.Bucket{
			{
				Uuid:    "did:sonr:1",
				Creator: sample.AccAddress(),
			},
			{
				Uuid:    "did:sonr:2",
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

	var weightMsgDefineBucket int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDefineBucket, &weightMsgDefineBucket, nil,
		func(_ *rand.Rand) {
			weightMsgDefineBucket = defaultWeightMsgDefineBucket
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDefineBucket,
		bucketsimulation.SimulateMsgDefineBucket(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgGenerateBucket int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgGenerateBucket, &weightMsgGenerateBucket, nil,
		func(_ *rand.Rand) {
			weightMsgGenerateBucket = defaultWeightMsgGenerateBucket
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgGenerateBucket,
		bucketsimulation.SimulateMsgGenerateBucket(am.accountKeeper, am.bankKeeper, am.keeper),
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

	var weightMsgBurnBucket int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgBurnBucket, &weightMsgBurnBucket, nil,
		func(_ *rand.Rand) {
			weightMsgBurnBucket = defaultWeightMsgBurnBucket
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBurnBucket,
		bucketsimulation.SimulateMsgBurnBucket(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
