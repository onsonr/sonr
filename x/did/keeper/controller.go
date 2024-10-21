package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onsonr/crypto/mpc"
	"github.com/onsonr/sonr/x/did/types"
)

func (k Keeper) NewController(ctx sdk.Context) (types.ControllerI, error) {
	shares, err := mpc.GenerateKeyshares()
	if err != nil {
		return nil, err
	}
	controller, err := types.NewController(shares)
	if err != nil {
		return nil, err
	}
	return controller, nil
}

func (k Keeper) ResolveController(ctx sdk.Context, did string) (types.ControllerI, error) {
	ct, err := k.OrmDB.ControllerTable().GetByDid(ctx, did)
	if err != nil {
		return nil, err
	}
	c, err := types.LoadControllerFromTableEntry(ctx, ct)
	if err != nil {
		return nil, err
	}
	return c, nil
}
