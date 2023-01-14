package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/sonr-hq/sonr/x/identity/keeper"
	"github.com/sonr-hq/sonr/x/identity/types"
)

func SimulateMsgCreateAccount(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCreateAccount{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CreateAccount simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CreateAccount simulation not implemented"), nil, nil
	}
}
