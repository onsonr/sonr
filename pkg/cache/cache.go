// Package cache provides a generic caching mechanism based on the
// github.com/patrickmn/go-cache package, allowing for type-safe caches
// with string-constrained keys.
package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

// stringConstraint is a type constraint that only allows types that are
// underlying strings.
type stringConstraint interface {
	~string
}

// Cache represents a generic cache with keys constrained to string types
// and values of any type.
type Cache[K stringConstraint, T any] struct {
	cache *cache.Cache
}

// New creates a new Cache instance with the specified default expiration
// duration and cleanup interval.
//
// Parameters:
// - defaultExpiration: the default time-to-live for cache entries.
// - cleanupInterval: the interval at which expired entries are removed.
//
// Returns:
// A pointer to the newly created Cache instance.
func New[K stringConstraint, T any](defaultExpiration time.Duration, cleanupInterval time.Duration) *Cache[K, T] {
	c := cache.New(defaultExpiration, cleanupInterval)

	return &Cache[K, T]{
		cache: c,
	}
}

// Set adds a value to the cache with the specified key. The value will
// be stored using the default expiration time of the cache.
//
// Parameters:
// - key: the key under which the value will be stored.
// - val: the value to be stored.
func (c *Cache[K, T]) Set(key K, val any) {
	c.cache.Set(string(key), val, cache.DefaultExpiration)
}

// Get retrieves a value from the cache based on the specified key.
//
// Parameters:
// - key: the key associated with the value to be retrieved.
//
// Returns:
// - A pointer to the value if it exists and is of the correct type, or nil otherwise.
// - A boolean indicating whether the value was found and of the expected type.
func (c *Cache[K, T]) Get(key K) (*T, bool) {
	v, ok := c.cache.Get(string(key))
	if !ok {
		return nil, false
	}

	vTyped, ok := v.(T)
	if !ok {
		return nil, false
	}

	return &vTyped, true
}
