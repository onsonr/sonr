package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type stringConstraint interface {
	~string
}

type Cache[K stringConstraint, T any] struct {
	cache *cache.Cache
}

func New[K stringConstraint, T any](defaultExpiration time.Duration, cleanupInterval time.Duration) *Cache[K, T] {
	c := cache.New(defaultExpiration, cleanupInterval)

	return &Cache[K, T]{
		cache: c,
	}
}

func (c *Cache[K, T]) Set(key K, val any) {
	c.cache.Set(string(key), val, cache.DefaultExpiration)
}

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
