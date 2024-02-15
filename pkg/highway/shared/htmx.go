package shared

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// ShowTempl renders a templ.Component
func ShowTempl(cmp templ.Component) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Render the templ component to a `template.HTML` value.
		html, err := templ.ToGoHTML(context.Background(), cmp)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.HTML(http.StatusOK, string(html))
	}
}

// RenderPage renders a templ.Component
func RenderPage(c echo.Context, cmp templ.Component) error {
	return cmp.Render(c.Request().Context(), c.Response())
}
