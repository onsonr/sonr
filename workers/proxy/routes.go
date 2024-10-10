package proxy

import (
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/nebula/components/auth"
	"github.com/onsonr/sonr/nebula/components/home"
)

func RegisterViews(e *echo.Echo) {
	e.GET("/", home.Route)
	e.GET("/login", auth.LoginRoute)
	e.GET("/register", auth.RegisterRoute)
}
