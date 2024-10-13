package ctx

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/segmentio/ksuid"
)

func GetSession(c echo.Context) *Session {
	return c.(*Session)
}

// HighwaySessionMiddleware establishes a Session Cookie.
func HighwaySessionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sessionID := getSessionIDFromCookie(c)
		cc := &Session{
			Context: c,
			id:      sessionID,
			address: c.Request().Header.Get("X-Sonr-Address"),
			chainID: "",
		}
		return next(cc)
	}
}

func getSessionIDFromCookie(c echo.Context) string {
	// Attempt to read the session ID from the "session" cookie
	sessionID, err := readSessionIDFromCookie(c)
	if err != nil {
		// Generate a new KSUID if the session cookie is missing or invalid
		sessionID = ksuid.New().String()
		// Write the new session ID to the "session" cookie
		writeSessionIDToCookie(c, sessionID)
	}
	return sessionID
}

func readSessionIDFromCookie(c echo.Context) (string, error) {
	cookie, err := c.Cookie("session")
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

func writeSessionIDToCookie(c echo.Context, sessionID string) error {
	cookie := &http.Cookie{
		Name:     "session",
		Value:    sessionID,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Path:     "/",
		// Add Secure and SameSite attributes as needed
	}
	c.SetCookie(cookie)
	return nil
}
