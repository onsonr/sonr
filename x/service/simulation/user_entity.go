package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/sonrhq/core/x/service/keeper"
	"github.com/sonrhq/core/x/service/types"
)

func SimulateMsgRegisterUserEntity(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgRegisterUserEntity{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the RegisterUserEntity simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "RegisterUserEntity simulation not implemented"), nil, nil
	}
}

func SimulateMsgAuthenticateUserEntity(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgAuthenticateUserEntity{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the AuthenticateUserEntity simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "AuthenticateUserEntity simulation not implemented"), nil, nil
	}
}
