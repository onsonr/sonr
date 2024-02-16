package service

import (
	"github.com/go-webauthn/webauthn/protocol"

	servicev1 "github.com/sonrhq/sonr/api/sonr/service/v1"
)

// ConvertCredentialToDescriptor converts a credential to a descriptor
func ConvertCredentialToDescriptor(credential *servicev1.WebauthnCredential) CredentialDescriptor {
	transports := make([]protocol.AuthenticatorTransport, 0)
	for _, transport := range credential.Transports {
		transports = append(transports, protocol.AuthenticatorTransport(transport))
	}
	return CredentialDescriptor{
		Type:            "public-key",
		CredentialID:    credential.Id,
		Transport:       transports,
		AttestationType: credential.AssertionType,
	}
}

// ExtractCredentialDescriptors extracts the credential descriptors from the given credentials
func ExtractCredentialDescriptors(credentials []*servicev1.WebauthnCredential) (descriptors []CredentialDescriptor) {
	for _, credential := range credentials {
		descriptors = append(descriptors, ConvertCredentialToDescriptor(credential))
	}
	return
}

// NewQueryServiceRequest creates a new query service request
func NewQueryServiceRequest(origin string) *servicev1.QueryServiceRecordRequest {
	return &servicev1.QueryServiceRecordRequest{Origin: origin}
}

// NewQueryWebauthnRecord creates a new query service request
func NewQueryWebauthnRecord(origin string, handle string) *servicev1.QueryWebauthnCredentialRequest {
	return &servicev1.QueryWebauthnCredentialRequest{Origin: origin, Handle: handle}
}
