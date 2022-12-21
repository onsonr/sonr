package wallet

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"

	"github.com/libp2p/go-libp2p/core/crypto"
	crypto_pb "github.com/libp2p/go-libp2p/core/crypto/pb"

	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
)

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
