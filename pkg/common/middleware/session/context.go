package session

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/common"
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

// WithTheme sets the theme in the context
func WithTheme(ctx context.Context, theme string) context.Context {
	return context.WithValue(ctx, ThemeKey, theme)
}

// GetTheme gets the theme from the context
func GetTheme(ctx context.Context) string {
	if theme, ok := ctx.Value(ThemeKey).(string); ok {
		return theme
	}
	return ""
}
