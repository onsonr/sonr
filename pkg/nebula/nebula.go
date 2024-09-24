package nebula

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/nebula/pages"
)

func RegisterHandlers(e *echo.Echo) {
	e.GET("/home", pages.Home)
	e.GET("/login", pages.Login)
	e.GET("/register", pages.Register)
	e.GET("/profile", pages.Profile)
	e.GET("/authorize", pages.Authorize)
}
