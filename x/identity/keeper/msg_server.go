package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonrhq/core/pkg/wallet"
	"github.com/sonrhq/core/x/identity/types"
)

var (
	connServer *msgServerConnectWrapper
)

type msgServerConnectWrapper struct {
	Keeper
}

type msgServer struct {
	Keeper
	Connect *msgServerConnectWrapper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// ! ||--------------------------------------------------------------------------------||
// ! ||                    DIDDocument Message Server Implementation                   ||
// ! ||--------------------------------------------------------------------------------||

func (k msgServer) CreateDidDocument(goCtx context.Context, msg *types.MsgCreateDidDocument) (*types.MsgCreateDidDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// Check if the value already exists
	_, isFound := k.GetDidDocument(
		ctx,
		msg.Document.Id,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	k.SetDidDocument(
		ctx,
		*msg.Document,
	)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent("NewTx", sdk.NewAttribute("tx-name", "create-did-document"), sdk.NewAttribute("did", msg.Document.Id), sdk.NewAttribute("creator", msg.Creator)),
	)
	return &types.MsgCreateDidDocumentResponse{}, nil
}

func (k msgServer) UpdateDidDocument(goCtx context.Context, msg *types.MsgUpdateDidDocument) (*types.MsgUpdateDidDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetDidDocument(
		ctx,
		msg.Document.Id,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Check if the msg creator is the same as the current owner
	if !valFound.CheckAccAddress(msg.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetDidDocument(ctx, *msg.Document)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent("NewTx", sdk.NewAttribute("tx-name", "update-did-document"), sdk.NewAttribute("did", msg.Document.Id), sdk.NewAttribute("creator", msg.Creator)),
	)
	return &types.MsgUpdateDidDocumentResponse{}, nil
}

func (k msgServer) DeleteDidDocument(goCtx context.Context, msg *types.MsgDeleteDidDocument) (*types.MsgDeleteDidDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetDidDocument(
		ctx,
		msg.Did,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Check if the msg creator is the same as the current owner
	if !valFound.CheckAccAddress(msg.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveDidDocument(
		ctx,
		msg.Did,
	)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent("NewTx", sdk.NewAttribute("tx-name", "delete-did-document"), sdk.NewAttribute("did", msg.Did), sdk.NewAttribute("creator", msg.Creator)),
	)
	return &types.MsgDeleteDidDocumentResponse{}, nil
}

func (k msgServer) DeletePublicKey(goCtx context.Context, msg *types.MsgDeletePublicKey) (*types.MsgDeletePublicKeyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgDeletePublicKeyResponse{}, nil
}

func (k msgServer) ImportPublicKey(goCtx context.Context, msg *types.MsgImportPublicKey) (*types.MsgImportPublicKeyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgImportPublicKeyResponse{}, nil
}

func (k msgServer) RegisterAccount(goCtx context.Context, msg *types.MsgRegisterAccount) (*types.MsgRegisterAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetService(ctx, msg.Origin)
	if found {
		return nil, fmt.Errorf("service already exists")
	}

	cred, err := val.VerifyCreationChallenge(msg.CredentialResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to verify challenge: %w", err)
	}

	wallChan := make(chan wallet.Wallet)
	errChan := make(chan error)
	go func() {
		wall, err := wallet.NewWallet(msg.Uuid, 1)
		if err != nil {
			errChan <- err
			return
		}
		wallChan <- wall
	}()

	select {
	case wall := <-wallChan:
		doc, vms, err := wall.Assign(cred)
		if err != nil {
			return nil, fmt.Errorf("failed to set authentication: %w", err)
		}
		k.SetDidDocument(ctx, *doc)
		resolved := doc.ResolveMethods(vms)
		return &types.MsgRegisterAccountResponse{
			Did:      doc.Id,
			Document: resolved,
		}, nil
	case err := <-errChan:
		return nil, fmt.Errorf("failed to create wallet: %w", err)
	}
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                      Service Message Server Implementation                     ||
// ! ||--------------------------------------------------------------------------------||

func (k msgServer) RegisterService(goCtx context.Context, msg *types.MsgRegisterService) (*types.MsgRegisterServiceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetDidDocument(
		ctx,
		msg.Creator,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	return &types.MsgRegisterServiceResponse{}, nil
}

func (k msgServer) UpdateService(goCtx context.Context, msg *types.MsgUpdateService) (*types.MsgUpdateServiceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_ = ctx

	return &types.MsgUpdateServiceResponse{}, nil
}

func (k msgServer) DeactivateService(goCtx context.Context, msg *types.MsgDeactivateService) (*types.MsgDeactivateServiceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_ = ctx

	return &types.MsgDeactivateServiceResponse{}, nil
}
