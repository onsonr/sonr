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
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		createMsg := types.MsgCreateSchema{
			Creator: simAccount.Address.String(),
			Label:   "test schema",
			Fields:  make([]*types.SchemaField, 0),
		}

		createMsg.Fields = append(createMsg.Fields, &types.SchemaField{
			Name: "message",
			FieldKind: &types.SchemaFieldKind{
				Kind: types.Kind_STRING,
			},
		})
		createMsg.Fields = append(createMsg.Fields, &types.SchemaField{
			Name: "Icon",
			FieldKind: &types.SchemaFieldKind{
				Kind: types.Kind_INT,
			},
		})
		createMsg.Fields = append(createMsg.Fields, &types.SchemaField{
			Name: "type",
			FieldKind: &types.SchemaFieldKind{
				Kind: types.Kind_STRING,
			},
		})

		createMsg.Fields = append(createMsg.Fields, &types.SchemaField{
			Name: "comment",
			FieldKind: &types.SchemaFieldKind{
				Kind:    types.Kind_LINK,
				LinkDid: "QmZcGZYuoff9BQSqhzR9aqWfQBHU6bCMzKH7u25xZAijZB",
			},
		})

		createMsg.Fields = append(createMsg.Fields, &types.SchemaField{
			Name: "attributes",
			FieldKind: &types.SchemaFieldKind{
				Kind: types.Kind_LIST,
				ListKind: &types.SchemaFieldKind{
					Kind: types.Kind_STRING,
				},
			},
		})

		createMsg.Metadata = []*types.MetadataDefintion{
			{
				Key:   "image",
				Value: `<meta property="og:image" content="https://ahrefs.com/blog/wp-content/uploads/2019/12/fb-how-to-become-an-seo-expert.png" />`,
			},
		}

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
