package landing

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/webui/landing/handlers"
	// "github.com/onsonr/sonr/pkg/common/response"
	// "github.com/onsonr/sonr/pkg/webui/landing/pages"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/", handlers.HandleIndex)
}
