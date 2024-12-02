package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RedirectLanding(c echo.Context) error {
	return c.Redirect(http.StatusFound, "http://localhost:3000")
}

func RedirectIPFS(c echo.Context, cid string) error {
	return c.Redirect(http.StatusFound, cid)
}

func RedirectAuth(c echo.Context) error {
	return c.Redirect(http.StatusFound, "http://auth.localhost:3000")
}
