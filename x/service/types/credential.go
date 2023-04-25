package types

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/protocol/webauthncose"
	"github.com/sonrhq/core/internal/crypto"
	idtypes "github.com/sonrhq/core/x/identity/types"
)

type Credential interface {
	// Controller returns the credential's controller
	GetController() string

	// Get the credential's DID
	Did() string

	// Descriptor returns the credential's descriptor
	Descriptor() protocol.CredentialDescriptor

	// GetWebauthnCredential returns the webauthn credential instance
	GetWebauthnCredential() *WebauthnCredential

	// Convert the credential to a DID VerificationMethod
	ToVerificationMethod() *idtypes.VerificationMethod

	// Encrypt is used to encrypt a message for the credential
	Encrypt(msg []byte) ([]byte, error)

	// Decrypt is used to decrypt a message for the credential
	Decrypt(msg []byte) ([]byte, error)

	// Marshal is used to marshal the credential to JSON
	Marshal() ([]byte, error)
}

type DidCredential struct {
	*WebauthnCredential `json:"credential,omitempty"`
	Controller             string `json:"controller,omitempty"`
}

func NewCredential(cred *WebauthnCredential, controller string) Credential {
	return &DidCredential{
		WebauthnCredential: cred,
		Controller:            controller,
	}
}

func LoadCredential(didCred *DidCredential) (Credential, error) {
	if didCred.WebauthnCredential == nil {
		return nil, errors.New("invalid credential")
	}

	return &DidCredential{
		WebauthnCredential: didCred.WebauthnCredential,
		Controller:            didCred.Controller,
	}, nil
}

func (c *DidCredential) GetController() string {
	return c.Controller
}

// Descriptor returns the credential's descriptor
func (c *DidCredential) Descriptor() protocol.CredentialDescriptor {
	return c.WebauthnCredential.ToStdCredential().Descriptor()
}

func (c *DidCredential) GetWebauthnCredential() *WebauthnCredential {
	return c.WebauthnCredential
}

// MarshalJSON is used to marshal the credential to JSON
func (c *DidCredential) Marshal() ([]byte, error) {
	vm := c.WebauthnCredential
	bz, err := json.Marshal(vm)
	if err != nil {
		return nil, err
	}
	return bz, nil
}

// ToVerificationMethod converts the credential to a DID VerificationMethod
func (c *DidCredential) ToVerificationMethod() *idtypes.VerificationMethod {
	return &idtypes.VerificationMethod{
		Id:                 fmt.Sprintf("did:key:%s", crypto.Base58Encode(c.WebauthnCredential.Id)),
		Type:               "webauthn/alg-es256",
		PublicKeyMultibase: crypto.Base58Encode(c.WebauthnCredential.PublicKey),
		Controller:         c.Controller,
	}
}

// Did returns the credential's DID
func (c *DidCredential) Did() string {
	return c.ToVerificationMethod().Id
}

// Encrypt is used to encrypt a message for the credential
func (c *DidCredential) Encrypt(data []byte) ([]byte, error) {
	// Get the public key from the credential
	keyFace, err := webauthncose.ParsePublicKey(c.WebauthnCredential.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}
	publicKey, ok := keyFace.(webauthncose.EC2PublicKeyData)
	if !ok {
		return nil, errors.New("public key is not an EC2 key")
	}
	// Derive a shared secret using ECDH
	privateKey, err := derivePrivateKey(c.WebauthnCredential)
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
func (c *DidCredential) Decrypt(data []byte) ([]byte, error) {
	// Get the public key from the credential
	keyFace, err := webauthncose.ParsePublicKey(c.WebauthnCredential.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}
	publicKey, ok := keyFace.(webauthncose.EC2PublicKeyData)
	if !ok {
		return nil, errors.New("public key is not an EC2 key")
	}
	// Derive a shared secret using ECDH
	privateKey, err := derivePrivateKey(c.WebauthnCredential)
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
	return NewCredential(credential, controller), nil
}
