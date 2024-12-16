package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/context"
	"github.com/onsonr/sonr/internal/database/repository"
)

func CheckHandleUnique(c echo.Context, handle string) bool {
	ctx, ok := c.(*GatewayContext)
	if !ok {
		return false
	}
	ok, err := ctx.dbq.CheckHandleExists(bgCtx(), handle)
	if err != nil {
		return false
	}
	if ok {
		return false
	}
	context.WriteCookie(c, context.UserHandle, handle)
	return true
}

func CreateProfile(c echo.Context) (*repository.Profile, error) {
	ctx, ok := c.(*GatewayContext)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Profile Context not found")
	}
	address := c.FormValue("address")
	handle := c.FormValue("handle")
	origin := c.FormValue("origin")
	name := c.FormValue("name")
	profile, err := ctx.dbq.InsertProfile(bgCtx(), repository.InsertProfileParams{
		Address: address,
		Handle:  handle,
		Origin:  origin,
		Name:    name,
	})
	if err != nil {
		return nil, err
	}
	// Update session with profile id
	sid := GetSessionID(c)
	_, err = ctx.dbq.UpdateSessionWithProfileID(bgCtx(), repository.UpdateSessionWithProfileIDParams{
		ProfileID: profile.ID,
		ID:        sid,
	})
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func UpdateProfile(c echo.Context) (*repository.Profile, error) {
	ctx, ok := c.(*GatewayContext)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Profile Context not found")
	}
	address := c.FormValue("address")
	handle := c.FormValue("handle")
	name := c.FormValue("name")
	profile, err := ctx.dbq.UpdateProfile(bgCtx(), repository.UpdateProfileParams{
		Address: address,
		Handle:  handle,
		Name:    name,
	})
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func ReadProfile(c echo.Context) (*repository.Profile, error) {
	ctx, ok := c.(*GatewayContext)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Profile Context not found")
	}
	handle := c.Param("handle")
	profile, err := ctx.dbq.GetProfileByHandle(bgCtx(), handle)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func DeleteProfile(c echo.Context) error {
	ctx, ok := c.(*GatewayContext)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Profile Context not found")
	}
	address := c.Param("address")
	err := ctx.dbq.SoftDeleteProfile(bgCtx(), address)
	if err != nil {
		return err
	}
	return nil
}
