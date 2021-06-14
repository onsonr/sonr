package models

import (
	"log"

	"github.com/prologic/bitcask"
	"google.golang.org/protobuf/proto"
)

type Store interface {
	Handle(req *StoreRequest) *StoreResponse
	Has(key []byte) bool
	Get(key []byte) (*StoreEntry, *SonrError)
	Put(value *StoreEntry) *SonrError
}

type store struct {
	Store
	database *bitcask.Bitcask
	device   *Device
}

// Initializes new memory store
func InitStore(d *Device) (Store, *SonrError) {
	// Open Store
	db, err := bitcask.Open(d.WorkingSupportPath("mem-store-db"))
	if err != nil {
		log.Println(err)
		return nil, NewError(err, ErrorMessage_STORE_INIT)
	}

	// Return Store
	return &store{
		database: db,
		device:   d,
	}, nil
}

// Checks if Store Has given Key
func (s *store) Has(key []byte) bool {
	if s.database.Has(key) {
		return true
	}
	return false
}

// Return Value From Store
func (s *store) Get(key []byte) (*StoreEntry, *SonrError) {
	// Get Value
	val, err := s.database.Get(key)
	if err != nil {
		log.Println(err)
		return nil, NewError(err, ErrorMessage_STORE_GET)
	}

	// Unmarshal Data
	entry := &StoreEntry{}
	err = proto.Unmarshal(val, entry)
	if err != nil {
		return nil, NewUnmarshalError(err)
	}
	return entry, nil
}

// Puts Value in Store
func (s *store) Put(entry *StoreEntry) *SonrError {
	// Get Key
	key := entry.KeyBytes()

	// Marshal Data
	value, err := proto.Marshal(entry)
	if err != nil {
		return NewMarshalError(err)
	}

	// Put Value
	err = s.database.Put(key, value)
	if err != nil {
		log.Println(err)
		return NewError(err, ErrorMessage_STORE_PUT)
	}
	return nil
}

// @ Handle Store Request
func (s *store) Handle(req *StoreRequest) *StoreResponse {
	// Check Request Type
	if req.IsGet() {
		// Get Request Method
		entry, err := s.Get(req.KeyBytes())
		if err != nil {
			return NewStoreGetResponse(nil, err)
		}
		return NewStoreGetResponse(entry, nil)
	} else if req.IsPut() {
		// Put Request Method
		if err := s.Put(req.ValueToEntry()); err != nil {
			return NewStorePutResponse(false, err)
		} else {
			return NewStorePutResponse(true, nil)
		}
	} else {
		// Has Request Method
		return NewStoreHasResponse(s.Has(req.KeyBytes()))
	}
}
