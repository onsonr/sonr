package models

import (
	"github.com/prologic/bitcask"
)

type Store interface {
	Has(key StoreKeys) bool
	Get(key StoreKeys) ([]byte, *SonrError)
	Put(key StoreKeys, value []byte) *SonrError
}

type store struct {
	Store
	database *bitcask.Bitcask
	device   *Device
}

// Initializes new memory store
func InitStore(cr *ConnectionRequest) (Store, *SonrError) {
	// Open Store
	db, err := bitcask.Open(cr.GetDevice().WorkingFilePath("mem-store-db"))
	if err != nil {
		return nil, NewError(err, ErrorMessage_STORE_INIT)
	}

	// Return Store
	return &store{
		database: db,
		device:   cr.GetDevice(),
	}, nil
}

// Checks if Store Has given Key
func (s *store) Has(key StoreKeys) bool {
	if s.database.Has(key.Bytes()) {
		return true
	}
	return false
}

// Return Value From Store
func (s *store) Get(key StoreKeys) ([]byte, *SonrError) {
	val, err := s.database.Get(key.Bytes())
	if err != nil {
		return nil, NewError(err, ErrorMessage_STORE_GET)
	}
	return val, nil
}

// Puts Value in Store
func (s *store) Put(key StoreKeys, value []byte) *SonrError {
	err := s.database.Put(key.Bytes(), value)
	if err != nil {
		return NewError(err, ErrorMessage_STORE_PUT)
	}
	return nil
}
