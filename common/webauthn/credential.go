package webauthn

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
	"github.com/yoseplee/vrf"

	"github.com/sonrhq/sonr/common/crypto"
)

// Serialize the credential to JSON
func (c *Credential) Serialize() ([]byte, error) {
	return json.Marshal(c)
}

// Deserialize the credential from JSON
func (c *Credential) Deserialize(data []byte) error {
	return json.Unmarshal(data, c)
}

func (c *Credential) DID() string {
	did := fmt.Sprintf("did:%s:%s", "webauthn", crypto.Base64Encode(c.Id))
	return did
}

// Encrypt is used to encrypt a message for the credential
func (c *Credential) Encrypt(data []byte) ([]byte, error) {
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
	gcm, err := cipher.NewGCMWithNonceSize(block, aes.BlockSize)
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
func (c *Credential) Decrypt(data []byte) ([]byte, error) {
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


// ShortID returns the first 8 characters of the base58 encoded credential id
func (c *Credential) ShortID() string {
	return crypto.Base58Encode(c.Id)[0:8]
}

// ToCredentialDescriptor converts a VerificationMethod to a CredentialDescriptor if the VerificationMethod uses the `did:webauthn` method
func (vm *Credential) GetDescriptor() protocol.CredentialDescriptor {
	transport := make([]protocol.AuthenticatorTransport, 0)
	for _, t := range vm.Transport {
		transport = append(transport, protocol.AuthenticatorTransport(t))
	}

	return protocol.CredentialDescriptor{
		CredentialID:    protocol.URLEncodedBase64(vm.Id),
		Type:            protocol.PublicKeyCredentialType,
		Transport:       transport,
		AttestationType: vm.AttestationType,
	}
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
