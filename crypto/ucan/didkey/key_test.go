package didkey

import (
	"log"
	"testing"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/onsonr/sonr/crypto/mpc"
)

func TestID_Parse(t *testing.T) {
	keyStrED := "did:key:z6MkpTHR8VNsBxYAAWHut2Geadd9jSwuBV8xRoAnwWsdvktH"
	id, err := Parse(keyStrED)
	if err != nil {
		t.Fatal(err)
	}

	if id.String() != keyStrED {
		t.Errorf("string mismatch.\nwant: %q\ngot:  %q", keyStrED, id.String())
	}

	keyStrRSA := "did:key:z2MGw4gk84USotaWf4AkJ83DcnrfgGaceF86KQXRYMfQ7xqnUFp38UZ6Le8JPfkb4uCLGjHBzKpjEXb9hx9n2ftecQWCHXKtKszkke4FmENdTZ7i9sqRmL3pLnEEJ774r3HMuuC7tNRQ6pqzrxatXx2WinCibdhUmvh3FobnA9ygeqkSGtV6WLa7NVFw9cAvnv8Y6oHcaoZK7fNP4ASGs6AHmSC6ydSR676aKYMe95QmEAj4xJptDsSxG7zLAGzAdwCgm56M4fTno8GdWNmU6Pdghnuf6fWyYus9ASwdfwyaf3SDf4uo5T16PRJssHkQh6DJHfK4Rka7RNQLjzfGBPjFLHbUSvmf4EdbHasbVaveAArD68ZfazRCCvjdovQjWr6uyLCwSAQLPUFZBTT8mW"

	id, err = Parse(keyStrRSA)
	if err != nil {
		t.Fatal(err)
	}

	if id.String() != keyStrRSA {
		t.Errorf("string mismatch.\nwant: %q\ngot:  %q", keyStrRSA, id.String())
	}
}

func TestID_FromMPCKey(t *testing.T) {
	// Generate new MPC keyset
	ks, err := mpc.NewKeyset()
	if err != nil {
		t.Fatalf("failed to generate MPC keyset: %v", err)
	}

	// Get public key from validator share
	pubKey := ks.Val().PublicKey()
	if len(pubKey) != 65 {
		t.Fatalf("expected 65-byte uncompressed public key, got %d bytes", len(pubKey))
	}

	// Create crypto.PubKey from raw bytes
	cryptoPubKey, err := crypto.UnmarshalSecp256k1PublicKey(pubKey)
	if err != nil {
		t.Fatalf("failed to unmarshal public key: %v", err)
	}

	// Create DID Key ID
	id, err := NewID(cryptoPubKey)
	if err != nil {
		t.Fatalf("failed to create DID Key ID: %v", err)
	}
	log.Printf("%s\n", id.String())

	// Verify the key can be parsed back
	parsed, err := Parse(id.String())
	if err != nil {
		t.Fatalf("failed to parse DID Key string: %v", err)
	}

	// Verify the parsed key matches original
	if parsed.String() != id.String() {
		t.Errorf("parsed key doesn't match original.\nwant: %q\ngot:  %q",
			id.String(), parsed.String())
	}

	// Verify we can get back a valid verify key
	verifyKey, err := id.VerifyKey()
	if err != nil {
		t.Fatalf("failed to get verify key: %v", err)
	}

	// Verify the key is the right type and length
	rawKey, ok := verifyKey.([]byte)
	if !ok {
		t.Fatalf("expected []byte verify key, got %T", verifyKey)
	}
	if len(rawKey) != 65 && len(rawKey) != 33 {
		t.Errorf("invalid key length %d, expected 65 or 33 bytes", len(rawKey))
	}
}
