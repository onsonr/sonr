package controller

import (
	"context"
	"testing"

	"github.com/sonrhq/core/pkg/crypto"
	"github.com/sonrhq/core/pkg/crypto/mpc"
	"github.com/tendermint/tendermint/libs/rand"
)

func TestTree(t *testing.T) {
	randUuid := rand.Str(4)

	// Call Handler for keygen
	confs, err := mpc.Keygen(crypto.PartyID(randUuid), 1, []crypto.PartyID{"vault"})
	if err != nil {
		t.Fatal(err)
	}

	var kss []KeyShare
	for _, conf := range confs {
		ksb, err := conf.MarshalBinary()
		if err != nil {
			t.Fatal(err)
		}
		ks, err := NewKeyshare(string(conf.ID), ksb, crypto.SONRCoinType, "test")
		if err != nil {
			t.Fatal(err)
		}
		kss = append(kss, ks)
	}

	for _, ks := range kss {
		t.Logf("keyshare key id: %s", ks.Did())
	}
	acc := NewAccount(kss, crypto.SONRCoinType)
	if err != nil {
		t.Fatal(err)
	}
	msg := []byte("hello world")
	sig, err := acc.Sign(msg)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("signature: %x", sig)
	ok, err := acc.Verify(msg, sig)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("verify: %v", ok)
}

func TestController(t *testing.T) {
	randUuid := rand.Str(4)
	cred := &crypto.WebauthnCredential{
		Id: []byte(randUuid),
	}

	controller, err := NewController(context.Background(), cred)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("controller: %v", controller.Did())

	err = controller.CreateAccount("ethTest", crypto.ETHCoinType)
	if err != nil {
		t.Fatal(err)
	}

	acc, err := controller.GetAccount("ethTest", crypto.ETHCoinType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("account: %v", acc.Address())

	msg := []byte("hello world")
	sig, err := acc.Sign(msg)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("signature: %x", sig)

	ok, err := acc.Verify(msg, sig)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("verify: %v", ok)

	sig2, err := controller.Sign("ethTest", crypto.ETHCoinType, msg)
	if err != nil {
		t.Fatal(err)
	}

	ok, err = controller.Verify("ethTest", crypto.ETHCoinType, msg, sig2)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("verify: %v", ok)

	err = controller.CreateAccount("btcTest", crypto.BTCCoinType)
	if err != nil {
		t.Fatal(err)
	}

	acc2, err := controller.GetAccount("btcTest", crypto.BTCCoinType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("account: %v", acc2.Address())

	did := controller.Did()
	t.Logf("did: %s", did)

	didDoc := controller.DidDocument()
	t.Logf("did doc: %v", didDoc.String())
}

func TestNewLoad(t *testing.T) {
	randUuid := rand.Str(4)
	cred := &crypto.WebauthnCredential{
		Id: []byte(randUuid),
	}

	cn, err := NewController(context.Background(), cred)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("controller: %v", cn.Did())
	didDoc := cn.DidDocument()

	err = cn.CreateAccount("ethTest", crypto.ETHCoinType)
	if err != nil {
		t.Fatal(err)
	}

	cn2, err := LoadController(context.Background(), cred, didDoc)
	if err != nil {
		t.Fatal(err)
	}

	acc, err := cn2.GetAccount("ethTest", crypto.ETHCoinType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("account: %v", acc.Address())
}
