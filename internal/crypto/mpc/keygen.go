package mpc

import (
	"sync"

	"github.com/sonrhq/core/internal/crypto"
	"github.com/sonrhq/core/internal/crypto/mpc/algorithm"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/pkg/pool"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

// Keygen Generates a new ECDSA private key shared among all the given participants.
func Keygen(current crypto.PartyID, option ...KeygenOption) ([]*cmp.Config, error) {
	opts := defaultKeygenOpts(current)
	opts.Apply(option...)
	net := opts.getOfflineNetwork()

	var mtx sync.Mutex
	var wg sync.WaitGroup
	confs := make([]*cmp.Config, 0)
	for _, id := range net.Ls() {
		wg.Add(1)
		go func(id party.ID) {
			pl := pool.NewPool(0)
			defer pl.TearDown()
			conf, err := algorithm.CmpKeygen(id, net.Ls(), net, opts.Threshold, &wg, pl)
			if err != nil {
				return
			}
			mtx.Lock()
			if opts.Handlers != nil {
				manageHandlers(opts.Handlers, conf)
			}
			confs = append(confs, conf)
			mtx.Unlock()
		}(id)
	}
	wg.Wait()
	return confs, nil
}

func manageHandlers(handlers []OnConfigGenerated, conf *cmp.Config) error {
	for _, h := range handlers {
		if h == nil {
			continue
		}
		if err := h(conf); err != nil {
			return err
		}
	}
	return nil
}
