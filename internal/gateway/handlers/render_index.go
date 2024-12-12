package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/gateway/context"
	"github.com/onsonr/sonr/internal/gateway/views"
	"github.com/onsonr/sonr/pkg/common/response"
)

func RenderIndex(c echo.Context) error {
	if isReturning(c) {
		return response.TemplEcho(c, views.ReturningView())
	}
	return response.TemplEcho(c, views.InitialView())
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
