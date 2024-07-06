package handlers

import (
	pages "github.com/onsonr/hway/pkg/vault/components"
	"github.com/onsonr/hway/pkg/vault/middleware"
	"github.com/labstack/echo/v4"
)

var Login = loginHandler{}

type loginHandler struct{}

func (h loginHandler) Page(e echo.Context) error {
	return middleware.Render(e, pages.Login(middleware.SessionID(e)))
}
