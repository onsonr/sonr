package mpc

import (
	"fmt"

	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/pkg/protocol"
)

// Algorithm is an enum that represents the different MPC algorithms that are supported.
type Algorithm int

const (
	// CMP is the CMP algorithm.
	CMP Algorithm = iota
	// Doerner is the Doerner algorithm.
	Doerner
)

// Name returns the name of the MPC algorithm.
func (a Algorithm) Name() string {
	switch a {
	case CMP:
		return "cmp"
	case Doerner:
		return "doerner"
	default:
		panic("unknown algorithm")
	}
}

// TopicKey takes a method string and returns the topic key for the method.
func (a Algorithm) TopicKey(id party.ID, method string) string {
	return fmt.Sprintf("%s/mpc/%s-%s", string(id), a.Name(), method)
}

func peerIdToPartyId(id peer.ID) party.ID {
	return party.ID(id)
}

func partyIdToPeerId(id party.ID) peer.ID {
	return peer.ID(id)
}

func peerIdListToPartyIdList(ids []peer.ID) []party.ID {
	partyIds := make([]party.ID, len(ids))
	for i, id := range ids {
		partyIds[i] = peerIdToPartyId(id)
	}
	return partyIds
}

func partyIdListToPeerIdList(ids []party.ID) []peer.ID {
	peerIds := make([]peer.ID, len(ids))
	for i, id := range ids {
		peerIds[i] = partyIdToPeerId(id)
	}
	return peerIds
}

// HandlerLoop blocks until the handler has finished. The result of the execution is given by Handler.Result().
func HandlerLoop(h protocol.Handler, network *Network) {
	for {
		select {
		// outgoing messages
		case msg, ok := <-h.Listen():
			if !ok {
				<-network.Done()
				// the channel was closed, indicating that the protocol is done executing.
				return
			}
			go network.Send(msg)

		// incoming messages
		case msg := <-network.Next():
			h.Accept(msg)
		}
	}
}
