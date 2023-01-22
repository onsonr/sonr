package keeper

import (
	"errors"
)

func (v *VaultBank) getEntryFromCache(id string) (*Session, error) {
	val, ok := v.cache.Get(id)
	if !ok {
		return nil, errors.New("Failed to find entry for ID")
	}
	e, ok := val.(*Session)
	if !ok {
		return nil, errors.New("Invalid type for session entry")
	}
	return e, nil
}

func (v *VaultBank) putEntryIntoCache(entry *Session) error {
	if entry == nil {
		return errors.New("Entry cannot be nil to put into cache")
	}
	return v.cache.Add(entry.ID, entry, -1)
}
