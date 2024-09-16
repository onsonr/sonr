//go:build js && wasm
// +build js,wasm

package mdw

import (
	"github.com/donseba/go-htmx"
	"github.com/labstack/echo/v4"
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
	credentials    CredentialsAPI
	indexedDB      IndexedDBAPI
	localStorage   LocalStorageAPI
	push           PushAPI
	sessionStorage SessionStorageAPI
}
