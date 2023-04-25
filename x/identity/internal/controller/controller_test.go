package controller_test

import (
	"testing"

	"github.com/sonrhq/core/internal/crypto"
	"github.com/sonrhq/core/internal/crypto/mpc"
	"github.com/sonrhq/core/x/identity/internal/controller"
	"github.com/sonrhq/core/x/identity/types/models"
	servicetypes "github.com/sonrhq/core/x/service/types"
	"github.com/stretchr/testify/assert"
	"github.com/tendermint/tendermint/libs/rand"
)

func TestKeyshare(t *testing.T) {
	randUuid := rand.Str(4)

	// Call Handler for keygen
	confs, err := mpc.Keygen(crypto.PartyID(randUuid), mpc.WithThreshold(1))
	if err != nil {
		t.Fatal(err)
	}

	var kss []models.KeyShare
	for _, conf := range confs {
		ksb, err := conf.MarshalBinary()
		if err != nil {
			t.Fatal(err)
		}
		ks, err := models.NewKeyshare(string(conf.ID), ksb, crypto.SONRCoinType)
		if err != nil {
			t.Fatal(err)
		}
		kss = append(kss, ks)
	}

	for _, ks := range kss {
		t.Logf("keyshare key id: %s", ks.Did())
	}
	acc := models.NewAccount(kss, crypto.SONRCoinType)
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
	t.Log("create controller with initial accounts: bitcoin, ethereum")
	controller, err := controller.NewController(defaultOfflineOptions()...)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("controller: %v", controller.Did())
	t.Log("list accounts")
	accs, err := controller.ListAccounts()
	if err != nil {
		t.Fatal(err)
	}
	primAcc := accs[0]
	t.Logf("primary account: %v", primAcc.Did())
	t.Log("create new account: ethereum")
	ethAcc, err := controller.CreateAccount("ethereum", crypto.ETHCoinType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("ethereum account: %v", ethAcc.Did())
	msg := []byte("hello world")
	sig, err := ethAcc.Sign(msg)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("signature: %x", sig)
	ok, err := ethAcc.Verify(msg, sig)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("verify: %v", ok)

	t.Log("create new account: bitcoin")
	btcAcc, err := controller.CreateAccount("bitcoin", crypto.BTCCoinType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("bitcoin account: %v", btcAcc.Did())
	msg = []byte("hello world")
	sig, err = btcAcc.Sign(msg)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("signature: %x", sig)
	ok, err = btcAcc.Verify(msg, sig)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("verify: %v", ok)
}

func TestControllerMail(t *testing.T) {
	cn, err := controller.NewController(defaultOptions()...)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("controller: %v", cn.Did())

	ethAcc1, err := cn.CreateAccount("ethTest1", crypto.ETHCoinType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Acc #1: %v", cn.Address())

	ethAcc2, err := cn.CreateAccount("ethTest2", crypto.ETHCoinType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Acc #2: %v", cn.Address())

	t.Logf("Create message 'hello world' from %s to %s", ethAcc1.Address(), ethAcc2.Address())
	err = cn.SendMail(ethAcc1.Address(), ethAcc2.Address(), "hello world")
	if err != nil {
		t.Fatal(err)
	}

	mails, err := cn.ReadMail(ethAcc2.Address())
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, len(mails))
	t.Logf("\n[MAIL]: \n\t- ID: %s\n\t- Content: %s \n\t- To:%s\n\t- From: %s", mails[0].Id, mails[0].Content, mails[0].Receiver, mails[0].Sender)
}

func TestControllerLoad(t *testing.T) {
	cn, err := controller.NewController(defaultOfflineOptions()...)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("controller: %v", cn.Did())

	_, err = cn.CreateAccount("ethTest", crypto.ETHCoinType)
	if err != nil {
		t.Fatal(err)
	}
	cn2, err := controller.LoadController(cn.PrimaryIdentity())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("controller: %v", cn2.Did())

	assert.Equal(t, cn.Did(), cn2.Did())
}

func TestBroadcast(t *testing.T) {
	opts := defaultOfflineOptions()
	//opts = append(opts, controller.WithBroadcastTx())
	cn, err := controller.NewController(opts...)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("controller: %v", cn.Did())
}

func defaultOfflineOptions() []controller.Option {
	cred := &servicetypes.WebauthnCredential{
		Id:        []byte("test"),
		PublicKey: []byte("-----BEGIN PUBLIC KEY----- MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEQ9Z0Z8Z0Z8Z0Z8Z0Z8Z0Z8Z0Z8Z0Z-----END PUBLIC KEY-----"),
	}
	return []controller.Option{
		controller.WithWebauthnCredential(cred),
	}
}

func defaultOptions() []controller.Option {
	cred := &servicetypes.WebauthnCredential{
		Id:        []byte("test"),
		PublicKey: []byte("-----BEGIN PUBLIC KEY----- MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEQ9Z0Z8Z0Z8Z0Z8Z0Z8Z0Z8Z0Z8Z0Z-----END PUBLIC KEY-----"),
	}
	return []controller.Option{
		controller.WithWebauthnCredential(cred),
	}
}
