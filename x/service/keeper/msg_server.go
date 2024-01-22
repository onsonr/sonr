package keeper

import (
	"context"
	"fmt"
	"strings"

	modulev1 "github.com/sonrhq/sonr/api/service/module/v1"
	"github.com/sonrhq/sonr/x/service"
)

type msgServer struct {
	k Keeper
}

var _ service.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the module MsgServer interface.
func NewMsgServerImpl(keeper Keeper) service.MsgServer {
	return &msgServer{k: keeper}
}

// UpdateParams params is defining the handler for the MsgUpdateParams message.
func (ms msgServer) UpdateParams(ctx context.Context, msg *service.MsgUpdateParams) (*service.MsgUpdateParamsResponse, error) {
	if _, err := ms.k.addressCodec.StringToBytes(msg.Authority); err != nil {
		return nil, fmt.Errorf("invalid authority address: %w", err)
	}

	if authority := ms.k.GetAuthority(); !strings.EqualFold(msg.Authority, authority) {
		return nil, fmt.Errorf("unauthorized, authority does not match the module's authority: got %s, want %s", msg.Authority, authority)
	}

	if err := msg.Params.Validate(); err != nil {
		return nil, err
	}

	if err := ms.k.Params.Set(ctx, msg.Params); err != nil {
		return nil, err
	}

	return &service.MsgUpdateParamsResponse{}, nil
}

// CreateRecord params is defining the handler for the MsgCreateRecord message.
func (ms msgServer) CreateRecord(ctx context.Context, msg *service.MsgCreateRecord) (*service.MsgCreateRecordResponse, error) {
	err := ms.k.db.ServiceRecordTable().Insert(ctx, &modulev1.ServiceRecord{
		Name:        msg.Name,
		Origin:      msg.Origin,
		Description: msg.Description,
		Controller:  msg.Owner,
		Permissions: modulev1.ServicePermissions_SERVICE_PERMISSIONS_OWN,
	})
	if err != nil {
		return nil, err
	}
	return &service.MsgCreateRecordResponse{}, nil
}

// UpdateRecord params is defining the handler for the MsgUpdateRecord message.
func (ms msgServer) UpdateRecord(ctx context.Context, msg *service.MsgUpdateRecord) (*service.MsgUpdateRecordResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

// DeleteRecord params is defining the handler for the MsgDeleteRecord message.
func (ms msgServer) DeleteRecord(ctx context.Context, msg *service.MsgDeleteRecord) (*service.MsgDeleteRecordResponse, error) {
	rec, err := ms.k.db.ServiceRecordTable().Get(ctx, msg.RecordId)
	if err != nil {
		return nil, err
	}
	if rec == nil {
		return nil, fmt.Errorf("record does not exist")
	}
	if rec.Controller != msg.Owner {
		return nil, fmt.Errorf("unauthorized, record owner does not match the module's authority: got %s, want %s", msg.Owner, rec.Controller)
	}

	if err := ms.k.db.ServiceRecordTable().Delete(ctx, rec); err != nil {
		return nil, err
	}
	return &service.MsgDeleteRecordResponse{}, nil
}
