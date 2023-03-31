package types

import (
	"encoding/base64"
	"errors"
	"strconv"
	"strings"

	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/shengdoushi/base58"
	"github.com/sonr-io/sonr/pkg/crypto"
)

func ConvertBoolToString(v bool) string {
	if v {
		return "TRUE"
	} else {
		return "FALSE"
	}
}

func ConvertStringToBool(v string) bool {
	if v == "TRUE" {
		return true
	}
	return false
}

func (wvm *VerificationMethod) WebAuthnCredential() (*crypto.WebauthnCredential, error) {
	if wvm.Type != crypto.WebAuthnKeyType.PrettyString() {
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
	return &crypto.WebauthnCredential{
		Id:              credId,
		Transport:       transport,
		PublicKey:       pubBz,
		AttestationType: attestionType,
		Authenticator: &crypto.WebauthnAuthenticator{
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

func (d *DidDocument) GetCommonWebauthCredentials() []*crypto.WebauthnCredential {
	creds := []*crypto.WebauthnCredential{}
	for _, vm := range d.VerificationMethod {
		if vm.Type == crypto.WebAuthnKeyType.PrettyString() {
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
func MapToKeyValueList(m map[string]string) []*KeyValuePair {
	kvs := make([]*KeyValuePair, 0)
	for k, v := range m {
		kvs = append(kvs, &KeyValuePair{Key: k, Value: v})
	}
	return kvs
}

func KeyValueListToMap(kvs []*KeyValuePair) map[string]string {
	m := make(map[string]string)
	for _, kv := range kvs {
		m[kv.Key] = kv.Value
	}
	return m
}
