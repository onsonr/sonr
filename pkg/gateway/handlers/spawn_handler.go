package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/crypto/mpc"
	"github.com/onsonr/sonr/pkg/common/session"
	"github.com/onsonr/sonr/pkg/gateway/middleware"
	"github.com/onsonr/sonr/pkg/vault/config"
)

type SpawnVaultRequest struct {
	Name string `json:"name"`
}

func SpawnVault(c echo.Context) error {
	ks, err := mpc.NewKeyset()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	src, err := mpc.NewSource(ks)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	tk, err := src.OriginToken()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	// Create the vault keyshare auth token
	kscid, err := tk.CID()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	// Create the vault config
	dir, err := config.NewFS(session.GetVaultConfig(c, src.Address(), kscid.String()))
	path, err := middleware.IPFSAdd(c, dir)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.Redirect(http.StatusFound, path)
}
