package proxy

import "time"

// Session is a proxy session.
type Session struct {
	// Address is the address of the session.
	Address string `json:"address"`

	// Token is the token of the session.
	Token string `json:"token"`

	// Expires is the expiration time of the session.
	Expires time.Time `json:"expires"`
}
