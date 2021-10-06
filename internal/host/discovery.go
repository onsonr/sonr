package host

import (
	"context"
	"errors"
	"fmt"
	"time"

	dscl "github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/peer"
	dsc "github.com/libp2p/go-libp2p-discovery"
	psub "github.com/libp2p/go-libp2p-pubsub"

	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/tools/logger"
)

var (
	ErrDHTNotFound = errors.New("DHT has not been set by Routing Function")
)

// discoveryNotifee is a Notifee for the Discovery Service
type discoveryNotifee struct {
	PeerChan chan peer.AddrInfo
}

// HandlePeerFound is to be called when new  peer is found
func (n *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	n.PeerChan <- pi
}

// ** ─── HostNode Connection Methods ────────────────────────────────────────────────────────
func (hn *SNRHost) checkDhtSet() error {
	if hn.IpfsDHT == nil {
		return ErrDHTNotFound
	}
	return nil
}

// Bootstrap begins bootstrap with peers
func (h *SNRHost) Bootstrap() error {
	// // Add Host Address to Peerstore
	// h.Peerstore().AddAddrs(h.ID(), h.Addrs(), peerstore.PermanentAddrTTL)

	// Check DHT Set
	retryFunc := common.NewRetryFunc(h.checkDhtSet, 3, time.Second*3)
	if err := retryFunc(); err != nil {
		return logger.Error("Host DHT was never set", err)
	}

	// Bootstrap DHT
	if err := h.IpfsDHT.Bootstrap(h.ctx); err != nil {
		return logger.Error("Failed to Bootstrap KDHT to Host", err)
	}

	// Connect to bootstrap nodes, if any
	for _, pi := range h.opts.bootstrapPeers {
		if err := h.Connect(pi); err != nil {
			continue
		} else {
			break
		}
	}

	// Set Routing Discovery, Find Peers
	routingDiscovery := dsc.NewRoutingDiscovery(h.IpfsDHT)
	dsc.Advertise(h.ctx, routingDiscovery, h.opts.rendezvous, dscl.TTL(h.opts.ttl))
	h.disc = routingDiscovery

	// Create Pub Sub
	ps, err := psub.NewGossipSub(h.ctx, h.Host, psub.WithDiscovery(routingDiscovery))
	if err != nil {
		return logger.Error("Failed to Create new Gossip Sub", err)
	}

	// Handle DHT Peers
	h.PubSub = ps
	peersChan, err := routingDiscovery.FindPeers(h.ctx, h.opts.rendezvous, dscl.TTL(h.opts.ttl))
	if err != nil {
		return logger.Error("Failed to create FindPeers Discovery channel", err)
	}
	go h.handleDiscoveredPeers(peersChan)
	return nil
}

func (h *SNRHost) Connect(pi peer.AddrInfo) error {
	ctxCancel, cancel := context.WithDeadline(h.ctx, <-time.After(10*time.Second))
	defer cancel()

	// Create Connect Func
	connectFunc := func() error {
		// Validate not Self
		if h.checkUnknown(pi) {
			return h.Host.Connect(ctxCancel, pi)
		}
		return nil
	}
	// Attempt Connect
	if err := common.NewRetryFunc(connectFunc, 3, 5*time.Second)(); err != nil {
		msg := fmt.Sprintf("Failed to connect to Peer %s \n Clearing from PeerStore and Adding to Ignored", pi.ID)
		logger.Error(msg, err)
		h.Peerstore().ClearAddrs(pi.ID)
		return err
	}
	return nil
}

// handleDiscoveredPeers Connect to Peers that are discovered
func (h *SNRHost) handleDiscoveredPeers(peerChan <-chan peer.AddrInfo) {
	for {
		select {
		case pi := <-peerChan:
			if err := h.Connect(pi); err != nil {
				continue
			}
		case <-h.ctx.Done():
			return
		}
	}
}

// checkUnknown is a Helper Method checks if Peer AddrInfo is Unknown
func (h *SNRHost) checkUnknown(pi peer.AddrInfo) bool {
	// Check if Peer is Self
	if h.ID() == pi.ID {
		return false
	}

	// Iterate and Check
	if len(h.Peerstore().Addrs(pi.ID)) > 0 {
		return false
	}
	return true
}
