package bank

import (
	"errors"

	"github.com/sonr-hq/sonr/pkg/vault/internal/session"
)

func (v *VaultBank) getEntryFromCache(id string) (*session.SessionEntry, error) {
	val, ok := v.cache.Get(id)
	if !ok {
		return nil, errors.New("Failed to find entry for ID")
	}
	e, ok := val.(*session.SessionEntry)
	if !ok {
		return nil, errors.New("Invalid type for session entry")
	}
	return e, nil
}

func (v *VaultBank) putEntryIntoCache(entry *session.SessionEntry) error {
	if entry == nil {
		return errors.New("Entry cannot be nil to put into cache")
	}
	return v.cache.Add(entry.ID, entry, -1)
}
