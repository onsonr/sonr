package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onsonr/crypto/mpc"
	"github.com/onsonr/sonr/x/did/types"
)

func (k Keeper) NewController(ctx sdk.Context) (uint64, types.ControllerI, error) {
	shares, err := mpc.GenerateKeyshares()
	if err != nil {
		return 0, nil, err
	}
	controller, err := types.NewController(shares)
	if err != nil {
		return 0, nil, err
	}
	entry, err := controller.GetTableEntry()
	if err != nil {
		return 0, nil, err
	}
	num, err := k.OrmDB.ControllerTable().InsertReturningNumber(ctx, entry)
	if err != nil {
		return 0, nil, err
	}
	return num, controller, nil
}
