//go:build js && wasm
// +build js,wasm

package server

import (
	"github.com/labstack/echo/v4"
	"github.com/syumai/workers"

	"github.com/onsonr/sonr/pkg/common/middleware/session"
)

// Server is the interface that wraps the Serve function.
type Server interface {
	Ctx(c echo.Context) session.Context
	Serve()
}

type HwayServer struct {
	e *echo.Echo

	WasmPath       string
	WasmExecPath   string
	HTTPServerPath string
	CacheVersion   string
	IsDev          bool
}

func New() Server {
	s := &HwayServer{e: echo.New()}

	s.e.Use(session.HwayMiddleware())

	// Add WASM-specific routes
	RegisterGatewayAPI(s.e)
	RegisterFrontendViews(s.e)
	return s
}

func (s *HwayServer) Ctx(c echo.Context) session.Context {
	cc, _ := session.Get(c)
	return cc
}

func (s *HwayServer) Serve() {
	workers.Serve(s.e)
}
