package ctx

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/labstack/echo/v4"
)

type WebBytes = protocol.URLEncodedBase64

type Session struct {
	echo.Context

	// Defaults
	id string // Generated ksuid http cookie; Initialized on first request

	// Initialization
	address string // Webauthn mapping to User ID; Supplied by DWN frontend
	chainID string // Macaroon mapping to location; Supplied by DWN frontend

	// Authentication
	challenge WebBytes // Webauthn mapping to Challenge; Per session based on origin
}

func (s *Session) ID() string {
	return s.id
}

func (s *Session) Address() string {
	return s.address
}

func (s *Session) ChainID() string {
	return s.chainID
}
