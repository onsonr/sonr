package keeper

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/sonr-hq/sonr/x/identity/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) WebauthnRegistrationBegin(goCtx context.Context, req *types.QueryWebauthnRegisterBeginRequest) (*types.QueryWebauthnRegisterBeginResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	user := &types.VerificationMethod{
		ID: req.Uuid,
	}
	options, sessionData, err := k.web.BeginRegistration(user)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	k.userCache.Set(fmt.Sprintf("%s-registration", req.Uuid), sessionData, 0)

	// Marshal the options into a JSON string
	dj, err := json.Marshal(options)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.QueryWebauthnRegisterBeginResponse{
		Response: dj,
	}, nil
}

func (k Keeper) WebauthnRegistrationFinish(goCtx context.Context, req *types.QueryWebauthnRegisterFinishRequest) (*types.QueryWebauthnRegisterFinishResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	user := &types.VerificationMethod{
		ID: req.Uuid,
	}
	sessionDataRaw, ok := k.userCache.Get(fmt.Sprintf("%s-registration", req.Uuid))
	if !ok {
		return nil, status.Error(codes.Internal, "session data not found")
	}
	sessionData, ok := sessionDataRaw.(webauthn.SessionData)
	if !ok {
		return nil, status.Error(codes.Internal, "session data not found")
	}
	parsedResponse, err := protocol.ParseCredentialCreationResponseBody(bytes.NewReader(req.GeneratedCredential))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	credential, err := k.web.CreateCredential(user, sessionData, parsedResponse)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.QueryWebauthnRegisterFinishResponse{
		Credential: types.ConvertFromWebauthnCredential(credential),
	}, nil
}

func (k Keeper) WebauthnLoginBegin(goCtx context.Context, req *types.QueryWebauthnLoginBeginRequest) (*types.QueryWebauthnLoginBeginResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return nil, errors.New("not implemented")
}

func (k Keeper) WebauthnLoginFinish(goCtx context.Context, req *types.QueryWebauthnLoginFinishRequest) (*types.QueryWebauthnLoginFinishResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return nil, errors.New("not implemented")
}
