package front

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/dwn/front/views"
)

func RegisterViews(e *echo.Echo) {
	e.GET("/", views.HomeView)
	e.GET("/login", views.LoginView)
	e.GET("/register", views.RegisterView)
	e.GET("/profile", views.ProfileView)
	e.GET("/authorize", views.AuthorizeView)
}
