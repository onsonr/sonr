package keeper

import (
	"context"
	"fmt"
	"strings"

	modulev1 "github.com/sonrhq/sonr/api/sonr/service/module/v1"
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
	// Get permissions from int32
	if msg.Authority == "" {
		return nil, fmt.Errorf("owner cannot be empty")
	}

	if msg.Name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}

	if msg.Origin == "" {
		return nil, fmt.Errorf("origin cannot be empty")
	}

	err := ms.k.db.ServiceRecordTable().Insert(ctx, &modulev1.ServiceRecord{
		Name:        msg.Name,
		Origin:      msg.Origin,
		Description: msg.Description,
		Controller:  msg.Authority,
		Permissions: getPermissionsFromInt32(msg.Permissions),
	})
	if err != nil {
		return nil, err
	}
	return &service.MsgCreateRecordResponse{}, nil
}

// UpdateRecord params is defining the handler for the MsgUpdateRecord message.
func (ms msgServer) UpdateRecord(ctx context.Context, msg *service.MsgUpdateRecord) (*service.MsgUpdateRecordResponse, error) {
	rec, err := ms.k.db.ServiceRecordTable().Get(ctx, msg.RecordId)
	if err != nil {
		return nil, err
	}
	if rec == nil {
		return nil, fmt.Errorf("record does not exist")
	}
	if rec.Controller != msg.Authority {
		return nil, fmt.Errorf("unauthorized, record owner does not match the module's authority: got %s, want %s", msg.Authority, rec.Controller)
	}
	err = ms.k.db.ServiceRecordTable().Update(ctx, &modulev1.ServiceRecord{
		Id:          msg.RecordId,
		Name:        msg.Name,
		Origin:      msg.Origin,
		Description: msg.Description,
		Controller:  msg.Authority,
		Permissions: getPermissionsFromInt32(msg.Permissions),
	})
	if err != nil {
		return nil, err
	}
	return &service.MsgUpdateRecordResponse{}, nil
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
	if rec.Controller != msg.Authority {
		return nil, fmt.Errorf("unauthorized, record owner does not match the module's authority: got %s, want %s", msg.Authority, rec.Controller)
	}

	if err := ms.k.db.ServiceRecordTable().Delete(ctx, rec); err != nil {
		return nil, err
	}
	return &service.MsgDeleteRecordResponse{}, nil
}

func (ms msgServer) LoginAccount(ctx context.Context, msg *service.MsgLoginAccount) (*service.MsgLoginAccountResponse, error) {
	return nil, nil
}

func (ms msgServer) RegisterAccount(ctx context.Context, msg *service.MsgRegisterAccount) (*service.MsgRegisterAccountResponse, error) {
	return nil, nil
}

// getPermissionsFromInt32 converts an int32 to a ServicePermissions enum.
func getPermissionsFromInt32(value int32) modulev1.ServicePermissions {
	switch value {
	case int32(modulev1.ServicePermissions_SERVICE_PERMISSIONS_BASE):
		return modulev1.ServicePermissions_SERVICE_PERMISSIONS_BASE
	case int32(modulev1.ServicePermissions_SERVICE_PERMISSIONS_READ):
		return modulev1.ServicePermissions_SERVICE_PERMISSIONS_READ
	case int32(modulev1.ServicePermissions_SERVICE_PERMISSIONS_WRITE):
		return modulev1.ServicePermissions_SERVICE_PERMISSIONS_WRITE
	case int32(modulev1.ServicePermissions_SERVICE_PERMISSIONS_OWN):
		return modulev1.ServicePermissions_SERVICE_PERMISSIONS_OWN
	default:
		return modulev1.ServicePermissions_SERVICE_PERMISSIONS_BASE
	}
}
