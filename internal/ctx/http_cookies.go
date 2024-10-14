package ctx

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	dwngen "github.com/onsonr/sonr/internal/dwn/gen"
	"github.com/segmentio/ksuid"
)

type CookieKey string

func (c CookieKey) String() string {
	return string(c)
}

const (
	CookieKeySessionID CookieKey = "session.id"
	CookieKeyConfig    CookieKey = "dwn.config"
)

func GetSessionID(c echo.Context) string {
	// Attempt to read the session ID from the "session" cookie
	sessionID, err := readCookie(c, CookieKeySessionID)
	if err != nil {
		// Generate a new KSUID if the session cookie is missing or invalid
		writeCookie(c, CookieKeySessionID, ksuid.New().String())
	}
	return sessionID
}

func GetConfig(c echo.Context) (*dwngen.Config, error) {
	cnfg := new(dwngen.Config)
	// Attempt to read the session ID from the "session" cookie
	cnfgJSON, err := readCookie(c, CookieKeyConfig)
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
	cnfgBz, err := json.Marshal(config)
	if err != nil {
		c.Logger().Error(err)
	}
	writeCookie(c, CookieKeyConfig, string(cnfgBz))
}

// ╭───────────────────────────────────────────────────────────╮
// │                  Registration Components                  │
// ╰───────────────────────────────────────────────────────────╯

func readCookie(c echo.Context, key CookieKey) (string, error) {
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

func writeCookie(c echo.Context, key CookieKey, value string) error {
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
