package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/sonr-io/sonr/x/registry/keeper"
	"github.com/sonr-io/sonr/x/registry/types"
)

func SimulateMsgUpdateName(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgUpdateName{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the UpdateName simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "UpdateName simulation not implemented"), nil, nil
	}
}
