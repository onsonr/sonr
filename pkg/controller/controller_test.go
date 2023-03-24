package controller_test

import (
	"context"
	"testing"

	"github.com/sonrhq/core/pkg/controller"
	"github.com/sonrhq/core/pkg/crypto"
	"github.com/sonrhq/core/pkg/crypto/mpc"
	"github.com/sonrhq/core/pkg/resolver"
	"github.com/sonrhq/core/pkg/tx/cosmos"
	"github.com/sonrhq/core/x/identity/types"
	"github.com/tendermint/tendermint/libs/rand"
)

func TestTree(t *testing.T) {
	randUuid := rand.Str(4)

	// Call Handler for keygen
	confs, err := mpc.Keygen(crypto.PartyID(randUuid), 1, []crypto.PartyID{"vault"})
	if err != nil {
		t.Fatal(err)
	}

	var kss []controller.KeyShare
	for _, conf := range confs {
		ksb, err := conf.MarshalBinary()
		if err != nil {
			t.Fatal(err)
		}
		ks, err := controller.NewKeyshare(string(conf.ID), ksb, crypto.SONRCoinType, "test")
		if err != nil {
			t.Fatal(err)
		}
		kss = append(kss, ks)
	}

	for _, ks := range kss {
		t.Logf("keyshare key id: %s", ks.Did())
	}
	acc := controller.NewAccount(kss, crypto.SONRCoinType)
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

	controller, _, err := controller.NewController(context.Background(), cred)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("controller: %v", controller.Did())

	_, err = controller.CreateAccount("ethTest", crypto.ETHCoinType)
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

	acc2, err := controller.CreateAccount("btcTest", crypto.BTCCoinType)
	if err != nil {
		t.Fatal(err)
	}
	did := acc2.DID()
	t.Logf("did: %s", did)

	didDoc := controller.DidDocument()
	t.Logf("did doc: %v", didDoc.String())
}

func TestNewLoad(t *testing.T) {
	randUuid := rand.Str(4)
	cred := &crypto.WebauthnCredential{
		Id: []byte(randUuid),
	}

	cn, _, err := controller.NewController(context.Background(), cred)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("controller: %v", cn.Did())
	didDoc := cn.DidDocument()

	_, err = cn.CreateAccount("ethTest", crypto.ETHCoinType)
	if err != nil {
		t.Fatal(err)
	}

	cn2, err := controller.LoadController(context.Background(), didDoc)
	if err != nil {
		t.Fatal(err)
	}

	acc, err := cn2.GetAccount("ethTest", crypto.ETHCoinType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("account: %v", acc.Address())
}

func TestControllerCreateBroadcastTx(t *testing.T) {
	randUuid := rand.Str(4)
	cred := &crypto.WebauthnCredential{
		Id: []byte(randUuid),
	}

	cn, prim, err := controller.NewController(context.Background(), cred)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("controller: %v", cn.Did())
	didDoc := cn.DidDocument()

	_, err = cn.CreateAccount("ethTest", crypto.ETHCoinType)
	if err != nil {
		t.Fatal(err)
	}

	msg := types.NewMsgCreateDidDocument(cn.Address(), didDoc)
	txBz, err := cosmos.SignAnyTransactions(prim, msg)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("tx: %x", txBz)

	res, err := resolver.BroadcastTx(context.TODO(), txBz)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("res: %v", res)
}
