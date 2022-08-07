package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/sonr-io/sonr/x/schema/keeper"
	"github.com/sonr-io/sonr/x/schema/types"
)

func SimulateMsgCreateScehma(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		createMsg := types.MsgCreateSchema{
			Definition: &types.SchemaDefinition{
				Creator: simAccount.Address.String(),
				Label:   "test schema",
				Fields:  make([]*types.SchemaKindDefinition, 0),
			},
		}

		createMsg.Definition.Fields = append(createMsg.Definition.Fields, &types.SchemaKindDefinition{
			Name:  "message",
			Field: types.SchemaKind_STRING,
		})
		createMsg.Definition.Fields = append(createMsg.Definition.Fields, &types.SchemaKindDefinition{
			Name:  "Icon",
			Field: types.SchemaKind_INT,
		})
		createMsg.Definition.Fields = append(createMsg.Definition.Fields, &types.SchemaKindDefinition{
			Name:  "type",
			Field: types.SchemaKind_STRING,
		})

		createMsg.Definition.Fields = append(createMsg.Definition.Fields, &types.SchemaKindDefinition{
			Name:     "comment",
			Field:    types.SchemaKind_LINK,
			LinkKind: types.LinkKind_SCHEMA,
			Link:     "QmZcGZYuoff9BQSqhzR9aqWfQBHU6bCMzKH7u25xZAijZB",
		})

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             &createMsg,
			MsgType:         createMsg.Type(),
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
