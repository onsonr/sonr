package keeper

import (
	"context"
	"strings"

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

// RegisterIdentity registers a new identity with the provided Identity and Verification Relationships. Fails if not at least one Authentication relationship is provided.
func (k Keeper) RegisterIdentity(goCtx context.Context, msg *types.MsgRegisterIdentity) (*types.MsgRegisterIdentityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !strings.Contains(msg.DidDocument.Id, msg.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "identity owner does not match creator")
	}

	// Set the identity
	err := k.SetIdentity(ctx, *msg.DidDocument)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "failed to set identity, sequence cannot proceed.")
	}
	k.Logger(ctx).Info("(x/identity) Account registered", "did", msg.DidDocument.Id, "owner", msg.Creator)
	k.vaultKeeper.RemoveClaimableWallet(ctx, msg.GetWalletId())

	// err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, "identity", sdk.AccAddress(msg.Creator), sdk.NewCoins(sdk.NewCoin("snr", sdk.NewInt(1))))
	// if err != nil {
	// 	k.Logger(ctx).Error("failed to send coins", "error", err)
	// 	return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "failed to send coins")
	// }

	return &types.MsgRegisterIdentityResponse{
		Success:     true,
		DidDocument: msg.DidDocument,
	}, nil
}
