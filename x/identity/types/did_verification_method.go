package types

import (
	"crypto"
	"crypto/ed25519"
	"encoding/json"
	fmt "fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/shengdoushi/base58"
	common "github.com/sonr-hq/sonr/pkg/common"
	"github.com/sonr-hq/sonr/pkg/common/crypto/jwx"
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
		vm.ID = fmt.Sprintf("%s#%s", vm.ID, v)
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

// NewWebAuthnVM creates a new WebAuthn VerificationMethod
func NewWebAuthnVM(webauthnCredential *common.WebauthnCredential, options ...VerificationMethodOption) (*VerificationMethod, error) {
	// Configure base Verification MEthod
	vm := &VerificationMethod{
		ID:                 webauthnCredential.Did(),
		Type:               KeyType_KeyType_WEB_AUTHN_AUTHENTICATION_2018,
		PublicKeyMultibase: webauthnCredential.PublicKeyMultibase(),
		Metadata:           webauthnCredential.ToMetadata(),
	}

	// Apply VerificationMethod Options
	for _, opt := range options {
		err := opt(vm)
		if err != nil {
			return nil, err
		}
	}
	return vm, nil
}

// NewJWKVM creates a new JWK VerificationMethod
func NewJWKVM(id string, key crypto.PublicKey, options ...VerificationMethodOption) (*VerificationMethod, error) {
	_, err := ParseDID(id)
	if err != nil {
		return nil, err
	}
	// Check if Key is crypto.PublicKey
	vm := &VerificationMethod{
		ID:   id,
		Type: KeyType_KeyType_JSON_WEB_KEY_2020,
	}

	// Check if Key is crypto.PublicKey
	if _, ok := key.(crypto.PublicKey); !ok {
		return nil, fmt.Errorf("key is not a crypto.PublicKey")
	}

	// Convert to jwk.Key
	keyAsJWK, err := jwx.New(key).CreateEncJWK()
	if err != nil {
		return nil, err
	}
	// Convert to JSON and back to fix encoding of key material to make sure
	// an unmarshalled and newly created VerificationMethod are equal on object level.
	// The format of PublicKeyJwk in verificationMethod is a map[string]interface{}.
	// We can't use the Key.AsMap since the values of the map will all be internal jwk lib structs.
	// After unmarshalling all the fields will be map[string]string.
	keyAsJSON, err := json.Marshal(keyAsJWK)
	if err != nil {
		return nil, err
	}
	keyAsMap := map[string]string{}
	json.Unmarshal(keyAsJSON, &keyAsMap)
	vm.PublicKeyJwk = keyAsMap

	// Apply VerificationMethod Options
	for _, opt := range options {
		err := opt(vm)
		if err != nil {
			return nil, err
		}
	}
	return vm, nil
}

// NewEd25519VM creates a new Ed25519 VerificationMethod
func NewEd25519VM(id string, key ed25519.PublicKey, options ...VerificationMethodOption) (*VerificationMethod, error) {
	_, err := ParseDID(id)
	if err != nil {
		return nil, err
	}
	// Check if Key is crypto.PublicKey
	vm := &VerificationMethod{
		ID:                 id,
		Type:               KeyType_KeyType_ED25519_VERIFICATION_KEY_2018,
		PublicKeyMultibase: base58.Encode(key, base58.BitcoinAlphabet),
	}

	// Apply VerificationMethod Options
	for _, opt := range options {
		err := opt(vm)
		if err != nil {
			return nil, err
		}
	}
	return vm, nil
}

// NewSecp256k1VM creates a new Secp256k1 VerificationMethod
func NewSecp256k1VM(key *secp256k1.PubKey, options ...VerificationMethodOption) (*VerificationMethod, error) {
	prefix, addr := findAddrPrefix(key.Address().String())
	did := fmt.Sprintf("%s:%s", prefix, addr)
	bz, err := key.Marshal()
	if err != nil {
		return nil, err
	}
	// Configure base Verification MEthod
	vm := &VerificationMethod{
		ID:                 did,
		Type:               KeyType_KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019,
		PublicKeyMultibase: base58.Encode(bz, base58.BitcoinAlphabet),
	}

	// Apply VerificationMethod Options
	for _, opt := range options {
		err := opt(vm)
		if err != nil {
			return nil, err
		}
	}
	return vm, nil
}

// NewVerificationMethod is a convenience method to easily create verificationMethods based on a set of given params.
// It automatically encodes the provided public key based on the keyType.
func NewVerificationMethod(id string, keyType KeyType, controller string, key interface{}) (*VerificationMethod, error) {
	vm := &VerificationMethod{
		ID:         id,
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

func (v *VerificationMethod) Address() string {
	ptrs := strings.Split(v.ID, ":")
	return ptrs[len(ptrs)-1]
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

func (d *DidDocument) GetVerificationMethods() *VerificationMethods {
	return d.VerificationMethod
}

// FindByID find the first VerificationMethod which matches the provided DID.
// Returns nil when not found
func (vms VerificationMethods) FindByID(id string) *VerificationMethod {
	for _, vm := range vms.Data {
		if vm.ID == id {
			return vm
		}
	}
	return nil
}

// Remove removes a VerificationMethod from the slice.
// If a verificationMethod was removed with the given DID, it will be returned
func (vms *VerificationMethods) Remove(id string) *VerificationMethod {
	var (
		filteredVMS []*VerificationMethod
		foundVM     *VerificationMethod
	)
	for _, vm := range vms.Data {
		if vm.ID != id {
			filteredVMS = append(filteredVMS, vm)
		} else {
			foundVM = vm
		}
	}
	vms.Data = filteredVMS
	return foundVM
}

// Add adds a verificationMethod to the verificationMethods if it not already present.
func (vms *VerificationMethods) Add(v *VerificationMethod) {
	for _, ptr := range vms.Data {
		// check if the pointer is already in the list
		if ptr == v {
			return
		}
		// check if the actual ids match?
		if ptr.ID == v.ID {
			return
		}
	}
	vms.Data = append(vms.Data, v)
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
	var (
		validKey         bool
		validPrefix      bool
		containsMetadata bool
	)
	validKey = vm.Type == KeyType_KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019
	wkp := NewWalletPrefix(vm.BlockchainAccountId)
	validPrefix = wkp != ChainWalletPrefixNone
	_, containsMetadata = vm.Metadata["blockchain"]
	return validKey && validPrefix && containsMetadata
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
