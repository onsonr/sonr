package keeper

import (
	didv1 "github.com/onsonr/sonr/api/did/v1"
	"github.com/onsonr/sonr/x/did/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) convertProfileToVerification() {
}

func (k Keeper) insertAuthenticationFromCredential() {
}

func (k Keeper) insertControllerFromMotrVault() {
}

func (k Keeper) insertDelegationFromAccount(ctx sdk.Context, address string, label string) (*didv1.Delegation, error) {
	del, err := k.OrmDB.DelegationTable().GetByControllerAccountLabel(ctx, address, label)
	if err != nil {
		return nil, err
	}
	return del, nil
}

func (k Keeper) insertServiceRecord(ctx sdk.Context, msg *types.MsgRegisterService) error {
	record, err := msg.ExtractServiceRecord()
	if err != nil {
		return err
	}
	err = k.OrmDB.ServiceRecordTable().Insert(ctx, record)
	if err != nil {
		return err
	}
	return nil
}
