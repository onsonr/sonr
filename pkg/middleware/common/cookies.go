package common

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func Cookies(ctx echo.Context) *cookies {
	return &cookies{
		Context: ctx,
	}
}

type cookies struct {
	echo.Context
}

func (c *cookies) SetHandle(handle string) {
	cookie := new(http.Cookie)
	cookie.Name = "sonr-handle"
	cookie.Value = handle
	cookie.Expires = time.Now().Add(4 * time.Hour)
	c.SetCookie(cookie)
}

func (c *cookies) GetHandle() string {
	cookie, err := c.Cookie("sonr-handle")
	if err != nil {
		return ""
	}
	return cookie.Value
}

func (c *cookies) HasHandle() bool {
	cookie, err := c.Cookie("sonr-handle")
	if err != nil {
		return false
	}
	return cookie.Value != ""
}

// SetAddress sets the address cookie
func (c *cookies) SetAddress(address string) {
	cookie := new(http.Cookie)
	cookie.Name = "sonr-address"
	cookie.Value = address
	cookie.Expires = time.Now().Add(1 * time.Hour)
	c.SetCookie(cookie)
}

// GetAddress gets the address cookie
func (c *cookies) GetAddress() string {
	cookie, err := c.Cookie("sonr-address")
	if err != nil {
		return ""
	}
	return cookie.Value
}

// HasAddress checks if the address cookie exists
func (c *cookies) HasAddress() bool {
	cookie, err := c.Cookie("sonr-address")
	if err != nil {
		return false
	}
	return cookie.Value != ""
}

// SetOrigin sets the origin cookie
func (c *cookies) SetOrigin(origin string) {
	cookie := new(http.Cookie)
	cookie.Name = "sonr-origin"
	cookie.Value = origin
	cookie.Expires = time.Now().Add(1 * time.Hour)
	c.SetCookie(cookie)
}

// GetOrigin gets the origin cookie
func (c *cookies) GetOrigin() string {
	cookie, err := c.Cookie("sonr-origin")
	if err != nil {
		return ""
	}
	return cookie.Value
}

// HasOrigin checks if the origin cookie exists
func (c *cookies) HasOrigin() bool {
	cookie, err := c.Cookie("sonr-origin")
	if err != nil {
		return false
	}
	return cookie.Value != ""
}

// GetSessionID gets the session id cookie
func (c *cookies) GetSessionID() string {
	cookie, err := c.Cookie("session-id")
	if err != nil {
		return ""
	}
	return cookie.Value
}

// HasSessionID checks if the session id cookie exists
func (c *cookies) HasSessionID() bool {
	cookie, err := c.Cookie("session-id")
	if err != nil {
		return false
	}
	return cookie.Value != ""
}
