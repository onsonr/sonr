package channel

import (
	"context"

	"github.com/kataras/golog"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/pkg/errors"
	"github.com/sonr-io/core/node"
	channelV1 "github.com/sonr-io/core/types/go/channel/v1"
)

var (
	logger            *golog.Logger
	ErrNotOwner       = errors.New("Not owner of key - (Beam)")
	ErrNotFound       = errors.New("Key not found in store - (Beam)")
	ErrInvalidMessage = errors.New("Invalid message received in Pubsub Topic - (Beam)")
)

// Channel is a pubsub based Key-Value store for Libp2p nodes.
type Channel interface {
	// Get returns the value for the given key.
	Get(key string) ([]byte, error)

	// Put stores the value for the given key.
	Put(key string, value []byte) error

	// Delete removes the value for the given key.
	Delete(key string) error
}

// channel is the implementation of the Beam interface.
type channel struct {
	Channel
	ctx context.Context
	n   node.NodeImpl
	id  ID

	events  chan *channelV1.Event
	handler *pubsub.TopicEventHandler
	sub     *pubsub.Subscription
	topic   *pubsub.Topic
	store   *channelV1.Store
}

// New creates a new beam with the given name and options.
func New(ctx context.Context, n node.NodeImpl, id ID, options ...Option) (Channel, error) {
	logger = golog.Default.Child(id.Prefix())
	opts := defaultOptions()
	for _, option := range options {
		option(opts)
	}

	topic, err := n.Join(id.String())
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

	b := &channel{
		ctx:     ctx,
		n:       n,
		id:      id,
		topic:   topic,
		sub:     sub,
		handler: handler,
		store:   NewStore(opts),
	}

	// Start the event handler.
	go b.handleEvents()
	go b.handleMessages()
	go b.serve()
	return b, nil
}

// Delete removes the key in the beam store.
func (b *channel) Delete(key string) error {
	return DeleteStoreKey(b.store, b.id.Key(key), b)
}

// Get returns the value for the given key in the beam store.
func (b *channel) Get(key string) ([]byte, error) {
	return GetKey(b.store, b.id.Key(key))
}

// Put stores the value for the given key in the beam store.
func (b *channel) Put(key string, value []byte) error {
	return PutStoreKey(b.store, b.id.Key(key), value, b)
}
