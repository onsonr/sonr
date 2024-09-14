//go:build wasm

package main

import (
	"github.com/labstack/echo/v4"
	"github.com/syumai/workers"

	"github.com/onsonr/sonr/internal/db"

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
	registerFrontend(e)
	registerOpenID(e.Group("/authorize"))
	registerVault(e.Group("/vault"))

	// Serve Worker
	workers.Serve(e)
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

func registerVault(g *echo.Group) {
	// Add Authenticated Pages
	g.Use(mdw.MacaroonMiddleware("test", "test"))
	vault, err := db.New(db.WitDir("vault"))
	if err != nil {
		// panic(err)
	}
	vault.ServeEcho(g)
}
