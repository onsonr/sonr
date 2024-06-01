package session

import (
	"errors"
	"time"

	"github.com/go-webauthn/webauthn/protocol"
)

// session is a proxy session.
type session struct {
	// ID is the ksuid of the Session
	ID string `json:"id"`

	// Address is the address of the session.
	Address string `json:"address"`

	// Validator is the address of the associated validator node address for the session.
	Validator string `json:"address"`

	// ChainID is the current sonr blockchain network chain ID for the session.
	ChainID string `json:"address"`

	// Token is the token of the session.
	Token string `json:"token"`

	// Expires is the expiration time of the session.
	Expires time.Time `json:"expires"`

	// Challenge is used for authenticating credentials for the Session
	Challenge []byte `json:"challenge"`
}

// GetAddress returns the session address
func (s session) GetAddress() (string, error) {
	if s.Address == "" {
		return "", errors.New("session does not have attached address")
	}
	return s.Address, nil
}

// GetValidator returns the associated validator address
func (s session) GetValidator() (string, error) {
	if s.Address == "" {
		return "", errors.New("session does not have attached address")
	}
	return s.Address, nil
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
func (s session) IsAuthorized() (bool, error) {
	if s.Token == "" {
		return false, errors.New("session does not have attached address")
	}
	return true, nil
}

// SessionID returns the string ksuid for the session
func (s session) SessionID() string {
	return s.ID
}
