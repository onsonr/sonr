//go:build js && wasm
// +build js,wasm

package middleware

import (
	"github.com/donseba/go-htmx"
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/dwn/middleware/jsexc"
)

type Browser struct {
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
	push           jsexc.PushAPI
	sessionStorage jsexc.SessionStorageAPI
}

func UseNavigator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := jsexc.NewNavigator(c)
		return next(cc)
	}
}
