package sfs

import (
	"bytes"
	crypto_rand "crypto/rand"
	"encoding/base64"
	fmt "fmt"
	"io"

	"golang.org/x/crypto/nacl/box"
)

type mbKey struct {
	sk   []byte
	pub  *[32]byte
	priv *[32]byte
}

func NewMailboxKey(secretKey []byte) (*mbKey, error) {
	if len(secretKey) != 32 {
		return nil, fmt.Errorf("invalid secret key length: %d, needs to be 32", len(secretKey))
	}
	pub, priv, err := box.GenerateKey(bytes.NewReader(secretKey))
	if err != nil {
		return nil, err
	}
	return &mbKey{
		sk:   secretKey,
		pub:  pub,
		priv: priv,
	}, nil
}

// Seal message with the public key of the recipient.
func (k *mbKey) SealMessage(message []byte, recipient []byte) ([]byte, error) {
	peersPublicKey := bytesToPointer(recipient)
	var nonce [24]byte
	if _, err := io.ReadFull(crypto_rand.Reader, nonce[:]); err != nil {
		return nil, err
	}
	// This encrypts msg and appends the result to the nonce.
	encrypted := box.Seal(nonce[:], message, &nonce, peersPublicKey, k.priv)
	return encrypted, nil
}

func (k *mbKey) PublicKeyBase64() string {
	return base64.StdEncoding.EncodeToString(k.NormalizePublicKey())
}

func (k *mbKey) Type() string {
	return "mailbox"
}


func (k *mbKey) NormalizePrivateKey() []byte {
	return k.priv[:]
}

func (k *mbKey) NormalizePublicKey() []byte {
	return k.pub[:]
}

func bytesToPointer(b []byte) *[32]byte {
	var a [32]byte
	copy(a[:], b)
	return &a
}
