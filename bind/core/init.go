package core

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p"
	connmgr "github.com/libp2p/go-libp2p-connmgr"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p/p2p/discovery"
)

// initMDNSDiscovery creates an mDNS discovery service and attaches it to the libp2p Host.
func initMDNSDiscovery(ctx context.Context, sn SonrNode, call SonrCallback) error {
	// setup mDNS discovery to find local peers
	disc, err := discovery.NewMdnsService(ctx, sn.Host, discoveryInterval, discoveryServiceTag)
	if err != nil {
		return err
	}

	// Create Discovery Notifier
	n := discoveryNotifee{sn: sn, call: call}
	disc.RegisterNotifee(&n)
	return nil
}

// initBasicHost creates basic host for mdns
func initBasicHost(ctx context.Context) (host.Host, error) {
	// Create Host
	host, err := libp2p.New(ctx, libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"),
		libp2p.ConnectionManager(connmgr.NewConnManager(
			100,         // Lowwater
			400,         // HighWater,
			time.Minute, // GracePeriod
		)))

	// Check for Error
	if err != nil {
		return nil, err
	}

	// Return Host
	return host, nil
}
