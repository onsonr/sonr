package resolver

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/sonrhq/core/pkg/node"
)

// ! ||--------------------------------------------------------------------------------||
// ! ||                          Global Resolver Store Methods                         ||
// ! ||--------------------------------------------------------------------------------||

// InsertRecord inserts a record into the IPFS store for the given controller
func InsertRecord(controller string, key string, value interface{}) error {
	ds, err := node.OpenKeyValueStore(context.Background(), controller)
	defer ds.Close()
	if err != nil {
		return err
	}
	var vBiz []byte
	switch value.(type) {
	case string:
		v, err := base64.StdEncoding.DecodeString(value.(string))
		if err != nil {
			return err
		}
		vBiz = v
	case []byte:
		vBiz = value.([]byte)
	default:
		return fmt.Errorf("value must be a string or []byte")
	}
	store := makeIpfsStore(ds, controller)
	return store.Put(key, vBiz)
}

// GetRecord gets a record from the IPFS store for the given controller
func GetRecord(controller string, key string) ([]byte, error) {
	ds, err := node.OpenKeyValueStore(context.Background(), controller)
	defer ds.Close()
	if err != nil {
		return nil, err
	}
	store := makeIpfsStore(ds, controller)
	vBiz, err := store.Get(key)
	if err != nil {
		return nil, err
	}
	return vBiz, nil
}

// DeleteRecord deletes a record from the IPFS store for the given controller
func DeleteRecord(controller string, key string) error {
	ds, err := node.OpenKeyValueStore(context.Background(), controller)
	defer ds.Close()
	if err != nil {
		return err
	}
	store := makeIpfsStore(ds, controller)
	return store.Delete(key)
}

// ListRecords lists all records in the IPFS store for the given controller
func ListRecords(controller string) (map[string][]byte, error) {
	ds, err := node.OpenKeyValueStore(context.Background(), controller)
	defer ds.Close()
	if err != nil {
		return nil, err
	}
	store := makeIpfsStore(ds, controller)
	m := make(map[string][]byte)
	for k, v := range store.All() {
		m[k] = v
	}
	return m, nil
}

// ! ||--------------------------------------------------------------------------------||
// ! ||              IPFS Based Wallet Store Implementation using OrbitDB              ||
// ! ||--------------------------------------------------------------------------------||

type ipfsStore struct {
	controller string
	node.IPFSKVStore
}

func makeIpfsStore(store node.IPFSKVStore, controller string) *ipfsStore {
	return &ipfsStore{
		IPFSKVStore: store,
		controller:  controller,
	}
}

func (s *ipfsStore) All() map[string][]byte {
	return s.IPFSKVStore.All()
}

func (s *ipfsStore) Address() string {
	return s.IPFSKVStore.Address().String()
}

func (s *ipfsStore) DBName() string {
	return s.IPFSKVStore.DBName()
}

func (s *ipfsStore) Identity() string {
	return s.IPFSKVStore.Identity().ID
}

func (s *ipfsStore) PublicKey() []byte {
	return s.IPFSKVStore.Identity().PublicKey
}

func (s *ipfsStore) Type() string {
	return s.IPFSKVStore.Type()
}

func (s *ipfsStore) Close() error {
	return s.Close()
}

func (s *ipfsStore) Get(key string) ([]byte, error) {
	return s.IPFSKVStore.Get(context.Background(), key)
}

func (s *ipfsStore) Put(key string, value []byte) error {
	_, err := s.IPFSKVStore.Put(context.Background(), key, value)
	return err
}

func (s *ipfsStore) Delete(key string) error {
	_, err := s.IPFSKVStore.Delete(context.Background(), key)
	return err
}
