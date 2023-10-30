package zk

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"

	"github.com/sonr-io/kryptology/pkg/accumulator"
	"github.com/sonr-io/kryptology/pkg/core/curves"
	"github.com/sonrhq/core/internal/crypto"
)

type Accumulator = accumulator.Accumulator
type SecretKey = accumulator.SecretKey
type Element = accumulator.Element

func (zk ZKEphemeralKey) String() string {
	return crypto.Base64Encode(zk)
}

type ZKSet string

var emptyZKSet = ZKSet("")

// String returns the string representation of the ZKSet
func (s ZKSet) String() string {
	return string(s)
}

// GetAccumulator returns the accumulator for the ZKSet
func (s ZKSet) GetAccumulator() *ZKAccumulator {
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	accBz, err := crypto.Base64Decode(s.String())
	if err != nil {
		panic(err)
	}
	e, err := new(ZKAccumulator).New(curve)
	if err != nil {
		panic(err)
	}
	err = e.UnmarshalBinary(accBz)
	if err != nil {
		panic(err)
	}
	return e
}

// SetAccumulator sets the accumulator for the ZKSet
func (s ZKSet) SetAccumulator(a *ZKAccumulator) error {
	bz, err := a.MarshalBinary()
	if err != nil {
		return err
	}
	str := crypto.Base64Encode(bz)
	s = ZKSet(str)
	return nil
}

// CreateZkSet creates a new ZKSet from a public key and a list of initial ids
func CreateZkSet(publicKey PubKey, initialIds ...string) (ZKSet, error) {
	if publicKey == nil {
		return emptyZKSet, errors.New("public key cannot be nil")
	}
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	key, err := GetSecretFromSecp256k1(publicKey)
	if err != nil {
		return emptyZKSet, err
	}
	acc, err := new(ZKAccumulator).WithElements(curve, key, StringListToZkElements(initialIds...))
	if err != nil {
		return emptyZKSet, err
	}
	bz, err := acc.MarshalBinary()
	if err != nil {
		return emptyZKSet, err
	}
	str := crypto.Base64Encode(bz)
	return ZKSet(str), nil
}

// OpenZkSet takes a plain string and returns a ZkSet on success
func OpenZkSet(str string) (ZKSet, error) {
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	accBz, err := crypto.Base64Decode(str)
	if err != nil {
		return emptyZKSet, err
	}
	e, err := new(ZKAccumulator).New(curve)
	if err != nil {
		return emptyZKSet, err
	}
	err = e.UnmarshalBinary(accBz)
	if err != nil {
		return emptyZKSet, err
	}
	return ZKSet(str), nil
}

// AddElement adds an element to the accumulator
func (c ZKSet) AddElement(pk PubKey, elem string) error {
	if pk == nil {
		return errors.New("no secret key provided")
	}
	key, err := GetSecretFromSecp256k1(pk)
	if err != nil {
		return err
	}
	newAcc, err := c.GetAccumulator().Add(key, StringToZkElement(elem))
	if err != nil {
		return err
	}
	err = c.SetAccumulator(newAcc)
	if err != nil {
		return err
	}
	return nil
}

// RemoveElement removes an element from the accumulator
func (c ZKSet) RemoveElement(pk PubKey, elem string) error {
	if pk == nil {
		return errors.New("no secret key provided")
	}
	key, err := GetSecretFromSecp256k1(pk)
	if err != nil {
		return err
	}
	newAcc, err := c.GetAccumulator().Remove(key, StringToZkElement(elem))
	if err != nil {
		return err
	}
	c.SetAccumulator(newAcc)
	return nil
}

// Encrypt encrypts a message with the ZKSet
func (c ZKSet) Encrypt(key PubKey, bz []byte) ([]byte, error) {
	if key == nil {
		return nil, errors.New("no secret key provided")
	}
	// Derive AES key from public key
	aesKey, err := GetEncryptionKeyFromSecp256k1(key)
	if err != nil {
		return nil, err
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

// Decrypt decrypts a ciphertext using the public key
func (c ZKSet) Decrypt(key PubKey, ciphertext []byte) ([]byte, error) {
	if key == nil {
		return nil, errors.New("no secret key provided")
	}
	// Derive AES key from public key
	aesKey, err := GetEncryptionKeyFromSecp256k1(key)
	if err != nil {
		return nil, err
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
		return nil, errors.New("ciphertext too short")
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	ksbz, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return ksbz, nil
}

// ValidateMembership validates that an element is a member of the set
func (c ZKSet) ValidateMembership(spk PubKey, elem string) bool {
	if spk == nil {
		return false
	}
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	key, err := GetSecretFromSecp256k1(spk)
	if err != nil {
		return false
	}
	wit, err := new(accumulator.MembershipWitness).New(curve.Scalar.Hash([]byte(elem)), c.GetAccumulator(), key)
	if err != nil {
		return false
	}

	pk, err := key.GetPublicKey(curve)
	if err != nil {
		return false
	}
	err = wit.Verify(pk, c.GetAccumulator())
	if err != nil {
		return false
	}
	return true
}

// NewZKSet creates a new zero-knowledge set from a list of zero-knowledge proofs.
func NewZKSet(pubKey *crypto.Secp256k1PubKey, initialIds ...string) (ZKSet, error) {
	return CreateZkSet(pubKey, initialIds...)
}

// LoadZKSet loads a zero-knowledge set from a list of zero-knowledge proofs.
func LoadZKSet(str string) (ZKSet, error) {
	return OpenZkSet(str)
}
