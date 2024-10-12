package ctx

import (
	"errors"
	"fmt"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/labstack/echo/v4"
)

type WebBytes = protocol.URLEncodedBase64

type Session struct {
	// Defaults
	ID        string // Generated ksuid http cookie; Initialized on first request
	Origin    string // Webauthn mapping to Relaying Party ID; Initialized on first request
	UserAgent string
	Platform  string

	// Initialization
	Address string // Webauthn mapping to User ID; Supplied by DWN frontend
	ChainID string // Macaroon mapping to location; Supplied by DWN frontend

	Subject string // Webauthn mapping to User Displayable Name; Supplied by DWN frontend

	// Authentication
	challenge WebBytes // Webauthn mapping to Challenge; Per session based on origin
}

func (s *Session) GetChallenge(subject string) (WebBytes, error) {
	// Check if challenge is already set and subject matches
	if s.Subject != "" && s.Subject != subject {
		return nil, errors.New("challenge already set, and subject does not match")
	} else if s.Subject == "" {
		s.Subject = subject
	} else {
		return s.challenge, nil
	}

	if s.challenge == nil {
		chl, err := protocol.CreateChallenge()
		if err != nil {
			return nil, err
		}
		s.challenge = chl
	}
	return s.challenge, nil
}

func (s *Session) ValidateChallenge(challenge WebBytes, subject string) error {
	if s.challenge == nil {
		return nil
	}
	if s.challenge.String() != challenge.String() {
		return fmt.Errorf("invalid challenge")
	}
	s.Subject = subject
	return nil
}

func GetSession(c echo.Context) *Session {
	id, _ := getSessionID(c.Request().Context())
	return buildSession(c, id)
}

func SetAddress(c echo.Context, address string) *Session {
	// Write address to X-Sonr-Address header
	c.Response().Header().Set("X-Sonr-Address", address)
	return buildSession(c, "")
}
