package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/sonrhq/core/x/service/types"
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

func (k msgServer) RegisterServiceRecord(goCtx context.Context, msg *types.MsgRegisterServiceRecord) (*types.MsgRegisterServiceRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetServiceRecord(
		ctx,
		msg.Record.Id,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Id already set")
	}
	k.SetServiceRecord(
		ctx,
		*msg.Record,
	)
	return &types.MsgRegisterServiceRecordResponse{}, nil
}

func (k msgServer) UpdateServiceRecord(goCtx context.Context, msg *types.MsgUpdateServiceRecord) (*types.MsgUpdateServiceRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetServiceRecord(
		ctx,
		msg.Record.Id,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "Id not set")
	}

	// Checks if the the msg Controller is the same as the current owner
	if msg.Controller != valFound.Controller {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	k.SetServiceRecord(ctx, *msg.Record)
	return &types.MsgUpdateServiceRecordResponse{}, nil
}

func (k msgServer) BurnServiceRecord(goCtx context.Context, msg *types.MsgBurnServiceRecord) (*types.MsgBurnServiceRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetServiceRecord(
		ctx,
		msg.Id,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "Id not set")
	}

	// Checks if the the msg Controller is the same as the current owner
	if msg.Controller != valFound.Controller {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveServiceRecord(
		ctx,
		msg.Id,
	)

	return &types.MsgBurnServiceRecordResponse{}, nil
}

// RegisterUserEntity registers a new user entity using the provided attestation and challenge
func (k msgServer) RegisterUserEntity(goCtx context.Context, msg *types.MsgRegisterUserEntity) (*types.MsgRegisterUserEntityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the service exists
	service, ok := k.GetServiceRecord(ctx, msg.ServiceOrigin)
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "Service not found")
	}

	// Check if desired alias is already taken
	err := k.identityKeeper.CheckAlsoKnownAs(ctx, msg.DesiredAlias)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Desired alias already taken")
	}

	ucw, ok := k.identityKeeper.GetClaimableWallet(ctx, msg.UcwId)
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "Wallet Claims not found")
	}


	// Verify both attestion and challenge are valid
	cred, err := service.VerifyCreationChallenge(msg.Attestation, msg.Challenge)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Attestation or challenge invalid")
	}

	// Assign identity to user entity
	id, err := k.identityKeeper.AssignIdentity(ctx, ucw, cred, msg.DesiredAlias)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Identity could not be assigned")
	}

	// Set service relationship
	k.SetServiceRelationship(ctx, types.ServiceRelationship{
		Reference: service.Id,
		Did: id.Id,
		Count: 0,
	})

	// Return response
	return &types.MsgRegisterUserEntityResponse{
		Identity: id,
		Success:  true,
		Did:      id.Id,
	}, nil
}

// AuthenticateUserEntity authenticates a user entity using the provided attestation and challenge
func (k msgServer) AuthenticateUserEntity(goCtx context.Context, msg *types.MsgAuthenticateUserEntity) (*types.MsgAuthenticateUserEntityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx

	// Example of gofiber implementation for loading a user
	// ---
	//
	// 	q := middleware.ParseQuery(c)
	// if !q.HasAssertion() {
	// 	return c.Status(400).SendString("Missing assertion.")
	// }
	// _, err := q.GetService()
	// if err != nil {
	// 	return c.Status(404).SendString(err.Error())
	// }

	// doc, err := q.GetDID()
	// if err != nil {
	// 	return c.Status(405).SendString(err.Error())
	// }
	// if err := service.VerifyAssertionChallenge(q.Assertion(), doc.ListCredentialVerificationMethods()...); err != nil {
	// 	return c.Status(403).SendString(err.Error())
	// }

	// cont, err := identity.LoadControllerWithDid(doc)
	// if err != nil {
	// 	return c.Status(412).SendString(err.Error())
	// }
	// usr := middleware.NewUser(cont, doc.FindUsername())
	// // Create the Claims
	// jwt, err := usr.JWT()
	// if err != nil {
	// 	return c.Status(401).SendString(err.Error())
	// }
	return &types.MsgAuthenticateUserEntityResponse{}, nil
}
