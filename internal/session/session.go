package session

import (
	"context"
	"time"

	"github.com/bool64/cache"
	"github.com/di-dao/sonr/internal/local"
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

// Initialize configures cache and inital settings for proxy.
func Initialize() {
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

// Get returns a session from cache given a key.
func Get(ctx context.Context) (Session, error) {
	id := unwrapFromContext(ctx)
	snrCtx := local.UnwrapContext(ctx)

	return baseSessionCache.Get(
		context.Background(),
		[]byte(snrCtx.SessionID),
		func(ctx context.Context) (session, error) {
			// Build value or return error on failure.
			return session{
				ID: id,
			}, nil
		},
	)
}
