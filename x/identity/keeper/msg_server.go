package keeper

import (
	"context"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/sonr-io/core/pkg/crypto"
	"github.com/sonr-io/core/x/identity/types"
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

// RegisterIdentity registers a new identity with the provided Identity and Verification Relationships. Fails if not at least one Authentication relationship is provided.
func (k Keeper) RegisterIdentity(goCtx context.Context, msg *types.MsgRegisterIdentity) (*types.MsgRegisterIdentityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if !strings.Contains(msg.DidDocument.Id, msg.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "identity owner does not match creator")
	}
	doc := msg.GetDidDocument()
	addr, err := doc.SDKAddress()
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "failed to get sdk address")
	}
	// Set the identity
	k.SetDIDDocument(ctx, *doc)
	k.Logger(ctx).Info("(x/identity) Account registered", "did", msg.DidDocument.Id, "owner", msg.Creator)

	// acc := k.accountKeeper.NewAccountWithAddress(ctx, addr)
	// k.accountKeeper.SetAccount(ctx, acc)
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, "identity", addr, crypto.NewSNRCoins(1))
	if err != nil {
		k.Logger(ctx).Error("failed to send coins", "error", err)
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "failed to send coins")
	}
	k.Logger(ctx).Info("(x/identity) Sent Claims Reward", "did", msg.DidDocument.Id, "denom", "snr", "amount", 1)
	return &types.MsgRegisterIdentityResponse{
		Success:     true,
		DidDocument: msg.DidDocument,
	}, nil
}

func (k msgServer) CreateEscrowAccount(goCtx context.Context, msg *types.MsgCreateEscrowAccount) (*types.MsgCreateEscrowAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var escrowAccount = types.EscrowAccount{
		Creator:          msg.Creator,
		Address:          msg.Address,
		PublicKey:        msg.PublicKey,
		LockupUsdBalance: msg.LockupUsdBalance,
	}

	id := k.SetEscrowAccount(
		ctx,
		escrowAccount,
	)

	return &types.MsgCreateEscrowAccountResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateEscrowAccount(goCtx context.Context, msg *types.MsgUpdateEscrowAccount) (*types.MsgUpdateEscrowAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var escrowAccount = types.EscrowAccount{
		Creator:          msg.Creator,
		Id:               msg.Id,
		Address:          msg.Address,
		PublicKey:        msg.PublicKey,
		LockupUsdBalance: msg.LockupUsdBalance,
	}

	// Checks that the element exists
	val, found := k.GetEscrowAccount(ctx, msg.Address)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetEscrowAccount(ctx, escrowAccount)

	return &types.MsgUpdateEscrowAccountResponse{}, nil
}

func (k msgServer) DeleteEscrowAccount(goCtx context.Context, msg *types.MsgDeleteEscrowAccount) (*types.MsgDeleteEscrowAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetEscrowAccount(ctx, msg.Address)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Address))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveEscrowAccount(ctx, msg.Address)

	return &types.MsgDeleteEscrowAccountResponse{}, nil
}

func (k msgServer) CreateControllerAccount(goCtx context.Context, msg *types.MsgCreateControllerAccount) (*types.MsgCreateControllerAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var controllerAccount = types.ControllerAccount{
		Address:        msg.Address,
		PublicKey:      msg.PublicKey,
		Authenticators: msg.Authenticators,
		Wallets:        msg.Wallets,
	}

	id := k.SetControllerAccount(
		ctx,
		controllerAccount,
	)

	return &types.MsgCreateControllerAccountResponse{
		Account: &controllerAccount,
		Id:      id,
	}, nil
}

func (k msgServer) UpdateControllerAccount(goCtx context.Context, msg *types.MsgUpdateControllerAccount) (*types.MsgUpdateControllerAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var controllerAccount = types.ControllerAccount{
		Address: msg.Address,
	}

	// Checks that the element exists
	val, found := k.GetControllerAccount(ctx, msg.Address)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Address != val.Address {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetControllerAccount(ctx, controllerAccount)

	return &types.MsgUpdateControllerAccountResponse{}, nil
}

func (k msgServer) DeleteControllerAccount(goCtx context.Context, msg *types.MsgDeleteControllerAccount) (*types.MsgDeleteControllerAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetControllerAccount(ctx, msg.Address)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Address))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Address {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveControllerAccount(ctx, msg.Address)

	return &types.MsgDeleteControllerAccountResponse{}, nil
}
