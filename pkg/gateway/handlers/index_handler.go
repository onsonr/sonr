package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/gateway/context"
)

func HandleIndex(c echo.Context) error {
	id := context.GetSessionID(c)
	if id == "" {
		context.NewSession(c)
	}
	return context.RenderInitial(c)
}
