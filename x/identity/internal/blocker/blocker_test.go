package blocker_test

import (
	"testing"
	"time"

	"github.com/sonrhq/core/x/identity/internal/blocker"
)

func TestBlocker(t *testing.T) {
	sq := blocker.NewBlocker()
	for i := 0; i < 5; i++ {
		sq.Next()
		w := sq.Pop()
		if w == nil {
			t.Log("No wallet available")
			time.Sleep(time.Second)
			continue
		}
		t.Logf("Wallet %d: %s", w.Id, w.PublicKey)
	}
}
