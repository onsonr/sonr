package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/common/response"
	"github.com/onsonr/sonr/pkg/webui/landing/components/index"
)

func HandleDash(c echo.Context) error {
	if isInitial(c) {
		return response.TemplEcho(c, index.InitialView())
	}
	if isExpired(c) {
		return response.TemplEcho(c, index.ReturningView())
	}
	return c.Render(http.StatusOK, "index.templ", nil)
}
