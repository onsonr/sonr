package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/sonrhq/core/x/identity/keeper"
	"github.com/sonrhq/core/x/identity/types"
)

func SimulateMsgImportPublicKey(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgImportPublicKey{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the ImportPublicKey simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "ImportPublicKey simulation not implemented"), nil, nil
	}
}
