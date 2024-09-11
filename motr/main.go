//go:build wasm

package main

import (
	"github.com/labstack/echo/v4"
	"github.com/syumai/workers"

	"github.com/onsonr/sonr/internal/gui/views"
	"github.com/onsonr/sonr/internal/mdw"
	"github.com/onsonr/sonr/internal/svc"
)

func main() {
	// Configure the server
	e := echo.New()

	// Use Middlewares
	e.Use(mdw.UseSession)

	// Setup routes
	registerRoutes(e)

	// Serve Worker
	workers.Serve(e)
}

func registerRoutes(e *echo.Echo) {
	// Add Public Pages
	e.GET("/", views.HomeView)
	e.GET("/login", views.LoginView)
	e.POST("/login/:identifier", svc.HandleCredentialAssertion)
	e.GET("/register", views.RegisterView)
	e.POST("/register/:subject", svc.HandleCredentialCreation)
	e.POST("/register/:subject/check", svc.CheckSubjectIsValid)

	// Add Authenticated Pages
	e.GET("/authorize", views.AuthorizeView)
	e.GET("/authorize/discovery", svc.GetDiscovery)
	e.GET("/authorize/jwks", svc.GetJWKS)
	e.GET("/authorize/token", svc.GetToken)
	e.POST("/authorize/:origin/grant/:subject", svc.GrantAuthorization)
}
