package algorithm

import (
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/sonrhq/core/internal/crypto"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/pkg/protocol"
)

// HandleNetworkProtocol is a helper function that loops over all the parties and calls the given handler.
func HandleNetworkProtocol(id party.ID, h protocol.Handler, network crypto.Network) {
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

// EnsureSelfIDInGroup ensures that the given self ID is in the given group.
func EnsureSelfIDInGroup(selfID party.ID, group []party.ID) []party.ID {
	if len(selfID) == 0 {
		selfID = party.ID(peer.ID("user1"))
	}
	for _, id := range group {
		if id == selfID {
			return group
		}
	}
	return append(group, selfID)
}

// It converts a peer.ID to a party.ID
func PeerIdToPartyId(id crypto.PeerID) crypto.PartyID {
	return party.ID(id)
}

// It converts a party ID to a peer ID
func PartyIdToPeerId(id crypto.PartyID) crypto.PeerID {
	return peer.ID(id)
}

// It converts a list of peer IDs to a list of party IDs
func PeerIdListToPartyIdList(ids []crypto.PeerID) []crypto.PartyID {
	partyIds := make([]party.ID, len(ids))
	for i, id := range ids {
		partyIds[i] = PeerIdToPartyId(id)
	}
	return partyIds
}

// It converts a list of party IDs to a list of peer IDs
func PartyIdListToPeerIdList(ids []crypto.PartyID) []crypto.PeerID {
	peerIds := make([]crypto.PeerID, len(ids))
	for i, id := range ids {
		peerIds[i] = PartyIdToPeerId(id)
	}
	return peerIds
}
