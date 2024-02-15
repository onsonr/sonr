package service

import (
	"github.com/go-webauthn/webauthn/protocol"

	modulev1 "github.com/sonrhq/sonr/api/sonr/service/module/v1"
)

func ConvertCredentialToDescriptor(credential *modulev1.Credential) CredentialDescriptor {
	transports := make([]protocol.AuthenticatorTransport, 0)
	for _, transport := range credential.Transport {
		transports = append(transports, protocol.AuthenticatorTransport(transport))
	}
	return CredentialDescriptor{
		Type:            "public-key",
		CredentialID:    credential.CredentialId,
		Transport:       transports,
		AttestationType: credential.AttestationType,
	}
}

// ExtractCredentialDescriptors extracts the credential descriptors from the given credentials
func ExtractCredentialDescriptors(credentials []*modulev1.Credential) (descriptors []CredentialDescriptor) {
	for _, credential := range credentials {
		descriptors = append(descriptors, ConvertCredentialToDescriptor(credential))
	}
	return
}

func NewQueryCredentialsRequest(origin, identifier string) *modulev1.GetCredentialByOriginHandleRequest {
	return &modulev1.GetCredentialByOriginHandleRequest{Origin: origin, Handle: identifier}
}

func NewQueryServiceRequest(origin string) *modulev1.GetServiceRecordByOriginRequest {
	return &modulev1.GetServiceRecordByOriginRequest{Origin: origin}
}
