package middleware

import (
	"context"

	"github.com/go-webauthn/webauthn/protocol"

	authenticationpb "github.com/sonrhq/core/types/highway/authentication/v1"
	domainproxy "github.com/sonrhq/core/x/domain/client/proxy"
	serviceproxy "github.com/sonrhq/core/x/service/client/proxy"
)

// GetCredentialAttestationParams returns the credential attestation options to start wallet registration.
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
		Success:            true,
		AttestationOptions: creOpts,
		Challenge:          chal.String(),
		Origin:             origin,
		Alias:              alias,
		Existing:           false,
	}, nil
}

// GetCredentialAssertionOptions returns the credential assertion options to start wallet authentication.
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
		Success:          true,
		AssertionOptions: assertionOpts,
		Challenge:        chal.String(),
		Origin:           origin,
		Alias:            alias,
		Address:          addr,
		Existing:         true,
	}, nil
}
