package common

import (
	"encoding/base64"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// CookieKey is a type alias for string.
type CookieKey string

const (
	// SessionID is the key for the session ID cookie.
	SessionID CookieKey = "session.id"

	// SessionChallenge is the key for the session challenge cookie.
	SessionChallenge CookieKey = "session.challenge"

	// SessionRole is the key for the session role cookie.
	SessionRole CookieKey = "session.role"

	// SonrAddress is the key for the Sonr address cookie.
	SonrAddress CookieKey = "sonr.address"

	// SonrDID is the key for the Sonr DID cookie.
	SonrDID CookieKey = "sonr.did"

	// UserAvatar is the key for the User Avatar cookie.
	UserAvatar CookieKey = "user.avatar"

	// UserHandle is the key for the User Handle cookie.
	UserHandle CookieKey = "user.handle"

	// UserName is the key for the User Name cookie.
	UserName CookieKey = "user.full_name"

	// VaultAddress is the key for the Vault address cookie.
	VaultAddress CookieKey = "vault.address"

	// VaultCID is the key for the Vault CID cookie.
	VaultCID CookieKey = "vault.cid"

	// VaultSchema is the key for the Vault schema cookie.
	VaultSchema CookieKey = "vault.schema"
)

// String returns the string representation of the CookieKey.
func (c CookieKey) String() string {
	return string(c)
}

// ╭───────────────────────────────────────────────────────────╮
// │                      Utility Methods                      │
// ╰───────────────────────────────────────────────────────────╯

func CookieExists(c echo.Context, key CookieKey) bool {
	ck, err := c.Cookie(key.String())
	if err != nil {
		return false
	}
	return ck != nil
}

func ReadCookie(c echo.Context, key CookieKey) (string, error) {
	cookie, err := c.Cookie(key.String())
	if err != nil {
		// Cookie not found or other error
		return "", err
	}
	if cookie == nil || cookie.Value == "" {
		// Cookie is empty
		return "", http.ErrNoCookie
	}
	return cookie.Value, nil
}

func ReadCookieBytes(c echo.Context, key CookieKey) ([]byte, error) {
	cookie, err := c.Cookie(key.String())
	if err != nil {
		// Cookie not found or other error
		return nil, err
	}
	if cookie == nil || cookie.Value == "" {
		// Cookie is empty
		return nil, http.ErrNoCookie
	}
	return base64.RawURLEncoding.DecodeString(cookie.Value)
}

func ReadCookieUnsafe(c echo.Context, key CookieKey) string {
	ck, err := c.Cookie(key.String())
	if err != nil {
		return ""
	}
	return ck.Value
}

func WriteCookie(c echo.Context, key CookieKey, value string) error {
	cookie := &http.Cookie{
		Name:     key.String(),
		Value:    value,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Path:     "/",
		// Add Secure and SameSite attributes as needed
	}
	c.SetCookie(cookie)
	return nil
}

func WriteCookieBytes(c echo.Context, key CookieKey, value []byte) error {
	cookie := &http.Cookie{
		Name:     key.String(),
		Value:    base64.RawURLEncoding.EncodeToString(value),
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Path:     "/",
		// Add Secure and SameSite attributes as needed
	}
	c.SetCookie(cookie)
	return nil
}
