package channel

import (
	"context"
	"time"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	v1 "go.buf.build/grpc/go/sonr-io/core/host/channel/v1"
	"google.golang.org/protobuf/proto"
)

// NewStore creates a new store
func NewStore(opts *options) *v1.ChannelStore {
	// Create a new store
	return &v1.ChannelStore{
		Entries:  make(map[string]*v1.ChannelStoreRecord),
		Capacity: int32(opts.capacity),
		Modified: time.Now().Unix(),
		Ttl:      opts.ttl.Milliseconds(),
	}
}

// Delete deletes an entry from the store and publishes an event
func DeleteStoreKey(s *v1.ChannelStore, key string, b *channel) error {
	// Fetch the entry
	entry := s.Entries[key]
	if entry == nil {
		return ErrNotFound
	}

	// Check if the entry is owned by this node
	if entry.Owner != b.n.HostID().String() {
		return ErrNotOwner
	}

	// Delete the entry
	delete(s.Entries, key)
	s.Modified = time.Now().Unix()

	// Create Delete Event
	event := b.NewDeleteEvent(key)
	return PublishEvent(b.ctx, b.storeEventsTopic, event)
}

// Get returns the value of the entry
func GetKey(s *v1.ChannelStore, key string) ([]byte, error) {
	entry := s.Entries[key]
	if entry == nil {
		return nil, ErrNotFound
	}
	return entry.Value, nil
}

// Handle checks the event type and handles it with the store
func HandleStore(s *v1.ChannelStore, e *v1.ChannelEvent, b *channel) error {
	// Check if the event is valid
	if b.n.HostID().String() == e.Owner {
		return nil
	}

	switch e.Type {
	case v1.ChannelEventType_CHANNEL_EVENT_TYPE_DELETE:
		delete(s.Entries, e.Record.Key)
		return nil
	case v1.ChannelEventType_CHANNEL_EVENT_TYPE_SET:
		if e.Record != nil {
			s.Entries[e.Record.Key] = e.Record
			s.Modified = time.Now().Unix()
			logger.Debug("Store - Set Updated Store Entry")
		}
		return nil
	}
	return nil
}

// Put puts an entry into the store and publishes an event
func PutStoreKey(s *v1.ChannelStore, key string, value []byte, b *channel) error {
	// Fetch the entry
	entry := s.Entries[key]
	if entry == nil {
		// Create new entry with Event
		event, entry := b.NewPutEvent(key, value)
		s.Entries[key] = entry
		s.Modified = time.Now().Unix()
		return PublishEvent(b.ctx, b.storeEventsTopic, event)
	}

	// Get existing entry and update it
	event, err := SetStoreEntry(entry, value, b.n.HostID().String())
	if err != nil {
		return err
	}
	s.Modified = time.Now().Unix()
	return PublishEvent(b.ctx, b.storeEventsTopic, event)
}

// Set updates the entry in the store and publishes an event
func SetStoreEntry(se *v1.ChannelStoreRecord, value []byte, selfID string) (*v1.ChannelEvent, error) {
	if se.Owner != selfID {
		return nil, ErrNotOwner
	}
	se.Value = value
	se.Modified = time.Now().Unix()
	return &v1.ChannelEvent{
		Type:   v1.ChannelEventType_CHANNEL_EVENT_TYPE_SET,
		Record: se,
		Owner:  se.GetOwner(),
	}, nil
}

// NewPutEvent creates a new put event
func (b *channel) NewPutEvent(key string, value []byte) (*v1.ChannelEvent, *v1.ChannelStoreRecord) {
	entry := &v1.ChannelStoreRecord{
		Key:      key,
		Value:    value,
		Owner:    b.n.HostID().String(),
		Created:  time.Now().Unix(),
		Modified: time.Now().Unix(),
	}
	event := &v1.ChannelEvent{
		Type:   v1.ChannelEventType_CHANNEL_EVENT_TYPE_SET,
		Owner:  b.n.HostID().String(),
		Record: entry,
	}
	return event, entry
}

// NewDeleteEvent creates a new delete event
func (b *channel) NewDeleteEvent(key string) *v1.ChannelEvent {
	entry := &v1.ChannelStoreRecord{
		Key:      key,
		Owner:    b.n.HostID().String(),
		Modified: time.Now().Unix(),
	}
	event := &v1.ChannelEvent{
		Type:   v1.ChannelEventType_CHANNEL_EVENT_TYPE_DELETE,
		Owner:  b.n.HostID().String(),
		Record: entry,
	}
	return event
}

// Publish publishes the event to the topic
func PublishEvent(ctx context.Context, t *pubsub.Topic, e *v1.ChannelEvent) error {
	buf, err := proto.Marshal(e)
	if err != nil {
		return err
	}
	return t.Publish(ctx, buf)
}
