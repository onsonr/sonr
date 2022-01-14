package channel

import (
	"context"

	"github.com/kataras/golog"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/pkg/errors"
	"github.com/sonr-io/core/node"
	v1 "github.com/sonr-io/core/protocols/channel/v1"
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

	// Listen subscribes to the beam topic and returns a channel that will
	// receive events.
	Listen() (<-chan *v1.ChannelMessage, error)

	// Close closes the channel.
	Close() error
}

// channel is the implementation of the Beam interface.
type channel struct {
	Channel
	ctx context.Context
	n   node.NodeImpl
	id  ID

	// Channel Messages
	messages        chan *v1.ChannelMessage
	messagesHandler *pubsub.TopicEventHandler
	messagesSub     *pubsub.Subscription
	messagesTopic   *pubsub.Topic

	// Store Events
	storeEvents        chan *v1.ChannelEvent
	storeEventsHandler *pubsub.TopicEventHandler
	storeEventsSub     *pubsub.Subscription
	storeEventsTopic   *pubsub.Topic
	store              *v1.ChannelStore
}

// New creates a new beam with the given name and options.
func New(ctx context.Context, n node.NodeImpl, id ID, options ...Option) (Channel, error) {
	logger = golog.Default.Child(id.Prefix())
	opts := defaultOptions()
	for _, option := range options {
		option(opts)
	}

	mTopic, err := n.Join(id.String())
	if err != nil {
		return nil, err
	}

	mSub, err := mTopic.Subscribe()
	if err != nil {
		return nil, err
	}

	mHandler, err := mTopic.EventHandler()
	if err != nil {
		return nil, err
	}

	evTopic, err := n.Join(id.String())
	if err != nil {
		return nil, err
	}

	evSub, err := evTopic.Subscribe()
	if err != nil {
		return nil, err
	}

	evHandler, err := evTopic.EventHandler()
	if err != nil {
		return nil, err
	}

	b := &channel{
		ctx:                ctx,
		n:                  n,
		id:                 id,
		messages:           make(chan *v1.ChannelMessage),
		messagesHandler:    mHandler,
		messagesSub:        mSub,
		messagesTopic:      mTopic,
		storeEvents:        make(chan *v1.ChannelEvent),
		storeEventsTopic:   evTopic,
		storeEventsSub:     evSub,
		storeEventsHandler: evHandler,
		store:              NewStore(opts),
	}

	// Start the event handler.
	go b.handleChannelEvents()
	go b.handleChannelMessages()
	go b.handleStoreEvents()
	go b.handleStoreMessages()
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

// Listen subscribes to the beam topic and returns a channel that will
func (b *channel) Listen() (<-chan *v1.ChannelMessage, error) {
	return b.messages, nil
}

// Close closes the channel.
func (b *channel) Close() error {
	err := b.storeEventsTopic.Close()
	if err != nil {
		return err
	}

	err = b.messagesTopic.Close()
	if err != nil {
		return err
	}
	return nil
}
