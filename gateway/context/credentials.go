package context

import (
	"github.com/go-webauthn/webauthn/protocol"
)

func (c *GatewayContext) NewChallenge() string {
	chal, _ := protocol.CreateChallenge()
	chalStr := chal.String()
	return chalStr
}

func (cc *GatewayContext) ListCredentials(handle string) ([]*CredentialDescriptor, error) {
	creds, err := cc.GetCredentialsByHandle(bgCtx(), handle)
	if err != nil {
		return nil, err
	}
	return CredentialArrayToDescriptors(creds), nil
}
