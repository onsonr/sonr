package types

import "github.com/sonr-io/sonr/internal/crypto"

// DIDMethod is a DID method
type DIDMethod string

// store returns the store for the DID method
func (d DIDMethod) store() *DIDStore {
	return GetMethodStore(d)
}

// CoinType returns the coin type for the DID method
func (d DIDMethod) CoinType() crypto.CoinType {
	return crypto.CoinTypeFromDidMethod(d.String())
}

// Equals returns true if the DID methods are equal
func (d DIDMethod) Equals(other DIDMethod) bool {
	return d.String() == other.String()
}

// HasKey returns true if the key exists in the store
func (d DIDMethod) HasKey(key string) bool {
	ok, _ := d.store().HasKey(key)
	return ok
}

// GetKey returns the value for the key
func (d DIDMethod) GetKey(key string) string {
	v, err := d.store().GetKey(key)
	if err != nil {
		return ""
	}
	return v
}

// SetKey sets the value for the key
func (d DIDMethod) SetKey(key string, value string) {
	err := d.store().SetKey(key, value)
	if err != nil {
		panic(err)
	}
}

// String returns the string representation of the DID method
func (d DIDMethod) String() string {
	return string(d)
}

// AnySignerEntity is an entity that can sign and verify messages
type AnySignerEntity interface {
	Sign([]byte) ([]byte, error)
	PublicKey() (*crypto.Secp256k1PubKey, error)
	Verify(msg []byte, sig []byte) (bool, error)
}
