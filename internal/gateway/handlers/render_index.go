package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/gateway/context"
	"github.com/onsonr/sonr/internal/gateway/views"
	"github.com/onsonr/sonr/pkg/common/response"
)

func RenderIndex(c echo.Context) error {
	return response.TemplEcho(c, views.InitialView(isUnavailableDevice(c)))
}

// isUnavailableDevice returns true if the device is unavailable
func isUnavailableDevice(c echo.Context) bool {
	s, err := context.Get(c)
	if err != nil {
		return true
	}
	return s.IsBot() || s.IsTV()
}
