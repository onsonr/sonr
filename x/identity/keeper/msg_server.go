package keeper

import (
	"context"

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
	_, ok := k.GetDidDocument(ctx, msg.Primary.Id)
	if ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}
	_, found := k.GetDidDocumentByAlsoKnownAs(ctx, msg.Primary.AlsoKnownAs[0])
	if found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}
	// Set the value
	k.SetDidDocument(
		ctx,
		*msg.Primary,
	)

	ucw, found := k.GetClaimableWallet(ctx, uint64(msg.WalletId))
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "unclaimed wallet index not set")
	}

	ucw.Claimed = true
	k.RemoveClaimableWallet(ctx, uint64(msg.WalletId))
	return &types.MsgCreateDidDocumentResponse{}, nil
}

func (k msgServer) UpdateDidDocument(goCtx context.Context, msg *types.MsgUpdateDidDocument) (*types.MsgUpdateDidDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetDidDocument(
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

	k.SetDidDocument(ctx, *msg.Primary)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent("NewTx", sdk.NewAttribute("tx-name", "update-did-document"), sdk.NewAttribute("did", msg.Primary.Id), sdk.NewAttribute("creator", msg.Creator)),
	)
	return &types.MsgUpdateDidDocumentResponse{}, nil
}

// RegisterIdentity registers a new identity with the provided Identity and Verification Relationships. Fails if not at least one Authentication relationship is provided.
func (k Keeper) RegisterIdentity(goCtx context.Context, msg *types.MsgRegisterIdentity) (*types.MsgRegisterIdentityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if len(msg.Authentication) == 0 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "at least one authentication relationship must be provided")
	}

	// Check if the value already exists
	if ok := k.HasIdentity(ctx, msg.Identity.Id); ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "identity already registered")
	}
	if msg.Creator != msg.Identity.Owner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "identity owner does not match creator")
	}

	// Remove the unclaimed wallet
	ucw, found := k.GetClaimableWallet(ctx, uint64(msg.WalletId))
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "unclaimed wallet index not set")
	}
	ucw.Claimed = true
	k.RemoveClaimableWallet(ctx, uint64(msg.WalletId))

	// Set the identity
	err := k.SetIdentity(ctx, *msg.Identity)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "failed to set identity, sequence cannot proceed.")
	}

	// Iteratively set authentication relations
	for _, auth := range msg.Authentication {
		auth.Owner = msg.Identity.Owner
		auth.Reference = auth.VerificationMethod.Id
		auth.Type = "Authentication"
		k.SetAuthentication(ctx, *auth)
		if err != nil {
			k.Logger(ctx).Error("failed to set authentication", "error", err)
		}
	}

	// Iteratively set assertion relations
	for _, assertion := range msg.Assertion {
		assertion.Owner = msg.Identity.Owner
		assertion.Reference = assertion.VerificationMethod.Id
		assertion.Type = "AssertionMethod"
		k.SetAssertion(ctx, *assertion)
		if err != nil {
			k.Logger(ctx).Error("failed to set assertion", "error", err)
		}
	}

	// Iteratively set capability delegation relations
	for _, capability := range msg.CapabilityDelegation {
		capability.Owner = msg.Identity.Owner
		capability.Reference = capability.VerificationMethod.Id
		capability.Type = "CapabilityDelegation"
		k.SetCapabilityDelegation(ctx, *capability)
		if err != nil {
			k.Logger(ctx).Error("failed to set capability delegation", "error", err)
		}
	}

	// Iteratively set capability invocation relations
	for _, capability := range msg.CapabilityInvocation {
		capability.Owner = msg.Identity.Owner
		capability.Reference = capability.VerificationMethod.Id
		capability.Type = "CapabilityInvocation"
		k.SetCapabilityInvocation(ctx, *capability)
		if err != nil {
			k.Logger(ctx).Error("failed to set capability invocation", "error", err)
		}
	}

	// Iteratively set key agreement relations
	for _, keyAgreement := range msg.KeyAgreement {
		keyAgreement.Owner = msg.Identity.Owner
		k.SetKeyAgreement(ctx, *keyAgreement)
		if err != nil {
			k.Logger(ctx).Error("failed to set key agreement", "error", err)
		}
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, "identity", sdk.AccAddress(msg.Creator), sdk.NewCoins(sdk.NewCoin("snr", sdk.NewInt(1))))
	if err != nil {
		k.Logger(ctx).Error("failed to send coins", "error", err)
	}

	return &types.MsgRegisterIdentityResponse{
		Success: true,
		DidDocument: &types.DIDDocument{
			Context:              []string{types.DefaultParams().DidBaseContext, types.DefaultParams().AccountDidMethodContext},
			Id:                   msg.Identity.Id,
			Authentication:       msg.Authentication,
			AssertionMethod:      msg.Assertion,
			CapabilityDelegation: msg.CapabilityDelegation,
			CapabilityInvocation: msg.CapabilityInvocation,
			Metadata:             msg.Identity.Metadata,
		},
	}, nil
}
