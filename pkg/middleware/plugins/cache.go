package plugins

import (
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

type cache struct {
	redisClient *redis.Client
}

func Cache(ctx echo.Context) *cache {
	return ctx.Get("cache").(*cache)
}
