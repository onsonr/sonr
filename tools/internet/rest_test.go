package internet_test

import (
	"context"
	"crypto/rand"
	"os"
	"testing"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/sonr-io/core/internal/keychain"
	"github.com/sonr-io/core/tools/internet"
)

func TestNewAuthRecord(t *testing.T) {
	// Create Protocol
	testName := "testerLope"
	nbClient := internet.NewNamebaseClient(context.Background(), os.Getenv("HANDSHAKE_KEY"), os.Getenv("HANDSHAKE_SECRET"))

	// Generate Keys
	_, pubKey, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		t.Error(err)
	}

	// Convert to SNRKey
	snrPubKey := keychain.NewSnrPubKey(pubKey)
	pubStr, err := snrPubKey.String()
	if err != nil {
		t.Error(err)
	}

	// Create Record
	record := internet.NewNBNameRecord(pubStr, testName)
	println(record.Name())
	println(record.Prefix())
	println(record.Fingerprint())
	t.Log(record.Fingerprint())
	t.Log(record.Name())
	t.Log(record.Prefix())

	// Create New Add NamebaseRequest
	ok, err := nbClient.PutRecords(internet.NewNBAddRequest(record))
	if err != nil {
		t.Error(err)
	}
	if !ok {
		t.Error("AddNamebaseRequest failed")
	}

	// Verify ADD result by ListRecords
	putSuccess, err := nbClient.HasRecords(testName)
	if err != nil {
		t.Error(err)
	}
	if !putSuccess {
		t.Error("PutNamebaseRecords failed")
	} else {
		println("PutNamebaseRecords success")
	}

	// Create New Delete NamebaseRequest
	ok, err = nbClient.PutRecords(internet.NewNBDeleteRequest(record.ToDeleteRecord()))
	if err != nil {
		t.Error(err)
	}
	if !ok {
		t.Error("DeleteNamebaseRequest failed")
	}

	// Verify Delete result by ListRecords
	hasRec, err := nbClient.HasRecords(testName)
	if err != nil {
		t.Error(err)
	}

	if hasRec {
		t.Error("DeleteNamebaseRequest failed")
	} else {
		println("Namebase DELETE Succeeded")
	}
	t.Log(hasRec)
}
