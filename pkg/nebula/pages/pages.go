package pages

import (
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/nebula/components/grant"
	"github.com/onsonr/sonr/pkg/nebula/components/home"
	"github.com/onsonr/sonr/pkg/nebula/components/login"
	"github.com/onsonr/sonr/pkg/nebula/components/profile"
	"github.com/onsonr/sonr/pkg/nebula/components/register"
	"github.com/onsonr/sonr/pkg/nebula/models"
)

func Authorize(c echo.Context) error {
	return echoResponse(c, grant.View(c))
}

func Home(c echo.Context) error {
	mdls, err := models.GetModels()
	if err != nil {
		return err
	}
	return echoResponse(c, home.View(mdls.Home))
}

func Login(c echo.Context) error {
	return echoResponse(c, login.Modal(c))
}

func Profile(c echo.Context) error {
	return echoResponse(c, profile.View(c))
}

func Register(c echo.Context) error {
	return echoResponse(c, register.Modal(c))
}
