package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/nebula/components/authentication"
	"github.com/onsonr/sonr/pkg/workers/handlers"
)

func RegisterClientAPI(e *echo.Echo) {
	g1 := e.Group("api")
	g1.GET("/register/:subject/start", handlers.RegisterSubjectStart)
	g1.POST("/register/:subject/check", handlers.RegisterSubjectCheck)
	g1.POST("/register/:subject/finish", handlers.RegisterSubjectFinish)

	g1.GET("/login/:subject/start", handlers.LoginSubjectStart)
	g1.POST("/login/:subject/check", handlers.LoginSubjectCheck)
	g1.POST("/login/:subject/finish", handlers.LoginSubjectFinish)

	g1.GET("/jwks", handlers.GetJWKS)
	g1.GET("/token", handlers.GetToken)
	g1.POST("/:origin/grant/:subject", handlers.GrantAuthorization)
}

func RegisterClientViews(e *echo.Echo) {
	// Static Routes
	e.File("/", "index.html")
	e.File("/config.json", "config.json")

	// DWN-Side Routes
	e.GET("/#", authentication.CurrentRoute)
	e.GET("/login", authentication.LoginRoute)
	e.GET("/register", authentication.RegisterRoute)
}
