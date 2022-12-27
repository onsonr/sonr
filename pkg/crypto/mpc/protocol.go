package mpc

import (
	"errors"
	"fmt"
	"sync"

	"github.com/sonr-hq/sonr/pkg/common"
	"github.com/sonr-hq/sonr/pkg/wallet"
	"github.com/taurusgroup/multi-party-sig/pkg/ecdsa"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/pkg/pool"
	"github.com/taurusgroup/multi-party-sig/pkg/protocol"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

// `MPCProtocol` is a struct that contains a `pool` of `*cmp.Config`s, a `currentId` of type
// `party.ID`, a `participants` slice of type `party.IDSlice`, a `pubKey` of type `[]byte`, a `configs`
// map of type `map[party.ID]*cmp.Config`, a `mtx` of type `sync.Mutex`, a `threshold` of type `int`,
// and a `callback` of type `common.NodeCallback`.
// @property pool - the pool of parties that we can communicate with
// @property currentId - The ID of the current party.
// @property participants - A list of all the parties involved in the protocol.
// @property {[]byte} pubKey - the public key of the current party
// @property configs - a map of party IDs to the configuration of the party.
// @property mtx - a mutex to protect the configs map
// @property {int} threshold - The minimum number of parties required to complete the protocol.
// @property callback - This is a callback function that will be called when the protocol is finished.
type MPCProtocol struct {
	currentId party.ID

	pubKey    []byte
	configs   map[party.ID]*cmp.Config
	mtx       sync.Mutex
	threshold int
	callback  common.NodeCallback
}

// Initialize takes in a list of `WalletOption`s and returns a `MPCProtocol` object
func Initialize(options ...Option) *MPCProtocol {
	opt := defaultConfig()
	w := opt.Apply(options...)
	return w
}

// Keygen Generates a new ECDSA private key shared among all the given participants.
func (p *MPCProtocol) Keygen(current party.ID, net Network) (common.WalletShare, error) {
	p.currentId = current
	if len(p.configs) > 0 {
		return nil, fmt.Errorf("wallet already initialized")
	}
	p.currentId = current
	var wg sync.WaitGroup
	for _, id := range net.Ls() {
		wg.Add(1)
		go func(id party.ID) {
			pl := pool.NewPool(0)
			defer pl.TearDown()
			conf, err := cmpKeygen(id, net.Ls(), net, p.threshold, &wg, pl)
			if err != nil {
				return
			}
			p.mtx.Lock()
			p.configs[conf.ID] = conf
			p.mtx.Unlock()
		}(id)
	}
	wg.Wait()
	// conf := <-doneChan
	return wallet.NewWalletImpl("snr", p.configs[p.currentId]), nil
}

// Refreshes all shares of an existing ECDSA private key.
func (w *MPCProtocol) Refresh(current party.ID, net Network) (common.WalletShare, error) {
	w.currentId = current
	var wg sync.WaitGroup
	for _, id := range net.Ls() {
		wg.Add(1)
		go func(id party.ID) {
			pl := pool.NewPool(0)
			conf, err := cmpRefresh(w.configs[id], net, &wg, pl)
			if err != nil {
				return
			}

			w.mtx.Lock()
			w.configs[conf.ID] = conf
			w.mtx.Unlock()
		}(id)
	}
	wg.Wait()
	return wallet.NewWalletImpl("snr", w.configs[w.currentId]), nil
}

// Sign Generates an ECDSA signature for messageHash.
func (w *MPCProtocol) Sign(current party.ID, m []byte, net Network) (*ecdsa.Signature, error) {
	w.currentId = current
	doneChan := make(chan *ecdsa.Signature)
	var wg sync.WaitGroup
	for _, id := range net.Ls() {
		wg.Add(1)
		go func(id party.ID) {
			pl := pool.NewPool(0)
			defer pl.TearDown()
			sig, err := cmpSign(w.configs[id], m, net.Ls(), net, &wg, pl)
			if err != nil {
				return
			}
			if id == w.currentId {
				doneChan <- sig
			}
		}(id)
	}
	wg.Wait()
	return <-doneChan, nil
}

//
// Private methods
//

// It creates a new handler for the keygen protocol, runs the handler loop, and returns the result
func cmpKeygen(id party.ID, ids party.IDSlice, n Network, threshold int, wg *sync.WaitGroup, pl *pool.Pool) (*cmp.Config, error) {
	defer wg.Done()
	h, err := protocol.NewMultiHandler(cmp.Keygen(curve.Secp256k1{}, id, ids, threshold, pl), nil)
	if err != nil {
		return nil, err
	}

	handlerLoop(id, h, n)
	r, err := h.Result()
	if err != nil {
		return nil, err
	}
	conf := r.(*cmp.Config)
	return conf, nil
}

// It creates a new handler for the refresh protocol, runs the handler loop, and returns the result
func cmpRefresh(c *cmp.Config, n Network, wg *sync.WaitGroup, pl *pool.Pool) (*cmp.Config, error) {
	defer wg.Done()
	h, err := protocol.NewMultiHandler(cmp.Refresh(c, pl), nil)
	if err != nil {
		return nil, err
	}

	handlerLoop(c.ID, h, n)
	r, err := h.Result()
	if err != nil {
		return nil, err
	}
	conf := r.(*cmp.Config)
	return conf, nil
}

// It creates a new `protocol.MultiHandler` for the `cmp.Sign` protocol, and then runs the handler loop
func cmpSign(c *cmp.Config, m []byte, signers party.IDSlice, n Network, wg *sync.WaitGroup, pl *pool.Pool) (*ecdsa.Signature, error) {
	defer wg.Done()
	h, err := protocol.NewMultiHandler(cmp.Sign(c, signers, m, pl), nil)
	if err != nil {
		return nil, err
	}
	handlerLoop(c.ID, h, n)

	r, err := h.Result()
	if err != nil {
		return nil, err
	}
	sig := r.(*ecdsa.Signature)
	if !sig.Verify(c.PublicPoint(), m) {
		return nil, errors.New("failed to verify cmp signature")
	}
	return sig, nil
}

// handlerLoop is a helper function that loops over all the parties and calls the given handler.
func handlerLoop(id party.ID, h protocol.Handler, network Network) {
	for {
		select {

		// outgoing messages
		case msg, ok := <-h.Listen():
			if !ok {
				<-network.Done(id)
				// the channel was closed, indicating that the protocol is done executing.
				return
			}
			go network.Send(msg)

			// incoming messages
		case msg := <-network.Next(id):
			h.Accept(msg)
		}
	}
}
