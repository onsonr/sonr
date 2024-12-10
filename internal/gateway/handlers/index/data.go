package index

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/gateway/context"
)

// Initial users have no authorization, user handle, or vault address
func isInitial(c echo.Context) bool {
	sess, err := context.Get(c)
	if err != nil {
		return false
	}
	data := sess.Session()
	return data.UserHandle == "" && data.VaultAddress == ""
}

// Expired users have either a user handle or vault address
func isExpired(c echo.Context) bool {
	sess, err := context.Get(c)
	if err != nil {
		return false
	}
	data := sess.Session()
	return data.UserHandle != "" || data.VaultAddress != ""
}

// Returning users have a valid authorization, and either a user handle or vault address
func isReturning(c echo.Context) bool {
	sess, err := context.Get(c)
	if err != nil {
		return false
	}
	data := sess.Session()
	return data.UserHandle != "" && data.VaultAddress != ""
}
