package middleware

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/ipfs/boxo/path"
	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
)

func SessionID(c echo.Context) string {
	return readCookie(c, "session")
}

var Cache = cacheHandler{}

type cacheHandler struct{}

func (c cacheHandler) GetChallenge(e echo.Context) protocol.URLEncodedBase64 {
	key := cacheKey(e, "challenge")
	if x, found := ccref.Challenges.Get(key); found {
		chalbz := x.([]byte)
		chal := protocol.URLEncodedBase64{}
		err := chal.UnmarshalJSON(chalbz)
		if err != nil {
			return chal
		}
		return chal
	}
	chal, _ := protocol.CreateChallenge()
	// Save challenge
	str, err := chal.MarshalJSON()
	ccref.Challenges.Set(key, str, cache.DefaultExpiration)
	if err != nil {
		return chal
	}
	return chal
}

func (c cacheHandler) GetLocalPath(e echo.Context) string {
	key := cacheKey(e, "localPath")
	if x, found := ccref.Paths.Get(key); found {
		return x.(string)
	}
	return ""
}

func (c cacheHandler) GetRemoteCID(e echo.Context) path.Path {
	key := cacheKey(e, "remoteCID")
	if x, found := ccref.CIDs.Get(key); found {
		return x.(path.Path)
	}
	return nil
}

func (c cacheHandler) SetLocalPath(e echo.Context, path string) {
	key := cacheKey(e, "localPath")
	ccref.Paths.Set(key, path, cache.DefaultExpiration)
}

func (c cacheHandler) SetRemoteCID(e echo.Context, path path.Path) {
	key := cacheKey(e, "remoteCID")
	ccref.CIDs.Set(key, path, cache.DefaultExpiration)
}
