package host

import (
	"context"
	"fmt"

	p2p "github.com/libp2p/go-libp2p"
	p2p_connmgr "github.com/libp2p/go-libp2p-core/connmgr"
	p2p_event "github.com/libp2p/go-libp2p-core/event"
	p2p_host "github.com/libp2p/go-libp2p-core/host"
	p2p_network "github.com/libp2p/go-libp2p-core/network"
	p2p_peer "github.com/libp2p/go-libp2p-core/peer"
	p2p_pstore "github.com/libp2p/go-libp2p-core/peerstore"
	p2p_protocol "github.com/libp2p/go-libp2p-core/protocol"
	p2p_config "github.com/libp2p/go-libp2p/config"

	ipfs_p2p "github.com/ipfs/go-ipfs/core/node/libp2p"

	ma "github.com/multiformats/go-multiaddr"
)

// MobileHost is an host
var _ p2p_host.Host = (*MobileHost)(nil)

type MobileHost struct {
	p2p_host.Host
}

func NewMobileHostOption(mcfg *MobileConfig) ipfs_p2p.HostOption {
	return func(ctx context.Context, id p2p_peer.ID, ps p2p_pstore.Peerstore, options ...p2p.Option) (p2p_host.Host, error) {
		pkey := ps.PrivKey(id)
		if pkey == nil {
			return nil, fmt.Errorf("missing private key for node ID: %s", id.Pretty())
		}

		options = append([]p2p.Option{p2p.Identity(pkey), p2p.Peerstore(ps)}, options...)

		cfg := &p2p_config.Config{}
		if err := cfg.Apply(options...); err != nil {
			return nil, err
		}

		return NewMobileHost(ctx, mcfg, cfg)
	}
}

func NewMobileHost(ctx context.Context, _ *MobileConfig, cfg *p2p.Config) (p2p_host.Host, error) {
	host, err := cfg.NewNode(ctx)
	if err != nil {
		return nil, err
	}

	// @TODO: MobileHost custom config

	return &MobileHost{
		Host: host,
	}, nil
}

// ID returns the (local) peer.ID associated with this Host
func (mh *MobileHost) ID() p2p_peer.ID {
	return mh.Host.ID()
}

// Peerstore returns the Host's repository of Peer Addresses and Keys.
func (mh *MobileHost) Peerstore() p2p_pstore.Peerstore {
	return mh.Host.Peerstore()
}

// Returns the listen addresses of the Host
func (mh *MobileHost) Addrs() []ma.Multiaddr {
	return mh.Host.Addrs()
}

// Networks returns the Network interface of the Host
func (mh *MobileHost) Network() p2p_network.Network {
	return mh.Host.Network()
}

// Mux returns the Mux multiplexing incoming streams to protocol handlers
func (mh *MobileHost) Mux() p2p_protocol.Switch {
	return mh.Host.Mux()
}

// Connect ensures there is a connection between this host and the peer with
// given peer.ID. Connect will absorb the addresses in pi into its internal
// peerstore. If there is not an active connection, Connect will issue a
// h.Network.Dial, and block until a connection is open, or an error is
// returned. // TODO: Relay + NAT.
func (mh *MobileHost) Connect(ctx context.Context, pi p2p_peer.AddrInfo) error {
	return mh.Host.Connect(ctx, pi)
}

// SetStreamHandler sets the protocol handler on the Host's Mux.
// This is equivalent to:
//   host.Mux().SetHandler(proto, handler)
// (Threadsafe)
func (mh *MobileHost) SetStreamHandler(pid p2p_protocol.ID, handler p2p_network.StreamHandler) {
	mh.Host.SetStreamHandler(pid, handler)
}

// SetStreamHandlerMatch sets the protocol handler on the Host's Mux
// using a matching function for protocol selection.
func (mh *MobileHost) SetStreamHandlerMatch(pid p2p_protocol.ID, m func(string) bool, h p2p_network.StreamHandler) {
	mh.Host.SetStreamHandlerMatch(pid, m, h)
}

// RemoveStreamHandler removes a handler on the mux that was set by
// SetStreamHandler
func (mh *MobileHost) RemoveStreamHandler(pid p2p_protocol.ID) {
	mh.Host.RemoveStreamHandler(pid)
}

// NewStream opens a new stream to given peer p, and writes a p2p/protocol
// header with given ProtocolID. If there is no connection to p, attempts
// to create one. If ProtocolID is "", writes no header.
// (Threadsafe)
func (mh *MobileHost) NewStream(ctx context.Context, p p2p_peer.ID, pids ...p2p_protocol.ID) (p2p_network.Stream, error) {
	return mh.Host.NewStream(ctx, p, pids...)
}

// Close shuts down the host, its Network, and services.
func (mh *MobileHost) Close() error {
	return mh.Host.Close()
}

// ConnManager returns this hosts connection manager
func (mh *MobileHost) ConnManager() p2p_connmgr.ConnManager {
	return mh.Host.ConnManager()
}

// EventBus returns the hosts eventbus
func (mh *MobileHost) EventBus() p2p_event.Bus {
	return mh.Host.EventBus()
}
