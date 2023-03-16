package wallet

import (
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
			t.Logf(" \t↪ PubKey: %s", acc.PubKey().Base64())
		}
		t.Logf("")
	}
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

func TestExportImport(t *testing.T) {
	w, err := LoadWallet()
	if err != nil {
		t.Fatal(err)
	}

	// Export
	enc, err := w.Export()
	if err != nil {
		t.Fatal(err)
	}

	// Import
	w2, err := Import(enc)
	if err != nil {
		t.Fatal(err)
	}

	// Compare
	if w2 == nil {
		t.Fatal("wallet is nil")
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
