package channel

import (
	"context"
	"time"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	channelV1 "github.com/sonr-io/core/protocols/channel/v1"
	"google.golang.org/protobuf/proto"
)

// NewStore creates a new store
func NewStore(opts *options) *channelV1.Store {
	// Create a new store
	return &channelV1.Store{
		Data:     make(map[string]*channelV1.StoreEntry),
		Capacity: int32(opts.capacity),
		Modified: time.Now().Unix(),
		Ttl:      opts.ttl.Milliseconds(),
	}
}

// Delete deletes an entry from the store and publishes an event
func DeleteStoreKey(s *channelV1.Store, key string, b *channel) error {
	// Fetch the entry
	entry := s.Data[key]
	if entry == nil {
		return ErrNotFound
	}

	// Check if the entry is owned by this node
	if entry.Peer != b.n.HostID().String() {
		return ErrNotOwner
	}

	// Delete the entry
	delete(s.Data, key)
	s.Modified = time.Now().Unix()

	// Create Delete Event
	event := b.NewDeleteEvent(key)
	return PublishEvent(b.ctx, b.topic, event)
}

// Get returns the value of the entry
func GetKey(s *channelV1.Store, key string) ([]byte, error) {
	entry := s.Data[key]
	if entry == nil {
		return nil, ErrNotFound
	}
	return entry.Value, nil
}

// Handle checks the event type and handles it with the store
func HandleStore(s *channelV1.Store, e *channelV1.Event, b *channel) error {
	// Check if the event is valid
	if b.n.HostID().String() == e.Peer {
		return nil
	}

	switch e.Type {
	case channelV1.EventType_EVENT_TYPE_DELETE:
		delete(s.Data, e.Entry.Key)
		return nil
	case channelV1.EventType_EVENT_TYPE_SYNC:
		if e.Store != nil {
			if s.Modified > e.Store.Modified && len(s.Data) < int(e.Store.Capacity) {
				s.Data = e.Store.Data
				s.Modified = e.Store.Modified
				logger.Debug("Store - Updated store to pushed earlier version")
			}
		}
		return nil
	case channelV1.EventType_EVENT_TYPE_PUT:
		if e.Entry != nil {
			s.Data[e.Entry.Key] = e.Entry
			s.Modified = time.Now().Unix()
			logger.Debug("Store - Added new Store Entry")
		}
		return nil
	case channelV1.EventType_EVENT_TYPE_SET:
		if e.Entry != nil {
			s.Data[e.Entry.Key] = e.Entry
			s.Modified = time.Now().Unix()
			logger.Debug("Store - Set Updated Store Entry")
		}
		return nil
	}
	return nil
}

// Put puts an entry into the store and publishes an event
func PutStoreKey(s *channelV1.Store, key string, value []byte, b *channel) error {
	// Fetch the entry
	entry := s.Data[key]
	if entry == nil {
		// Create new entry with Event
		event, entry := b.NewPutEvent(key, value)
		s.Data[key] = entry
		s.Modified = time.Now().Unix()
		return PublishEvent(b.ctx, b.topic, event)
	}

	// Get existing entry and update it
	event, err := SetStoreEntry(entry, value, b.n.HostID().String())
	if err != nil {
		return err
	}
	s.Modified = time.Now().Unix()
	return PublishEvent(b.ctx, b.topic, event)
}

// Set updates the entry in the store and publishes an event
func SetStoreEntry(se *channelV1.StoreEntry, value []byte, selfID string) (*channelV1.Event, error) {
	if se.Peer != selfID {
		return nil, ErrNotOwner
	}
	se.Value = value
	se.Modified = time.Now().Unix()
	return &channelV1.Event{
		Type:  channelV1.EventType_EVENT_TYPE_SET,
		Entry: se,
		Peer:  se.GetPeer(),
	}, nil
}

// NewPutEvent creates a new put event
func (b *channel) NewPutEvent(key string, value []byte) (*channelV1.Event, *channelV1.StoreEntry) {
	entry := &channelV1.StoreEntry{
		Key:      key,
		Value:    value,
		Peer:     b.n.HostID().String(),
		Created:  time.Now().Unix(),
		Modified: time.Now().Unix(),
	}
	event := &channelV1.Event{
		Type:  channelV1.EventType_EVENT_TYPE_PUT,
		Peer:  b.n.HostID().String(),
		Entry: entry,
	}
	return event, entry
}

// NewSyncEvent creates a new sync event
func (b *channel) NewSyncEvent() *channelV1.Event {
	return &channelV1.Event{
		Type:  channelV1.EventType_EVENT_TYPE_SYNC,
		Peer:  b.n.HostID().String(),
		Store: b.store,
	}
}

// NewDeleteEvent creates a new delete event
func (b *channel) NewDeleteEvent(key string) *channelV1.Event {
	entry := &channelV1.StoreEntry{
		Key:      key,
		Peer:     b.n.HostID().String(),
		Modified: time.Now().Unix(),
	}
	event := &channelV1.Event{
		Type:  channelV1.EventType_EVENT_TYPE_DELETE,
		Peer:  b.n.HostID().String(),
		Entry: entry,
	}
	return event
}

// Publish publishes the event to the topic
func PublishEvent(ctx context.Context, t *pubsub.Topic, e *channelV1.Event) error {
	buf, err := proto.Marshal(e)
	if err != nil {
		return err
	}
	return t.Publish(ctx, buf)
}
