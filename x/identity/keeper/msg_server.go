package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonrhq/core/x/identity/types"
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

// ! ||--------------------------------------------------------------------------------||
// ! ||                    DIDDocument Message Server Implementation                   ||
// ! ||--------------------------------------------------------------------------------||

func (k msgServer) CreateDidDocument(goCtx context.Context, msg *types.MsgCreateDidDocument) (*types.MsgCreateDidDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// Check if the value already exists
	_, ok := k.GetPrimaryIdentity(ctx, msg.Primary.Id)
	if ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}
	_, found := k.GetPrimaryIdentityByAlias(ctx, msg.Primary.AlsoKnownAs[0])
	if found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}
	// Set the value
	k.SetPrimaryIdentity(
		ctx,
		*msg.Primary,
	)

	ucw, found := k.GetClaimableWallet(ctx, uint64(msg.WalletId))
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "unclaimed wallet index not set")
	}

	ucw.Claimed = true
	k.RemoveClaimableWallet(ctx, uint64(msg.WalletId))

	// Set the blockchain identities
	k.SetBlockchainIdentities(ctx, msg.Blockchains...)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent("NewTx", sdk.NewAttribute("tx-name", "create-did-document"), sdk.NewAttribute("did", msg.Primary.Id), sdk.NewAttribute("creator", msg.Creator), sdk.NewAttribute("alias", msg.Alias)),
	)
	return &types.MsgCreateDidDocumentResponse{}, nil
}

func (k msgServer) UpdateDidDocument(goCtx context.Context, msg *types.MsgUpdateDidDocument) (*types.MsgUpdateDidDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetPrimaryIdentity(
		ctx,
		msg.Primary.Id,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Check if the msg creator is the same as the current owner
	if !valFound.CheckAccAddress(msg.Creator) {
		return nil, types.ErrUnauthorized
	}

	k.SetPrimaryIdentity(ctx, *msg.Primary)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent("NewTx", sdk.NewAttribute("tx-name", "update-did-document"), sdk.NewAttribute("did", msg.Primary.Id), sdk.NewAttribute("creator", msg.Creator)),
	)
	return &types.MsgUpdateDidDocumentResponse{}, nil
}

func (k msgServer) DeleteDidDocument(goCtx context.Context, msg *types.MsgDeleteDidDocument) (*types.MsgDeleteDidDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetPrimaryIdentity(
		ctx,
		msg.Did,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Check if the msg creator is the same as the current owner
	if !valFound.CheckAccAddress(msg.Creator) {
		return nil, types.ErrUnauthorized
	}

	k.RemovePrimaryIdentity(
		ctx,
		msg.Did,
	)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent("NewTx", sdk.NewAttribute("tx-name", "delete-did-document"), sdk.NewAttribute("did", msg.Did), sdk.NewAttribute("creator", msg.Creator)),
	)
	return &types.MsgDeleteDidDocumentResponse{}, nil
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                                  Wallet Claims                                 ||
// ! ||--------------------------------------------------------------------------------||

func (k msgServer) CreateClaimableWallet(goCtx context.Context, msg *types.MsgCreateClaimableWallet) (*types.MsgCreateClaimableWalletResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var claimableWallet = types.ClaimableWallet{
		Creator: msg.Creator,
	}

	id := k.AppendClaimableWallet(
		ctx,
		claimableWallet,
	)

	return &types.MsgCreateClaimableWalletResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateClaimableWallet(goCtx context.Context, msg *types.MsgUpdateClaimableWallet) (*types.MsgUpdateClaimableWalletResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var claimableWallet = types.ClaimableWallet{
		Creator: msg.Creator,
		Id:      msg.Id,
	}

	// Checks that the element exists
	val, found := k.GetClaimableWallet(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetClaimableWallet(ctx, claimableWallet)

	return &types.MsgUpdateClaimableWalletResponse{}, nil
}

func (k msgServer) DeleteClaimableWallet(goCtx context.Context, msg *types.MsgDeleteClaimableWallet) (*types.MsgDeleteClaimableWalletResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetClaimableWallet(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveClaimableWallet(ctx, msg.Id)

	return &types.MsgDeleteClaimableWalletResponse{}, nil
}
