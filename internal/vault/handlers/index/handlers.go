package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/vault/context"
)

func Handler(c echo.Context) error {
	// TODO: Create views
	if isInitial(c) {
		// return response.TemplEcho(c, index.InitialView())
	}
	// TODO: Add authorization check
	if isExpired(c) {
		// return response.TemplEcho(c, index.ReturningView())
	}
	return c.Render(http.StatusOK, "index.templ", nil)
}

// ╭─────────────────────────────────────────────────────────╮
// │                    Utility Functions                    │
// ╰─────────────────────────────────────────────────────────╯

// Initial users have no authorization, user handle, or vault address
func isInitial(c echo.Context) bool {
	noAuth := !context.HasAuthorization(c)
	noUserHandle := !context.HasUserHandle(c)
	noVaultAddress := !context.HasVaultAddress(c)
	return noUserHandle && noVaultAddress && noAuth
}

// Expired users have either a user handle or vault address
func isExpired(c echo.Context) bool {
	noAuth := !context.HasAuthorization(c)
	hasUserHandle := context.HasUserHandle(c)
	hasVaultAddress := context.HasVaultAddress(c)
	return noAuth && hasUserHandle || noAuth && hasVaultAddress
}

// Returning users have a valid authorization, and either a user handle or vault address
func isReturning(c echo.Context) bool {
	hasAuth := context.HasAuthorization(c)
	hasUserHandle := context.HasUserHandle(c)
	hasVaultAddress := context.HasVaultAddress(c)
	return hasAuth && (hasUserHandle || hasVaultAddress)
}
