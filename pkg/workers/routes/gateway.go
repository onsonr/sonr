package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/nebula/routes"
)

func RegisterGatewayAPI(e *echo.Echo) {
}

func RegisterGatewayViews(e *echo.Echo) {
	e.GET("/", routes.HomeRoute)
	e.GET("/login", routes.LoginModalRoute)
	e.GET("/register", routes.RegisterModalRoute)
}
