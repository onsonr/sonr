package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/database/repository"
	"github.com/onsonr/sonr/pkg/gateway/types"
)

func ListCredentials(c echo.Context, handle string) ([]*types.CredentialDescriptor, error) {
	cc, ok := c.(*GatewayContext)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Credentials Context not found")
	}
	creds, err := cc.dbq.GetCredentialsByHandle(bgCtx(), handle)
	if err != nil {
		return nil, err
	}
	return types.CredentialArrayToDescriptors(creds), nil
}

func SubmitCredential(c echo.Context, cred *types.CredentialDescriptor) error {
	origin := GetOrigin(c)
	handle := GetHandle(c)
	md := cred.ToModel(handle, origin)

	cc, ok := c.(*GatewayContext)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Credentials Context not found")
	}

	_, err := cc.dbq.InsertCredential(bgCtx(), repository.InsertCredentialParams{
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
