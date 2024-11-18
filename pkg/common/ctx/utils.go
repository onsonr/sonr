package ctx

import (
	"bytes"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// ╭───────────────────────────────────────────────────────────╮
// │                  Template Rendering                       │
// ╰───────────────────────────────────────────────────────────╯

func RenderTempl(c echo.Context, cmp templ.Component) error {
	// Create a buffer to store the rendered HTML
	buf := &bytes.Buffer{}
	// Render the component to the buffer
	err := cmp.Render(c.Request().Context(), buf)
	if err != nil {
		return err
	}

	// Set the content type
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	// Write the buffered content to the response
	_, err = c.Response().Write(buf.Bytes())
	return err
}

// ╭──────────────────────────────────────────────────────────╮
// │                  Cookie Management                       │
// ╰──────────────────────────────────────────────────────────╯

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

// ╭────────────────────────────────────────────────────────╮
// │                  HTTP Headers                          │
// ╰────────────────────────────────────────────────────────╯

func WriteHeader(c echo.Context, key HeaderKey, value string) {
	c.Response().Header().Set(key.String(), value)
}

func ReadHeader(c echo.Context, key HeaderKey) string {
	return c.Response().Header().Get(key.String())
}
