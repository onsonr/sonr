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
