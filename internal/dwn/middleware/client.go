//go:build js && wasm
// +build js,wasm

package middleware

import (
	"github.com/donseba/go-htmx"
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/config/dwn"
	"github.com/onsonr/sonr/internal/dwn/middleware/jsexc"
)

type Client struct {
	echo.Context
	isMobile  bool
	userAgent string
	width     int
	olc       string
	ksuid     string

	// HTMX Specific
	htmx *htmx.HTMX

	// WebAPIs
	indexedDB      jsexc.IndexedDBAPI
	localStorage   jsexc.LocalStorageAPI
	sessionStorage jsexc.SessionStorageAPI
}

func UseNavigator(next echo.HandlerFunc, cnfg *dwn.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := jsexc.NewNavigator(c, cnfg)
		return next(cc)
	}
}
