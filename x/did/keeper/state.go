package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/onsonr/sonr/x/did/builder"
	"github.com/onsonr/sonr/x/did/types"
)

// insertService inserts a service record into the database
func (k Keeper) insertService(
	ctx sdk.Context,
	svc *types.Service,
) (*types.MsgRegisterServiceResponse, error) {
	record := builder.APIFormatServiceRecord(svc)
	err := k.OrmDB.ServiceRecordTable().Insert(ctx, record)
	if err != nil {
		return nil, err
	}
	return &types.MsgRegisterServiceResponse{
		Success: true,
		Did:     record.Id,
	}, nil
func (k Keeper) insertAliasFromDisplayName() {
}

func (k Keeper) insertAssertionFromIdentity() {
}

func (k Keeper) insertAuthenticationFromCredential() {
}

func (k Keeper) insertControllerFromMotrVault() {
}

func (k Keeper) insertDelegationFromAccount() {
}
