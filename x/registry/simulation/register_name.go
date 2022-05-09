package simulation

import (
	"encoding/hex"
	"math/rand"
	"strconv"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/duo-labs/webauthn/webauthn"
	"github.com/sonr-io/sonr/x/registry/keeper"
	"github.com/sonr-io/sonr/x/registry/types"
)

func SimulateMsgRegisterName(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		// Example ECDSA key
		// https://datatracker.ietf.org/doc/html/rfc8392#appendix-A.2.3
		cborData, _ := hex.DecodeString("a72358206c1382765aec5358f117733d281c1c7bdc39884d04a45a1e6c67c858bc206c1922582060f7f1a780d8a783bfb7a2dd6b2796e8128dbbcef9d3d168db9529971a36e7b9215820143329cce7868e416927599cf65a34f3ce2ffda55a7eca69ed8919a394d42f0f2001010202524173796d6d657472696345434453413235360326")
		msg := types.NewMsgRegisterName(
			simAccount.Address.String(),
			"test"+strconv.Itoa(r.Int()),
			webauthn.Credential{
				PublicKey: cborData,
			},
		)

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

		// TODO: Handling the RegisterName simulation

		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}
