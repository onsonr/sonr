package sfs

import (
	"sonr.io/core/internal/sfs/base/redis"
	"sonr.io/core/internal/sfs/types"
)

type Map = types.SFSMap

func InitMap(key string) Map {
	return redis.NewMap(key)
}

type Set = types.SFSSet

func InitSet(key string) Set {
	return redis.NewSet(key)
}
