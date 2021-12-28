package node

import (
	"context"
	"net"
	"sync/atomic"

	"git.mills.io/prologic/bitcask"
	"github.com/kataras/golog"
	common "github.com/sonr-io/core/common"
	"github.com/sonr-io/core/host"
	"github.com/sonr-io/core/wallet"
)

var (
	logger   = golog.Default.Child("core/node")
	ctx      context.Context
	instance NodeImpl
	sockets  *host.SockManager
)

// NodeImpl returns the NodeImpl for the Main Node
type NodeImpl interface {
	Host() *host.SNRHost

	// Profile returns the profile of the node from Local Store
	Profile() (*common.Profile, error)

	// Peer returns the peer of the node
	Peer() (*common.Peer, error)

	// NeedsWait checks if state is Resumed or Paused and blocks channel if needed
	NeedsWait()

	// Resume tells all of goroutines to resume execution
	Resume()

	// Pause tells all of goroutines to pause execution
	Pause()

	// Role returns the role of the node
	Role() Role

	// Close closes the node
	Close()
}

// node type - a p2p host implementing one or more p2p protocols
type node struct {
	// Standard Node Implementation
	*host.SNRHost
	NodeImpl
	mode Role

	// Host and context
	listener net.Listener

	// Properties
	ctx   context.Context
	store *bitcask.Bitcask

	flag uint64
	Chn  chan bool
}

// NewMotor Creates a node with its implemented protocols
func NewMotor(ctx context.Context, l net.Listener, options ...Option) (NodeImpl, error) {
	// Set Node Options
	opts := defaultOptions()
	for _, opt := range options {
		opt(opts)
	}

	// Initialize Host
	host, err := host.NewHost(ctx, host.WithConnection(opts.connection))
	if err != nil {
		logger.Errorf("%s - Failed to initialize host", err)
		return nil, err
	}

	// Open Store with profileBuf
	// Create Node
	node := &node{
		ctx:      ctx,
		listener: l,
		SNRHost:  host,
		mode:     Role_MOTOR,
	}
	return node, nil
}

// NewHighway Creates a node with its implemented protocols
func NewHighway(ctx context.Context, l net.Listener, options ...Option) (NodeImpl, error) {
	// Set Node Options
	opts := defaultOptions()
	for _, opt := range options {
		opt(opts)
	}

	// Initialize Host
	host, err := host.NewHost(ctx, host.WithConnection(opts.connection))
	if err != nil {
		logger.Errorf("%s - Failed to initialize host", err)
		return nil, err
	}

	// Open Store with profileBuf
	// Create Node
	node := &node{
		ctx:      ctx,
		listener: l,
		SNRHost:  host,
		mode:     Role_HIGHWAY,
	}
	return node, nil
}

func (n *node) Role() Role {
	return n.mode
}

// Host returns the underlying host
func (n *node) Host() *host.SNRHost {
	return n.SNRHost
}

// Close closes the node
func (n *node) Close() {
	// Close Store
	if err := n.store.Close(); err != nil {
		logger.Errorf("%s - Failed to close store, ", err)
	}

	// Close Host
	if err := n.SNRHost.Close(); err != nil {
		logger.Errorf("%s - Failed to close host, ", err)
	}
}

// NeedsWait checks if state is Resumed or Paused and blocks channel if needed
func (c *node) NeedsWait() {
	<-c.Chn
}

// Resume tells all of goroutines to resume execution
func (c *node) Resume() {
	if atomic.LoadUint64(&c.flag) == 1 {
		close(c.Chn)
		atomic.StoreUint64(&c.flag, 0)
	}
}

// Pause tells all of goroutines to pause execution
func (c *node) Pause() {
	if atomic.LoadUint64(&c.flag) == 0 {
		atomic.StoreUint64(&c.flag, 1)
		c.Chn = make(chan bool)
	}
}

// SignedMetadataToProto converts a SignedMetadata to a protobuf.
func SignedMetadataToProto(m *wallet.SignedMetadata) *common.Metadata {
	return &common.Metadata{
		Timestamp: m.Timestamp,
		NodeId:    m.NodeId,
		PublicKey: m.PublicKey,
	}
}
