package host

import (
	"time"

	dscl "github.com/libp2p/go-libp2p-core/discovery"
	dsc "github.com/libp2p/go-libp2p-discovery"
	psub "github.com/libp2p/go-libp2p-pubsub"
	cfg "github.com/sonr-io/core/internal/network"
	md "github.com/sonr-io/core/pkg/models"
)

// ^ Bootstrap begins bootstrap with peers ^
func (h *HostNode) Bootstrap() *md.SonrError {
	// Create Bootstrapper Info
	bootstrappers, err := cfg.GetBootstrapAddrInfo()
	if err != nil {
		return md.NewError(err, md.ErrorMessage_BOOTSTRAP)
	}

	// Bootstrap DHT
	if err := h.KDHT.Bootstrap(h.ctx); err != nil {
		return md.NewError(err, md.ErrorMessage_BOOTSTRAP)
	}

	// Connect to bootstrap nodes, if any
	for _, pi := range bootstrappers {
		if err := h.Host.Connect(h.ctx, pi); err != nil {
			continue
		} else {
			break
		}
	}

	// Set Routing Discovery, Find Peers
	routingDiscovery := dsc.NewRoutingDiscovery(h.KDHT)
	dsc.Advertise(h.ctx, routingDiscovery, h.Point, dscl.TTL(time.Second*4))
	h.Discovery = routingDiscovery

	// Create Pub Sub
	ps, err := psub.NewGossipSub(h.ctx, h.Host, psub.WithDiscovery(routingDiscovery))
	if err != nil {
		return md.NewError(err, md.ErrorMessage_HOST_PUBSUB)
	}
	h.Pubsub = ps
	go h.handleDHTPeers(routingDiscovery)
	return nil
}
