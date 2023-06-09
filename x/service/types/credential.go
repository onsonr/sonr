package types

import (
	"crypto/ed25519"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/sonrhq/core/internal/crypto"
	idtypes "github.com/sonrhq/core/x/identity/types"
	"github.com/yoseplee/vrf"
)

func LoadCredentialFromVerificationMethod(vm *idtypes.VerificationMethod) (*WebauthnCredential, error) {
	if vm.Metadata == "" {
		return nil, errors.New("no credential metadata")
	}
	cred := &WebauthnCredential{}
	err := json.Unmarshal([]byte(vm.Metadata), cred)
	if err != nil {
		return nil, err
	}
	return cred, nil
}

// Serialize the credential to JSON
func (c *WebauthnCredential) Serialize() ([]byte, error) {
	return json.Marshal(c)
}

// Descriptor returns the credential's descriptor
func (c *WebauthnCredential) CredentialDescriptor() protocol.CredentialDescriptor {
	transport := make([]protocol.AuthenticatorTransport, 0)
	for _, t := range c.Transport {
		transport = append(transport, protocol.AuthenticatorTransport(t))
	}

	return protocol.CredentialDescriptor{
		CredentialID:    protocol.URLEncodedBase64(c.Id),
		Type:            protocol.PublicKeyCredentialType,
		Transport:       transport,
		AttestationType: c.AttestationType,
	}
}

func (c *WebauthnCredential) GetWebauthnCredential() *WebauthnCredential {
	return c
}

// ToVerificationMethod converts the credential to a DID VerificationMethod
func (c *WebauthnCredential) ToVerificationMethod() *idtypes.VerificationMethod {
	vm := &idtypes.VerificationMethod{
		Id:                 fmt.Sprintf("did:%s:%s", idtypes.WEBAUTHN_DID_METHOD, protocol.URLEncodedBase64(c.Id).String()),
		Type:               "webauthn/alg-es256",
		PublicKeyMultibase: crypto.Base64Encode(c.PublicKey),
		Controller:         c.Controller,
	}
	jsonCred, err := json.Marshal(c)
	if err != nil {
		return vm
	}
	vm.Metadata = string(jsonCred)
	return vm
}

// Did returns the credential's DID
func (c *WebauthnCredential) Did() string {
	return c.ToVerificationMethod().Id
}


func computeVRF(secretKey ed25519.PrivateKey, message []byte) ([]byte, []byte, error) {
	publicKey, _ := secretKey.Public().(ed25519.PublicKey)

	// generate proof and hash using the Prove function
	proof, hash, err := vrf.Prove(publicKey, secretKey, message)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate VRF proof: %v", err)
	}

	// verify the proof using the Verify function
	ok, err := vrf.Verify(publicKey, proof, message)
	if err != nil || !ok {
		return nil, nil, fmt.Errorf("failed to verify VRF proof: %v", err)
	}

	return proof, hash, nil
}
