//go:build js && wasm
// +build js,wasm

package server

import (
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/cmd/motr/server/bridge"
	"github.com/onsonr/sonr/cmd/motr/server/routes"
	"github.com/onsonr/sonr/pkg/common/middleware/session"
	"github.com/onsonr/sonr/pkg/core/dwn"
)

// Server is the interface that wraps the Serve function.
type Server interface {
	Ctx(c echo.Context) session.Context
	Serve() func()

	loadEnv(e *dwn.Environment) Server
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
	routes.RegisterServerAPI(s.e)
	routes.RegisterFrontendViews(s.e)
	return s.loadEnv(env)
}

func (s *MotrServer) loadEnv(e *dwn.Environment) Server {
	s.WasmPath = e.WasmPath
	s.WasmExecPath = e.WasmExecPath
	s.HTTPServerPath = e.HttpserverPath
	s.CacheVersion = e.CacheVersion
	s.IsDev = e.IsDevelopment
	return s
}

func (s *MotrServer) Ctx(c echo.Context) session.Context {
	cc, _ := session.Get(c)
	return cc
}

func (s *MotrServer) Serve() func() {
	return bridge.ServeFetch(s.e)
}
