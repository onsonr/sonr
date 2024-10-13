package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/nebula/components/authentication"
	"github.com/onsonr/sonr/pkg/nebula/components/marketing"
)

func RegisterGatewayAPI(e *echo.Echo) {
}

func RegisterGatewayViews(e *echo.Echo) {
	e.GET("/", marketing.HomeRoute)
	e.GET("/login", authentication.LoginModalRoute)
	e.GET("/register", authentication.RegisterModalRoute)
}
