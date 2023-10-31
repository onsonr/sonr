package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonrhq/sonr/x/service/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) CreateServiceRecord(goCtx context.Context, msg *types.MsgCreateServiceRecord) (*types.MsgCreateServiceRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var serviceRecord = types.ServiceRecord{}

	k.SetServiceRecord(
		ctx,
		serviceRecord,
	)

	return &types.MsgCreateServiceRecordResponse{}, nil
}

func (k msgServer) UpdateServiceRecord(goCtx context.Context, msg *types.MsgUpdateServiceRecord) (*types.MsgUpdateServiceRecordResponse, error) {
	_ = sdk.UnwrapSDKContext(goCtx)

	// var serviceRecord = types.ServiceRecord{
	// 	Controller: msg.Creator,

	// }

	// // Checks that the element exists
	// val, found := k.GetServiceRecord(ctx, msg.Id)
	// if !found {
	// 	return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	// }

	// // Checks if the msg creator is the same as the current owner
	// if msg.Creator != val.Creator {
	// 	return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	// }

	// k.SetServiceRecord(ctx, serviceRecord)

	return &types.MsgUpdateServiceRecordResponse{}, nil
}

func (k msgServer) DeleteServiceRecord(goCtx context.Context, msg *types.MsgDeleteServiceRecord) (*types.MsgDeleteServiceRecordResponse, error) {
	_ = sdk.UnwrapSDKContext(goCtx)

	// // Checks that the element exists
	// val, found := k.GetServiceRecord(ctx, msg.Id)
	// if !found {
	// 	return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	// }

	// // Checks if the msg creator is the same as the current owner
	// if msg.Creator != val.Creator {
	// 	return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	// }

	// k.RemoveServiceRecord(ctx, msg.Id)

	return &types.MsgDeleteServiceRecordResponse{}, nil
}
