package ctx

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
	dwngen "github.com/onsonr/sonr/internal/dwn/gen"
)

type HeaderKey string

const (
	HeaderIPFSGatewayURL HeaderKey = "X-IPFS-Gateway-URL"
	HeaderVaultCID       HeaderKey = "X-Vault-CID"
	HeaderOriginSubject  HeaderKey = "X-Origin-Subject"
	HeaderSonrChainID    HeaderKey = "X-Sonr-Chain-ID"
)

func (h HeaderKey) String() string {
	return string(h)
}

func GetConfig(c echo.Context) (*dwngen.Config, error) {
	cnfg := new(dwngen.Config)
	// Attempt to read the session ID from the "session" cookie
	cnfgJSON, err := ReadCookie(c, CookieKeyConfig)
	if err != nil {
		c.Logger().Error(err)
		return nil, err
	}

	err = json.Unmarshal([]byte(cnfgJSON), cnfg)
	if err != nil {
		c.Logger().Error(err)
		return nil, err
	}
	return cnfg, nil
}

func SetConfig(c echo.Context, config *dwngen.Config) {
	WriteHeader(c, HeaderIPFSGatewayURL, config.ProxyUrl)
	WriteHeader(c, HeaderSonrChainID, config.Sonr.ChainId)
	cnfgBz, err := json.Marshal(config)
	if err != nil {
		c.Logger().Error(err)
	}
	WriteCookie(c, CookieKeyConfig, string(cnfgBz))
}
