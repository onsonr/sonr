package service

import (
	"context"

	"github.com/go-webauthn/webauthn/protocol"
	grpc "google.golang.org/grpc"
)

// CredentialDescriptor is a descriptor for a credential
type CredentialDescriptor = protocol.CredentialDescriptor

// PublicKeyCredentialCreationOptions is the options for creating a public key credential
type PublicKeyCredentialCreationOptions = protocol.PublicKeyCredentialCreationOptions

// PublicKeyCredentialRequestOptions is the options for requesting a public key credential
type PublicKeyCredentialRequestOptions = protocol.PublicKeyCredentialRequestOptions

// UserEntity is the entity of a user
type UserEntity = protocol.UserEntity

// GetCredentialsByHandle returns the credentials for the given handle
func GetCredentialsByHandle(conn *grpc.ClientConn, handle, origin string) ([]CredentialDescriptor, error) {
	res, err := getQueryServiceClient(conn).WebauthnCredential(context.Background(), NewQueryWebauthnRecord(origin, handle))
	if err != nil {
		return nil, err
	}
	return ExtractCredentialDescriptors(res.GetWebauthnCredential()), nil
}

// GetCredentialCreationOptions returns the PublicKeyCredentialCreationOptions for the given service record and user entity.
func GetCredentialCreationOptions(record *ServiceRecord, entity protocol.UserEntity) PublicKeyCredentialCreationOptions {
	return protocol.PublicKeyCredentialCreationOptions{
		RelyingParty: protocol.RelyingPartyEntity{
			ID: record.Origin,
		},
		Challenge: GenerateChallenge(),
		AuthenticatorSelection: protocol.AuthenticatorSelection{
			AuthenticatorAttachment: protocol.CrossPlatform,
			UserVerification:        protocol.VerificationPreferred,
			ResidentKey:             protocol.ResidentKeyRequirementPreferred,
		},
		Attestation: protocol.PreferIndirectAttestation,
		User:        entity,
	}
}

// GetCredentialRequestOptions returns the PublicKeyCredentialRequestOptions for the given service record and credentials.
func GetCredentialRequestOptions(record *ServiceRecord, creds []CredentialDescriptor) PublicKeyCredentialRequestOptions {
	return protocol.PublicKeyCredentialRequestOptions{
		Challenge:          GenerateChallenge(),
		UserVerification:   protocol.VerificationPreferred,
		RelyingPartyID:     record.Origin,
		AllowedCredentials: creds,
	}
}

// GetRecordByOrigin returns the service record for the given origin
func GetRecordByOrigin(conn *grpc.ClientConn, origin string) (*ServiceRecord, error) {
	res, err := getQueryServiceClient(conn).ServiceRecord(context.Background(), NewQueryServiceRequest(origin))
	if err != nil {
		return nil, err
	}
	srv := res.GetServiceRecord()
	return &ServiceRecord{
		Origin:      srv.Origin,
		Name:        srv.Name,
		Description: srv.Description,
		Permissions: srv.Permissions,
		Authority:   srv.Authority,
	}, nil
}

// GetUserEntity returns the user entity for the given address and username
func GetUserEntity(address, username string) protocol.UserEntity {
	return protocol.UserEntity{
		ID:          []byte(address),
		DisplayName: username,
	}
}
