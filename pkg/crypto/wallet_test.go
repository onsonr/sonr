package crypto

import (
	"fmt"
	"testing"

	"github.com/taurusgroup/multi-party-sig/pkg/pool"
)

func Test_MPCCreate(t *testing.T) {
	pl := pool.NewPool(0)
	defer pl.TearDown()
	_, err := Generate(pl, WithParticipants("bio1"))
	if err != nil {
		t.Error(err)
	}
}

func Test_MPCSign(t *testing.T) {
	pl := pool.NewPool(0)
	w, err := Generate(pl, WithParticipants("bio1"))
	if err != nil {
		t.Error(err)
	}

	sig, err := w.CMPSign([]byte("test"), w.GetSigners("bio1"), pl)
	if err != nil {
		fmt.Println(err)
		return
	}

	if !sig.Verify(w.Config.PublicPoint(), []byte("test")) {
		fmt.Println("failed to verify cmp signature")
		return
	}
	t.Log("success")
}
