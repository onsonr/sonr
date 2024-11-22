package session

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/common"
)

type contextKey string

// Context keys
const (
	HTTPContextKey contextKey = "http_context"
)

type Context = common.SessionCtx

// Get returns the session.Context from the echo context.
func Get(c echo.Context) (Context, error) {
	ctx, ok := c.(*HTTPContext)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "DWN Context not found")
	}
	return loadHTTPContext(ctx), nil
}

// WithHTTPContext sets the HTTP context in the context
func WithHTTPContext(ctx context.Context, httpCtx *HTTPContext) context.Context {
	return context.WithValue(ctx, HTTPContextKey, httpCtx)
}

// GetHTTPContext gets the HTTP context from the context
func GetHTTPContext(ctx context.Context) *HTTPContext {
	if httpCtx, ok := ctx.Value(HTTPContextKey).(*HTTPContext); ok {
		return httpCtx
	}
	return nil
}
