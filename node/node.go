package node

import (
	"context"
	"net"
	"sync/atomic"

	"git.mills.io/prologic/bitcask"
	common "github.com/sonr-io/core/common"
	"github.com/sonr-io/core/host"
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

	// Close closes the node
	Close()
}

// Node type - a p2p host implementing one or more p2p protocols
type Node struct {
	// Standard Node Implementation
	*host.SNRHost
	NodeImpl
	mode StubMode

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
	opts := defaultNodeOptions()
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
	node := &Node{
		ctx:      ctx,
		listener: l,
		SNRHost:  host,
		mode:     StubMode_LIB,
	}
	return node, nil
}

// NewHighway Creates a node with its implemented protocols
func NewHighway(ctx context.Context, l net.Listener, options ...Option) (NodeImpl, error) {
	// Set Node Options
	opts := defaultNodeOptions()
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
	node := &Node{
		ctx:      ctx,
		listener: l,
		SNRHost:  host,
		mode:     StubMode_FULL,
	}
	return node, nil
}

// Host returns the underlying host
func (n *Node) Host() *host.SNRHost {
	return n.SNRHost
}

// Close closes the node
func (n *Node) Close() {
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
func (c *Node) NeedsWait() {
	<-c.Chn
}

// Resume tells all of goroutines to resume execution
func (c *Node) Resume() {
	if atomic.LoadUint64(&c.flag) == 1 {
		close(c.Chn)
		atomic.StoreUint64(&c.flag, 0)
	}
}

// Pause tells all of goroutines to pause execution
func (c *Node) Pause() {
	if atomic.LoadUint64(&c.flag) == 0 {
		atomic.StoreUint64(&c.flag, 1)
		c.Chn = make(chan bool)
	}
}
