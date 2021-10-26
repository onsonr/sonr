package beam

import (
	"context"
	"time"

	"github.com/kataras/golog"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/pkg/errors"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/go-gun"
)

var (
	logger            *golog.Logger
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
}

// beam is the implementation of the Beam interface.
type beam struct {
	Beam
	ctx context.Context
	h   *host.SNRHost
	id  ID

	events  chan *Event
	handler *pubsub.TopicEventHandler
	sub     *pubsub.Subscription
	topic   *pubsub.Topic

	gun   *gun.Gun
	store *Store
}

// New creates a new beam with the given name and options.
func New(ctx context.Context, h *host.SNRHost, id ID, options ...Option) (Beam, error) {
	logger = golog.Default.Child(id.Prefix())
	opts := defaultOptions()
	for _, option := range options {
		option(opts)
	}

	topic, err := h.Join(id.String())
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
		id:      id,
		topic:   topic,
		sub:     sub,
		handler: handler,
	}

	// Initialize the store.
	if err := b.initStore(ctx, opts); err != nil {
		return nil, err
	}

	// Start the event handler.
	go b.handleEvents()
	go b.handleMessages()
	go b.serve()
	return b, nil
}

// Delete removes the key in the beam store.
func (b *beam) Delete(key string) error {
	return b.store.Delete(b.id.Key(key), b)
}

// Get returns the value for the given key in the beam store.
func (b *beam) Get(key string) ([]byte, error) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancelFn()

	fetchCh := b.gun.Scoped(ctx, b.id.String(), key).Fetch(ctx)
	// Log all updates and exit when context times out
	logger.Print("Waiting for value")
	for {
		select {
		case <-ctx.Done():
			logger.Print("Time's up")
			return b.store.Get(b.id.Key(key))
		case fetchResult := <-fetchCh:
			if fetchResult.Err != nil {
				logger.Printf("Error fetching: %v", fetchResult.Err)
			} else if fetchResult.ValueExists {
				logger.Printf("Got value: %v", fetchResult.Value)
			}
			return b.store.Get(b.id.Key(key))
		}
	}
}

// Put stores the value for the given key in the beam store.
func (b *beam) Put(key string, value []byte) error {
	// Give a 1 minute timeout, but shouldn't get hit
	ctx, cancelFn := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancelFn()
	// Issue a simple put and wait for a single peer ack
	putScope := b.gun.Scoped(ctx, b.id.String(), key)
	logger.Print("Sending first value")
	putCh := putScope.Put(ctx, gun.ValueString(string(value)))
	for {
		if result := <-putCh; result.Err != nil {
			logger.Printf("Error putting: %v", result.Err)
		} else if result.Peer != nil {
			logger.Printf("Got ack from %v", result.Peer)
			break
		}
	}
	return b.store.Put(b.id.Key(key), value, b)
}
