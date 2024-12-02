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
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Session Context not found")
	}
	return ctx, nil
}

// WithData sets the session data in the context
func WithData(ctx context.Context, data *types.Session) context.Context {
	return context.WithValue(ctx, DataContextKey, data)
}

// GetData gets the session data from any context type
func GetData(ctx interface{}) *types.Session {
	switch c := ctx.(type) {
	case *HTTPContext:
		if c != nil {
			return c.sessionData
		}
	case context.Context:
		if c != nil {
			if val := c.Value(DataContextKey); val != nil {
				if httpCtx, ok := val.(*types.Session); ok {
					return httpCtx
				}
			}
		}
	case echo.Context:
		if c != nil {
			if httpCtx, ok := c.(*HTTPContext); ok && httpCtx != nil {
				return httpCtx.sessionData
			}
		}
	}
	// Return empty session rather than nil to prevent nil pointer panics
	return &types.Session{}
}
