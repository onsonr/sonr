package didstore

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
	"lukechampine.com/blake3"

	"github.com/sonrhq/sonr/internal/sfs"
)

// Store is a store for a DID
type Store struct {
	Name       string
	Store      sfs.Map
	Method    string
	Identifier string
	IsMethod   bool
}

// GetMethod returns a store for a DID method
func GetMethod(name string) *Store {
	hash := blake3.Sum256([]byte(name))
	id := fmt.Sprintf("%s://store/%s", EnvChainID(), hash)
	return &Store{
		Name:     id,
		Store:    sfs.InitMap(id),
		Method:   name,
		IsMethod: true,
	}
}

// GetIdentifier returns a store for a DID identifier
func GetIdentifier(name string) *Store {
	hash := blake3.Sum256([]byte(name))
	id := fmt.Sprintf("%s://store/%s", EnvChainID(), hash)
	return &Store{
		Name:       id,
		Store:      sfs.InitMap(id),
		Identifier: name,
		IsMethod:   false,
	}
}

// HasKey returns whether a key exists in the store
func (g *Store) HasKey(k string) (bool, error) {
	return g.Store.Contains(k)
}

// GetKey returns the value of a key in the store
func (g *Store) GetKey(k string) (string, error) {
	return g.Store.Get(k)
}

// SetKey sets the value of a key in the store
func (g *Store) SetKey(k string, v string) error {
	return g.Store.Add(k, v)
}

// AppendList appends a list of values to a key in the store
func (g *Store) AppendList(key string, values ...string) error {
	joined := strings.Join(values, ",")
	val, _ := g.Store.Get(key)
	if val == "" {
		return g.Store.Add(key, joined)
	}
	existing := strings.Split(val, ",")
	// ensure no duplicates
	for _, v := range values {
		for _, e := range existing {
			if e == v {
				continue
			}
		}
		existing = append(existing, v)
	}
	return g.Store.Add(key, strings.Join(existing, ","))
}

// RemoveList removes a list of values from a key in the store
func (g *Store) RemoveList(key string, values ...string) error {
	val, _ := g.Store.Get(key)
	if val == "" {
		return nil
	}
	existing := strings.Split(val, ",")
	// ensure no duplicates
	for _, v := range values {
		for i, e := range existing {
			if e == v {
				existing = append(existing[:i], existing[i+1:]...)
			}
		}
	}
	return g.Store.Add(key, strings.Join(existing, ","))
}

// GetList returns a list of values from a key in the store
func (g *Store) GetList(key string) ([]string, error) {
	val, _ := g.Store.Get(key)
	if val == "" {
		return []string{}, nil
	}
	return strings.Split(val, ","), nil
}

// StoreKey returns the store key for the store
func (g *Store) StoreKey() string {
	id := blake3.Sum256([]byte(g.Name))
	return fmt.Sprintf("%s://store/%s", EnvChainID(), id)
}

// EnvChainID returns the chain ID from the configuration. (default: sonr-localnet-1)
func EnvChainID() string {
	return viper.GetString("launch.chain-id")
}

