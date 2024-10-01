package pages

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/nebula/components/home"
	"github.com/onsonr/sonr/pkg/nebula/models"
)

func Home(c echo.Context) error {
	mdls, err := models.GetModels()
	if err != nil {
		return err
	}
	return echoResponse(c, home.View(mdls.Hero))
}
