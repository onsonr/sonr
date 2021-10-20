package beam

import (
	"time"

	"google.golang.org/protobuf/proto"
)

// newStore creates a new store
func newStore(cap int, ttl time.Duration) *Store {
	s := &Store{
		Data:     make(map[string]*StoreEntry),
		Capacity: int32(cap),
		Modified: time.Now().Unix(),
		Ttl:      ttl.Milliseconds(),
	}
	go s.handleStoreTTL()
	return s
}

// Delete deletes an entry from the store and publishes an event
func (s *Store) Delete(key string, b *beam) error {
	// Fetch the entry
	entry := s.Data[key]
	if entry == nil {
		return ErrNotFound
	}

	// Check if the entry is owned by this node
	if entry.Peer != b.h.ID().String() {
		return ErrNotOwner
	}

	// Create Delete Event
	event := b.newDeleteEvent(key)
	eventBuf, err := event.Marshal()
	if err != nil {
		return err
	}
	return b.topic.Publish(b.ctx, eventBuf)
}

// Get returns the value of the entry
func (s *Store) Get(key string) ([]byte, error) {
	entry := s.Data[key]
	if entry == nil {
		return nil, ErrNotFound
	}
	return entry.Value, nil
}

// Handle checks the event type and handles it with the store
func (s *Store) Handle(e *Event, b *beam) error {
	// Check if the event is valid
	if b.h.ID().String() == e.Peer {
		return nil
	}

	switch e.Type {
	case EventType_DELETE:
		delete(s.Data, e.Entry.Key)
	default:
		s.Data[e.Entry.Key] = e.Entry
	}
	return nil
}

// Put puts an entry into the store and publishes an event
func (s *Store) Put(key string, value []byte, b *beam) error {
	// Fetch the entry
	entry := s.Data[key]
	if entry == nil {
		// Create new entry with Event
		event, entry := b.newPutEvent(key, value)
		s.Data[key] = entry

		// Publish event
		eventBuf, err := event.Marshal()
		if err != nil {
			return err
		}
		return b.topic.Publish(b.ctx, eventBuf)
	}

	// Get existing entry and update it
	event, err := entry.Set(value, b.h.ID().String())
	if err != nil {
		return err
	}
	eventBuf, err := event.Marshal()
	if err != nil {
		return err
	}
	return b.topic.Publish(b.ctx, eventBuf)
}

// Set updates the entry in the store and publishes an event
func (se *StoreEntry) Set(value []byte, selfID string) (*Event, error) {
	if se.Peer != selfID {
		return nil, ErrNotOwner
	}
	se.Value = value
	se.Modified = time.Now().Unix()
	return &Event{
		Type:  EventType_SET,
		Entry: se,
		Peer:  se.GetPeer(),
	}, nil
}

func (b *beam) newPutEvent(key string, value []byte) (*Event, *StoreEntry) {
	entry := &StoreEntry{
		Key:      key,
		Value:    value,
		Peer:     b.h.ID().String(),
		Created:  time.Now().Unix(),
		Modified: time.Now().Unix(),
	}
	event := &Event{
		Type:  EventType_PUT,
		Peer:  b.h.ID().String(),
		Entry: entry,
	}
	return event, entry
}

func (b *beam) newDeleteEvent(key string) *Event {
	entry := &StoreEntry{
		Key:      key,
		Peer:     b.h.ID().String(),
		Modified: time.Now().Unix(),
	}
	event := &Event{
		Type:  EventType_DELETE,
		Peer:  b.h.ID().String(),
		Entry: entry,
	}
	return event
}

func (e *Event) Marshal() ([]byte, error) {
	return proto.Marshal(e)
}
