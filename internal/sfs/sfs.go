package sfs

import (
	"github.com/sonr-io/sonr/internal/sfs/base/redis"
	"github.com/sonr-io/sonr/internal/sfs/types"
)

type Map = types.SFSMap

func InitMap(key string) Map {
	return redis.NewMap(key)
}

type Set = types.SFSSet

func InitSet(key string) Set {
	return redis.NewSet(key)
}
