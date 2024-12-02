package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RedirectLanding(c echo.Context) error {
	return c.Redirect(http.StatusFound, "http://localhost:3000")
}

func RedirectVaultCID(c echo.Context, cid string) error {
	return c.Redirect(http.StatusFound, cid)
}

func RedirectVaultIPNS(c echo.Context, ipns string) error {
	return c.Redirect(http.StatusFound, ipns)
}
