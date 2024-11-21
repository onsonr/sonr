package gateway

import (
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/webapp"
)

func RegisterFrontendViews(e *echo.Echo) {
	e.GET("/", webapp.HomePage)
	e.GET("/login", webapp.LoginModal)
	e.GET("/register", webapp.RegisterModal)
}
