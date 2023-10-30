package identity

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
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
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateDIDDocument = "op_weight_msg_did_document"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateDIDDocument int = 100

	opWeightMsgUpdateDIDDocument = "op_weight_msg_did_document"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateDIDDocument int = 100

	opWeightMsgDeleteDIDDocument = "op_weight_msg_did_document"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteDIDDocument int = 100

	opWeightMsgCreateControllerAccount = "op_weight_msg_controller_account"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateControllerAccount int = 100

	opWeightMsgUpdateControllerAccount = "op_weight_msg_controller_account"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateControllerAccount int = 100

	opWeightMsgDeleteControllerAccount = "op_weight_msg_controller_account"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteControllerAccount int = 100

	opWeightMsgCreateEscrowAccount = "op_weight_msg_escrow_account"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateEscrowAccount int = 100

	opWeightMsgUpdateEscrowAccount = "op_weight_msg_escrow_account"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateEscrowAccount int = 100

	opWeightMsgDeleteEscrowAccount = "op_weight_msg_escrow_account"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteEscrowAccount int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	identityGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		DIDDocumentList: []types.DIDDocument{
			{
				Id: sample.AccAddress(),
			},
			{
				Id: sample.AccAddress(),
			},
		},
		ControllerAccountList: []types.ControllerAccount{
			{

				Address: sample.AccAddress(),
			},
			{

				Address: sample.AccAddress(),
			},
		},
		ControllerAccountCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&identityGenesis)
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

	var weightMsgCreateDIDDocument int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateDIDDocument, &weightMsgCreateDIDDocument, nil,
		func(_ *rand.Rand) {
			weightMsgCreateDIDDocument = defaultWeightMsgCreateDIDDocument
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateDIDDocument,
		identitysimulation.SimulateMsgCreateDidDocument(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateDIDDocument int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateDIDDocument, &weightMsgUpdateDIDDocument, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateDIDDocument = defaultWeightMsgUpdateDIDDocument
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateDIDDocument,
		identitysimulation.SimulateMsgUpdateDidDocument(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteDIDDocument int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteDIDDocument, &weightMsgDeleteDIDDocument, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteDIDDocument = defaultWeightMsgDeleteDIDDocument
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteDIDDocument,
		identitysimulation.SimulateMsgRegisterIdentity(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateControllerAccount int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateControllerAccount, &weightMsgCreateControllerAccount, nil,
		func(_ *rand.Rand) {
			weightMsgCreateControllerAccount = defaultWeightMsgCreateControllerAccount
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateControllerAccount,
		identitysimulation.SimulateMsgCreateControllerAccount(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateControllerAccount int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateControllerAccount, &weightMsgUpdateControllerAccount, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateControllerAccount = defaultWeightMsgUpdateControllerAccount
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateControllerAccount,
		identitysimulation.SimulateMsgUpdateControllerAccount(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteControllerAccount int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteControllerAccount, &weightMsgDeleteControllerAccount, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteControllerAccount = defaultWeightMsgDeleteControllerAccount
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteControllerAccount,
		identitysimulation.SimulateMsgDeleteControllerAccount(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateDIDDocument,
			defaultWeightMsgCreateDIDDocument,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				identitysimulation.SimulateMsgCreateDidDocument(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateDIDDocument,
			defaultWeightMsgUpdateDIDDocument,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				identitysimulation.SimulateMsgUpdateDidDocument(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteDIDDocument,
			defaultWeightMsgDeleteDIDDocument,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				identitysimulation.SimulateMsgRegisterIdentity(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateControllerAccount,
			defaultWeightMsgCreateControllerAccount,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				identitysimulation.SimulateMsgCreateControllerAccount(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateControllerAccount,
			defaultWeightMsgUpdateControllerAccount,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				identitysimulation.SimulateMsgUpdateControllerAccount(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteControllerAccount,
			defaultWeightMsgDeleteControllerAccount,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				identitysimulation.SimulateMsgDeleteControllerAccount(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
