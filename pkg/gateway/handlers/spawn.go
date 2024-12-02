package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/crypto/mpc"
	"github.com/onsonr/sonr/pkg/common/session"
	"github.com/onsonr/sonr/pkg/gateway/embed"
	"github.com/onsonr/sonr/pkg/gateway/middleware"
	"github.com/onsonr/sonr/pkg/vault/types"
)

func SpawnVault(c echo.Context) error {
	ipfs, err := middleware.GetIPFSClient(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
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
	cfg := session.GetVaultConfig(c, src.Address(), kscid.String())
	cnf, err := json.Marshal(cfg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Create the vault app manifest
	manifest := types.NewWebManifest()
	manifestBz, err := json.Marshal(manifest)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	dir := embed.NewFS(cnf, manifestBz)
	path, err := ipfs.Unixfs().Add(c.Request().Context(), dir)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.Redirect(http.StatusFound, path.String())
}
