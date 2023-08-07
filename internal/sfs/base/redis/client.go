package redis

import (
	"context"

	"github.com/go-redis/redis/v8"

	"github.com/sonrhq/core/config"
)

var ifr *IceFireRedis

// IceFireRedis is a wrapper around the redis client
type IceFireRedis struct {
	rdb *redis.Client
	ctx context.Context
}

func init() {
	ifr = &IceFireRedis{
		ctx: context.Background(),
		rdb: redis.NewClient(&redis.Options{
			Addr:     config.IceFireKVHost(),
			Password: "", // no password set
			DB:       0,  // use default DB
		}),
	}
}

func existsInSet(key string, member string) (bool, error) {
	return ifr.rdb.SIsMember(ifr.ctx, key, member).Result()
}

func getSet(key string) ([]string, error) {
	return ifr.rdb.SMembers(ifr.ctx, key).Result()
}

func delSetItem(key string, member string) (int64, error) {
	return ifr.rdb.SRem(ifr.ctx, key, member).Result()
}

func addSetItem(key string, member string) error {
	return ifr.rdb.SAdd(ifr.ctx, key, member).Err()
}

func existsInMap(key string, field string) (bool, error) {
	return ifr.rdb.HExists(ifr.ctx, key, field).Result()
}

func getMap(key string) (map[string]string, error) {
	return ifr.rdb.HGetAll(ifr.ctx, key).Result()
}

func delMapItem(key string, field string) (int64, error) {
	return ifr.rdb.HDel(ifr.ctx, key, field).Result()
}

func addMapItem(key string, field string, value string) error {
	return ifr.rdb.HSet(ifr.ctx, key, field, value).Err()
}
