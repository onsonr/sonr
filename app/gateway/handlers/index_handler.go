package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/app/gateway/internal/pages/index"
	"github.com/onsonr/sonr/app/gateway/internal/session"
	"github.com/onsonr/sonr/pkg/common/response"
)

func HandleIndex(c echo.Context) error {
	if isExpired(c) {
		return response.TemplEcho(c, index.ReturningView())
	}
	return response.TemplEcho(c, index.InitialView())
}

// ╭─────────────────────────────────────────────────────────╮
// │                    Utility Functions                    │
// ╰─────────────────────────────────────────────────────────╯

// Initial users have no authorization, user handle, or vault address
func isInitial(c echo.Context) bool {
	sess, err := session.Get(c)
	if err != nil {
		return false
	}
	data := sess.Session()
	return data.UserHandle == "" && data.VaultAddress == ""
}

// Expired users have either a user handle or vault address
func isExpired(c echo.Context) bool {
	sess, err := session.Get(c)
	if err != nil {
		return false
	}
	data := sess.Session()
	return data.UserHandle != "" || data.VaultAddress != ""
}

// Returning users have a valid authorization, and either a user handle or vault address
func isReturning(c echo.Context) bool {
	sess, err := session.Get(c)
	if err != nil {
		return false
	}
	data := sess.Session()
	return data.UserHandle != "" && data.VaultAddress != ""
}
