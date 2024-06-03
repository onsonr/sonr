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

// GetAddress returns the session address
func (s authorizedSession) GetAddress() (string, error) {
	if s.Address == "" {
		return "", errors.New("session does not have attached address")
	}
	return s.Address, nil
}

// GetValidator returns the associated validator address
func (s authorizedSession) GetValidator() (string, error) {
	if s.Address == "" {
		return "", errors.New("session does not have attached address")
	}
	return s.Address, nil
}

// GetChallenge returns the URL Encoded byte challenge
func (s authorizedSession) GetChallenge() []byte {
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
func (s authorizedSession) IsAuthorized() bool {
	return true
}

// SessionID returns the string ksuid for the session
func (s authorizedSession) SessionID() string {
	return s.ID
}
