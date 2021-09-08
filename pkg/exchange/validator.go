package exchange

import (
	"bytes"

	peer "github.com/libp2p/go-libp2p-core/peer"
	ps "github.com/libp2p/go-libp2p-pubsub"
	record "github.com/libp2p/go-libp2p-record"
)

type ExchangeValidator struct {
}

func (ExchangeValidator) Validate(key string, value []byte) error {
	ns, k, err := record.SplitKey(key)
	if err != nil {
		return err
	}
	if ns != "store" {
		return record.ErrInvalidRecordType
	}
	if !bytes.Contains(value, []byte(k)) {
		return record.ErrInvalidRecordType
	}
	if bytes.Contains(value, []byte("invalid")) {
		return record.ErrInvalidRecordType
	}
	return nil
}

func (ExchangeValidator) Select(key string, vals [][]byte) (int, error) {
	if len(vals) == 0 {
		return -1, record.ErrInvalidRecordType
	}
	var best []byte
	idx := 0
	for i, val := range vals {
		if bytes.Compare(best, val) < 0 {
			best = val
			idx = i
		}
	}
	return idx, nil
}

// HasPeer Method Checks if Peer ID is Subscribed to Room
func (tm *ExchangeProtocol) HasPeerID(q peer.ID) bool {
	// Iterate through PubSub in room
	for _, id := range tm.topic.ListPeers() {
		// If Found Match
		if id == q {
			return true
		}
	}
	return false
}

// Check if PeerEvent is Join and NOT User
func (tm *ExchangeProtocol) isEventJoin(ev ps.PeerEvent) bool {
	return ev.Type == ps.PeerJoin && ev.Peer != tm.host.ID()
}

// Check if PeerEvent is Exit and NOT User
func (tm *ExchangeProtocol) isEventExit(ev ps.PeerEvent) bool {
	return ev.Type == ps.PeerLeave && ev.Peer != tm.host.ID()
}

// Check if Message is NOT from User
func (tm *ExchangeProtocol) isValidMessage(msg *ps.Message) bool {
	return tm.host.ID() != msg.ReceivedFrom && tm.HasPeerID(msg.ReceivedFrom)
}
