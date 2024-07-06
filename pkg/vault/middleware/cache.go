package middleware

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
)

var ccref *cacheStore

func CacheStores(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if ccref == nil {
			initCacheStore()
		}
		return next(c)
	}
}

type cacheStore struct {
	Challenges *cache.Cache
	Paths      *cache.Cache
	CIDs       *cache.Cache
}

func initCacheStore() {
	ccref = &cacheStore{
		Challenges: cache.New(5*time.Minute, 10*time.Minute),
		Paths:      cache.New(30*time.Minute, 1*time.Hour),
		CIDs:       cache.New(30*time.Minute, 1*time.Hour),
	}
}

func cacheKey(e echo.Context, key string) string {
	return fmt.Sprintf(SessionID(e), ".", key)
}
