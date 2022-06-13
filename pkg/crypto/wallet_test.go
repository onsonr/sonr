package crypto

import (
	"testing"

	"github.com/taurusgroup/multi-party-sig/pkg/pool"
)

func Test_MPCCreate(t *testing.T) {
	_, err := Generate()
	if err != nil {
		t.Error(err)
	}
}

func Test_MPCSign(t *testing.T) {
	pl := pool.NewPool(0)
	w, err := Generate()
	if err != nil {
		t.Error(err)
	}

	sig, err := w.Sign([]byte("test"), pl)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("success", sig)
}

func Test_MPCCosmosAddr(t *testing.T) {
	w, err := Generate()
	if err != nil {
		t.Error(err)
		return
	}

	pub, err := w.AccountAddress()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("success", pub)
}
