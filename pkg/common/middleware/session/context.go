package session

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/motr/config"
)

type Context interface {
	ID() string

	GetLoginParams() *LoginOptions
	GetRegisterParams() *RegisterOptions

	Address() string
	ChainID() string

	Schema() *config.Schema
}

// Get returns the session.Context from the echo context.
func Get(c echo.Context) (Context, error) {
	ctx, ok := c.(*HTTPContext)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "DWN Context not found")
	}
	return ctx, nil
}
