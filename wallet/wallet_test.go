package wallet_test

import (
	"log"
	"testing"

	"github.com/sonr-io/core/wallet"
)

func TestParseURI(t *testing.T) {

	t.Run("for VC types", func(t *testing.T) {
		err := wallet.Open(wallet.WithPassphrase("super-test-passphrase"), wallet.WithSName("test"), wallet.Reset())
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("Wallet created: %s", wallet.Info.String())
		key, err := wallet.DevicePubKey()
		if err != nil {
			t.Fatal(err)
		}
		log.Printf("Device KeyInfo: %v", key)

		pkh, err := wallet.DevicePrivKH()
		if err != nil {
			t.Fatal(err)
		}
		log.Printf("Device PrivKey: %v", pkh)

		privKey, err := wallet.DevicePrivKey()
		if err != nil {
			t.Fatal(err)
		}
		log.Printf("Device PrivKey: %v", privKey)
	})
}
