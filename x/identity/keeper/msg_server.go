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
	identity := msg.DidDocument.ToIdentification()
	err := k.SetIdentity(ctx, *identity)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "failed to set identity, sequence cannot proceed.")
	}

	// Iteratively set authentication relations
	for _, auth := range msg.DidDocument.Authentication {
		auth.Owner = msg.Creator
		auth.Reference = auth.VerificationMethod.Id
		auth.Type = "Authentication"
		k.SetAuthentication(ctx, *auth)
		if err != nil {
			k.Logger(ctx).Error("failed to set authentication", "error", err)
		}
	}

	// Iteratively set assertion relations
	for _, assertion := range msg.DidDocument.AssertionMethod {
		assertion.Owner = msg.Creator
		assertion.Reference = assertion.VerificationMethod.Id
		assertion.Type = "AssertionMethod"
		k.SetAssertion(ctx, *assertion)
		if err != nil {
			k.Logger(ctx).Error("failed to set assertion", "error", err)
		}
	}

	// Iteratively set capability delegation relations
	for _, capability := range msg.DidDocument.CapabilityDelegation {
		capability.Owner = msg.Creator
		capability.Reference = capability.VerificationMethod.Id
		capability.Type = "CapabilityDelegation"
		k.SetCapabilityDelegation(ctx, *capability)
		if err != nil {
			k.Logger(ctx).Error("failed to set capability delegation", "error", err)
		}
	}

	// Iteratively set capability invocation relations
	for _, capability := range msg.DidDocument.CapabilityInvocation {
		capability.Owner = msg.Creator
		capability.Reference = capability.VerificationMethod.Id
		capability.Type = "CapabilityInvocation"
		k.SetCapabilityInvocation(ctx, *capability)
		if err != nil {
			k.Logger(ctx).Error("failed to set capability invocation", "error", err)
		}
	}

	// Iteratively set key agreement relations
	for _, keyAgreement := range msg.DidDocument.KeyAgreement {
		keyAgreement.Owner = msg.Creator
		k.SetKeyAgreement(ctx, *keyAgreement)
		if err != nil {
			k.Logger(ctx).Error("failed to set key agreement", "error", err)
		}
	}

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
