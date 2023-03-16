package keeper

import (
	"context"

	"github.com/bufbuild/connect-go"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonrhq/core/x/identity/types"
)

func (s *msgServerConnectWrapper) CreateDidDocument(
	ctx context.Context,
	req *connect.Request[types.MsgCreateDidDocument],
) (*connect.Response[types.MsgCreateDidDocumentResponse], error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Check if the value already exists
	_, isFound := s.GetDidDocument(
		sdkCtx,
		req.Msg.Document.Id,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	s.SetDidDocument(
		sdkCtx,
		*req.Msg.Document,
	)
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent("create-did-document", sdk.NewAttribute("did", req.Msg.Document.Id), sdk.NewAttribute("creator", req.Msg.Creator), sdk.NewAttribute("address", req.Msg.Document.Address())),
	)

	res := connect.NewResponse(&types.MsgCreateDidDocumentResponse{})
	return res, nil
}


func (s *msgServerConnectWrapper) UpdateDidDocument(
	ctx context.Context,
	req *connect.Request[types.MsgUpdateDidDocument],
) (*connect.Response[types.MsgUpdateDidDocumentResponse], error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Check if the value exists
	valFound, isFound := s.GetDidDocument(
		sdkCtx,
		req.Msg.Document.Id,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if valFound.CheckAccAddress(req.Msg.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	s.SetDidDocument(sdkCtx, *req.Msg.Document)
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent("update-did-document", sdk.NewAttribute("did", req.Msg.Document.Id), sdk.NewAttribute("creator", req.Msg.Creator), sdk.NewAttribute("address", req.Msg.Document.Address())),
	)

	res := connect.NewResponse(&types.MsgUpdateDidDocumentResponse{})
	return res, nil
}



func (s *msgServerConnectWrapper) DeleteDidDocument(
	ctx context.Context,
	req *connect.Request[types.MsgDeleteDidDocument],
) (*connect.Response[types.MsgDeleteDidDocumentResponse], error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Check if the value exists
	valFound, isFound := s.GetDidDocument(
		sdkCtx,
		req.Msg.Did,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if valFound.CheckAccAddress(req.Msg.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	s.RemoveDidDocument(
		sdkCtx,
		req.Msg.Did,
	)
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent("delete-did-document", sdk.NewAttribute("document-did", req.Msg.Did), sdk.NewAttribute("creator", req.Msg.Creator)),
	)

	res := connect.NewResponse(&types.MsgDeleteDidDocumentResponse{})
	return res, nil
}
