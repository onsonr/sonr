package crypto

import (
	"errors"
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
func Generate(options ...WalletOption) (*MPCWallet, error) {
	opt := defaultConfig()
	wallet := opt.Apply(options...)
	pl := pool.NewPool(0)
	defer pl.TearDown()

	var wg sync.WaitGroup
	for _, id := range wallet.Participants {
		wg.Add(1)
		go func(id party.ID) {
			pl := pool.NewPool(0)
			defer pl.TearDown()
			if err := wallet.CMPKeygen(id, pl, &wg); err != nil {
				fmt.Println(err)
			} else {
				c, err := wallet.Config.DeriveBIP32(2)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(c.PublicPoint())
				fmt.Println("success")
				fmt.Printf("%+v\n", wallet.Config)
			}
		}(id)
	}

	fmt.Println("done.")
	return wallet, nil
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
func (w *MPCWallet) CMPSign(m []byte, signers party.IDSlice, pl *pool.Pool) error {
	h, err := protocol.NewMultiHandler(cmp.Sign(w.Config, signers, m, pl), nil)
	if err != nil {
		return err
	}
	handlerLoop(w.Config.ID, h, w.Network)

	signResult, err := h.Result()
	if err != nil {
		return err
	}
	signature := signResult.(*ecdsa.Signature)
	if !signature.Verify(w.Config.PublicPoint(), m) {
		return errors.New("failed to verify cmp signature")
	}
	return nil
}

// Generates a preprocessed ECDSA signature which does not depend on the message being signed.
func (w *MPCWallet) CMPPreSign(signers party.IDSlice, pl *pool.Pool) (*ecdsa.PreSignature, error) {
	h, err := protocol.NewMultiHandler(cmp.Presign(w.Config, signers, pl), nil)
	if err != nil {
		return nil, err
	}
	handlerLoop(w.Config.ID, h, w.Network)

	signResult, err := h.Result()
	if err != nil {
		return nil, err
	}

	preSignature := signResult.(*ecdsa.PreSignature)
	if err = preSignature.Validate(); err != nil {
		return nil, errors.New("failed to verify cmp presignature")
	}
	return preSignature, nil
}

// Combines each party's PreSignature share to create an ECDSA signature for messageHash.
func (w *MPCWallet) CMPPreSignOnline(preSignature *ecdsa.PreSignature, m []byte, pl *pool.Pool) error {
	h, err := protocol.NewMultiHandler(cmp.PresignOnline(w.Config, preSignature, m, pl), nil)
	if err != nil {
		return err
	}
	handlerLoop(w.Config.ID, h, w.Network)

	signResult, err := h.Result()
	if err != nil {
		return err
	}
	signature := signResult.(*ecdsa.Signature)
	if !signature.Verify(w.Config.PublicPoint(), m) {
		return errors.New("failed to verify cmp signature")
	}
	return nil
}
