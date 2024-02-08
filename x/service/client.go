package service

import (
	"context"

	"github.com/go-webauthn/webauthn/protocol"

	modulev1 "github.com/sonrhq/sonr/api/sonr/service/module/v1"
)

// FinishLogin returns the result of the credential request
func FinishLogin(ctx context.Context, record *modulev1.ServiceRecord, credential protocol.PublicKeyCredential) (bool, error) {
	return false, nil
}

// FinishRegistration returns the result of the credential creation
func FinishRegistration(ctx context.Context, record *modulev1.ServiceRecord, credential protocol.PublicKeyCredential) (bool, error) {
	return false, nil
}
