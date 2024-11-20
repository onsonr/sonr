package macaroon

import (
	"crypto/rand"
	"fmt"
	"testing"

	qt "github.com/frankban/quicktest"
	"golang.org/x/crypto/nacl/secretbox"
)

var testCryptKey = &[hashLen]byte{'k', 'e', 'y'}
var testCryptText = &[hashLen]byte{'t', 'e', 'x', 't'}

func TestEncDec(t *testing.T) {
	c := qt.New(t)
	b, err := encrypt(testCryptKey, testCryptText, rand.Reader)
	c.Assert(err, qt.Equals, nil)
	p, err := decrypt(testCryptKey, b)
	c.Assert(err, qt.Equals, nil)
	c.Assert(string(p[:]), qt.Equals, string(testCryptText[:]))
}

func TestUniqueNonces(t *testing.T) {
	c := qt.New(t)
	nonces := make(map[string]struct{})
	for i := 0; i < 100; i++ {
		nonce, err := newNonce(rand.Reader)
		c.Assert(err, qt.Equals, nil)
		nonces[string(nonce[:])] = struct{}{}
	}
	c.Assert(nonces, qt.HasLen, 100, qt.Commentf("duplicate nonce detected"))
}

type ErrorReader struct{}

func (*ErrorReader) Read([]byte) (int, error) {
	return 0, fmt.Errorf("fail")
}

func TestBadRandom(t *testing.T) {
	c := qt.New(t)
	_, err := newNonce(&ErrorReader{})
	c.Assert(err, qt.ErrorMatches, "^cannot generate random bytes:.*")

	_, err = encrypt(testCryptKey, testCryptText, &ErrorReader{})
	c.Assert(err, qt.ErrorMatches, "^cannot generate random bytes:.*")
}

func TestBadCiphertext(t *testing.T) {
	c := qt.New(t)
	buf := randomBytes(nonceLen + secretbox.Overhead)
	for i := range buf {
		_, err := decrypt(testCryptKey, buf[0:i])
		c.Assert(err, qt.ErrorMatches, "message too short")
	}
	_, err := decrypt(testCryptKey, buf)
	c.Assert(err, qt.ErrorMatches, "decryption failure")
}

func randomBytes(n int) []byte {
	buf := make([]byte, n)
	if _, err := rand.Reader.Read(buf); err != nil {
		panic(err)
	}
	return buf
}
