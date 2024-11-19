package ctx

import (
	"encoding/json"

	"github.com/labstack/echo/v4"

	dwngen "github.com/onsonr/sonr/pkg/motr/config"
)

type HeaderKey string

const (
	HeaderAuthorization HeaderKey = "Authorization"

	HeaderIPFSGatewayURL HeaderKey = "X-IPFS-Gateway"
	HeaderSonrChainID    HeaderKey = "X-Sonr-ChainID"
	HeaderSonrKeyshare   HeaderKey = "X-Sonr-Keyshare"
)

func (h HeaderKey) String() string {
	return string(h)
}

func injectConfig(c echo.Context, config *dwngen.Config) {
	WriteHeader(c, HeaderIPFSGatewayURL, config.IpfsGatewayUrl)
	WriteHeader(c, HeaderSonrChainID, config.SonrChainId)
	WriteHeader(c, HeaderSonrKeyshare, config.MotrKeyshare)
	WriteCookie(c, CookieKeySonrAddr, config.MotrAddress)

	schemaBz, err := json.Marshal(config.VaultSchema)
	if err != nil {
		c.Logger().Error(err)
		return
	}

	WriteCookie(c, CookieKeyVaultSchema, string(schemaBz))
}
