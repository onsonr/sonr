package ctx

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/segmentio/ksuid"
)

var store sessions.Store

type ctxKeySessionID struct{}

// SessionMiddleware establishes a Session Cookie.
func SessionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	store = sessions.NewCookieStore([]byte("SESSION_KEY"))
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		// Attempt to read the session ID from the "session" cookie
		sessionID, err := readSessionIDFromCookie(c)
		if err != nil {
			// Generate a new KSUID if the session cookie is missing or invalid
			sessionID = ksuid.New().String()
			// Write the new session ID to the "session" cookie
			err = writeSessionIDToCookie(c, sessionID)
			if err != nil {
				return c.JSON(
					http.StatusInternalServerError,
					map[string]string{"error": "Failed to set session cookie"},
				)
			}
		}

		// Inject the session ID into the context
		ctx = context.WithValue(ctx, ctxKeySessionID{}, sessionID)
		// Update the request with the new context
		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}

func defaultSession(id string, s *sessions.Session) *session {
	return &session{
		session: s,
		id:      id,
		origin:  "",
		address: "",
		chainID: "",
	}
}

func getSessionID(ctx context.Context) (string, error) {
	sessionID, ok := ctx.Value(ctxKeySessionID{}).(string)
	if !ok || sessionID == "" {
		return "", errors.New("session ID not found in context")
	}
	return sessionID, nil
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
