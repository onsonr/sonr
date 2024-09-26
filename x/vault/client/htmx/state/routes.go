package state

import (
	"github.com/labstack/echo/v4"
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
