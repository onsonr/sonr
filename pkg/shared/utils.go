package shared

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// RenderPage renders a templ.Component
func RenderPage(c echo.Context, cmp templ.Component) error {
	return cmp.Render(c.Request().Context(), c.Response())
}
