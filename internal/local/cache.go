package local

import (
	"context"
	"time"

	"github.com/bool64/cache"
)

var (
	baseSessionCache       *cache.FailoverOf[session]
	authorizedSessionCache *cache.FailoverOf[authorizedSession]
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

// Get returns a session from cache given a key.
func GetSession(ctx context.Context) (Session, error) {
	snrCtx := UnwrapContext(ctx)
	return baseSessionCache.Get(
		context.Background(),
		[]byte(snrCtx.SessionID),
		func(ctx context.Context) (session, error) {
			// Build value or return error on failure.
			return session{
				ID:        snrCtx.SessionID,
				Validator: snrCtx.ValidatorAddress,
				ChainID:   snrCtx.ChainID,
			}, nil
		},
	)
}
