package index

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/common/response"
)

func RenderHandler(c echo.Context) error {
	if isExpired(c) {
		return response.TemplEcho(c, ReturningView())
	}
	return response.TemplEcho(c, InitialView())
}
