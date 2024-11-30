package gateway

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/common/session"
	"github.com/onsonr/sonr/pkg/crypto/mpc"
)

func spawnVault(c echo.Context) error {
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
	kscid, err := tk.CID()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	cfg := session.GetVaultConfig(c, src.Address(), kscid.String())
	cnf, err := json.Marshal(cfg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	dir := setupVaultDirectory(cnf)
	path, err := ipfs.Unixfs().Add(c.Request().Context(), dir)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.Redirect(http.StatusFound, path.String())
}
