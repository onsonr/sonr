package types_test

import (
	"testing"

	secp256k1 "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/sonr-io/core/pkg/mpc/types/v1"
	"github.com/stretchr/testify/assert"
)

func TestZkSet(t *testing.T) {
	priv := secp256k1.GenPrivKey()
	pub := priv.PubKey()
	zkset, err := types.CreateZkSet(pub)
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

	zkset2, err := types.OpenZkSet(raw)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(zkset2)
	ok1 := zkset2.ValidateMembership(pub, "test")
	assert.True(t, ok1)
}
