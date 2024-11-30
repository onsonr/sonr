package gateway

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/common/session"
	"github.com/onsonr/sonr/pkg/crypto/mpc"
	"github.com/onsonr/sonr/web/vault/types"
)

func spawnVault(client IPFSClient) echo.HandlerFunc {
	return func(c echo.Context) error {
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
		dwnCfg := &types.Config{
			MotrKeyshare:   kscid.String(),
			MotrAddress:    src.Address(),
			IpfsGatewayUrl: "https://ipfs.sonr.land",
			SonrApiUrl:     "https://api.sonr.land",
			SonrRpcUrl:     "https://rpc.sonr.land",
			SonrChainId:    session.GetChainID(c),
			VaultSchema:    session.GetVaultSchema(c),
		}
		cnf, err := json.Marshal(dwnCfg)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		dir := setupVaultDirectory(cnf)
		path, err := client.Unixfs().Add(c.Request().Context(), dir)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.Redirect(http.StatusFound, path.String())
	}
}
