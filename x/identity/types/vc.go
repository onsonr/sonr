package types

import (
	"encoding/base64"
	"errors"
	"strconv"
	"strings"

	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/shengdoushi/base58"
	common "github.com/sonrhq/core/pkg/common"
)

func (wvm *VerificationMethod) WebAuthnCredential() (*common.WebauthnCredential, error) {
	if wvm.Type != KeyType_KeyType_WEB_AUTHN_AUTHENTICATION_2018 {
		return nil, errors.New("VerificationMethod is not a WebAuthn Credential")
	}
	data := KeyValueListToMap(wvm.GetMetadata())
	// Fetch Properties from map
	if data == nil {
		return nil, errors.New("Failed to find metadata for VerificationMethod")
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
	credId, err := base58.Decode(fetchFinalDidPath(wvm.Id), base58.BitcoinAlphabet)
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

func (wvm *DidDocument) WebAuthnID() []byte {
	return []byte(wvm.Id)
}

func (wvm *DidDocument) WebAuthnDisplayName() string {
	if len(wvm.AlsoKnownAs) == 0 {
		return wvm.Id
	}
	return wvm.AlsoKnownAs[0]
}

func (wvm *DidDocument) WebAuthnName() string {
	return "Sonr"
}

func (wvm *DidDocument) WebAuthnIcon() string {
	return "https://raw.githubusercontent.com/sonrhq/core/master/docs/static/favicon.png"
}

func (wvm *DidDocument) WebAuthnCredentials() []webauthn.Credential {
	ccreds := wvm.GetCommonWebauthCredentials()
	wac := []webauthn.Credential{}
	for _, wc := range ccreds {
		wac = append(wac, *wc.ToStdCredential())
	}
	return wac
}

func (d *DidDocument) GetCommonWebauthCredentials() []*common.WebauthnCredential {
	creds := []*common.WebauthnCredential{}
	for _, vm := range d.VerificationMethod {
		if vm.Type == KeyType_KeyType_WEB_AUTHN_AUTHENTICATION_2018 {
			vmcr, err := vm.WebAuthnCredential()
			if err != nil {
				return nil
			}
			creds = append(creds, vmcr)
		}
	}
	return creds
}

func fetchFinalDidPath(path string) string {
	parts := strings.Split(path, ":")
	return parts[len(parts)-1]
}
