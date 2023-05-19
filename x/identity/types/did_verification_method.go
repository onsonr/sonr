package types

import (
	fmt "fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/shengdoushi/base58"
	"github.com/sonrhq/core/internal/crypto"
)

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

// Equal returns true if the verification method is equal to the given verification method
func (d *VerificationMethod) Equal(other *VerificationMethod) bool {
	return d.Id == other.Id &&
		d.Type == other.Type &&
		d.BlockchainAccountId == other.BlockchainAccountId &&
		d.PublicKeyMultibase == other.PublicKeyMultibase &&
		d.Metadata == other.Metadata
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
	credId, err := crypto.Base64Decode(id)
	if err != nil {
		return nil
	}
	return protocol.URLEncodedBase64(credId)
}
