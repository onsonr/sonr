package nebula

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/nebula/routes"
	"github.com/onsonr/sonr/pkg/nebula/view"
)

func RouteViews(e *echo.Echo) {
	e.GET("/home", view.Render(routes.HomeRoute))
	e.GET("/login", view.Render(routes.LoginRoute))
	e.GET("/register", view.Render(routes.RegisterRoute))
	e.GET("/profile", view.Render(routes.ProfileRoute))
	e.GET("/authorize", view.Render(routes.AuthorizeRoute))
}
