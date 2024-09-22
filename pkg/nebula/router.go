package nebula

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/nebula/views"
)

func RouteViews(e *echo.Echo) {
	e.GET("/home", views.HomeView)
	e.GET("/login", views.LoginView)
	e.GET("/register", views.RegisterView)
	e.GET("/profile", views.ProfileView)
	e.GET("/authorize", views.AuthorizeView)
}
