package authr

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"

	"github.com/sonr-io/core/services/did/types"
	"golang.org/x/crypto/hkdf"
	"lukechampine.com/blake3"
)

type AuthData struct {
	// SecretKey is the secret key used to encrypt the data
	SecretKey []byte `json:"secret_key"`

	// Data is the encrypted data
	Data []byte `json:"data"`

	// Type is the type of data
	Type string `json:"type"`

	// Origin is the origin of the data
	Origin string `json:"origin"`
}

// NewEmailAuthData creates a new email auth data
func NewEmailAuthData(email string, key types.DIDSecretKey) (*AuthData, error) {
	hash := blake3.Sum256([]byte(email))
	enc, err := encryptDataWithBz(hash[:], key.Bytes())
	if err != nil {
		return nil, err
	}
	return &AuthData{
		SecretKey: enc,
		Data:      hash[:],
		Type:      "email",
	}, nil
}

// NewCredentialAuthData creates a new credential auth data
func NewCredentialAuthData(cred *types.Credential, origin string, key types.DIDSecretKey) (*AuthData, error) {
	bz, err := cred.Serialize()
	if err != nil {
		return nil, err
	}
	enc, err := encryptDataWithBz(bz, key.Bytes())
	if err != nil {
		return nil, err
	}
	return &AuthData{
		SecretKey: enc,
		Data:      bz,
		Type:      "credentials/" + origin,
		Origin:    origin,
	}, nil
}

// Key returns the key of the auth data
func (a *AuthData) Key() string {
	return a.Type
}

// Marshal marshals the auth data
func (a *AuthData) Marshal() ([]byte, error) {
	return json.Marshal(a)
}

// Unmarshal unmarshals the auth data
func (a *AuthData) Unmarshal(bz []byte) error {
	return json.Unmarshal(bz, a)
}

// GetSecretKey returns the secret key
func (a *AuthData) GetSecretKeyFromEmail(email string) (types.DIDSecretKey, error) {
	hash := blake3.Sum256([]byte(email))
	dec, err := decryptDataWithBz(hash[:], a.SecretKey)
	if err != nil {
		return nil, err
	}
	return AuthSecretKey(dec), nil
}

// encryptDataWithBz derives a hkdf key from bytes and encrypts the data
func encryptDataWithBz(sk []byte, data []byte) ([]byte, error) {
	hkdf := hkdf.New(sha256.New, sk, nil, nil)
	aesKey := make([]byte, 32) // 32 bytes for AES-256
	if _, err := io.ReadFull(hkdf, aesKey); err != nil {
		return nil, fmt.Errorf("error derive AES key: %v", err)
	}
	// AES-GCM encryption
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

// decryptDataWithBz derives a hkdf key from bytes and decrypts the data
func decryptDataWithBz(sk []byte, ciphertext []byte) ([]byte, error) {
	hkdf := hkdf.New(sha256.New, sk, nil, nil)
	aesKey := make([]byte, 32) // 32 bytes for AES-256
	if _, err := io.ReadFull(hkdf, aesKey); err != nil {
		return nil, fmt.Errorf("error derive AES key: %v", err)
	}
	// AES-GCM decryption
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
