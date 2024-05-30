package cache

import (
	"context"
	"time"

	"github.com/bool64/cache"
	"github.com/segmentio/ksuid"
)

var (
	sessionCache *cache.FailoverOf[Session]
)

// Initialize configures cache and inital settings for proxy.
func Initialize() {
	// Setup cache for session.
	sessionCache = cache.NewFailoverOf(func(cfg *cache.FailoverConfigOf[Session]) {
		// Using last 30 seconds of 5m TTL for background update.
		cfg.MaxStaleness = 30 * time.Minute
		cfg.BackendConfig.TimeToLive = 1*time.Hour - cfg.MaxStaleness
	})
}

// GetSession returns a session from cache given a key.
func GetSession(id string) (Session, error) {
	return sessionCache.Get(context.Background(), []byte(id), func(ctx context.Context) (Session, error) {
		// Build value or return error on failure.
		return Session{
			ID: ksuid.New().String(),
		}, nil
	})
}
