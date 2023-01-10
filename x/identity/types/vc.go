package types

import (
	"github.com/go-webauthn/webauthn/webauthn"
	common "github.com/sonr-hq/sonr/pkg/common"
)

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
