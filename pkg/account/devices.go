package account

import (
	"context"
	"errors"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/sonr-io/core/pkg/data"
	"github.com/sonr-io/core/tools/logger"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

type GetRoomFunc func() *data.Room

// AddDevice adds a Device Peer to the Room
func (tm *userLinker) AddDevice(peerID peer.ID, d *data.Device) {
	// Add Device to Map
	tm.activeDevices[peerID] = d
}

// FindPeerreturns Pointer to Peer.ID and Peer
func (tm *userLinker) FindPeer(q string) (peer.ID, error) {
	// Iterate through Room Peers
	for _, id := range tm.topic.ListPeers() {
		// If Found Match
		if id.String() == q {
			return id, nil
		}
	}
	return "", errors.New("Peer ID was not found in room")
}

// RemoveDevice removes a Device Peer from the Room
func (tm *userLinker) RemoveDevice(peerID peer.ID) {
	// Add Device to Map
	delete(tm.activeDevices, peerID)
}

// Sync Publishes message to User Device room
func (tm *userLinker) Sync(msg *data.SyncEvent) error {
	if tm.room.IsDevices() {
		// Convert Event to Proto Binary
		bytes, err := proto.Marshal(msg)
		if err != nil {
			logger.Error("Marshal Sync Event", zap.Error(err))
			return err
		}

		// Publish to Room
		err = tm.topic.Publish(tm.ctx, bytes)
		if err != nil {
			logger.Error("Failed to Publish Topic Message", zap.Error(err))
			return err
		}

	}
	return nil
}

// HasPeer Method Checks if Peer ID String is Subscribed to Room
func (tm *userLinker) HasPeer(q string) bool {
	// Iterate through PubSub in room
	for _, id := range tm.topic.ListPeers() {
		// If Found Match
		if id.String() == q {
			return true
		}
	}
	return false
}

// HasPeerID Method Checks if Peer ID is Subscribed to Room
func (tm *userLinker) HasPeerID(q peer.ID) bool {
	// Iterate through PubSub in room
	for _, id := range tm.topic.ListPeers() {
		// If Found Match
		if id == q {
			return true
		}
	}
	return false
}

// handleTopicEvents listens to Pubsub Events for room
func (rm *userLinker) handleTopicEvents(ctx context.Context) {
	// Loop Events
	for {
		// Get next event
		event, err := rm.eventHandler.NextPeerEvent(ctx)
		if err != nil {
			logger.Error("Failed to Get next peer event.", zap.Error(err))
			rm.eventHandler.Cancel()
			return
		}

		// Check Event and Validate not User
		if rm.isEventJoin(event) {
			err = rm.Verify(event.Peer)
			if err != nil {
				logger.Error("Failed to Verify peer.", zap.Error(err))
				continue
			}
		} else if rm.isEventExit(event) {

		}
		data.GetState().NeedsWait()
	}
}

// handleTopicMessages listens for messages in room subscription
func (rm *userLinker) handleTopicMessages(ctx context.Context) {
	for {
		// Get next msg from pub/sub
		msg, err := rm.subscription.Next(ctx)
		if err != nil {
			logger.Error("Failed to get next subcription message.", zap.Error(err))
			return
		}

		// Only forward messages delivered by others
		if rm.isValidMessage(msg) {
			// Unmarshal RoomEvent
			m := &data.SyncEvent{}
			err = proto.Unmarshal(msg.Data, m)
			if err != nil {
				logger.Error("Failed to Unmarshal Sync Event", zap.Error(err))
				continue
			}

			// Check Peer is Online, if not ignore
			rm.OnSyncEvent(m)
		}
		data.GetState().NeedsWait()
	}
}

// Returns RoomData Data instance
func (tm *userLinker) Room() *data.Room {
	return tm.room
}

func (al *userLinker) OnSyncEvent(*data.SyncEvent) {

}
