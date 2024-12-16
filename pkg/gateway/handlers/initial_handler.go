package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/gateway/middleware"
)

func RenderIndex(c echo.Context) error {
	// Initialize the session
	err := middleware.NewSession(c)
	if err != nil {
		return middleware.RenderError(c, err)
	}
	// Render the initial view
	return middleware.RenderInitial(c)
}
