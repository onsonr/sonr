package session

import (
	"net/http"

	"github.com/labstack/echo/v4"

	commonv1 "github.com/onsonr/sonr/pkg/common/types"
)

type Context interface {
	ID() string

	LoginOptions(credentials []commonv1.CredDescriptor) *commonv1.LoginOptions
	RegisterOptions(subject string) *commonv1.RegisterOptions

	ClientConfig() *commonv1.ClientConfig
	UserAgent() *commonv1.UserAgent
	VaultDetails() *commonv1.VaultDetails
}

// Get returns the session.Context from the echo context.
func Get(c echo.Context) (Context, error) {
	ctx, ok := c.(*HTTPContext)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "DWN Context not found")
	}
	return loadHTTPContext(ctx), nil
}
