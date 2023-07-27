package sfs

import (
	"github.com/sonrhq/core/internal/sfs/base"
	"github.com/sonrhq/core/internal/sfs/types"
)

type Map = types.SFSMap

func InitMap(key string) Map {
	return base.NewMap(key)
}

type Set = types.SFSSet

func InitSet(key string) Set {
	return base.NewSet(key)
}
