package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MPCCreate(t *testing.T) {
	w, err := GenerateWallet()
	assert.NoError(t, err, "wallet generation succeeds")
	_, err = w.PublicKey()
	assert.NoError(t, err, "public key creation succeeds")
}

func Test_MPCDID(t *testing.T) {
	w, err := GenerateWallet()
	assert.NoError(t, err, "wallet generation succeeds")

	_, err = w.Address()
	assert.NoError(t, err, "Bech32Address successfully created")

	_, err = w.DIDDocument()
	assert.NoError(t, err, "DID Document creation succeeds")
}

func Test_MPCSignMessage(t *testing.T) {
	m := []byte("sign this message")
	w, err := GenerateWallet()
	assert.NoError(t, err, "wallet generation succeeds")

	sig, err := w.Sign(m)
	assert.NoError(t, err, "signing succeeds")

	sigRaw, err := SerializeSignature(sig)
	assert.NoError(t, err, "signature serialization succeeds")

	rawIsVerified := w.Verify(m, sigRaw)
	assert.True(t, rawIsVerified, "raw signature is verified")

	sig2, err := SignatureFromBytes(sigRaw)
	assert.NoError(t, err, "deserialization succeeds")
	deserializedSigVerified := sig2.Verify(w.Config().PublicPoint(), m)
	assert.True(t, deserializedSigVerified, "deserialized signature is verified")
}
