package webapp

import (
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/common/middleware/render"
	home "github.com/onsonr/sonr/pkg/webapp/landing"
	vault "github.com/onsonr/sonr/pkg/webapp/vault"
)

func HomePage(c echo.Context) error {
	return render.Templ(c, home.View())
}

// ServiceWorkerHandler is an Echo handler that serves the service worker
func IndexPage(c echo.Context) error {
	// Set appropriate headers for service worker
	c.Response().Header().Set("Content-Type", "text/html")
	c.Response().Header().Set("Service-Worker-Allowed", "/")

	// Generate and write the service worker JavaScript
	return render.Templ(c, vault.IndexFile())
}
