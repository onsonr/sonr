package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/nebula/components/authentication"
	"github.com/onsonr/sonr/pkg/workers/handlers"
)

func RegisterWebNodeAPI(e *echo.Echo) {
	g1 := e.Group("api")
	g1.GET("/register/:subject/start", handlers.Auth.RegisterSubjectStart)
	g1.POST("/register/:subject/check", handlers.Auth.RegisterSubjectCheck)
	g1.POST("/register/:subject/finish", handlers.Auth.RegisterSubjectFinish)

	g1.GET("/login/:subject/start", handlers.Auth.LoginSubjectStart)
	g1.POST("/login/:subject/check", handlers.Auth.LoginSubjectCheck)
	g1.POST("/login/:subject/finish", handlers.Auth.LoginSubjectFinish)

	g1.GET("/:origin/grant/jwks", handlers.OpenID.GetJWKS)
	g1.GET("/:origin/grant/token", handlers.OpenID.GetToken)
	g1.POST("/:origin/grant/:subject", handlers.OpenID.GrantAuthorization)
}

func RegisterWebNodeViews(e *echo.Echo) {
	e.File("/", "index.html")
	e.GET("/#", authentication.CurrentViewRoute)
	e.GET("/login", authentication.LoginModalRoute)
	e.File("/config", "config.json")
	e.GET("/register", authentication.RegisterModalRoute)
}
