package channel

import (
	"github.com/libp2p/go-libp2p-core/peer"
	ps "github.com/libp2p/go-libp2p-pubsub"
	"github.com/pkg/errors"
	v1 "go.buf.build/grpc/go/sonr-io/core/host/channel/v1"
	"google.golang.org/protobuf/proto"
)

// handleStoreEvents method listens to Pubsub Events for room
func (b *channel) handleChannelEvents() {
	// Loop Events
	for {
		// Get next event
		event, err := b.messagesHandler.NextPeerEvent(b.ctx)
		if err != nil {
			return
		}

		// Check Event and Validate not User
		switch event.Type {
		case ps.PeerJoin:
			// event := b.NewSyncEvent()
			// err = PublishEvent(b.ctx, b.topic, event)
			// if err != nil {
			// 	logger.Error(err)
			// 	continue
			// }
		default:
			continue
		}
	}
}

// handleStoreMessages method listens to Pubsub Messages for room
func (b *channel) handleChannelMessages() {
	// Loop Messages
	for {
		// Get next message
		buf, err := b.messagesSub.Next(b.ctx)
		if err != nil {
			return
		}

		// Unmarshal Message Data
		msg := &v1.ChannelMessage{}
		err = proto.Unmarshal(buf.Data, msg)
		if err != nil {
			logger.Errorf("failed to Unmarshal Message from pubsub.Message")
			return
		}

		// Push Message to Channel
		b.messages <- msg
	}
}

// handleStoreEvents method listens to Pubsub Events for room
func (b *channel) handleStoreEvents() {
	// Loop Events
	for {
		// Get next event
		event, err := b.storeEventsHandler.NextPeerEvent(b.ctx)
		if err != nil {
			return
		}

		// Check Event and Validate not User
		switch event.Type {
		case ps.PeerJoin:
			// event := b.NewSyncEvent()
			// err = PublishEvent(b.ctx, b.topic, event)
			// if err != nil {
			// 	logger.Error(err)
			// 	continue
			// }
		default:
			continue
		}
	}
}

// handleStoreMessages method listens to Pubsub Messages for room
func (b *channel) handleStoreMessages() {
	// Loop Messages
	for {
		// Get next message
		msg, err := b.storeEventsSub.Next(b.ctx)
		if err != nil {
			return
		}

		// Check Message and Validate not User
		e, err := eventFromMsg(msg, b.n.HostID())
		if err != nil {
			continue
		}

		// Handle Event
		err = HandleStore(b.store, e, b)
		if err != nil {
			logger.Error(err)
			continue
		}
	}
}

// serve handles the serving of the beam
func (b *channel) serve() {
	for {
		select {
		case <-b.ctx.Done():
			logger.Debugf("Closing Beam (%s)", b.label)
			b.storeEventsHandler.Cancel()
			b.storeEventsSub.Cancel()
			if err := b.storeEventsTopic.Close(); err != nil {
				logger.Errorf("%s - Failed to Close Beam", err)
			}
			return
		}
	}
}

// isEventJoin Checks if PeerEvent is Join and NOT User
func isEventJoin(ev ps.PeerEvent, selfID peer.ID) bool {
	return ev.Type == ps.PeerJoin && ev.Peer != selfID
}

// isEventExit Checks if PeerEvent is Exit and NOT User
func isEventExit(ev ps.PeerEvent) bool {
	return ev.Type == ps.PeerLeave
}

// eventFromMsg converts a message to an event
func eventFromMsg(msg *ps.Message, selfID peer.ID) (*v1.ChannelEvent, error) {
	// Check Message
	if msg.ReceivedFrom == selfID {
		return nil, errors.Wrap(ErrInvalidMessage, "Same Peer as Node")
	}

	// Check Message Data
	if len(msg.Data) == 0 {
		return nil, errors.Wrap(ErrInvalidMessage, "Invalid Data Length")
	}

	// Unmarshal Message Data
	e := &v1.ChannelEvent{}
	err := proto.Unmarshal(msg.Data, e)
	if err != nil {
		logger.Errorf("failed to Unmarshal Event from pubsub.Message")
		return nil, err
	}
	return e, nil
}
