package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/gateway/models"
	"github.com/onsonr/sonr/internal/gateway/models/repository"
)

type CredentialsContext struct {
	echo.Context
	dbq *repository.Queries
}

func Middleware(dbq *repository.Queries) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := &CredentialsContext{
				Context: c,
				dbq:     dbq,
			}
			return next(ctx)
		}
	}
}

func ListCredentials(c echo.Context, handle string) ([]*models.CredentialDescriptor, error) {
	cc, ok := c.(*CredentialsContext)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Credentials Context not found")
	}
	creds, err := cc.dbq.GetCredentialsByHandle(bgCtx(), handle)
	if err != nil {
		return nil, err
	}
	return models.CredentialArrayToDescriptors(creds), nil
}

func SubmitCredential(c echo.Context, cred *models.CredentialDescriptor) error {
	origin := GetOrigin(c)
	handle := GetHandle(c)
	md := cred.ToModel(handle, origin)

	cc, ok := c.(*CredentialsContext)
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
