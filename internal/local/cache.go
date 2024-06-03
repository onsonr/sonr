package local

import (
	"time"

	"github.com/bool64/cache"
)

var (
	baseSessionCache       *cache.FailoverOf[session]
	authorizedSessionCache *cache.FailoverOf[authorizedSession]
)

// setupCache configures cache and inital settings for proxy.
func setupCache() {
	// Setup cache for session.
	baseSessionCache = cache.NewFailoverOf(func(cfg *cache.FailoverConfigOf[session]) {
		// Using last 30 seconds of 5m TTL for background update.
		cfg.MaxStaleness = 1 * time.Hour
		cfg.BackendConfig.TimeToLive = 2*time.Hour - cfg.MaxStaleness
	})
	authorizedSessionCache = cache.NewFailoverOf(func(cfg *cache.FailoverConfigOf[authorizedSession]) {
		// Using last 30 seconds of 5m TTL for background update.
		cfg.MaxStaleness = 30 * time.Minute
		cfg.BackendConfig.TimeToLive = 1*time.Hour - cfg.MaxStaleness
	})
}

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
