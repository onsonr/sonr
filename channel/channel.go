package channel

import (
	"context"
	"strings"
	"time"

	"github.com/kataras/golog"
	"github.com/libp2p/go-libp2p-core/peer"
	ps "github.com/libp2p/go-libp2p-pubsub"
	"github.com/pkg/errors"
	ct "github.com/sonr-io/blockchain/x/channel/types"
	ot "github.com/sonr-io/blockchain/x/object/types"
	nh "github.com/sonr-io/core/host"
)

var (
	logger            *golog.Logger
	ErrNotOwner       = errors.New("Not owner of key - (Beam)")
	ErrNotFound       = errors.New("Key not found in store - (Beam)")
	ErrInvalidMessage = errors.New("Invalid message received in Pubsub Topic - (Beam)")
)

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

// Channel is a pubsub based Key-Value store for Libp2p nodes.
type Channel interface {
	// Did returns the DID of the channel.
	Did() string

	// Read returns a list of all peers subscribed to the channel topic.
	Read() []peer.ID

	// Publish publishes the given message to the channel topic.
	Publish(obj *ot.ObjectDoc) error

	// Listen subscribes to the beam topic and returns a channel that will
	// receive events.
	Listen() <-chan *ct.ChannelMessage

	// Close closes the channel.
	Close() error
}

// channel is the implementation of the Beam interface.
type channel struct {
	Channel
	ctx   context.Context
	n     nh.HostImpl
	label string
	did   string

	// Channel Messages
	config          *ct.ChannelDoc
	messages        chan *ct.ChannelMessage
	messagesHandler *ps.TopicEventHandler
	messagesSub     *ps.Subscription
	messagesTopic   *ps.Topic
}

// New creates a new beam with the given name and options.
func New(ctx context.Context, n nh.HostImpl, config *ct.ChannelDoc, options ...Option) (Channel, error) {
	logger = golog.Default.Child(config.Label)
	opts := defaultOptions()
	for _, option := range options {
		option(opts)
	}

	mTopic, mHandler, mSub, err := n.NewTopic(config.Did)
	if err != nil {
		return nil, err
	}

	b := &channel{
		ctx:             ctx,
		n:               n,
		config:          config,
		did:             config.Did,
		messages:        make(chan *ct.ChannelMessage),
		messagesHandler: mHandler,
		messagesSub:     mSub,
		messagesTopic:   mTopic,
	}

	// Start the event handler.
	go b.handleChannelMessages()
	go b.serve()
	return b, nil
}

// Read lists all peers subscribed to the beam topic.
func (b *channel) Read() []peer.ID {
	messagesPeers := b.messagesTopic.ListPeers()

	// filter out duplicates
	peers := make(map[peer.ID]struct{})
	for _, p := range messagesPeers {
		peers[p] = struct{}{}
	}

	// convert to slice
	var result []peer.ID
	for p := range peers {
		result = append(result, p)
	}
	return result
}

// Publish publishes the given message to the beam topic.
func (b *channel) Publish(obj *ot.ObjectDoc) error {
	// Check if both text and data are empty.
	if obj == nil {
		return errors.New("text and data cannot be empty")
	}

	// Check if passed object is one registered in the channel.
	if !strings.EqualFold(b.config.RegisteredObject.Did, obj.Did) {
		return errors.New("object not registered in channel")
	}

	// Create the message.
	msg := &ct.ChannelMessage{
		Object: obj,
		Did:    b.did,
	}

	// Encode the message.
	buf, err := msg.Marshal()
	if err != nil {
		return err
	}

	// Publish the message to the beam topic.
	return b.messagesTopic.Publish(b.ctx, buf)
}

// Listen subscribes to the beam topic and returns a channel that will
func (b *channel) Listen() <-chan *ct.ChannelMessage {
	return b.messages
}

// Close closes the channel.
func (b *channel) Close() error {
	err := b.messagesTopic.Close()
	if err != nil {
		return err
	}
	return nil
}
