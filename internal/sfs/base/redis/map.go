package redis

import (
	"github.com/sonr-io/core/internal/sfs/types"
)

type sfsMap struct {
	key           string
	backupEntries map[string]string
	isOnline      bool
}

// NewMap creates a new map
func NewMap(key string) types.SFSMap {
	return &sfsMap{
		key:           key,
		isOnline:      ifr.rdb != nil,
		backupEntries: make(map[string]string),
	}
}

func (st *sfsMap) Add(key string, value string) error {
	if st.isOnline {
		return addMapItem(st.key, key, value)
	}
	st.backupEntries[key] = value
	return nil
}

func (st *sfsMap) Remove(key string) error {
	if st.isOnline {
		_, err := delMapItem(st.key, key)
		return err
	}
	delete(st.backupEntries, key)
	return nil
}

func (st *sfsMap) Contains(key string) (bool, error) {
	if !st.isOnline {
		_, ok := st.backupEntries[key]
		return ok, nil
	}
	return existsInMap(st.key, key)
}

func (st *sfsMap) GetAll() (map[string]string, error) {
	if !st.isOnline {
		return st.backupEntries, nil
	}
	return getMap(st.key)
}

func (st *sfsMap) Get(key string) (string, error) {
	if !st.isOnline {
		return st.backupEntries[key], nil
	}
	sti, err := st.GetAll()
	if err != nil {
		return "", err
	}
	return sti[key], nil
}
