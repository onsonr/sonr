package landing

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/common/response"
	"github.com/onsonr/sonr/pkg/webui/landing/pages"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/", response.Templ(pages.HomePage()))
	e.GET("/register", response.Templ(pages.RegisterPage()))
}
