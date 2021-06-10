package models

import (
	"github.com/prologic/bitcask"
	"google.golang.org/protobuf/proto"
)

type Store interface {
	Has(key []byte) bool
	HasCrypto() bool
	HasSettings() bool
	Get(key []byte) ([]byte, *SonrError)
	GetCrypto() (*User_Crypto, *SonrError)
	GetSettings() (*User_Settings, *SonrError)
	Put(key []byte, value []byte) *SonrError
	PutCrypto(*User_Crypto) *SonrError
	PutSettings(*User_Settings) *SonrError
}

type store struct {
	Store
	database    *bitcask.Bitcask
	device      *Device
	cryptoKey   []byte
	settingsKey []byte
}

// Initializes new memory store
func InitStore(d *Device) (Store, *SonrError) {
	// Open Store
	db, err := bitcask.Open(d.WorkingFilePath("mem-store-db"))
	if err != nil {
		return nil, NewError(err, ErrorMessage_STORE_INIT)
	}

	// Return Store
	return &store{
		database:    db,
		device:      d,
		cryptoKey:   StoreKeys_CRYPTO.Bytes(),
		settingsKey: StoreKeys_SETTINGS.Bytes(),
	}, nil
}

// Checks if Store Has given Key
func (s *store) Has(key []byte) bool {
	if s.database.Has(key) {
		return true
	}
	return false
}

// Checks if Store Has User_Crypto
func (s *store) HasCrypto() bool {
	if s.database.Has(s.cryptoKey) {
		return true
	}
	return false
}

// Checks if Store Has User_Settings
func (s *store) HasSettings() bool {
	if s.database.Has(s.settingsKey) {
		return true
	}
	return false
}

// Return Value From Store
func (s *store) Get(key []byte) ([]byte, *SonrError) {
	val, err := s.database.Get(key)
	if err != nil {
		return nil, NewError(err, ErrorMessage_STORE_GET)
	}
	return val, nil
}

// Return Value From Store
func (s *store) GetCrypto() (*User_Crypto, *SonrError) {
	val, err := s.database.Get(s.cryptoKey)
	if err != nil {
		return nil, NewError(err, ErrorMessage_STORE_GET)
	}

	crypto := &User_Crypto{}
	serr := proto.Unmarshal(val, crypto)
	if serr != nil {
		return nil, NewError(err, ErrorMessage_UNMARSHAL)
	}
	return crypto, nil
}

// Return Value From Store
func (s *store) GetSettings() (*User_Settings, *SonrError) {
	val, err := s.database.Get(s.settingsKey)
	if err != nil {
		return nil, NewError(err, ErrorMessage_STORE_GET)
	}

	settings := &User_Settings{}
	serr := proto.Unmarshal(val, settings)
	if serr != nil {
		return nil, NewError(err, ErrorMessage_UNMARSHAL)
	}
	return settings, nil
}

// Puts Value in Store
func (s *store) Put(key []byte, value []byte) *SonrError {
	err := s.database.Put(key, value)
	if err != nil {
		return NewError(err, ErrorMessage_STORE_PUT)
	}
	return nil
}

// Puts Value in Store
func (s *store) PutCrypto(data *User_Crypto) *SonrError {
	// Marshal Object
	value, serr := proto.Marshal(data)
	if serr != nil {
		return NewMarshalError(serr)
	}

	// Place Item
	err := s.database.Put(s.cryptoKey, value)
	if err != nil {
		return NewError(err, ErrorMessage_STORE_PUT)
	}
	return nil
}

// Puts Value in Store
func (s *store) PutSettings(data *User_Settings) *SonrError {
	// Marshal Object
	value, serr := proto.Marshal(data)
	if serr != nil {
		return NewMarshalError(serr)
	}

	// Place Item
	err := s.database.Put(s.settingsKey, value)
	if err != nil {
		return NewError(err, ErrorMessage_STORE_PUT)
	}
	return nil
}
