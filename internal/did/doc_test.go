package did_test

import (
	"crypto/rand"
	"log"
	"testing"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/sonr-io/core/internal/did"
)

func TestParseURI(t *testing.T) {

	t.Run("for VC types", func(t *testing.T) {
		// Create New Account Key
		_, pub, err := crypto.GenerateEd25519Key(rand.Reader)
		if err != nil {
			log.Printf("%s - Failed to generate Account KeyPair", err)
			t.Fail()
		}

		_, err = did.NewIdWallet(pub, "test")
		if err != nil {
			log.Printf("%s - Failed to create new wallet", err)
			t.Fail()
		}
	})
}
