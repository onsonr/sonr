package context

import "github.com/labstack/echo/v4"

type HeaderKey string

const (
	Authorization HeaderKey = "Authorization"

	// User Agent
	Architecture    HeaderKey = "Sec-CH-UA-Arch"
	Bitness         HeaderKey = "Sec-CH-UA-Bitness"
	FullVersionList HeaderKey = "Sec-CH-UA-Full-Version-List"
	Mobile          HeaderKey = "Sec-CH-UA-Mobile"
	Model           HeaderKey = "Sec-CH-UA-Model"
	Platform        HeaderKey = "Sec-CH-UA-Platform"
	PlatformVersion HeaderKey = "Sec-CH-UA-Platform-Version"
	UserAgent       HeaderKey = "Sec-CH-UA"

	// Sonr Injected
	SonrAPIURL  HeaderKey = "X-Sonr-API"
	SonrgRPCURL HeaderKey = "X-Sonr-GRPC"
	SonrRPCURL  HeaderKey = "X-Sonr-RPC"
	SonrWSURL   HeaderKey = "X-Sonr-WS"
)

func (h HeaderKey) String() string {
	return string(h)
}

// ╭───────────────────────────────────────────────────────────╮
// │                      Utility Methods                      │
// ╰───────────────────────────────────────────────────────────╯

func HeaderEquals(c echo.Context, key HeaderKey, value string) bool {
	return c.Response().Header().Get(key.String()) == value
}

// HeaderExists returns true if the request has the header Key.
func HeaderExists(c echo.Context, key HeaderKey) bool {
	return c.Response().Header().Get(key.String()) != ""
}

// HeaderRead returns the header value for the Key.
func HeaderRead(c echo.Context, key HeaderKey) string {
	return c.Response().Header().Get(key.String())
}

// HeaderWrite sets the header value for the Key.
func HeaderWrite(c echo.Context, key HeaderKey, value string) {
	c.Response().Header().Set(key.String(), value)
}
