package crypto

import (
	"fmt"
	"sync"

	"github.com/taurusgroup/multi-party-sig/pkg/ecdsa"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/pkg/pool"
	"github.com/taurusgroup/multi-party-sig/pkg/protocol"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

type MPCWallet struct {
	Config       *cmp.Config
	Network      *Network
	Participants party.IDSlice
	Threshold    int
}

// Generate a new ECDSA private key shared among all the given participants.
func Generate(pl *pool.Pool, options ...WalletOption) (*MPCWallet, error) {
	opt := defaultConfig()
	wallet := opt.Apply(options...)

	var wg sync.WaitGroup
	for _, id := range wallet.Participants {
		wg.Add(1)
		go func(id party.ID) {
			pl := pool.NewPool(0)
			defer pl.TearDown()
			if err := wallet.CMPKeygen(id, pl, &wg); err != nil {
				fmt.Println(err)
				return
			}
		}(id)
	}
	wg.Wait()
	fmt.Println("done.")
	return wallet, nil
}

func (w *MPCWallet) GetSigners(id party.ID) party.IDSlice {
	signers := w.Participants[:w.Threshold+1]
	if !signers.Contains(id) {
		w.Network.Quit(id)
		return nil
	}
	return signers
}

// Generate a new ECDSA private key shared among all the given participants.
func (w *MPCWallet) CMPKeygen(id party.ID, pl *pool.Pool, wg *sync.WaitGroup) error {
	h, err := protocol.NewMultiHandler(cmp.Keygen(curve.Secp256k1{}, id, w.Participants, w.Threshold, pl), nil)
	if err != nil {
		return err
	}
	handlerLoop(id, h, w.Network)
	r, err := h.Result()
	if err != nil {
		return err
	}
	wg.Done()
	conf := r.(*cmp.Config)
	w.Config = conf
	return nil
}

// Refreshes all shares of an existing ECDSA private key.
func (w *MPCWallet) CMPRefresh(pl *pool.Pool) (*cmp.Config, error) {
	hRefresh, err := protocol.NewMultiHandler(cmp.Refresh(w.Config, pl), nil)
	if err != nil {
		return nil, err
	}
	handlerLoop(w.Config.ID, hRefresh, w.Network)

	r, err := hRefresh.Result()
	if err != nil {
		return nil, err
	}

	return r.(*cmp.Config), nil
}

// Generates an ECDSA signature for messageHash.
func (w *MPCWallet) CMPSign(m []byte, signers party.IDSlice, pl *pool.Pool) (*ecdsa.Signature, error) {
	h, err := protocol.NewMultiHandler(cmp.Sign(w.Config, signers, m, pl), nil)
	if err != nil {
		return nil, err
	}
	handlerLoop(w.Config.ID, h, NewNetwork(signers))
	signResult, err := h.Result()
	if err != nil {
		return nil, err
	}
	signature := signResult.(*ecdsa.Signature)
	if !signature.Verify(w.Config.PublicPoint(), m) {
		return nil, fmt.Errorf("failed to verify cmp signature")
	}
	return signature, nil
}
