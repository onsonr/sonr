package identity

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/sonr-io/sonr/testutil/sample"
	identitysimulation "github.com/sonr-io/sonr/x/identity/simulation"
	"github.com/sonr-io/sonr/x/identity/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = identitysimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateDidDocument = "op_weight_msg_did_document"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateDidDocument int = 100

	opWeightMsgUpdateDidDocument = "op_weight_msg_did_document"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateDidDocument int = 100

	opWeightMsgDeleteDidDocument = "op_weight_msg_did_document"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteDidDocument int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	identityGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		DidDocumentList: []types.DidDocument{
			{
				Creator: sample.AccAddress(),
				ID:     "0",
			},
			{
				Creator: sample.AccAddress(),
				ID:     "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&identityGenesis)
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

	var weightMsgCreateDidDocument int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateDidDocument, &weightMsgCreateDidDocument, nil,
		func(_ *rand.Rand) {
			weightMsgCreateDidDocument = defaultWeightMsgCreateDidDocument
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateDidDocument,
		identitysimulation.SimulateMsgCreateDidDocument(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateDidDocument int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateDidDocument, &weightMsgUpdateDidDocument, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateDidDocument = defaultWeightMsgUpdateDidDocument
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateDidDocument,
		identitysimulation.SimulateMsgUpdateDidDocument(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteDidDocument int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteDidDocument, &weightMsgDeleteDidDocument, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteDidDocument = defaultWeightMsgDeleteDidDocument
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteDidDocument,
		identitysimulation.SimulateMsgDeleteDidDocument(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
