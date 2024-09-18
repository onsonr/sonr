package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/dwn/handlers/state"
	middleware "github.com/onsonr/sonr/internal/dwn/middleware"
)

func RegisterState(e *echo.Echo) {
	g := e.Group("state")
	g.POST("/login/:identifier", state.HandleCredentialAssertion)
	//	g.GET("/discovery", state.GetDiscovery)
	g.GET("/jwks", state.GetJWKS)
	g.GET("/token", state.GetToken)
	g.POST("/:origin/grant/:subject", state.GrantAuthorization)
	g.POST("/register/:subject", state.HandleCredentialCreation)
	g.POST("/register/:subject/check", state.CheckSubjectIsValid)
}

func RegisterSync(e *echo.Echo) {
	g := e.Group("sync")
	g.Use(middleware.MacaroonMiddleware("test", "test"))
}
