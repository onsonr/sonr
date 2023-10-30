package types

import (
	"fmt"
	"strings"

	"lukechampine.com/blake3"

	"sonr.io/core/internal/sfs"
	"github.com/spf13/viper"
)

// DIDStore is a store for a DID
type DIDStore struct {
	Name       string
	Store      sfs.Map
	Method     DIDMethod
	Identifier DIDIdentifier
	IsMethod   bool
}

// GetMethodStore returns a store for a DID method
func GetMethodStore(name DIDMethod) *DIDStore {
	hash := blake3.Sum256([]byte(name.String()))
	id := fmt.Sprintf("%s://store/%s", EnvChainID(), hash)
	return &DIDStore{
		Name:     id,
		Store:    sfs.InitMap(id),
		Method:   name,
		IsMethod: true,
	}
}

// GetIdentifierStore returns a store for a DID identifier
func GetIdentifierStore(name DIDIdentifier) *DIDStore {
	hash := blake3.Sum256([]byte(name.String()))
	id := fmt.Sprintf("%s://store/%s", EnvChainID(), hash)
	return &DIDStore{
		Name:       id,
		Store:      sfs.InitMap(id),
		Identifier: name,
		IsMethod:   false,
	}
}

// HasKey returns whether a key exists in the store
func (g *DIDStore) HasKey(k string) (bool, error) {
	return g.Store.Contains(k)
}

// GetKey returns the value of a key in the store
func (g *DIDStore) GetKey(k string) (string, error) {
	return g.Store.Get(k)
}

// SetKey sets the value of a key in the store
func (g *DIDStore) SetKey(k string, v string) error {
	return g.Store.Add(k, v)
}

// AppendList appends a list of values to a key in the store
func (g *DIDStore) AppendList(key string, values ...string) error {
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
func (g *DIDStore) RemoveList(key string, values ...string) error {
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
func (g *DIDStore) GetList(key string) ([]string, error) {
	val, _ := g.Store.Get(key)
	if val == "" {
		return []string{}, nil
	}
	return strings.Split(val, ","), nil
}

// StoreKey returns the store key for the store
func (g *DIDStore) StoreKey() string {
	id := blake3.Sum256([]byte(g.Name))
	return fmt.Sprintf("%s://store/%s", EnvChainID(), id)
}

// EnvChainID returns the chain ID from the configuration. (default: sonr-localnet-1)
func EnvChainID() string {
	return viper.GetString("launch.chain-id")
}

// EnvNodeGrpcHostAddress returns the host and port of the Node P2P
func EnvNodeGrpcHostAddress() string {
	return fmt.Sprintf("%s:%d", viper.GetString("node.grpc.host"), viper.GetInt("node.grpc.port"))
}
