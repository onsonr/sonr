package webapp

import (
	"github.com/labstack/echo/v4"
)

func RegisterLandingRoutes(e *echo.Echo) {
	e.GET("/", HomePage)
	e.GET("/login", LoginModal)
	e.GET("/register", RegisterModal)
}

func RegisterVaultRoutes(e *echo.Echo) {
	e.GET("/", IndexPage)
	e.GET("/sw.js", IndexPage)
}
