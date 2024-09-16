//go:build js && wasm
// +build js,wasm

package main

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/dwn/wasm"

	svc "github.com/onsonr/sonr/internal/dwn/handlers"
	mdw "github.com/onsonr/sonr/internal/dwn/middleware"
	"github.com/onsonr/sonr/internal/dwn/views"
)

func main() {
	e := New()
	wasm.Serve(e)
}

func registerFrontend(e *echo.Echo) {
	// Add Public Pages
	e.GET("/", views.HomeView)
	e.GET("/login", views.LoginView)
	e.POST("/login/:identifier", svc.HandleCredentialAssertion)
	e.GET("/register", views.RegisterView)
	e.POST("/register/:subject", svc.HandleCredentialCreation)
	e.POST("/register/:subject/check", svc.CheckSubjectIsValid)
	e.GET("/profile", views.ProfileView)
}

func registerOpenID(g *echo.Group) {
	// Add Authenticated Pages
	g.Use(mdw.MacaroonMiddleware("test", "test"))
	g.GET("/", views.AuthorizeView)
	g.GET("/discovery", svc.GetDiscovery)
	g.GET("/jwks", svc.GetJWKS)
	g.GET("/token", svc.GetToken)
	g.POST("/:origin/grant/:subject", svc.GrantAuthorization)
}

func New() *echo.Echo {
	e := echo.New()
	e.Use(mdw.UseSession)
	registerFrontend(e)
	registerOpenID(e.Group("/authorize"))
	return e
}
