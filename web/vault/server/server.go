//go:build js && wasm
// +build js,wasm

package server

import (
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/common/middleware/session"
	"github.com/onsonr/sonr/pkg/core/dwn"

	"github.com/onsonr/sonr/pkg/core/dwn/bridge"
	"github.com/onsonr/sonr/pkg/core/dwn/handlers"
)

// Server is the interface that wraps the Serve function.
type Server interface {
	Serve() func()
}

type MotrServer struct {
	e *echo.Echo

	WasmPath       string
	WasmExecPath   string
	HTTPServerPath string
	CacheVersion   string
	IsDev          bool
}

func New(env *dwn.Environment, config *dwn.Config) Server {
	s := &MotrServer{e: echo.New()}

	s.e.Use(session.MotrMiddleware(config))
	s.e.Use(bridge.WasmContextMiddleware)

	// Add WASM-specific routes
	registerAPI(s.e)
	return s
}

func (s *MotrServer) Serve() func() {
	return bridge.ServeFetch(s.e)
}

// registerAPI registers the Decentralized Web Node API routes.
func registerAPI(e *echo.Echo) {
	g1 := e.Group("api")
	g1.GET("/register/:subject/start", handlers.RegisterSubjectStart)
	g1.POST("/register/:subject/check", handlers.RegisterSubjectCheck)
	g1.POST("/register/:subject/finish", handlers.RegisterSubjectFinish)

	g1.GET("/login/:subject/start", handlers.LoginSubjectStart)
	g1.POST("/login/:subject/check", handlers.LoginSubjectCheck)
	g1.POST("/login/:subject/finish", handlers.LoginSubjectFinish)

	g1.GET("/:origin/grant/jwks", handlers.GetJWKS)
	g1.GET("/:origin/grant/token", handlers.GetToken)
	g1.POST("/:origin/grant/:subject", handlers.GrantAuthorization)
}
