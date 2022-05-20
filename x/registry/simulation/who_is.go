package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/sonr-io/sonr/x/registry/keeper"
	"github.com/sonr-io/sonr/x/registry/types"
)

func SimulateMsgCreateWhoIs(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)

		// Creates a mock did document for the provided simulated account
		doc, err := CreateMockDidDocument(simAccount)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, "createWhoIs", "failed to create mock did document"), nil, err
		}

		// Marshal Json document
		docBytes, err := doc.MarshalJSON()
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, "createWhoIs", "failed to marshal json document"), nil, err
		}

		msg := &types.MsgCreateWhoIs{
			Creator:     simAccount.Address.String(),
			DidDocument: docBytes,
			WhoisType:   types.WhoIsType_USER,
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgUpdateWhoIs(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount = simtypes.Account{}
			whoIs      = types.WhoIs{}
			msg        = &types.MsgUpdateWhoIs{}
			allWhoIs   = k.GetAllWhoIs(ctx)
			found      = false
		)
		for _, obj := range allWhoIs {
			simAccount, found = FindAccount(accs, obj.Owner)
			if found {
				whoIs = obj
				break
			}
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "whoIs owner not found"), nil, nil
		}
		msg.Creator = simAccount.Address.String()
		msg.DidDocument = whoIs.DidDocument
		msg.Did = whoIs.Owner

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgDeleteWhoIs(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount = simtypes.Account{}
			whoIs      = types.WhoIs{}
			msg        = &types.MsgUpdateWhoIs{}
			allWhoIs   = k.GetAllWhoIs(ctx)
			found      = false
		)
		for _, obj := range allWhoIs {
			simAccount, found = FindAccount(accs, obj.Owner)
			if found {
				whoIs = obj
				break
			}
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "whoIs owner not found"), nil, nil
		}
		msg.Creator = simAccount.Address.String()
		msg.Did = whoIs.Owner
		msg.DidDocument = whoIs.DidDocument

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}
