package types

import (
	"github.com/sonr-io/kryptology/pkg/accumulator"

	"github.com/sonr-io/core/pkg/crypto"
	identitytypes "github.com/sonr-io/core/x/identity/types"
)

// DIDIdentifier is a DID identifier
type DIDIdentifier string

// AddResource adds a resource to the store
func (d DIDIdentifier) AddResource(k string, v []byte) (DIDResource, error) {
	err := d.store().SetKey(k, encodeResource(v))
	if err != nil {
		return "", err
	}
	did := NewResource(d, k)
	return did, nil
}

// FetchResource fetches a resource from the store
func (d DIDIdentifier) FetchResource(k string) ([]byte, error) {
	vstr, err := d.store().GetKey(k)
	if err != nil {
		return nil, err
	}
	biz := decodeResource(vstr)
	return biz, nil
}

// HasKey returns true if the key exists in the store
func (d DIDIdentifier) HasKey(key string) bool {
	v, err := d.store().HasKey(key)
	if err != nil {
		return false
	}
	return v
}

// GetKey returns the value for the key
func (d DIDIdentifier) GetKey(key string) string {
	v, err := d.store().GetKey(key)
	if err != nil {
		return ""
	}
	return v
}

// SetKey sets the value for the key
func (d DIDIdentifier) SetKey(key string, value string) {
	err := d.store().SetKey(key, value)
	if err != nil {
		panic(err)
	}
}

// AppendKeyList appends a list of values to the key
func (d DIDIdentifier) AppendKeyList(key string, values ...string) {
	err := d.store().AppendList(key, values...)
	if err != nil {
		panic(err)
	}
}

// RemoveKeyList removes a list of values from the key
func (d DIDIdentifier) RemoveKeyList(key string, values ...string) {
	err := d.store().RemoveList(key, values...)
	if err != nil {
		panic(err)
	}
}

// GetKeyList returns the list of values for the key
func (d DIDIdentifier) GetKeyList(key string) []string {
	vs, err := d.store().GetList(key)
	if err != nil {
		return []string{}
	}
	return vs
}

// String returns the string representation of the DID
func (d DIDIdentifier) String() string {
	return string(d)
}

// store returns the store for the DID
func (d DIDIdentifier) store() *DIDStore {
	return GetIdentifierStore(d)
}

// DIDSecretKey is an interface for a DID secret key
type DIDSecretKey interface {
	// AccumulatorKey returns the accumulator key for the DID
	AccumulatorKey() (*accumulator.SecretKey, error)

	// Bytes returns the bytes of the secret key
	Bytes() []byte

	// Encrypt encrypts a byte array
	Encrypt(bz []byte) ([]byte, error)

	// Decrypt decrypts a byte array
	Decrypt(ciphertext []byte) ([]byte, error)
}

// ControllerAccount is an interface for a controller account
type ControllerAccount = identitytypes.ControllerAccount

// WalletAccount is an interface that provides acces to a DID Wallet
type WalletAccount interface {
	// Address returns the address of the account
	Address() string
	Info() *crypto.AccountData
	Method() DIDMethod
	Sign(msg []byte) ([]byte, error)
	PublicKey() (*crypto.Secp256k1PubKey, error)
	Type() string
	Verify(msg []byte, sig []byte) (bool, error)
}
