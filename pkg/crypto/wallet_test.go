package crypto

import (
	"fmt"
	"testing"

	"github.com/sonr-io/multi-party-sig/pkg/math/curve"
	"github.com/sonr-io/multi-party-sig/protocols/cmp"
	"github.com/stretchr/testify/assert"
)

func Test_MPCCreate(t *testing.T) {
	w, err := GenerateWallet()
	assert.NoError(t, err, "wallet generation succeeds")
	_, err = w.PublicKey()
	assert.NoError(t, err, "public key creation succeeds")
}

func Test_EncryptConfig(t *testing.T) {
  w, err := GenerateWallet()
  assert.NoError(t, err, "wallet generation succeeds")

  cnfbin, err := w.Config().MarshalBinary()
  assert.NoError(t, err, "marshals binary")

  ciphercnf, err := AesEncryptWithPassword("password123", cnfbin)
  assert.NoError(t, err, "encrypts successfully")

  cnf := cmp.EmptyConfig(curve.Secp256k1{})
  plaincnfbin, err := AesDecryptWithPassword("password123", ciphercnf)
  assert.NoError(t, err, "decrypts successfully")

  err = cnf.UnmarshalBinary(plaincnfbin)
  assert.NoError(t, err, "unmarshals binary")

  fmt.Printf("%+v\n", cnf)
  fmt.Printf("%+v\n", w.Config())
}

func Test_MPCDID(t *testing.T) {
	w, err := GenerateWallet()
	assert.NoError(t, err, "wallet generation succeeds")

	_, err = w.Address()
	assert.NoError(t, err, "Bech32Address successfully created")
}

func Test_MPCSignMessage(t *testing.T) {
	m := []byte("sign this message")
	w, err := GenerateWallet()
	assert.NoError(t, err, "wallet generation succeeds")

	sig, err := w.Sign(m)
	assert.NoError(t, err, "signing succeeds")
	deserializedSigVerified := sig.Verify(w.Config().PublicPoint(), m)
	assert.True(t, deserializedSigVerified, "deserialized signature is verified")
}
