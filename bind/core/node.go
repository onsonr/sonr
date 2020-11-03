package core

import (
	"context"

	"github.com/libp2p/go-libp2p-core/connmgr"
	"github.com/libp2p/go-libp2p-core/event"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

// SonrNode contains all values for user
type SonrNode struct {
	PeerID  string
	PubSub  *pubsub.PubSub
	Host    host.Host
	Profile string
	OLC     string
}

// Connect ensures there is a connection between this host and the peer with
// given peer.ID. Connect will absorb the addresses in pi into its internal
// peerstore. If there is not an active connection, Connect will issue a
// h.Network.Dial, and block until a connection is open, or an error is
// returned. // TODO: Relay + NAT.
func (sn *SonrNode) connect(ctx context.Context, pi peer.AddrInfo) error {
	return sn.Host.Connect(ctx, pi)
}

// SetStreamHandler sets the protocol handler on the Host's Mux.
// This is equivalent to:
//   host.Mux().SetHandler(proto, handler)
// (Threadsafe)
func (sn *SonrNode) setStreamHandler(pid protocol.ID, handler network.StreamHandler) {
	sn.Host.SetStreamHandler(pid, handler)
}

// SetStreamHandlerMatch sets the protocol handler on the Host's Mux
// using a matching function for protocol selection.
func (sn *SonrNode) setStreamHandlerMatch(pid protocol.ID, m func(string) bool, h network.StreamHandler) {
	sn.Host.SetStreamHandlerMatch(pid, m, h)
}

// RemoveStreamHandler removes a handler on the mux that was set by
// SetStreamHandler
func (sn *SonrNode) removeStreamHandler(pid protocol.ID) {
	sn.Host.RemoveStreamHandler(pid)
}

// NewStream opens a new stream to given peer p, and writes a p2p/protocol
// header with given ProtocolID. If there is no connection to p, attempts
// to create one. If ProtocolID is "", writes no header.
// (Threadsafe)
func (sn *SonrNode) newStream(ctx context.Context, p peer.ID, pids ...protocol.ID) (network.Stream, error) {
	return sn.Host.NewStream(ctx, p, pids...)
}

// Close shuts down the host, its Network, and services.
func (sn *SonrNode) close() error {
	return sn.Host.Close()
}

// ConnManager returns this hosts connection manager
func (sn *SonrNode) connManager() connmgr.ConnManager {
	return sn.Host.ConnManager()
}

// EventBus returns the hosts eventbus
func (sn *SonrNode) eventBus() event.Bus {
	return sn.Host.EventBus()
}
