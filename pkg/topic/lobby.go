package topic

import (
	"log"

	"github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	md "github.com/sonr-io/core/internal/models"
	st "github.com/sonr-io/core/pkg/state"
	"google.golang.org/protobuf/proto"
)

type Lobby struct {
	OLC   string
	Size  int32
	Count int32
	Peers map[string]*md.Peer

	// Private Properties
	callback TopicHandler
	user     *md.Peer
}

// ^ Returns as Lobby Buffer ^
func (l *Lobby) Buffer() []byte {
	bytes, err := proto.Marshal(&md.Lobby{
		Olc:   l.OLC,
		Size:  l.Size,
		Count: l.Count,
		Peers: l.Peers,
	})
	if err != nil {
		log.Println(err)
		return nil
	}
	return bytes
}

// ^ Add/Update Peer in Lobby ^
func (l *Lobby) Add(peer *md.Peer) {
	// Update Peer with new data
	l.Peers[peer.Id.Peer] = peer
	l.Count = int32(len(l.Peers))
	l.Size = int32(len(l.Peers)) + 1 // Account for User
	l.Refresh()
}

// ^ Add/Update Peer in Lobby without Callback ^
func (l *Lobby) AddWithoutRefresh(peer *md.Peer) {
	// Update Peer with new data
	l.Peers[peer.Id.Peer] = peer
	l.Count = int32(len(l.Peers))
	l.Size = int32(len(l.Peers)) + 1 // Account for User
}

// ^ Delete Peer from Lobby ^
func (l *Lobby) Delete(id peer.ID) {
	// Update Peer with new data
	delete(l.Peers, id.String())
	l.Count = int32(len(l.Peers))
	l.Size = int32(len(l.Peers)) + 1 // Account for User
	l.Refresh()
}

// ^ Send Updated Lobby ^
func (l *Lobby) Refresh() {
	l.callback.OnRefresh(&md.Lobby{
		Olc:   l.OLC,
		Size:  l.Size,
		Count: l.Count,
		Peers: l.Peers,
	})
}

// ^ Sync Between Remote Peers Lobby ^
func (l *Lobby) Sync(ref *md.Lobby, remotePeer *md.Peer) {
	// Validate Lobbies are Different
	if l.Count != ref.Count {
		// Iterate Over List
		for id, peer := range ref.Peers {
			// Add all Peers NOT User
			if id != l.user.Id.Peer {
				l.AddWithoutRefresh(peer)
			}
		}
	}

	// Add Synced Peer to Lobby
	l.Add(remotePeer)
}

// ^ handleTopicEvents: listens to Pubsub Events for topic  ^
func (tm *TopicManager) handleTopicEvents() {
	// @ Loop Events
	for {
		// Get next event
		lobEvent, err := tm.handler.NextPeerEvent(tm.ctx)
		if err != nil {
			tm.handler.Cancel()
			return
		}

		if lobEvent.Type == pubsub.PeerJoin {
			p := tm.callback.GetPeer()
			buf, err := proto.Marshal(p)
			if err != nil {
				continue
			}
			err = tm.Exchange(lobEvent.Peer, buf)
			if err != nil {
				continue
			}
		}

		if lobEvent.Type == pubsub.PeerLeave {
			tm.Lobby.Delete(lobEvent.Peer)
		}

		st.GetState().NeedsWait()
	}
}

// ^ handleTopicMessages: listens for messages on pubsub topic subscription ^
func (tm *TopicManager) handleTopicMessages() {
	for {
		// Get next msg from pub/sub
		msg, err := tm.subscription.Next(tm.ctx)
		if err != nil {
			return
		}

		// Only forward messages delivered by others
		if msg.ReceivedFrom.String() == tm.callback.GetPeer().Id.Peer {
			continue
		}

		// Construct message
		m := &md.LobbyEvent{}
		err = proto.Unmarshal(msg.Data, m)
		if err != nil {
			continue
		}

		// Validate Peer in Lobby
		if tm.HasPeer(m.Id) {
			tm.Messages <- m
		}
		st.GetState().NeedsWait()
	}
}

// ^ processTopicMessages: pulls messages from channel that have been handled ^
func (tm *TopicManager) processTopicMessages() {
	for {
		select {
		// @ when we receive a message from the lobby room
		case m := <-tm.Messages:
			// Update Circle by event
			if m.Event == md.LobbyEvent_UPDATE {
				// Update Peer Data
				tm.Lobby.Add(m.From)
			} else if m.Event == md.LobbyEvent_MESSAGE {
				// Check is Message For Self
				if m.To == tm.callback.GetPeer().Id.Peer {
					// Call Event
					tm.callback.OnEvent(m)
				}
			}
		case <-tm.ctx.Done():
			return
		}
		st.GetState().NeedsWait()
	}
}
