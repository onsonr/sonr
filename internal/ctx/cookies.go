package ctx

import (
	"github.com/labstack/echo/v4"
	"github.com/segmentio/ksuid"
)

type CookieKey string

const (
	CookieKeySessionID   CookieKey = "session.id"
	CookieKeySonrAddr    CookieKey = "sonr.addr"
	CookieKeySonrDID     CookieKey = "sonr.did"
	CookieKeyVaultCID    CookieKey = "vault.cid"
	CookieKeyVaultSchema CookieKey = "vault.schema"
)

func (c CookieKey) String() string {
	return string(c)
}

func GetSessionID(c echo.Context) string {
	// Attempt to read the session ID from the "session" cookie
	sessionID, err := ReadCookie(c, CookieKeySessionID)
	if err != nil {
		// Generate a new KSUID if the session cookie is missing or invalid
		WriteCookie(c, CookieKeySessionID, ksuid.New().String())
	}
	return sessionID
}
