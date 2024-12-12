package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/gateway/views"
	"github.com/onsonr/sonr/pkg/common/response"
)

func RenderIndex(c echo.Context) error {
	if isExpired(c) {
		return response.TemplEcho(c, views.ReturningView())
	}
	return response.TemplEcho(c, views.InitialView())
}
