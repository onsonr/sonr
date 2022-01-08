package channel

import (
	"github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

// handleEvents method listens to Pubsub Events for room
func (b *channel) handleEvents() {
	// Loop Events
	for {
		// Get next event
		event, err := b.handler.NextPeerEvent(b.ctx)
		if err != nil {
			return
		}

		// Check Event and Validate not User
		switch event.Type {
		case pubsub.PeerJoin:
			event := b.NewSyncEvent()
			err = event.Publish(b.ctx, b.topic)
			if err != nil {
				logger.Error(err)
				continue
			}
		default:
			continue
		}
	}
}

// handleMessages method listens to Pubsub Messages for room
func (b *channel) handleMessages() {
	// Loop Messages
	for {
		// Get next message
		msg, err := b.sub.Next(b.ctx)
		if err != nil {
			return
		}

		// Check Message and Validate not User
		e, err := eventFromMsg(msg, b.n.HostID())
		if err != nil {
			continue
		}

		// Handle Event
		err = b.store.Handle(e, b)
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
			logger.Debugf("Closing Beam (%s)", b.id.Prefix())
			b.handler.Cancel()
			b.sub.Cancel()
			if err := b.topic.Close(); err != nil {
				logger.Errorf("%s - Failed to Close Beam", err)
			}
			return
		}
	}
}

// isEventJoin Checks if PeerEvent is Join and NOT User
func isEventJoin(ev pubsub.PeerEvent, selfID peer.ID) bool {
	return ev.Type == pubsub.PeerJoin && ev.Peer != selfID
}

// isEventExit Checks if PeerEvent is Exit and NOT User
func isEventExit(ev pubsub.PeerEvent) bool {
	return ev.Type == pubsub.PeerLeave
}

// eventFromMsg converts a message to an event
func eventFromMsg(msg *pubsub.Message, selfID peer.ID) (*Event, error) {
	// Check Message
	if msg.ReceivedFrom == selfID {
		return nil, errors.Wrap(ErrInvalidMessage, "Same Peer as Node")
	}

	// Check Message Data
	if len(msg.Data) == 0 {
		return nil, errors.Wrap(ErrInvalidMessage, "Invalid Data Length")
	}

	// Unmarshal Message Data
	e := &Event{}
	err := proto.Unmarshal(msg.Data, e)
	if err != nil {
		logger.Errorf("failed to Unmarshal Event from pubsub.Message")
		return nil, err
	}
	return e, nil
}
