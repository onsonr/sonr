package service

import (
	"context"

	"github.com/go-webauthn/webauthn/protocol"
	grpc "google.golang.org/grpc"

	modulev1 "github.com/sonrhq/sonr/api/sonr/service/module/v1"
)

// CredentialDescriptor is a descriptor for a credential
type CredentialDescriptor = protocol.CredentialDescriptor

// PublicKeyCredentialCreationOptions is the options for creating a public key credential
type PublicKeyCredentialCreationOptions = protocol.PublicKeyCredentialCreationOptions

// PublicKeyCredentialRequestOptions is the options for requesting a public key credential
type PublicKeyCredentialRequestOptions = protocol.PublicKeyCredentialRequestOptions

// ServiceRecord is a record of a service
type ServiceRecord = modulev1.ServiceRecord

// UserEntity is the entity of a user
type UserEntity = protocol.UserEntity

// GetCredentialCreationOptions returns the PublicKeyCredentialCreationOptions for the given service record and user entity.
func GetCredentialCreationOptions(record *modulev1.ServiceRecord, entity protocol.UserEntity) PublicKeyCredentialCreationOptions {
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
func GetCredentialRequestOptions(record *modulev1.ServiceRecord, creds []CredentialDescriptor) PublicKeyCredentialRequestOptions {
	return protocol.PublicKeyCredentialRequestOptions{
		Challenge:          GenerateChallenge(),
		UserVerification:   protocol.VerificationPreferred,
		RelyingPartyID:     record.Origin,
		AllowedCredentials: creds,
	}
}

// GetCredentialsByHandle returns the credentials for the given handle
func GetCredentialsByHandle(conn *grpc.ClientConn, handle, origin string) ([]CredentialDescriptor, error) {
	res, err := getStateServiceClient(conn).GetCredentialByOriginHandle(context.Background(), NewQueryCredentialsRequest(origin, handle))
	if err != nil {
		return nil, err
	}
	return []CredentialDescriptor{ConvertCredentialToDescriptor(res.GetValue())}, nil
}

// GetRecordByOrigin returns the service record for the given origin
func GetRecordByOrigin(conn *grpc.ClientConn, origin string) (*ServiceRecord, error) {
	res, err := getStateServiceClient(conn).GetServiceRecordByOrigin(context.Background(), NewQueryServiceRequest(origin))
	if err != nil {
		return nil, err
	}
	return res.GetValue(), nil
}

// GetUserEntity returns the user entity for the given address and username
func GetUserEntity(address, username string) protocol.UserEntity {
	return protocol.UserEntity{
		ID:          []byte(address),
		DisplayName: username,
	}
}
