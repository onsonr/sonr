package types

import (
	"encoding/json"
	fmt "fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/shengdoushi/base58"
	"github.com/sonrhq/core/pkg/crypto"
)

// VerificationMethodOption is used to define options that modify the creation of the verification method
type VerificationMethodOption func(vm *VerificationMethod, method DIDMethod) error

// WithController sets the controller of a verificationMethod
func WithController(v string) VerificationMethodOption {
	return func(vm *VerificationMethod, method DIDMethod) error {
		vm.Controller = v
		return nil
	}
}

// WithFragmentSuffix sets the fragment of the ID on a verificationMethod
func WithFragmentSuffix(v string) VerificationMethodOption {
	return func(vm *VerificationMethod, method DIDMethod) error {
		vm.Id = fmt.Sprintf("%s#%s", vm.Id, v)
		return nil
	}
}

// WithMetadataValues sets the metadata value of a verificationMethod
func WithMetadataValues(kvs ...KeyValuePair) VerificationMethodOption {
	return func(vm *VerificationMethod, method DIDMethod) error {
		for _, kv := range kvs {
			vm.SetMetadataValue(kv.Key, kv.Value)
		}
		return nil
	}
}

//
// VerificationMethod Creation Functions
//

// // VerificationMethod applies the given options and builds a verification method from this Key
func NewVerificationMethodFromPubKey(pk *crypto.PubKey, method DIDMethod, opts ...VerificationMethodOption) (*VerificationMethod, error) {
	vm := &VerificationMethod{
		Id:                 method.Format(pk.Multibase()),
		Type:               pk.KeyType,
		PublicKeyMultibase: pk.Multibase(),
		Metadata:           make([]*KeyValuePair, 0),
	}
	for _, opt := range opts {
		if err := opt(vm, method); err != nil {
			return nil, err
		}
	}
	return vm, nil
}

// NewVerificationMethodFromSonrAcc creates a verification method from the default wallet account
func NewVerificationMethodFromSonrAcc(pk *crypto.PubKey, options ...FormatOption) (*VerificationMethod, error) {
	accAddress, err := bech32.ConvertAndEncode("snr", pk.Bytes())
	if err != nil {
		return nil, err
	}
	vm := &VerificationMethod{
		Id:                  NewSonrID(accAddress),
		Type:                crypto.Secp256k1KeyType.PrettyString(),
		BlockchainAccountId: accAddress,
		PublicKeyMultibase:  pk.Multibase(),
		Metadata:            make([]*KeyValuePair, 0),
	}
	return vm, nil
}

// PubKey returns the public key of the verification method
func (v *VerificationMethod) PubKey() (*crypto.PubKey, error) {
	return crypto.PubKeyFromDID(v.Id)
}

// UnmarshalJSON implements the json.Unmarshaler interface for the VerificationMethod type.
func (v *VerificationMethod) UnmarshalJSON(bytes []byte) error {
	type Alias VerificationMethod
	tmp := Alias{}
	err := json.Unmarshal(bytes, &tmp)
	if err != nil {
		return err
	}
	*v = (VerificationMethod)(tmp)
	return nil
}

// CredentialDiscriptor is a descriptor for a credential for VerificationMethod which contains WebAuthnCredential
func (vm *VerificationMethod) CredentialDescriptor() (protocol.CredentialDescriptor, error) {
	if vm.Type != crypto.WebAuthnKeyType.PrettyString() {
		return protocol.CredentialDescriptor{}, fmt.Errorf("verification method is not of type WebAuthn")
	}
	cred, err := vm.WebAuthnCredential()
	if err != nil {
		return protocol.CredentialDescriptor{}, err
	}
	stdCred := cred.ToStdCredential()
	return stdCred.Descriptor(), nil
}

// IDFragmentSuffix returns the fragment of the ID of the VerificationMethod
func (vm *VerificationMethod) IDFragmentSuffix() string {
	ptrs := strings.Split(vm.Id, "#")
	return ptrs[len(ptrs)-1]
}

// IsBlockchainAccount returns true if the VerificationMethod is a blockchain account
func (vm *VerificationMethod) IsBlockchainAccount() bool {
	return vm.BlockchainAccountId != ""
}

// PublicKey returns the public key of the VerificationMethod
func (vm *VerificationMethod) PublicKey() ([]byte, error) {
	switch vm.Type {
	case crypto.Ed25519KeyType.PrettyString():
		return base58.Decode(vm.PublicKeyMultibase, base58.BitcoinAlphabet)
	case crypto.Secp256k1KeyType.PrettyString():
		_, bz, err := bech32.DecodeAndConvert(vm.BlockchainAccountId)
		if err != nil {
			return nil, err
		}
		return bz, nil
	default:
		return nil, fmt.Errorf("unsupported key type")
	}
}

func (vm *VerificationMethod) SetMetadata(data map[string]string) {
	vm.Metadata = MapToKeyValueList(data)
}

// SetMetadataValue sets the metadata value for the given key
func (vm *VerificationMethod) SetMetadataValue(key, value string) {
	for i, kv := range vm.Metadata {
		if kv.Key == key {
			vm.Metadata[i].Value = value
			return
		}
	}
	vm.Metadata = append(vm.Metadata, &KeyValuePair{Key: key, Value: value})
}

// GetMetadata returns the metadata value for the given key
func (vm *VerificationMethod) GetMetadataValue(key string) (string, bool) {
	ok := vm.HasMetadataValue(key)
	if !ok {
		return "", false
	}
	for _, kv := range vm.Metadata {
		if kv.Key == key {
			return kv.Value, true
		}
	}
	return "", false
}

// HasMetadata returns true if the VerificationMethod has the given metadata key
func (vm *VerificationMethod) HasMetadataValue(key string) bool {
	for _, kv := range vm.Metadata {
		if kv.Key == key {
			return true
		}
	}
	return false
}

// ToVerificationRelationship returns a VerificationRelationship from the VerificationMethod
func (vm *VerificationMethod) ToVerificationRelationship(controller string) VerificationRelationship {
	if vm.Controller == "" {
		vm.Controller = controller
	}
	return VerificationRelationship{
		VerificationMethod: vm,
		Reference:          vm.Id,
	}
}
