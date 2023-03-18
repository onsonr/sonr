package wallet

import (
	"crypto/ed25519"
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/sonrhq/core/pkg/crypto"
	"github.com/stretchr/testify/assert"
)

func TestNewWallet(t *testing.T) {
	w, err := NewWallet("test", 1)
	if err != nil {
		t.Fatal(err)
	}
	if w == nil {
		t.Fatal("wallet is nil")
	}
}

func TestCreateAccount(t *testing.T) {
	w, err := NewWallet("test", 1)
	if err != nil {
		t.Fatal(err)
	}

	for _, tt := range defaultCoinTestsSet() {
		_, err := w.CreateAccount(tt.coinType)
		if err != nil {
			t.Fatal(err)
		}

		accs, err := w.ListAccountsForCoin(tt.coinType)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("<%s> %s (%d)", tt.coinType.Ticker(), tt.coinType.Name(), tt.coinType.BipPath())
		for i, acc := range accs {
			t.Logf("- [%d] %s", i, acc.Name())
			t.Logf(" \t↪ Address: %s", acc.Address())
			t.Logf(" \t↪ Base64: %s", acc.PubKey().Base64())
			t.Logf(" \t↪ Blake3: %s", acc.PubKey().Blake3())
			t.Logf(" \t↪ Multibase: %s", acc.PubKey().Multibase())
		}
		t.Logf("")
	}
}

func TestLockUnlockWallet(t *testing.T) {
	// Sample public key generated from ed25519.GenerateKey(rand.Reader)
	var pub ed25519.PublicKey = []byte{0x7b, 0x88, 0x10, 0x24, 0xad, 0xc9, 0x82, 0xd3, 0x80, 0xb8, 0x77, 0x1e, 0x3b, 0x9b, 0xf8, 0xe4, 0xb3, 0x99, 0x8b, 0xc7, 0xd0, 0x58, 0x30, 0x66, 0x2, 0xce, 0x4d, 0xf, 0x2f, 0xe4, 0xb7, 0x81}
	credential := crypto.WebauthnCredential{
		Id:              []byte("some-probabilistically-unique-id"),
		PublicKey:       pub,
		AttestationType: "some-attestation-type",
		Transport:       []string{"usb", "ble"},
		Authenticator: &crypto.WebauthnAuthenticator{
			Aaguid:       []byte("some-aaguid"),
			CloneWarning: true,
			SignCount:    123,
		},
	}

	w, err := NewWallet("test", 1)
	assert.NoError(t, err)

	ethAcc, err := w.CreateAccount(crypto.ETHCoinType)
	assert.NoError(t, err)

	doc, _, err := w.Assign(&credential)
	assert.NoError(t, err)
	t.Logf("DID: %s", doc.String())

	// Attempt to sign a message with the wallet before it's unlocked
	_, err = ethAcc.Sign([]byte("some-message"))
	assert.Error(t, err)

	// Unlock the wallet
	err = w.Unlock(&credential)
	assert.NoError(t, err)

	// Sign a message with the wallet after it's unlocked
	sig, err := ethAcc.Sign([]byte("some-message"))
	assert.NoError(t, err)

	// Verify the signature
	okay, err := ethAcc.Verify([]byte("some-message"), sig)
	assert.NoError(t, err)
	assert.True(t, okay)
}

func TestGetAccount(t *testing.T) {
	w, err := LoadWallet()
	if err != nil {
		t.Fatal(err)
	}

	for _, tt := range defaultAccountTestsSet() {
		acc, err := w.GetAccount(tt.coinType, tt.index)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("<%s> %s (%d)", tt.coinType.Ticker(), tt.coinType.Name(), tt.coinType.BipPath())
		t.Logf("- [%d] %s", tt.index, acc.DID())
		t.Logf(" \t↪ Address: %s", acc.Address())
		// !!! This is what was used to connect two addresses from this test mpc wallet to the cosmos hub
		t.Logf(" \t↪ PubKey: %s", acc.PubKey().Base64())
		t.Logf("")
	}
}

