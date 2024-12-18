package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/gateway/middleware"
)

func HandleIndex(c echo.Context) error {
	id := middleware.GetSessionID(c)
	if id == "" {
		return startNewSession(c)
	}
	return middleware.RenderInitial(c)
}

func startNewSession(c echo.Context) error {
	// Initialize the session
	err := middleware.NewSession(c)
	if err != nil {
		return middleware.RenderError(c, err)
	}
	return middleware.RenderInitial(c)
}

func continueExistingSession(c echo.Context, id string) error {
	// Do some auth checks here
	return middleware.RenderInitial(c)
}
