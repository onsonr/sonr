package handlers

import (
	"github.com/onsonr/hway/pkg/vault/components"
	"github.com/onsonr/hway/pkg/vault/middleware"
	"github.com/labstack/echo/v4"
)

var Session = sessionHandler{}

type sessionHandler struct{}

func (h sessionHandler) Page(e echo.Context) error {
	return middleware.Render(e, components.Home(middleware.SessionID(e)))
}
