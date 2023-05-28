package types

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/protocol/webauthncose"
	"github.com/sonrhq/core/pkg/crypto"
	idtypes "github.com/sonrhq/core/x/identity/types"
	"github.com/yoseplee/vrf"
)

type Credential interface {
	// Get the credential's DID
	Did() string

	// Descriptor returns the credential's descriptor
	CredentialDescriptor() protocol.CredentialDescriptor

	// GetWebauthnCredential returns the webauthn credential instance
	GetWebauthnCredential() *WebauthnCredential

	// Convert the credential to a DID VerificationMethod
	ToVerificationMethod() *idtypes.VerificationMethod

	// Encrypt is used to encrypt a message for the credential
	Encrypt(msg []byte) ([]byte, error)

	// Decrypt is used to decrypt a message for the credential
	Decrypt(msg []byte) ([]byte, error)

	// Serialize the credential to JSON
	Serialize() ([]byte, error)
}

func NewCredential(cred *WebauthnCredential) Credential {
	return cred
}

func LoadCredential(didCred *WebauthnCredential) (Credential, error) {
	return didCred, nil
}

func LoadCredentialFromVerificationMethod(vm *idtypes.VerificationMethod) (Credential, error) {
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
		Id:                 fmt.Sprintf("did:key:%s", protocol.URLEncodedBase64(c.Id).String()),
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

// Encrypt is used to encrypt a message for the credential
func (c *WebauthnCredential) Encrypt(data []byte) ([]byte, error) {
	// Get the public key from the credential
	keyFace, err := webauthncose.ParsePublicKey(c.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}
	publicKey, ok := keyFace.(webauthncose.EC2PublicKeyData)
	if !ok {
		return nil, errors.New("public key is not an EC2 key")
	}
	// Derive a shared secret using ECDH
	privateKey, err := derivePrivateKey(c)
	if err != nil {
		return nil, fmt.Errorf("failed to derive private key: %w", err)
	}
	sharedSecret, err := sharedSecret(privateKey, publicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to derive shared secret: %w", err)
	}
	// Use the shared secret as the encryption key
	block, err := aes.NewCipher(sharedSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES cipher: %w", err)
	}

	// Generate a random IV
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, fmt.Errorf("failed to generate IV: %w", err)
	}

	// Encrypt the data using AES-GCM
	ciphertext := make([]byte, len(data))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM cipher: %w", err)
	}
	gcm.Seal(ciphertext[:0], iv, data, nil)

	// Encrypt the AES-GCM key using ECIES
	encryptedKey, err := eciesEncrypt(publicKey, sharedSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt key: %w", err)
	}

	// Concatenate the IV and ciphertext into a single byte slice
	result := make([]byte, len(iv)+len(ciphertext)+len(encryptedKey))
	copy(result[:len(iv)], iv)
	copy(result[len(iv):len(iv)+len(ciphertext)], ciphertext)
	copy(result[len(iv)+len(ciphertext):], encryptedKey)

	return result, nil
}

// Decrypt is used to decrypt a message for the credential
func (c *WebauthnCredential) Decrypt(data []byte) ([]byte, error) {
	// Get the public key from the credential
	keyFace, err := webauthncose.ParsePublicKey(c.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}
	publicKey, ok := keyFace.(webauthncose.EC2PublicKeyData)
	if !ok {
		return nil, errors.New("public key is not an EC2 key")
	}
	// Derive a shared secret using ECDH
	privateKey, err := derivePrivateKey(c)
	if err != nil {
		return nil, fmt.Errorf("failed to derive private key: %w", err)
	}

	// Derive the shared secret using ECDH and the WebAuthn credential
	sharedSecret, err := sharedSecret(privateKey, publicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to derive shared secret: %w", err)
	}

	// Use the shared secret as the decryption key
	block, err := aes.NewCipher(sharedSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES cipher: %w", err)
	}

	// Split the IV and ciphertext from the encrypted data
	iv := data[:aes.BlockSize]
	ciphertext := data[aes.BlockSize:]

	// Decrypt the ciphertext using AES-GCM
	plaintext := make([]byte, len(ciphertext))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM cipher: %w", err)
	}
	if _, err := gcm.Open(plaintext[:0], iv, ciphertext, nil); err != nil {
		return nil, fmt.Errorf("failed to decrypt data: %w", err)
	}
	return plaintext, nil
}

func ValidateWebauthnCredential(credential *WebauthnCredential, controller string) (Credential, error) {
	// Check for nil credential
	if credential == nil {
		return nil, errors.New("credential is nil")
	}

	// Check for nil credential id
	if credential.Id == nil {
		return nil, errors.New("credential id is nil")
	}
	return NewCredential(credential), nil
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
