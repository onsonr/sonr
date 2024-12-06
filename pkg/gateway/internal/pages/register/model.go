package register

import (
	"github.com/a-h/templ"
	"github.com/go-webauthn/webauthn/protocol"
)

type LinkCredentialRequest struct {
	Platform        string                                      `json:"platform"`
	Handle          string                                      `json:"handle"`
	DeviceModel     string                                      `json:"deviceModel"`
	Architecture    string                                      `json:"architecture"`
	Address         string                                      `json:"address"`
	RegisterOptions protocol.PublicKeyCredentialCreationOptions `json:"registerOptions"`
}

func (r LinkCredentialRequest) GetCredentialOptions() string {
	opts, _ := templ.JSONString(r.RegisterOptions)
	return opts
}
