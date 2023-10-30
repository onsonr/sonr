package zk_test

import (
	"testing"

	secp256k1 "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/sonr-io/kryptology/pkg/core/curves"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sonr-io/core/internal/zk"
)

func TestZkSet(t *testing.T) {
	priv := secp256k1.GenPrivKey()
	pub := priv.PubKey()
	zkset, err := zk.CreateZkSet(pub)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(zkset)
	err = zkset.AddElement(pub, "test")
	if err != nil {
		t.Fatal(err)
	}
	raw := zkset.String()
	t.Log(raw)

	zkset2, err := zk.OpenZkSet(raw)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(zkset2)
	ok1 := zkset2.ValidateMembership(pub, "test")
	assert.True(t, ok1)
}

func TestNewAccumulator100(t *testing.T) {
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	var seed [32]byte
	key, err := new(zk.SecretKey).New(curve, seed[:])
	require.NoError(t, err)
	require.NotNil(t, key)
	acc, err := new(zk.Accumulator).New(curve)
	require.NoError(t, err)
	require.NotNil(t, acc)
}

func TestNewAccumulator10K(t *testing.T) {
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	var seed [32]byte
	key, err := new(zk.SecretKey).New(curve, seed[:])
	require.NoError(t, err)
	require.NotNil(t, key)
	acc, err := new(zk.Accumulator).New(curve)
	require.NoError(t, err)
	require.NotNil(t, acc)
}

func TestNewAccumulator10M(t *testing.T) {
	// Initiating 10M values takes time
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	var seed [32]byte
	key, err := new(zk.SecretKey).New(curve, seed[:])
	require.NoError(t, err)
	require.NotNil(t, key)
	acc, err := new(zk.Accumulator).New(curve)
	require.NoError(t, err)
	require.NotNil(t, acc)
}

func TestWithElements(t *testing.T) {
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	var seed [32]byte
	key, _ := new(zk.SecretKey).New(curve, seed[:])
	element1 := curve.Scalar.Hash([]byte("value1"))
	element2 := curve.Scalar.Hash([]byte("value2"))
	elements := []zk.Element{element1, element2}
	newAcc, err := new(zk.Accumulator).WithElements(curve, key, elements)
	require.NoError(t, err)
	require.NotNil(t, newAcc)

	_, _ = newAcc.Remove(key, element1)
	_, _ = newAcc.Remove(key, element2)
}


func TestUpdate(t *testing.T) {
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	var seed [32]byte
	key, err := new(zk.SecretKey).New(curve, seed[:])
	require.NoError(t, err)
	require.NotNil(t, key)
	acc, err := new(zk.Accumulator).New(curve)
	require.NoError(t, err)
	require.NotNil(t, acc)

	element1 := curve.Scalar.Hash([]byte("value1"))
	element2 := curve.Scalar.Hash([]byte("value2"))
	element3 := curve.Scalar.Hash([]byte("value3"))
	elements := []zk.Element{element1, element2, element3}

	acc, _, err = acc.Update(key, elements, nil)
	require.NoError(t, err)

	acc, _, err = acc.Update(key, nil, elements)
	require.NoError(t, err)
}
