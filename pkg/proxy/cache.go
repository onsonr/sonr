package proxy

import (
	"context"
	"time"

	"github.com/bool64/cache"
)

var (
	challengeCache *cache.FailoverOf[Challenge]
	sessionCache   *cache.FailoverOf[Session]
)

// Initialize configures cache and inital settings for proxy.
func Initialize() {
	// Setup cache for session.
	sessionCache = cache.NewFailoverOf[Session](func(cfg *cache.FailoverConfigOf[Session]) {
		// Using last 30 seconds of 5m TTL for background update.
		cfg.MaxStaleness = 30 * time.Minute
		cfg.BackendConfig.TimeToLive = 1*time.Hour - cfg.MaxStaleness
	})

	challengeCache = cache.NewFailoverOf[Challenge](func(cfg *cache.FailoverConfigOf[Challenge]) {
		// Using last 30 seconds of 5m TTL for background update.
		cfg.MaxStaleness = 5 * time.Minute
		cfg.BackendConfig.TimeToLive = 20*time.Minute - cfg.MaxStaleness
	})
}

// GetChallenge returns a challenge from cache given a key.
func GetChallenge(key string) (Challenge, error) {
	return challengeCache.Get(context.Background(), []byte(key), func(ctx context.Context) (Challenge, error) {
		// Build value or return error on failure.
		return Challenge{}, nil
	})
}

// GetSession returns a session from cache given a key.
func GetSession(key string) (Session, error) {
	return sessionCache.Get(context.Background(), []byte(key), func(ctx context.Context) (Session, error) {
		// Build value or return error on failure.
		return Session{}, nil
	})
}
