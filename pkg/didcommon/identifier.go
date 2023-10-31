package didcommon

import (
	"github.com/sonr-io/kryptology/pkg/accumulator"

	"github.com/sonrhq/sonr/internal/crypto"
	"github.com/sonrhq/sonr/pkg/didstore"
	identitytypes "github.com/sonrhq/sonr/x/identity/types"
)

// Identifier is a DID identifier
type Identifier string


// AddResource adds a resource to the store
func (d Identifier) AddResource(k string, v []byte) (DIDResource, error) {
	err := d.store().SetKey(k, EncodeResource(v))
	if err != nil {
		return "", err
	}
	did := NewResource(d.String(), k)
	return did, nil
}

// FetchResource fetches a resource from the store
func (d Identifier) FetchResource(k string) ([]byte, error) {
	vstr, err := d.store().GetKey(k)
	if err != nil {
		return nil, err
	}
	biz := DecodeResource(vstr)
	return biz, nil
}

// HasKey returns true if the key exists in the store
func (d Identifier) HasKey(key string) bool {
	v, err := d.store().HasKey(key)
	if err != nil {
		return false
	}
	return v
}

// GetKey returns the value for the key
func (d Identifier) GetKey(key string) string {
	v, err := d.store().GetKey(key)
	if err != nil {
		return ""
	}
	return v
}

// SetKey sets the value for the key
func (d Identifier) SetKey(key string, value string) {
	err := d.store().SetKey(key, value)
	if err != nil {
		panic(err)
	}
}

// AppendKeyList appends a list of values to the key
func (d Identifier) AppendKeyList(key string, values ...string) {
	err := d.store().AppendList(key, values...)
	if err != nil {
		panic(err)
	}
}

// RemoveKeyList removes a list of values from the key
func (d Identifier) RemoveKeyList(key string, values ...string) {
	err := d.store().RemoveList(key, values...)
	if err != nil {
		panic(err)
	}
}

// GetKeyList returns the list of values for the key
func (d Identifier) GetKeyList(key string) []string {
	vs, err := d.store().GetList(key)
	if err != nil {
		return []string{}
	}
	return vs
}

// String returns the string representation of the DID
func (d Identifier) String() string {
	return string(d)
}

// store returns the store for the DID identifier
func (d Identifier) store() *didstore.Store {
	return didstore.GetIdentifier(d.String())
}

// SecretKey is an interface for a DID secret key
type SecretKey interface {
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
	Method() Method
	Sign(msg []byte) ([]byte, error)
	PublicKey() (*crypto.Secp256k1PubKey, error)
	Type() string
	Verify(msg []byte, sig []byte) (bool, error)
}
