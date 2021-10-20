package beam

import (
	"context"
	"fmt"
	"time"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/pkg/errors"
	"github.com/sonr-io/core/internal/host"
)

var (
	ErrNotOwner       = errors.New("Not owner of key - (Beam)")
	ErrNotFound       = errors.New("Key not found in store - (Beam)")
	ErrInvalidMessage = errors.New("Invalid message received in Pubsub Topic - (Beam)")
)

// Beam is a pubsub based Key-Value store for Libp2p nodes.
type Beam interface {
	// Get returns the value for the given key.
	Get(key string) ([]byte, error)

	// Put stores the value for the given key.
	Put(key string, value []byte) error

	// Delete removes the value for the given key.
	Delete(key string) error

	// Close closes the beam.
	Close() error
}

// beam is the implementation of the Beam interface.
type beam struct {
	Beam
	ctx  context.Context
	h    *host.SNRHost
	name string

	events  chan *Event
	handler *pubsub.TopicEventHandler
	sub     *pubsub.Subscription
	topic   *pubsub.Topic

	store *Store
}

// New creates a new beam with the given name and options.
func New(ctx context.Context, h *host.SNRHost, name string, options ...Option) (Beam, error) {
	opts := defaultOptions()
	for _, option := range options {
		option(opts)
	}

	topic, err := h.Join(name)
	if err != nil {
		return nil, err
	}

	sub, err := topic.Subscribe()
	if err != nil {
		return nil, err
	}

	handler, err := topic.EventHandler()
	if err != nil {
		return nil, err
	}

	b := &beam{
		ctx:     ctx,
		h:       h,
		name:    name,
		topic:   topic,
		sub:     sub,
		handler: handler,
		store:   newStore(opts.capacity, opts.ttl),
	}
	go b.handleEvents()
	go b.handleMessages()
	return b, nil
}

// Delete removes the key in the beam store.
func (b *beam) Delete(key string) error {
	return b.store.Delete(fmt.Sprintf("%s/%s", b.name, key), b)
}

// Get returns the value for the given key in the beam store.
func (b *beam) Get(key string) ([]byte, error) {
	return b.store.Get(fmt.Sprintf("%s/%s", b.name, key))
}

// Put stores the value for the given key in the beam store.
func (b *beam) Put(key string, value []byte) error {
	return b.store.Put(fmt.Sprintf("%s/%s", b.name, key), value, b)
}

// Close closes the beam.
func (b *beam) Close() error {
	b.handler.Cancel()
	b.sub.Cancel()
	return b.topic.Close()
}

// Option is a function that modifies the beam options.
type Option func(*options)

// WithTTL sets the time-to-live for the beam store entries
func WithTTL(ttl time.Duration) Option {
	return func(o *options) {
		o.ttl = ttl
	}
}

// WithCapacity sets the capacity of the beam store.
func WithCapacity(capacity int) Option {
	return func(o *options) {
		o.capacity = capacity
	}
}

// options is a collection of options for the beam.
type options struct {
	ttl      time.Duration
	capacity int
}

// defaultOptions is the default options for the beam.
func defaultOptions() *options {
	return &options{
		ttl:      time.Minute * 10,
		capacity: 4096,
	}
}
