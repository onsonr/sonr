package types

import (
	"github.com/go-webauthn/webauthn/webauthn"
	common "github.com/sonr-hq/sonr/pkg/common"
)

func (wvm *VerificationMethod) WebAuthnID() []byte {
	return []byte(wvm.ID)
}

func (wvm *VerificationMethod) WebAuthnName() string {
	return wvm.ID
}

func (wvm *VerificationMethod) WebAuthnDisplayName() string {
	return wvm.ID
}

func (wvm *VerificationMethod) WebAuthnIcon() string {
	return ""
}

func (wvm *VerificationMethod) WebAuthnCredentials() []webauthn.Credential {
	return []webauthn.Credential{
		common.ConvertToWebauthnCredential(wvm.WebauthnCredential),
	}
}
