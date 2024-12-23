package context

import (
	"net/http"

	"github.com/labstack/echo/v4"
	hwayorm "github.com/onsonr/sonr/internal/database/hwayorm"
)

func UpdateProfile(c echo.Context) (*hwayorm.Profile, error) {
	ctx, ok := c.(*GatewayContext)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Profile Context not found")
	}
	address := c.FormValue("address")
	handle := c.FormValue("handle")
	name := c.FormValue("name")
	profile, err := ctx.UpdateProfile(bgCtx(), hwayorm.UpdateProfileParams{
		Address: address,
		Handle:  handle,
		Name:    name,
	})
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func ReadProfile(c echo.Context) (*hwayorm.Profile, error) {
	ctx, ok := c.(*GatewayContext)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Profile Context not found")
	}
	handle := c.Param("handle")
	profile, err := ctx.GetProfileByHandle(bgCtx(), handle)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func DeleteProfile(c echo.Context) error {
	ctx, ok := c.(*GatewayContext)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Profile Context not found")
	}
	address := c.Param("address")
	err := ctx.SoftDeleteProfile(bgCtx(), address)
	if err != nil {
		return err
	}
	return nil
}
