package mpc

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
)

var defaultTestParticipants = party.IDSlice{"vault", "current"}

func TestSaveLoadWallet(t *testing.T) {
	w := EmptyWalletShare()
	path := filepath.Join(os.TempDir(), "test_wallet.json")
	err := SaveToPath(w, path)
	if err != nil {
		t.Fatal(err)
	}
	ws, err := LoadFromPath(path)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, w.Address(), ws.Address())
}
