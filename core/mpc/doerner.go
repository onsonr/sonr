package mpc

import (
	"sync"

	"github.com/sonr-hq/sonr/internal/node"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/pkg/pool"
	"github.com/taurusgroup/multi-party-sig/pkg/protocol"
	"github.com/taurusgroup/multi-party-sig/protocols/doerner"
)

func DoernerKeygen(id party.ID, ids party.IDSlice, topicHandler node.TopicHandler, threshold int, wg *sync.WaitGroup, pl *pool.Pool) (*doerner.ConfigReceiver, error) {
	tph, err := protocol.NewTwoPartyHandler(doerner.Keygen(curve.Secp256k1{}, false, id, ids[0], pl), []byte(topicHandler.Name()), true)
	if err != nil {
		return nil, err
	}
	// handlerLoopTopic(id, tph, topicHandler)
	r, err := tph.Result()
	if err != nil {
		return nil, err
	}
	return r.(*doerner.ConfigReceiver), nil
}
