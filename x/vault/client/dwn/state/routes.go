package state

import (
	"github.com/labstack/echo/v4"
	middleware "github.com/onsonr/sonr/x/vault/client/dwn/middleware"
)

func RegisterHandlers(e *echo.Echo) {
	g := e.Group("state")
	g.POST("/login/:identifier", handleCredentialAssertion)
	//	g.GET("/discovery", state.GetDiscovery)
	g.GET("/jwks", getJWKS)
	g.GET("/token", getToken)
	g.POST("/:origin/grant/:subject", grantAuthorization)
	g.POST("/register/:subject", handleCredentialCreation)
	g.POST("/register/:subject/check", checkSubjectIsValid)
}

func RegisterSync(e *echo.Echo) {
	g := e.Group("sync")
	g.Use(middleware.MacaroonMiddleware("test", "test"))
}
