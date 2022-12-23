package mpc

import (
	"context"
	"sync"

	"github.com/sonr-hq/sonr/pkg/common"
	"github.com/sonr-hq/sonr/pkg/wallet"
	"github.com/taurusgroup/multi-party-sig/pkg/ecdsa"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/pkg/pool"
	"github.com/taurusgroup/multi-party-sig/pkg/protocol"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

type MPCProtocol struct {
	ctx          context.Context
	pool         *pool.Pool
	currentId    party.ID
	participants party.IDSlice

	pubKey    []byte
	configs   map[party.ID]*cmp.Config
	mtx       sync.Mutex
	network   *Network
	threshold int
	callback  common.NodeCallback
}

func NewProtocol(ctx context.Context, options ...WalletOption) *MPCProtocol {
	opt := defaultConfig()
	w := opt.Apply(options...)
	w.ctx = ctx
	return w
}

// GenerateWallet a new ECDSA private key shared among all the given participants.
func (p *MPCProtocol) Keygen(target party.ID) (wallet.WalletShare, error) {
	doneChan := make(chan *cmp.Config)
	var wg sync.WaitGroup
	for _, id := range p.participants {
		wg.Add(1)
		go func(id party.ID) {
			pl := pool.NewPool(0)
			defer pl.TearDown()
			conf, err := cmpKeygen(id, p.participants, p.network, p.threshold, &wg, pl)
			if err != nil {
				return
			}
			if id == p.currentId {
				doneChan <- conf
			}
			p.mtx.Lock()
			p.configs[conf.ID] = conf
			p.mtx.Unlock()
		}(id)
	}
	wg.Wait()
	conf := <-doneChan
	return wallet.NewWalletImpl(conf), nil
}

// Config returns the configuration of this wallet.
func (w *MPCProtocol) Config() *cmp.Config {
	return w.configs[w.currentId]
}

// GetSigners returns the list of signers for the given message.
func (w *MPCProtocol) GetSigners() party.IDSlice {
	signers := party.IDSlice([]party.ID{"vault", "current"})
	// signers := w.Configs[w.ID].PartyIDs()[:w.Threshold+1]
	if !signers.Contains(w.currentId) {
		w.network.Quit(w.currentId)
		return nil
	}
	return party.NewIDSlice(signers)
}

// Refreshes all shares of an existing ECDSA private key.
func (w *MPCProtocol) Refresh(pl *pool.Pool) (*cmp.Config, error) {
	hRefresh, err := protocol.NewMultiHandler(cmp.Refresh(w.configs[w.currentId], pl), nil)
	if err != nil {
		return nil, err
	}
	handlerLoop(w.currentId, hRefresh, w.network)

	r, err := hRefresh.Result()
	if err != nil {
		return nil, err
	}
	handlerLoop(w.currentId, hRefresh, w.network)
	return r.(*cmp.Config), nil
}

// Generates an ECDSA signature for messageHash.
func (w *MPCProtocol) Sign(m []byte) (*ecdsa.Signature, error) {
	var wg sync.WaitGroup
	signers := w.GetSigners()
	net := NewNetwork(signers)

	var (
		sig *ecdsa.Signature
		err error
	)

	for _, id := range signers {
		wg.Add(1)
		go func(id party.ID) {
			pl := pool.NewPool(0)
			defer pl.TearDown()
			if sig, err = cmpSign(w.configs[id], m, signers, net, &wg, pl); err != nil {
				return
			}
		}(id)
	}
	wg.Wait()
	return sig, err
}
