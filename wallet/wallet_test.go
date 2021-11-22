package wallet_test

import (
	"log"
	"testing"

	"github.com/sonr-io/core/wallet"
)

func TestCreateWallet(t *testing.T) {
	t.Run("for VC types", func(t *testing.T) {
		pp := "super-test-passphrase"
		sname := "testSname"
		err := wallet.Create(pp, sname)
		if err != nil {
			log.Fatal(err)
		}
		key, err := wallet.DevicePubKey()
		if err != nil {
			t.Fatal(err)
		}
		log.Printf("Device PubKey: %v", key)

		privKey, err := wallet.DevicePrivKey()
		if err != nil {
			t.Fatal(err)
		}
		log.Printf("Device PrivKey: %v", privKey)
	})
}

func TestOpenWallet(t *testing.T) {
	t.Run("for VC types", func(t *testing.T) {
		err := wallet.Open(wallet.WithPassphrase("super-test-passphrase"), wallet.WithSName("testSname"))
		if err != nil {
			t.Fatal(err)
		}
		key, err := wallet.DevicePubKey()
		if err != nil {
			t.Fatal(err)
		}
		log.Printf("Device PubKey: %v", key)

		privKey, err := wallet.DevicePrivKey()
		if err != nil {
			t.Fatal(err)
		}
		log.Printf("Device PrivKey: %v", privKey)
	})
}
