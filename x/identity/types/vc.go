package types

import (
	"encoding/base64"
	"errors"
	"strconv"
	"strings"

	"github.com/go-webauthn/webauthn/webauthn"
	common "github.com/sonr-hq/sonr/pkg/common"
)

func (wvm *VerificationMethod) WebAuthnCredential() (*common.WebauthnCredential, error) {
	data := wvm.GetMetadata()
	// Fetch Properties from map
	if data == nil {
		return nil, errors.New("Failed to find metadata for VerificationMethod")
	}
	credIdRaw, ok := data["credential_id"]
	if !ok {
		return nil, errors.New("Failed to get authenticator aaguid")
	}
	authAaguidRaw, ok := data["authenticator.aaguid"]
	if !ok {
		return nil, errors.New("Failed to get authenticator aaguid")
	}
	authClonedRaw, ok := data["authenticator.clone_warning"]
	if !ok {
		return nil, errors.New("Failed to get authenticator clone_warning")
	}
	authSignCountRaw, ok := data["authenticator.sign_count"]
	if !ok {
		return nil, errors.New("Failed to get authenticator sign_count")
	}
	transportRaw, ok := data["transport"]
	if !ok {
		return nil, errors.New("Failed to get transport")
	}
	attestionType, ok := data["attestion_type"]
	if !ok {
		return nil, errors.New("Failed to get aattestion_type")
	}

	// Decode Cred ID
	credId, err := base64.StdEncoding.DecodeString(credIdRaw)
	if err != nil {
		return nil, err
	}

	// Decode PublicKeyMultibase
	pubBz, err := base64.StdEncoding.DecodeString(wvm.PublicKeyMultibase)
	if err != nil {
		return nil, err
	}

	// Convert clone warning
	cloneWarn := ConvertStringToBool(authClonedRaw)
	transport := strings.Split(transportRaw, ",")

	// Convert Aaguid from string to bytes
	aaguid, err := base64.StdEncoding.DecodeString(authAaguidRaw)
	if err != nil {
		return nil, err
	}

	// Convert Sign Count
	u64, err := strconv.ParseUint(authSignCountRaw, 10, 32)
	if err != nil {
		return nil, err
	}
	signCount := uint32(u64)
	return &common.WebauthnCredential{
		Id:              credId,
		Transport:       transport,
		PublicKey:       pubBz,
		AttestationType: attestionType,
		Authenticator: &common.WebauthnAuthenticator{
			SignCount:    signCount,
			CloneWarning: cloneWarn,
			Aaguid:       aaguid,
		},
	}, nil
}

func (wvm *VerificationMethod) WebAuthnID() []byte {
	if wvm.Type == KeyType_KeyType_WEB_AUTHN_AUTHENTICATION_2018 {
		return []byte(wvm.ID)
	}
	return nil
}

func (wvm *VerificationMethod) WebAuthnName() string {
	if wvm.Type == KeyType_KeyType_WEB_AUTHN_AUTHENTICATION_2018 {
		return "Sonr"
	}
	return ""
}

func (wvm *VerificationMethod) WebAuthnDisplayName() string {
	if wvm.Type == KeyType_KeyType_WEB_AUTHN_AUTHENTICATION_2018 {
		return wvm.ID
	}
	return ""
}

func (wvm *VerificationMethod) WebAuthnIcon() string {
	if wvm.Type == KeyType_KeyType_WEB_AUTHN_AUTHENTICATION_2018 {
		return "https://raw.githubusercontent.com/sonr-hq/sonr/master/docs/static/favicon.png"
	}
	return ""
}

func (wvm *VerificationMethod) WebAuthnCredentials() []webauthn.Credential {
	return []webauthn.Credential{
		common.ConvertToWebauthnCredential(wvm.WebauthnCredential),
	}
}
