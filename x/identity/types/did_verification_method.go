package types

import (
	"encoding/json"
	fmt "fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/shengdoushi/base58"
)

var (
	knownAddrPrefixes = []string{
		"snr",
		"btc",
		"0x",
		"cosmos",
		"fil",
	}
)

// VerificationMethodOption is used to define options that modify the creation of the verification method
type VerificationMethodOption func(vm *VerificationMethod) error

// WithController sets the controller of a verificationMethod
func WithController(v string) VerificationMethodOption {
	return func(vm *VerificationMethod) error {
		_, err := ParseDID(v)
		if err != nil {
			return err
		}
		vm.Controller = v
		return nil
	}
}

// WithIDFragmentSuffix sets the fragment of the ID on a verificationMethod
func WithIDFragmentSuffix(v string) VerificationMethodOption {
	return func(vm *VerificationMethod) error {
		vm.Id = fmt.Sprintf("%s#%s", vm.Id, v)
		return nil
	}
}

// WithBlockchainAccount sets the blockchain account of a verificationMethod
func WithBlockchainAccount(v string) VerificationMethodOption {
	return func(vm *VerificationMethod) error {
		vm.BlockchainAccountId = v
		return nil
	}
}

//
// VerificationMethod Creation Functions
//

// NewVerificationMethod is a convenience method to easily create verificationMethods based on a set of given params.
// It automatically encodes the provided public key based on the keyType.
func NewVerificationMethod(id string, keyType KeyType, controller string, key interface{}) (*VerificationMethod, error) {
	vm := &VerificationMethod{
		Id:         id,
		Type:       keyType,
		Controller: controller,
	}
	// Check for Secp256k1 key
	if keyType == KeyType_KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019 {
		// Switch Interface to *secp256k1.PublicKey or string
		switch key.(type) {
		case *secp256k1.PubKey:
			vm.BlockchainAccountId = key.(*secp256k1.PubKey).Address().String()
		case string:
			vm.BlockchainAccountId = key.(string)
		default:
			return nil, fmt.Errorf("key is not a secp256k1.PublicKey or string")
		}
	}
	return vm, nil
}

// PubKey returns the public key of the verification method
func (v *VerificationMethod) PubKey() (*PubKey, error) {
	return PubKeyFromDID(v.Id)
}

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

func findAddrPrefix(addr string) (string, string) {
	for _, prefix := range knownAddrPrefixes {
		if strings.HasPrefix(addr, prefix) {
			return prefix, strings.TrimPrefix(addr, prefix)
		}
	}
	return "", addr
}

// CredentialDiscriptor is a descriptor for a credential for VerificationMethod which contains WebAuthnCredential
func (vm *VerificationMethod) CredentialDescriptor() (protocol.CredentialDescriptor, error) {
	if vm.Type != KeyType_KeyType_WEB_AUTHN_AUTHENTICATION_2018 {
		return protocol.CredentialDescriptor{}, fmt.Errorf("verification method is not of type WebAuthn")
	}
	cred, err := vm.WebAuthnCredential()
	if err != nil {
		return protocol.CredentialDescriptor{}, err
	}
	stdCred := cred.ToStdCredential()
	return stdCred.Descriptor(), nil
}

// IsBlockchainAccount returns true if the VerificationMethod is a blockchain account
func (vm *VerificationMethod) IsBlockchainAccount() bool {
	return vm.BlockchainAccountId != ""
}

// PublicKey returns the public key of the VerificationMethod
func (vm *VerificationMethod) PublicKey() ([]byte, error) {
	switch vm.Type {
	case KeyType_KeyType_ED25519_VERIFICATION_KEY_2018:
		return base58.Decode(vm.PublicKeyMultibase, base58.BitcoinAlphabet)
	case KeyType_KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019:
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

// GetMetadata returns the metadata value for the given key
func (vm *VerificationMethod) GetMetadataValue(key string) string {
	for _, kv := range vm.Metadata {
		if kv.Key == key {
			return kv.Value
		}
	}
	return ""
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
