package session

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/common"
	"github.com/onsonr/sonr/pkg/common/types"
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
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "DWN Context not found")
	}
	return ctx, nil
}

// WithHTTPContext sets the HTTP context in the context
func WithData(ctx context.Context, data *types.Session) context.Context {
	return context.WithValue(ctx, DataContextKey, data)
}

// GetData gets the HTTP context from the context
func GetData(ctx context.Context) *types.Session {
	if httpCtx, ok := ctx.Value(DataContextKey).(*types.Session); ok {
		return httpCtx
	}
	return nil
}
