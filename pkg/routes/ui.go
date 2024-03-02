package routes

import (
	"github.com/labstack/echo/v4"

	handlers "github.com/sonrhq/sonr/pkg/handlers/ui"
	"github.com/sonrhq/sonr/pkg/middleware/common"
)

// RegisterUI registers the page routes for HTMX
func RegisterUI(e *echo.Echo) {
	e.GET("*", handlers.Pages.Index, common.UseHTMX)
}
