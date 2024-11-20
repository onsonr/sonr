//go:build js && wasm
// +build js,wasm

package server

import (
	"github.com/labstack/echo/v4"
	"github.com/syumai/workers"

	"github.com/onsonr/sonr/cmd/hway/routes"
	"github.com/onsonr/sonr/pkg/common/middleware/session"
	"github.com/onsonr/sonr/pkg/core/dwn"
)

// Server is the interface that wraps the Serve function.
type Server interface {
	Ctx(c echo.Context) session.Context
	Serve()

	loadEnv(e *dwn.Environment) Server
}

type HwayServer struct {
	e *echo.Echo

	WasmPath       string
	WasmExecPath   string
	HTTPServerPath string
	CacheVersion   string
	IsDev          bool
}

func New(env *dwn.Environment) Server {
	s := &HwayServer{e: echo.New()}

	s.e.Use(session.HwayMiddleware())

	// Add WASM-specific routes
	routes.RegisterGatewayAPI(s.e)
	routes.RegisterFrontendViews(s.e)
	return s.loadEnv(env)
}

func (s *HwayServer) loadEnv(e *dwn.Environment) Server {
	s.WasmPath = e.WasmPath
	s.WasmExecPath = e.WasmExecPath
	s.HTTPServerPath = e.HttpserverPath
	s.CacheVersion = e.CacheVersion
	s.IsDev = e.IsDevelopment
	return s
}

func (s *HwayServer) Ctx(c echo.Context) session.Context {
	cc, _ := session.Get(c)
	return cc
}

func (s *HwayServer) Serve() {
	workers.Serve(s.e)
}
