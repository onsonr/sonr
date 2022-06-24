package simulation

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/sonr-io/sonr/x/schema/keeper"
	"github.com/sonr-io/sonr/x/schema/types"
)

func SimulateMsgDeprecateSchema(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		what_is, err := CreatMockWhatIs(simAccount)

		if err != nil {
			k.Logger(ctx).Error("Could not create what is for simulation")
		}

		k.SetWhatIs(ctx, what_is)
		//Need to ensure WhatIs exists
		wi, foundWi := k.GetWhatIsFromCreator(ctx, simAccount.Address.String())
		if !foundWi || len(wi) < 1 {
			fmt.Println("Wi not found") //testing only
		}

		deprMsg := types.MsgDeprecateSchema{
			Creator: simAccount.Address.String(),
			Did:     wi[0].GetDid(),
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             &deprMsg,
			MsgType:         types.TypeMsgDeprecateSchema, //deprMsg.Type(), //Figure out discrepancy
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
