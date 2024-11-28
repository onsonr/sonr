package landing

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/common/middleware/response"
	"github.com/onsonr/sonr/web/landing/pages"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/", response.Templ(pages.HomePage()))
	e.GET("/register", response.Templ(pages.RegisterPage()))
}
