package local

import (
	"errors"
	"time"

	"github.com/go-webauthn/webauthn/protocol"
)

// Session is the reference to the clients current session over gRPC/HTTP in the local cache.
type Session interface {
	// GetAddress returns the currently authenticated Users Sonr Address for the Session.
	GetAddress() (string, error)

	// GetChallenge returns the existing challenge or a new challenge to use for validation
	GetChallenge() []byte

	// IsAuthorized returns true if the Session has an attached JWT Token
	IsAuthorized() bool

	// SessionID returns the ksuid for the current session
	SessionID() string
}

// session is a proxy session.
type session struct {
	// ID is the ksuid of the Session
	ID string `json:"id"`

	// Validator is the address of the associated validator node address for the session.
	Validator string `json:"validator"`

	// ChainID is the current sonr blockchain network chain ID for the session.
	ChainID string `json:"chain_id"`

	// Challenge is used for authenticating credentials for the Session
	Challenge []byte `json:"challenge"`
}

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

// session is a proxy session.
type authorizedSession struct {
	// ID is the ksuid of the Session
	ID string `json:"id"`

	// Address is the address of the session.
	Address string `json:"address"`

	// Validator is the address of the associated validator node address for the session.
	Validator string `json:"validator"`

	// ChainID is the current sonr blockchain network chain ID for the session.
	ChainID string `json:"chain_id"`

	// Token is the token of the session.
	Token string `json:"token"`

	// Expires is the expiration time of the session.
	Expires time.Time `json:"expires"`

	// Challenge is used for authenticating credentials for the Session
	Challenge []byte `json:"challenge"`
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
