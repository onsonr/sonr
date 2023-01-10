package session

import (
	"errors"

	gocache "github.com/patrickmn/go-cache"
)

func GetEntry(id string, cache *gocache.Cache) (*SessionEntry, error) {
	val, ok := cache.Get(id)
	if !ok {
		return nil, errors.New("Failed to find entry for ID")
	}
	e, ok := val.(*SessionEntry)
	if !ok {
		return nil, errors.New("Invalid type for session entry")
	}
	return e, nil
}

func PutEntry(entry *SessionEntry, cache *gocache.Cache) error {
	if entry == nil || cache == nil {
		return errors.New("Entry or Cache cannot be nil to put Entry")
	}
	return cache.Add(entry.ID, entry, -1)
}
