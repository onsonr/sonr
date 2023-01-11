// Utility functions for DID Capability Delegation - https://w3c.github.io/did-core/#capability-delegation
// I.e. Verification Material for IPFS Node which stores MPC Configurations
package types

import (
	fmt "fmt"

	crypto_pb "github.com/libp2p/go-libp2p/core/crypto/pb"
	"github.com/libp2p/go-libp2p/core/peer"
)

// AddCapabilityDelegation adds a VerificationMethod as CapabilityDelegation
// If the controller is not set, it will be set to the document's ID
func (d *DidDocument) AddCapabilityDelegation(v *VerificationMethod) {
	if v.Controller == "" {
		v.Controller = d.ID
	}
	d.VerificationMethod.Add(v)
	d.CapabilityDelegation.Add(v)
}

// FindCapabilityDelegation finds a VerificationMethod by its ID
func (d *DidDocument) FindCapabilityDelegation(id string) *VerificationMethod {
	return d.CapabilityDelegation.FindByID(id)
}

// FindCapabilityDelegationByFragment finds a VerificationMethod by its fragment
func (d *DidDocument) FindCapabilityDelegationByFragment(fragment string) *VerificationMethod {
	return d.CapabilityDelegation.FindByFragment(fragment)
}

// FindCapabilityInvocation finds a VerificationMethod by its ID
func (d *DidDocument) FindCapabilityInvocation(id string) *VerificationMethod {
	return d.CapabilityInvocation.FindByID(id)
}

// FindCapabilityInvocationByFragment finds a VerificationMethod by its fragment
func (d *DidDocument) FindCapabilityInvocationByFragment(fragment string) *VerificationMethod {
	return d.CapabilityInvocation.FindByFragment(fragment)
}

// AddCapabilityInvocation adds a VerificationMethod as CapabilityInvocation
// If the controller is not set, it will be set to the document's ID
func (d *DidDocument) AddCapabilityInvocation(v *VerificationMethod) {
	if v.Controller == "" {
		v.Controller = d.ID
	}
	d.VerificationMethod.Add(v)
	d.CapabilityInvocation.Add(v)
}

// AddIPFSNode adds a VerificationMethod as CapabilityInvocation
func (d *DidDocument) AddPeerID(peerID peer.ID) error {
	vm := &VerificationMethod{
		ID: fmt.Sprintf("did:p2p:%s", peerID.String()),
	}
	pubKey, err := peerID.ExtractPublicKey()
	if err != nil {
		return err
	}
	switch pubKey.Type() {
	case crypto_pb.KeyType_RSA:
		vm.Type = KeyType_KeyType_RSA_VERIFICATION_KEY_2018
		break
	case crypto_pb.KeyType_Ed25519:
		vm.Type = KeyType_KeyType_ED25519_VERIFICATION_KEY_2018
		break
	case crypto_pb.KeyType_Secp256k1:
		vm.Type = KeyType_KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019
		break
	case crypto_pb.KeyType_ECDSA:
		vm.Type = KeyType_KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019
		break
	default:
		return fmt.Errorf("unsupported key type: %s", pubKey.Type())
	}
	bz, err := pubKey.Raw()
	if err != nil {
		return err
	}
	vm.PublicKeyMultibase = fmt.Sprintf("z%s", string(bz))
	d.VerificationMethod.Add(vm)
	d.CapabilityInvocation.Add(vm)
	return nil
}
