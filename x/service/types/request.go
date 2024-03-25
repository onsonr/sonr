package types

// // ConvertCredentialToDescriptor converts a credential to a descriptor
// func ConvertCredentialToDescriptor(credential *WebauthnCredential) CredentialDescriptor {
// 	transports := make([]protocol.AuthenticatorTransport, 0)
// 	for _, transport := range credential.Transports {
// 		transports = append(transports, protocol.AuthenticatorTransport(transport))
// 	}
// 	return CredentialDescriptor{
// 		Type:            "public-key",
// 		CredentialID:    credential.Id,
// 		Transport:       transports,
// 		AttestationType: credential.AssertionType,
// 	}
// }

// // ExtractCredentialDescriptors extracts the credential descriptors from the given credentials
// func ExtractCredentialDescriptors(credentials []*WebauthnCredential) (descriptors []CredentialDescriptor) {
// 	for _, credential := range credentials {
// 		descriptors = append(descriptors, ConvertCredentialToDescriptor(credential))
// 	}
// 	return
// }

// NewQueryServiceRequest creates a new query service request
func NewQueryServiceRequest(origin string) *QueryServiceRecordRequest {
	return &QueryServiceRecordRequest{Origin: origin}
}

// // NewQueryWebauthnRecord creates a new query service request
// func NewQueryWebauthnRecord(origin string, handle string) *QueryWebauthnCredentialRequest {
// 	return &QueryWebauthnCredentialRequest{Origin: origin, Handle: handle}
// }
