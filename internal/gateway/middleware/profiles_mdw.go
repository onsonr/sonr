package middleware

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/gateway/models/repository"
)

type ProfilesContext struct {
	echo.Context
	dbq *repository.Queries
}

func UseProfiles(conn *sql.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := &ProfilesContext{Context: c, dbq: repository.New(conn)}
			return next(ctx)
		}
	}
}

func CreateProfile(c echo.Context) (*repository.Profile, error) {
	ctx, ok := c.(*ProfilesContext)
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
	return &profile, nil
}

func UpdateProfile(c echo.Context) (*repository.Profile, error) {
	ctx, ok := c.(*ProfilesContext)
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
	ctx, ok := c.(*ProfilesContext)
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
	ctx, ok := c.(*ProfilesContext)
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
