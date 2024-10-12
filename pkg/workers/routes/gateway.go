package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/nebula/components/authentication"
	"github.com/onsonr/sonr/pkg/nebula/components/marketing"
)

func RegisterProxyAPI(e *echo.Echo) {
}

func RegisterProxyViews(e *echo.Echo) {
	e.GET("/", marketing.Route)
	e.GET("/login", authentication.LoginRoute)
	e.GET("/register", authentication.RegisterRoute)
}
