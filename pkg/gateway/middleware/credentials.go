package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	hwayorm "github.com/onsonr/sonr/pkg/gateway/orm"
)

func ListCredentials(c echo.Context, handle string) ([]*CredentialDescriptor, error) {
	cc, ok := c.(*GatewayContext)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Credentials Context not found")
	}
	creds, err := cc.dbq.GetCredentialsByHandle(bgCtx(), handle)
	if err != nil {
		return nil, err
	}
	return CredentialArrayToDescriptors(creds), nil
}

func SubmitCredential(c echo.Context, cred *CredentialDescriptor) error {
	origin := GetOrigin(c)
	handle := GetHandle(c)
	md := cred.ToModel(handle, origin)

	cc, ok := c.(*GatewayContext)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Credentials Context not found")
	}

	_, err := cc.dbq.InsertCredential(bgCtx(), hwayorm.InsertCredentialParams{
		Handle:       handle,
		CredentialID: md.CredentialID,
		Origin:       origin,
		Type:         md.Type,
		Transports:   md.Transports,
	})
	if err != nil {
		return err
	}
	return nil
}

// Define the credential structure matching our frontend data
type CredentialDescriptor struct {
	ID                      string            `json:"id"`
	RawID                   string            `json:"rawId"`
	Type                    string            `json:"type"`
	AuthenticatorAttachment string            `json:"authenticatorAttachment"`
	Transports              string            `json:"transports"`
	ClientExtensionResults  map[string]string `json:"clientExtensionResults"`
	Response                struct {
		AttestationObject string `json:"attestationObject"`
		ClientDataJSON    string `json:"clientDataJSON"`
	} `json:"response"`
}

func (c *CredentialDescriptor) ToModel(handle, origin string) *hwayorm.Credential {
	return &hwayorm.Credential{
		Handle:                  handle,
		Origin:                  origin,
		CredentialID:            c.ID,
		Type:                    c.Type,
		Transports:              c.Transports,
		AuthenticatorAttachment: c.AuthenticatorAttachment,
	}
}

func CredentialArrayToDescriptors(credentials []hwayorm.Credential) []*CredentialDescriptor {
	var descriptors []*CredentialDescriptor
	for _, cred := range credentials {
		cd := &CredentialDescriptor{
			ID:                      cred.CredentialID,
			RawID:                   cred.CredentialID,
			Type:                    cred.Type,
			AuthenticatorAttachment: cred.AuthenticatorAttachment,
			Transports:              cred.Transports,
		}
		descriptors = append(descriptors, cd)
	}
	return descriptors
}
