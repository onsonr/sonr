package ctx

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/labstack/echo/v4"
	"github.com/segmentio/ksuid"
)

// CookieKey is a type alias for string.
type CookieKey string

const (
	// CookieKeySessionID is the key for the session ID cookie.
	CookieKeySessionID CookieKey = "session.id"

	// CookieKeySessionChal is the key for the session challenge cookie.
	CookieKeySessionChal CookieKey = "session.chal"

	// CookieKeySonrAddr is the key for the Sonr address cookie.
	CookieKeySonrAddr CookieKey = "sonr.addr"

	// CookieKeySonrDID is the key for the Sonr DID cookie.
	CookieKeySonrDID CookieKey = "sonr.did"

	// CookieKeyVaultCID is the key for the Vault CID cookie.
	CookieKeyVaultCID CookieKey = "vault.cid"

	// CookieKeyVaultSchema is the key for the Vault schema cookie.
	CookieKeyVaultSchema CookieKey = "vault.schema"
)

// String returns the string representation of the CookieKey.
func (c CookieKey) String() string {
	return string(c)
}

// GetSessionID returns the session ID from the cookies.
func GetSessionID(c echo.Context) string {
	// Attempt to read the session ID from the "session" cookie
	sessionID, err := ReadCookie(c, CookieKeySessionID)
	if err != nil {
		// Generate a new KSUID if the session cookie is missing or invalid
		WriteCookie(c, CookieKeySessionID, ksuid.New().String())
	}
	return sessionID
}

// GetSessionChallenge returns the session challenge from the cookies.
func GetSessionChallenge(c echo.Context) (*protocol.URLEncodedBase64, error) {
	chal := new(protocol.URLEncodedBase64)
	// Attempt to read the session challenge from the "session" cookie
	sessionChal, err := ReadCookie(c, CookieKeySessionChal)
	if err != nil {
		// Generate a new challenge if the session cookie is missing or invalid
		ch, errb := protocol.CreateChallenge()
		if errb != nil {
			return nil, err
		}
		WriteCookie(c, CookieKeySessionChal, ch.String())
		return &ch, nil
	}
	err = chal.UnmarshalJSON([]byte(sessionChal))
	if err != nil {
		return nil, err
	}
	return chal, nil
}
