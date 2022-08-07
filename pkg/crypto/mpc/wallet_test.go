package mpc

import (
	"encoding/base64"
	"fmt"
	"os"
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

	encodedcnf := base64.StdEncoding.EncodeToString(ciphercnf)

	/*
	 * Deserialize
	 */

	decodedcnf, err := base64.StdEncoding.DecodeString(encodedcnf)
	assert.NoError(t, err, "decodes from base64")

	cnf := cmp.EmptyConfig(curve.Secp256k1{})
	plaincnfbin, err := AesDecryptWithPassword("password123", decodedcnf)
	assert.NoError(t, err, "decrypts successfully")

	err = cnf.UnmarshalBinary(plaincnfbin)
	assert.NoError(t, err, "unmarshals binary")

	fmt.Printf("%+v\n", cnf)
	fmt.Printf("%+v\n", w.Config())
}

func Test_DecryptRecoveryShardFromFile(t *testing.T) {
	encodedShard, err := os.ReadFile("recovery.shard")
	if err != nil {
		fmt.Printf("didn't read: %s\n", err)
	}
	fmt.Printf("len: %d\n", len(encodedShard))

	//decodedcnf, err := base64.StdEncoding.DecodeString(string(encodedShard))
	decodedcnf := make([]byte, len(encodedShard))
	_, err = base64.StdEncoding.Decode(decodedcnf, encodedShard)
	if err != nil {
		fmt.Printf("err: %s\n", err)
	}
	assert.NoError(t, err, "decodes from base64")

	cnf := cmp.EmptyConfig(curve.Secp256k1{})
	plaincnfbin, err := AesDecryptWithPassword("password123", decodedcnf)
	assert.NoError(t, err, "decrypts successfully")

	err = cnf.UnmarshalBinary(plaincnfbin)
	assert.NoError(t, err, "unmarshals binary")

	fmt.Printf("%+v\n", cnf)
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
