package webapp

import (
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/common/middleware/render"
	home "github.com/onsonr/sonr/pkg/webapp/landing"
)

func HomePage(c echo.Context) error {
	return render.Templ(c, home.View())
}