func TestSignWithAccount(t *testing.T) {
	w, err := LoadWallet()
	if err != nil {
		t.Fatal(err)
	}
	for _, tt := range defaultAccountTestsSet() {
		acc, err := w.GetAccount(tt.coinType, tt.index)
		if err != nil {
			t.Fatal(err)
		}
		msg := []byte(fmt.Sprintf("Hello %s!", tt.coinType.Name()))
		t.Logf("- [%d] %s - %s", tt.index, acc.Name(), acc.Address())
		t.Logf(" \t↪ Message: %s", string(msg))
		sig, err := acc.Sign(msg)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf(" \t↪ Signature: %s", base64.StdEncoding.EncodeToString(sig))

		ok, err := acc.Verify(msg, sig)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("   Verify => %v", ok)
		t.Logf("")
	}
}

func TestSignWithDIDs(t *testing.T) {

	w, err := LoadWallet()
	if err != nil {
		t.Fatal(err)
	}
	for _, tt := range defaultAccountTestsSet() {
		acc, err := w.GetAccount(tt.coinType, tt.index)
		if err != nil {
			t.Fatal(err)
		}
		msg := []byte(fmt.Sprintf("Hello %s!", tt.coinType.Name()))
		t.Logf("- [%d] %s", tt.index, acc.DID())
		t.Logf(" \t↪ Message: %s", string(msg))
		sig, err := w.SignWithDID(acc.DID(), msg)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf(" \t↪ Signature: %s", base64.StdEncoding.EncodeToString(sig))

		ok, err := w.VerifyWithDID(acc.DID(), msg, sig)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("   Verify => %v", ok)
		t.Logf("")
	}
}

func TestMiscWalletActions(t *testing.T) {
	w, err := LoadWallet()
	if err != nil {
		t.Fatal(err)
	}
	assert.Contains(t, w.Controller(), "did:sonr:")
	t.Logf("Controller: %s", w.Controller())
	allCoins := crypto.AllCoinTypes()
	for _, coin := range allCoins {
		t.Logf("Coin: %s, %d accounts", coin.Name(), w.Count(coin))
		coinCount := w.Count(coin)
		t.Logf("Count: %d", coinCount)
		target := 1
		if coinCount < target {
			target = coinCount
		}
		for i := 0; i < target; i++ {
			acc, err := w.GetAccount(coin, i)
			if err != nil {
				t.Fatal(err)
			}
			t.Logf(" \t↪ Account: %s", acc.DID())

			// Test signing with the account
			msg := []byte(fmt.Sprintf("Hello %s!", coin.Name()))
			t.Logf(" \t↪ Message: %s", string(msg))
			sig, err := acc.Sign(msg)
			if err != nil {
				t.Fatal(err)
			}
			t.Logf(" \t↪ Signature: %s", base64.StdEncoding.EncodeToString(sig))

			ok, err := acc.Verify(msg, sig)
			if err != nil {
				t.Fatal(err)
			}
			t.Logf("   Verify => %v", ok)
			t.Logf("")
		}
	}
	size, err := w.Size()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Wallet Size: %d", size)
}

func defaultCoinTestsSet() []struct {
	coinType crypto.CoinType
} {
	return []struct {
		coinType crypto.CoinType
	}{
		{crypto.BTCCoinType},
		{crypto.ETHCoinType},
		{crypto.SONRCoinType},
		{crypto.FILCoinType},
		{crypto.DOGECoinType},
	}
}

func defaultAccountTestsSet() []struct {
	coinType crypto.CoinType
	index    int
} {
	return []struct {
		coinType crypto.CoinType
		index    int
	}{
		{crypto.BTCCoinType, 0},
		{crypto.ETHCoinType, 0},
		{crypto.SONRCoinType, 0},
		{crypto.FILCoinType, 0},
		{crypto.DOGECoinType, 0},
	}
}
