package common

import (
	"time"

	cache "github.com/SporkHubr/echo-http-cache"
	"github.com/SporkHubr/echo-http-cache/adapter/memory"
	"github.com/labstack/echo/v4"
)

var (
	// cacheClient is the cache client
	cacheClient *cache.Client

	// memCache is the memory cache
	memCache cache.Adapter
)

func UseCache(e *echo.Echo) error {
	// Setup the cache
	var err error
	memCache, err = initMemCache()
	if err != nil {
		return err
	}

	// Setup the cache client
	cacheClient, err = getCacheClient(memCache)
	if err != nil {
		return err
	}
	e.Use(cacheClient.Middleware())
	return nil
}

func initMemCache() (cache.Adapter, error) {
	memCache, err := memory.NewAdapter(
		memory.AdapterWithAlgorithm(memory.LRU),
		memory.AdapterWithCapacity(10000000),
	)
	if err != nil {
		return nil, err
	}
	return memCache, nil
}

func getCacheClient(a cache.Adapter) (*cache.Client, error) {
	cacheClient, err := cache.NewClient(
		cache.ClientWithAdapter(a),
		cache.ClientWithTTL(10*time.Minute),
	)
	if err != nil {
		return nil, err
	}
	return cacheClient, nil
}
