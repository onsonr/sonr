package middleware

import (
	"context"

	"github.com/go-webauthn/webauthn/protocol"

	authenticationpb "github.com/sonrhq/core/types/highway/authentication/v1"
	domainproxy "github.com/sonrhq/core/x/domain/client/proxy"
	serviceproxy "github.com/sonrhq/core/x/service/client/proxy"
)

// GetCredentialAttestationParams returns the credential creation options to start account registration.
//
// @Summary Get credential attestation parameters
// @Description Returns the credential creation options to start account registration.
// @Accept  json
// @Produce  json
// @Param origin path string true "Origin"
// @Param alias path string true "Alias"
// @Success 200 {object} map[string]interface{} "Success response"
// @Failure 500 {object} map[string]string "Error message"
// @Router /getCredentialAttestationParams/{origin}/{alias} [get]
func GetCredentialAttestationParams(ctx context.Context, origin string, alias string) (*authenticationpb.ParamsResponse, error) {
	ok, err := domainproxy.CheckAliasAvailable(ctx, alias)
	if err != nil || !ok {
		return nil, err
	}
	// Get the service record from the origin
	rec, err := serviceproxy.GetServiceRecord(ctx, origin)
	if err != nil {
		return nil, err
	}
	chal, err := protocol.CreateChallenge()
	if err != nil {
		return nil, err
	}
	creOpts, err := rec.GetCredentialCreationOptions(alias, chal)
	if err != nil {
		return nil, err
	}
	return &authenticationpb.ParamsResponse{
		AttestationOptions: creOpts,
		Challenge:          chal.String(),
		Origin:             origin,
		Alias:              alias,
		IsAuth:             false,
	}, nil
}

// GetCredentialAssertionParams returns the credential assertion options to start account login.
//
// @Summary Get credential assertion parameters
// @Description Returns the credential assertion options to start account login.
// @Accept  json
// @Produce  json
// @Param origin path string true "Origin"
// @Param alias path string true "Alias"
// @Success 200 {object} map[string]interface{} "Success response"
// @Failure 500 {object} map[string]string "Error message"
// @Router /getCredentialAssertionParams/{origin}/{alias} [get]
func GetCredentialAssertionOptions(ctx context.Context, origin string, alias string) (*authenticationpb.ParamsResponse, error) {
	record, err := serviceproxy.GetServiceRecord(ctx, origin)
	if err != nil {
		return nil, err
	}
	notok, err := domainproxy.CheckAliasUnavailable(ctx, alias)
	if err != nil && notok {
		return nil, err
	}
	assertionOpts, chal, addr, err := IssueCredentialAssertionOptions(alias, record)
	if err != nil {
		return nil, err
	}
	return &authenticationpb.ParamsResponse{
		AssertionOptions: assertionOpts,
		Challenge:        chal.String(),
		Origin:           origin,
		Alias:            alias,
		Address:          addr,
		IsAuth:           true,
	}, nil
}
