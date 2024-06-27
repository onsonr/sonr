package middleware

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// Render renders a templ.Component
func Render(c echo.Context, cmp templ.Component) error {
	c.Response().Writer.WriteHeader(http.StatusOK)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return cmp.Render(c.Request().Context(), c.Response())
}
