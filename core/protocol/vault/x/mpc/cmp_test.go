package mpc_test

import (
	"testing"

	"github.com/sonr-io/multi-party-sig/pkg/party"
	"github.com/sonr-io/sonr/core/protocol/vault/x/mpc"
)

const (
	path     = "/Users/pradn/.sonr/wallet/a_sonr_id.json"
	id       = "a"
	password = "goodpassword"
)

var (
	ids = party.IDSlice{"a", "b", "c", "d", "e", "f"}
	msg = []byte("hello world")
	sig []byte
)

func TestCreateSaveLoadWallet(t *testing.T) {
	w, err := mpc.NewWallet(id, ids)
	if err != nil {
		t.Fatal(err)
	}
	_, err = w.Save(password)
	if err != nil {
		t.Fatal(err)
	}
	_, err = mpc.NewWalletFromDisk(path, password)
	if err != nil {
		t.Fatal(err)
	}
}
