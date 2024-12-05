package session

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/common"
)

type contextKey string

// Context keys
const (
	DataContextKey contextKey = "http_session_data"
)

type Context = common.SessionCtx

// Get returns the session.Context from the echo context.
func Get(c echo.Context) (Context, error) {
	ctx, ok := c.(*HTTPContext)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Session Context not found")
	}
	return ctx, nil
}
