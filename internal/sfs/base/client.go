package base

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/sonrhq/core/internal/local"
)

var ins *IceFireInstance

type IceFireInstance struct {
	rdb *redis.Client
	ctx context.Context
}

func init() {
	ins = &IceFireInstance{
		ctx: context.Background(),
		rdb: redis.NewClient(&redis.Options{
			Addr:     local.IceFireHost(),
			Password: "", // no password set
			DB:       0,  // use default DB
		}),
	}
}

func existsInSet(key string, member string) (bool, error) {
	return ins.rdb.SIsMember(ins.ctx, key, member).Result()
}

func getSet(key string) ([]string, error) {
	return ins.rdb.SMembers(ins.ctx, key).Result()
}

func delSetItem(key string, member string) (int64, error) {
	return ins.rdb.SRem(ins.ctx, key, member).Result()
}

func addSetItem(key string, member string) error {
	return ins.rdb.SAdd(ins.ctx, key, member).Err()
}

func existsInMap(key string, field string) (bool, error) {
	return ins.rdb.HExists(ins.ctx, key, field).Result()
}

func getMap(key string) (map[string]string, error) {
	return ins.rdb.HGetAll(ins.ctx, key).Result()
}

func delMapItem(key string, field string) (int64, error) {
	return ins.rdb.HDel(ins.ctx, key, field).Result()
}

func addMapItem(key string, field string, value string) error {
	return ins.rdb.HSet(ins.ctx, key, field, value).Err()
}
