package types

import (
	"encoding/base64"
	fmt "fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/shengdoushi/base58"
	"github.com/sonrhq/core/internal/crypto"
)

type DIDParseResult struct {
	AccountName string
	Address     string
	CoinType    crypto.CoinType
}

// NewSonrID creates a new DID URI for the given Sonr Account address
func NewSonrID(addr string) string {
	return fmt.Sprintf("did:sonr:%s", addr)
}

// NewWebID creates a new DID URI for the given Sonr Account address
func NewWebID(addr string) string {
	return DIDMethod_DIDMethod_WEB.Format(addr)
}

// NewKeyID creates a new DID URI for the given Sonr Account address
func NewKeyID(addr string, keyName string) string {
	return DIDMethod_DIDMethod_KEY.Format(addr, WithFragment(keyName))
}

// NewIpfsID creates a new DID URI for the given Content ID
func NewIpfsID(addr string) string {
	return DIDMethod_DIDMethod_IPFS.Format(addr)
}

// NewPeerID creates a new DID URI for the given Peer ID
func NewPeerID(addr string) string {
	return DIDMethod_DIDMethod_PEER.Format(addr)
}

// Format returns a string representation of the DIDMethod that is on the DID spec
func (m DIDMethod) Format(val string, options ...FormatOption) string {
	r := fmt.Sprintf("did:%s:%s", m.PrettyString(), val)
	for _, opt := range options {
		r = opt(r)
	}
	return r
}

// PrettyString returns a string representation of the DIDMethod that is on the DID spec
func (m DIDMethod) PrettyString() string {
	prts := strings.Split(m.String(), "_")
	return strings.ToLower(prts[len(prts)-1])
}

// FormatOption is a function that can be used to format a DIDMethod
type FormatOption func(string) string

// WithFragment returns a FormatOption that will append a fragment to the DID
func WithFragment(frag string) FormatOption {
	return func(did string) string {
		return fmt.Sprintf("%s#%s", did, frag)
	}
}

// WithPath returns a FormatOption that will append a path to the DID
func WithPath(path string) FormatOption {
	return func(did string) string {
		return fmt.Sprintf("%s/%s", did, path)
	}
}

// WithQuery returns a FormatOption that will append a query to the DID
func WithQuery(query string) FormatOption {
	return func(did string) string {
		return fmt.Sprintf("%s?%s", did, query)
	}
}

///
/// Helper functions
///

// findCoinTypeFromAddress returns the CoinType for the given address
func findCoinTypeFromAddress(addr string) crypto.CoinType {
	for _, ct := range crypto.AllCoinTypes() {
		if strings.Contains(addr, ct.AddrPrefix()) {
			return ct
		}
	}
	return crypto.TestCoinType
}

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

// VerificationMethod applies the given options and builds a verification method from this Key
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

// Method returns the DID method of the document
func (d *VerificationMethod) DIDMethod() string {
	return strings.Split(d.Id, ":")[1]
}

// Identifier returns the DID identifier of the document
func (d *VerificationMethod) DIDIdentifier() string {
	return strings.Split(d.Id, ":")[2]
}

// Fragment returns the DID fragment of the document
func (d *VerificationMethod) DIDFragment() string {
	return strings.Split(d.Id, "#")[1]
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

func (d *VerificationMethod) WebauthnCredentialID() protocol.URLEncodedBase64 {
	ptrs := strings.Split(d.Id, ":")
	id := ptrs[len(ptrs)-1]
	// Decode the credential id
	credId, err := base64.RawURLEncoding.DecodeString(id)
	if err != nil {
		return nil
	}
	return protocol.URLEncodedBase64(credId)
}
