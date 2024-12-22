package context

import (
	"net/http"

	"github.com/labstack/echo/v4"
	hwayorm "github.com/onsonr/sonr/internal/database/hwayorm"
)

func ListCredentials(c echo.Context, handle string) ([]*CredentialDescriptor, error) {
	cc, ok := c.(*GatewayContext)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Gateway Context not found")
	}
	creds, err := cc.GetCredentialsByHandle(bgCtx(), handle)
	if err != nil {
		return nil, err
	}
	return CredentialArrayToDescriptors(creds), nil
}

func InsertCredential(c echo.Context, cred *CredentialDescriptor) error {
	origin := GetOrigin(c)
	handle := GetProfileHandle(c)
	md := cred.ToModel(handle, origin)
	cc, ok := c.(*GatewayContext)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Credentials Context not found")
	}

	_, err := cc.InsertCredential(bgCtx(), hwayorm.InsertCredentialParams{
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

