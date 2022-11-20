package mpc

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"

	"github.com/libp2p/go-libp2p/core/crypto"
	crypto_pb "github.com/libp2p/go-libp2p/core/crypto/pb"
	"github.com/sonr-io/multi-party-sig/pkg/ecdsa"
	"github.com/sonr-io/multi-party-sig/pkg/math/curve"
	"golang.org/x/crypto/scrypt"
)

func (w *Wallet) Equals(other crypto.Key) bool {
	wbz, err := w.Raw()
	if err != nil {
		return false
	}
	obz, err := other.Raw()
	if err != nil {
		return false
	}
	return string(wbz) == string(obz)
}

// Returns the Secp256k1 public key of the given party.
func (w *Wallet) PublicKey() ([]byte, error) {
	p := w.PublicPoint().(*curve.Secp256k1Point)
	buf, err := p.MarshalBinary()
	if err != nil {
		return nil, err
	}
	// Check length of the public key.
	if len(buf) != 33 {
		return nil, fmt.Errorf("invalid public key length")
	}
	return buf, nil
}

func (w *Wallet) GetPublic() crypto.PubKey {
	return makePublicKey(w.PublicPoint())
}

type PublicKey struct {
	// contains filtered or unexported fields
	p curve.Point
}

func makePublicKey(p curve.Point) *PublicKey {
	return &PublicKey{p: p}
}

func (p *PublicKey) Equals(k crypto.Key) bool {
	pbz, err := p.Raw()
	if err != nil {
		return false
	}

	kbz, err := k.Raw()
	if err != nil {
		return false
	}
	return string(pbz) == string(kbz)
}

func (p *PublicKey) Raw() ([]byte, error) {
	return p.p.MarshalBinary()
}

func (p *PublicKey) Type() crypto_pb.KeyType {
	return crypto_pb.KeyType_Secp256k1
}

func (w *PublicKey) Verify(data []byte, sig []byte) (bool, error) {
	signature, err := DeserializeSignature(sig)
	if err != nil {
		return false, err
	}
	return signature.Verify(w.p, data), nil
}

// AesEncryptWithKey uses the give 32-bit key to encrypt plaintext.
func AesEncryptWithKey(aesKey, plaintext []byte) ([]byte, error) {
	if len(aesKey) != 32 {
		return nil, errors.New("AES key must be 32 bytes")
	}

	blockCipher, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

	return ciphertext, nil
}

// AesDecryptWithKey uses the give 32-bit key to decrypt plaintext.
func AesDecryptWithKey(aesKey, ciphertext []byte) ([]byte, error) {
	if len(aesKey) != 32 {
		fmt.Printf("aesKey len: %d\n", len(aesKey))
		return nil, errors.New("AES key must be 32 bytes")
	}

	blockCipher, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}

	nonce, ct := ciphertext[:gcm.NonceSize()], ciphertext[gcm.NonceSize():]

	plaintext, err := gcm.Open(nil, nonce, ct, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// AesEncryptWithPassword uses the give password to generate an aes key and decrypt plaintext.
func AesEncryptWithPassword(password string, plaintext []byte) ([]byte, error) {
	key, err := deriveKey(password)
	if err != nil {
		return nil, err
	}

	return AesEncryptWithKey(key, plaintext)
}

// AesDecryptWithPassword uses the give password to generate an aes key and encrypt plaintext.
func AesDecryptWithPassword(password string, ciphertext []byte) ([]byte, error) {
	key, err := deriveKey(password)
	if err != nil {
		return nil, err
	}

	return AesDecryptWithKey(key, ciphertext)
}

func deriveKey(password string) ([]byte, error) {
	// including a salt would make it impossible to reliably login from other devices
	key, err := scrypt.Key([]byte(password), []byte(""), 1<<20, 8, 1, 32)
	if err != nil {
		return nil, err
	}

	return key, nil
}

// NewAesKey generates a new 32-bit key.
func NewAesKey() ([]byte, error) {
	key := make([]byte, 32)
	if n, err := rand.Read(key); err != nil {
		return nil, err
	} else if n != 32 {
		return nil, errors.New("could not create key at 32 bytes")
	}

	return key, nil
}

// SerializeSignature marshals an ECDSA signature to DER format for use with the CMP protocol
func SerializeSignature(sig *ecdsa.Signature) ([]byte, error) {
	rBytes, err := sig.R.MarshalBinary()
	if err != nil {
		return nil, err
	}
	sBytes, err := sig.S.MarshalBinary()
	if err != nil {
		return nil, err
	}

	sigBytes := make([]byte, 65)
	// 0 pad the byte arrays from the left if they aren't big enough.
	copy(sigBytes[33-len(rBytes):33], rBytes)
	copy(sigBytes[65-len(sBytes):65], sBytes)
	return sigBytes, nil
}

// - The R and S values must be in the valid range for secp256k1 scalars:
//   - Negative values are rejected
//   - Zero is rejected
//   - Values greater than or equal to the secp256k1 group order are rejected
func DeserializeSignature(sigStr []byte) (*ecdsa.Signature, error) {
	rBytes := sigStr[:33]
	sBytes := sigStr[33:65]

	sig := ecdsa.EmptySignature(curve.Secp256k1{})
	if err := sig.R.UnmarshalBinary(rBytes); err != nil {
		return nil, errors.New("malformed signature: R is not in the range [1, N-1]")
	}

	// S must be in the range [1, N-1].  Notice the check for the maximum number
	// of bytes is required because SetByteSlice truncates as noted in its
	// comment so it could otherwise fail to detect the overflow.
	if err := sig.S.UnmarshalBinary(sBytes); err != nil {
		return nil, errors.New("malformed signature: S is not in the range [1, N-1]")
	}

	// Create and return the signature.
	return &sig, nil
}
