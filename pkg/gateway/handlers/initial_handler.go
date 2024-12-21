package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/gateway/middleware"
)

func HandleIndex(c echo.Context) error {
	id := middleware.GetSessionID(c)
	if id == "" {
		middleware.NewSession(c)
	}
	return middleware.RenderInitial(c)
}
