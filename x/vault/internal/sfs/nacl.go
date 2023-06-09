package sfs

import (
	"bytes"
	"encoding/base64"
	fmt "fmt"

	"golang.org/x/crypto/nacl/box"
)

type mbKey struct {
	sk   []byte
	pub  []byte
	priv []byte
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
		pub:  pub[:],
		priv: priv[:],
	}, nil
}

func (k *mbKey) PublicKey() string {
	return base64.StdEncoding.EncodeToString(k.pub)
}

func (k *mbKey) Type() string {
	return "mailbox"
}
