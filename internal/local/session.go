package local

import (
	"errors"

	"github.com/go-webauthn/webauthn/protocol"
)

// GetAddress returns the session address
func (s session) GetAddress() (string, error) {
	return "", errors.New("session does not have attached address")
}

// GetValidator returns the associated validator address
func (s session) GetValidator() (string, error) {
	if s.Validator == "" {
		return "", errors.New("session does not have attached address")
	}
	return s.Validator, nil
}

// GetChallenge returns the URL Encoded byte challenge
func (s session) GetChallenge() []byte {
	if s.Challenge == nil {
		bz, err := protocol.CreateChallenge()
		if err != nil {
			panic(err)
		}
		s.Challenge = bz
	}
	return s.Challenge
}

// IsAuthorized returns true or false for if it is authorized
func (s session) IsAuthorized() bool {
	return false
}

// SessionID returns the string ksuid for the session
func (s session) SessionID() string {
	return s.ID
}

// SetUserAddress sets the user address for the session and persists it in cache
func (c SonrContext) SetUserAddress(address string) error {
	s, err := getSessionFromCache(c.Context, c.SessionID)
	if err != nil {
		return err
	}
	s.UserAddress = address

	return nil
}
