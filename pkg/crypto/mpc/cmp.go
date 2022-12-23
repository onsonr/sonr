package mpc

// import (
// 	"sync"

// 	"github.com/sonr-hq/sonr/pkg/node"

// 	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
// 	"github.com/taurusgroup/multi-party-sig/pkg/pool"
// 	"github.com/taurusgroup/multi-party-sig/pkg/protocol"
// 	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
// )

// // CMPKeygen is a function that creates a new network, starts the network, creates a new keygen session,
// func CMPKeygen(network *node.Network, wg *sync.WaitGroup) (*cmp.Config, error) {
// 	defer wg.Done()
// 	// Create a new network.
// 	go network.Start()

// 	// Create a new keygen session.
// 	pl := pool.NewPool(0)
// 	handler, err := protocol.NewMultiHandler(cmp.Keygen(curve.Secp256k1{}, network, network.PartyIDs(), 1, pl), nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	HandlerLoop(handler, network)

// 	// Wait for the network to finish.
// 	r, err := handler.Result()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return r.(*cmp.Config), nil
// }
