package client

import (
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/nebula/components/auth"
	"github.com/onsonr/sonr/nebula/components/home"
)

func RegisterAPI(e *echo.Echo) {
	g := e.Group("state")
	g.POST("/login/:identifier", handleCredentialAssertion)
	g.GET("/jwks", GetJWKS)
	g.GET("/token", GetToken)
	g.POST("/:origin/grant/:subject", GrantAuthorization)
	g.POST("/register/:subject", HandleCredentialCreation)
	g.POST("/register/:subject/check", CheckSubjectIsValid)
}

func RegisterViews(e *echo.Echo) {
	e.GET("/home", home.Route)
	e.GET("/login", auth.LoginRoute)
	e.GET("/register", auth.RegisterRoute)
}
