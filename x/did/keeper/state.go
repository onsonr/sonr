package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	didv1 "github.com/onsonr/sonr/api/did/v1"
	"github.com/onsonr/sonr/x/did/types"
)

// insertService inserts a service record into the database
func (k Keeper) insertService(
	ctx sdk.Context,
	svc *types.MsgRegisterService,
) (*types.MsgRegisterServiceResponse, error) {
	record := didv1.ServiceRecord{
		Id: svc.OriginUri,
	}
	err := k.OrmDB.ServiceRecordTable().Insert(ctx, &record)
	if err != nil {
		return nil, err
	}
	return &types.MsgRegisterServiceResponse{
		Success: true,
		Did:     record.Id,
	}, nil
}
