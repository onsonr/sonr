package simulation

import (
	"math/rand"
	"strconv"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/sonrhq/core/x/domain/keeper"
	"github.com/sonrhq/core/x/domain/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func SimulateMsgCreateSLDRecord(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)

		sld := types.SLDRecord{
			Creator: simAccount.Address.String(),
			Index:   strconv.Itoa(r.Intn(100)),
		}
		msg := &types.MsgCreateSLDRecord{
			Creator: simAccount.Address.String(),
			SldRecord: &sld,
		}

		_, found := k.GetSLDRecord(ctx, msg.SldRecord.Index)
		if found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "SLDRecord already exist"), nil, nil
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

func SimulateMsgUpdateSLDRecord(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount   = simtypes.Account{}
			sLDRecord    = types.SLDRecord{}
			msg          = &types.MsgUpdateSLDRecord{}
			allSLDRecord = k.GetAllSLDRecord(ctx)
			found        = false
		)
		for _, obj := range allSLDRecord {
			simAccount, found = FindAccount(accs, obj.Creator)
			if found {
				sLDRecord = obj
				break
			}
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "sLDRecord creator not found"), nil, nil
		}
		msg.Creator = simAccount.Address.String()

		msg.SldRecord.Index = sLDRecord.Index

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

func SimulateMsgDeleteSLDRecord(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount   = simtypes.Account{}
			sLDRecord    = types.SLDRecord{}
			msg          = &types.MsgUpdateSLDRecord{}
			allSLDRecord = k.GetAllSLDRecord(ctx)
			found        = false
		)
		for _, obj := range allSLDRecord {
			simAccount, found = FindAccount(accs, obj.Creator)
			if found {
				sLDRecord = obj
				break
			}
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "sLDRecord creator not found"), nil, nil
		}
		msg.Creator = simAccount.Address.String()

		msg.SldRecord.Index = sLDRecord.Index

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
