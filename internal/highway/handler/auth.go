package handler

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	mdw "github.com/sonrhq/core/internal/highway/middleware"
	"github.com/sonrhq/core/internal/highway/types"
	authenticationpb "github.com/sonrhq/core/types/highway/authentication/v1"
	domainproxy "github.com/sonrhq/core/x/domain/client/proxy"
	identityproxy "github.com/sonrhq/core/x/identity/client/proxy"
	serviceproxy "github.com/sonrhq/core/x/service/client/proxy"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// AuthenticationHandler is the handler for the authentication service
type AuthenticationHandler struct {
	cctx client.Context
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                                   API Methods                                  ||
// ! ||--------------------------------------------------------------------------------||

// Authenticate handles the authentication request
func (a *AuthenticationHandler) Login(ctx context.Context, req *authenticationpb.LoginRequest) (*authenticationpb.LoginResponse, error) {
	record, err := serviceproxy.GetServiceRecord(ctx, req.Origin)
	if err != nil {
		return nil, err
	}
	_, err = record.VerifyAssertionChallenge(req.Assertion)
	if err != nil {
		return nil, err
	}
	addr, err := domainproxy.GetEmailRecordCreator(ctx, req.Alias)
	if err != nil {
		return nil, err
	}
	contAcc, err := identityproxy.GetControllerAccount(ctx, addr)
	if err != nil {
		return nil, err
	}
	token, err := types.NewSessionJWTClaims(req.Alias, contAcc)
	if err != nil {
		return nil, err
	}
	return &authenticationpb.LoginResponse{
		Origin:  req.Origin,
		Address: contAcc.Address,
		Jwt:     token,
	}, nil
}

// CurrentUser returns the current user
func (a *AuthenticationHandler) CurrentUser(ctx context.Context, req *emptypb.Empty) (*authenticationpb.CurrentUserResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

// Params returns the parameters for the given request
func (a *AuthenticationHandler) Params(ctx context.Context, req *authenticationpb.ParamsRequest) (*authenticationpb.ParamsResponse, error) {
	if req.Alias == "" {
		return nil, fmt.Errorf("user provided identifier cannot be empty")
	}
	if req.GetExisting() {
		return mdw.GetCredentialAssertionOptions(ctx, req.Origin, req.Alias)
	}
	return mdw.GetCredentialAttestationParams(ctx, req.Origin, req.Alias)
}

// Register handles the registration request
func (a *AuthenticationHandler) Register(ctx context.Context, req *authenticationpb.RegisterRequest) (*authenticationpb.RegisterResponse, error) {
	// Get the service record from the origin
	record, err := serviceproxy.GetServiceRecord(ctx, req.Origin)
	if err != nil {
		return nil, err
	}
	credential, err := record.VerifyCreationChallenge(req.Attestation, req.Challenge)
	if err != nil && credential == nil {
		return nil, err
	}
	cont, resp, err := mdw.PublishControllerAccount(req.Username, credential, req.Origin)
	if err != nil {
		return nil, err
	}
	token, err := types.NewSessionJWTClaims(req.Username, cont.Account())
	if err != nil {
		return nil, err
	}
	return &authenticationpb.RegisterResponse{
		Origin:  req.Origin,
		Address: cont.Account().Address,
		Jwt:     token,
		TxHash:  resp.TxHash,
	}, nil
}

// RefreshToken handles the refresh token request
func (a *AuthenticationHandler) RefreshToken(ctx context.Context, req *authenticationpb.RefreshTokenRequest) (*authenticationpb.RefreshTokenResponse, error) {
	return nil, nil
}

// VerifyToken handles the verify token request
func (a *AuthenticationHandler) VerifyToken(ctx context.Context, req *authenticationpb.VerifyTokenRequest) (*authenticationpb.VerifyTokenResponse, error) {
	return nil, nil
}
