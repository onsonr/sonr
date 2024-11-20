package handlers

import (
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/app/nebula/views/auth"
	"github.com/onsonr/sonr/cmd/motr/handlers/view"
	"github.com/onsonr/sonr/pkg/common/middleware/render"
	"github.com/onsonr/sonr/pkg/common/middleware/session"
)

// ╭───────────────────────────────────────────────────────────╮
// │               DWN Routes - Authentication                 │
// ╰───────────────────────────────────────────────────────────╯

// ╭───────────────────────────────────────────────────────────╮
// │              Hway Routes - Authentication                 │
// ╰───────────────────────────────────────────────────────────╯

// AuthorizeModalRoute returns the Authorize Modal route.
func AuthorizeModalRoute(c echo.Context) error {
	cc, err := session.Get(c)
	if err != nil {
		return err
	}
	return render.Templ(c, auth.AuthorizeModal(cc))
}

// LoginModalRoute returns the Login Modal route.
func LoginModalRoute(c echo.Context) error {
	cc, err := session.Get(c)
	if err != nil {
		return err
	}
	return render.Templ(c, auth.LoginModal(cc))
}

// RegisterModalRoute returns the Register Modal route.
func RegisterModalRoute(c echo.Context) error {
	cc, err := session.Get(c)
	if err != nil {
		return err
	}
	return render.Templ(c, auth.RegisterModal(cc))
}

// ServiceWorkerHandler is an Echo handler that serves the service worker
func IndexFileHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Set appropriate headers for service worker
		c.Response().Header().Set("Content-Type", "text/html")
		c.Response().Header().Set("Service-Worker-Allowed", "/")

		// Generate and write the service worker JavaScript
		return render.Templ(c, view.IndexFile())
	}
}
