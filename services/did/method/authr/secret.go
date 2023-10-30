package authr

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"

	"github.com/google/uuid"
	"github.com/sonr-io/core/internal/crypto"
	"github.com/sonr-io/kryptology/pkg/accumulator"
	"github.com/sonr-io/kryptology/pkg/core/curves"
	"golang.org/x/crypto/hkdf"
	"lukechampine.com/blake3"
)

// AuthSecretKey is a secret key for an Authenticator
type AuthSecretKey []byte

// NewSecretKey creates a new secret key for an Authenticator
func NewSecretKey() (AuthSecretKey, error) {
	randPass := uuid.New().String()
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	sk, err := new(accumulator.SecretKey).New(curve, []byte(randPass))
	if err != nil {
		return nil, err
	}
	bz, err := sk.MarshalBinary()
	if err != nil {
		return nil, err
	}
	return AuthSecretKey(bz), nil
}

// AccumulatorKey returns the accumulator key
func (sk AuthSecretKey) AccumulatorKey() (*accumulator.SecretKey, error) {
	e := &accumulator.SecretKey{}
	err := e.UnmarshalBinary(sk)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// Bytes returns the bytes of the secret key
func (sk AuthSecretKey) Bytes() []byte {
	return []byte(sk)
}

// Encrypt encrypts a byte slice using AES-GCM
func (sk AuthSecretKey) Encrypt(bz []byte) ([]byte, error) {
	// Use the HKDF extractor to derive a 32-byte seed
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
	ciphertext := gcm.Seal(nonce, nonce, bz, nil)
	return ciphertext, nil
}

// Decrypt decrypts a byte slice using AES-GCM
func (sk AuthSecretKey) Decrypt(ciphertext []byte) ([]byte, error) {
	// Use the HKDF extractor to derive a 32-byte seed
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

// FormatIdentifier formats a public key into a DID
func FormatIdentifier(pub []byte) string {
	hex := crypto.Base64Encode(pub)
	return fmt.Sprintf("did:authr:%s", hex)
}

func emailKey(email string) string {
	hash := blake3.Sum256([]byte(email))
	return fmt.Sprintf("email/%s", hash)
}

func credentialKey(origin string, id []byte) string {
	return fmt.Sprintf("credentials/%s/%s", origin, crypto.Base64Encode(id))
}
