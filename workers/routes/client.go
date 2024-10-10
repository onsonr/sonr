package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/nebula/components/auth"
	"github.com/onsonr/sonr/nebula/components/home"
	"github.com/onsonr/sonr/workers/handlers"
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
	e.GET("/home", home.Route)
	e.GET("/login", auth.LoginRoute)
	e.GET("/register", auth.RegisterRoute)
}
