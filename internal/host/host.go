package host

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/libp2p/go-libp2p"
	connmgr "github.com/libp2p/go-libp2p-connmgr"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/libp2p/go-libp2p-core/routing"
	dsc "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	psub "github.com/libp2p/go-libp2p-pubsub"
	discovery "github.com/libp2p/go-libp2p/p2p/discovery"
	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/internal/keychain"
	"github.com/sonr-io/core/tools/logger"
	"github.com/sonr-io/core/tools/net"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

// SNRHostStat is the host stat info
type SNRHostStat struct {
	ID        peer.ID
	PublicKey string
	PeerID    string
	MultAddr  string
	Address   string
}

// SNRHost is the host wrapper for the Sonr Network
type SNRHost struct {
	host.Host

	// Properties
	ctxHost      context.Context
	ctxTileAuth  context.Context
	ctxTileToken context.Context

	// Libp2p
	disc   *dsc.RoutingDiscovery
	kdht   *dht.IpfsDHT
	mdns   discovery.Service
	pubsub *psub.PubSub

	// Textile
	// 	client  *client.Client
	// 	mail    *local.Mail
	// 	mailbox *local.Mailbox
}

type HostListenAddr struct {
	Addr   string
	Family string
}

// MultiAddrStr Returns MultiAddr from IP Address
func (la HostListenAddr) MultiAddrStr(port int) string {
	if strings.Contains(la.Family, "6") {
		return fmt.Sprintf("/ip6/%s/tcp/%d", la.Addr, port)
	}
	return fmt.Sprintf("/ip4/%s/tcp/%d", la.Addr, port)
}

// NewHost creates a new host
func NewHost(ctx context.Context, conn common.Connection, privKey crypto.PrivKey, listenAddrs ...HostListenAddr) (*SNRHost, error) {
	// Initialize DHT
	var kdhtRef *dht.IpfsDHT
	var listenAddresses []string

	// Add Listen Addresses
	if len(listenAddrs) > 0 {
		// Set Initial Port
		port, err := net.FreePort()
		if err != nil {
			logger.Warn("Failed to get free port", zap.Error(err))
			port = 600214
		}

		// Build MultAddr Address Strings
		for _, addr := range listenAddrs {
			listenAddresses = append(listenAddresses, addr.MultiAddrStr(port))
		}
	} else {
		addrs, err := net.PublicAddrStrs()
		if err != nil {
			logger.Warn("Failed to get public addresses", zap.Error(err))
			listenAddresses = []string{}
		} else {
			listenAddresses = addrs
		}
	}

	// Start Host
	h, err := libp2p.New(
		ctx,
		libp2p.Identity(privKey),
		libp2p.ListenAddrStrings(listenAddresses...),
		libp2p.DefaultTransports,
		libp2p.ConnectionManager(connmgr.NewConnManager(
			25,            // Lowwater
			50,            // HighWater,
			time.Minute*5, // GracePeriod
		)),
		libp2p.DefaultStaticRelays(),
		libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
			// Create DHT
			kdht, err := dht.New(ctx, h)
			if err != nil {
				return nil, err
			}

			// Set DHT
			kdhtRef = kdht
			return kdht, nil
		}),
		libp2p.NATPortMap(),
		libp2p.EnableAutoRelay(),
	)
	if err != nil {
		return nil, logger.Error("Failed to initialize libp2p host", err)
	}

	// Create Pub Sub
	ps, err := psub.NewGossipSub(ctx, h, psub.WithDiscovery(dsc.NewRoutingDiscovery(kdhtRef)))
	if err != nil {
		return nil, logger.Error("Failed to Create new Gossip Sub", err)
	}

	// Create Host
	hn := &SNRHost{
		ctxHost: ctx,
		Host:    h,
		kdht:    kdhtRef,
		pubsub:  ps,
	}

	// Bootstrap Host
	err = hn.Bootstrap()
	if err != nil {
		return nil, logger.Error("Failed to bootstrap libp2p Host", err)
	}

	// Check for Wifi/Ethernet for MDNS
	if conn == common.Connection_WIFI || conn == common.Connection_ETHERNET {
		// err = hn.MDNS()
		// if err != nil {
		// 	return nil, logger.Error("Failed to start MDNS Discovery", err)
		// }
		logger.Debug("MDNS Discovery Disabled: Temp")
	}
	return hn, nil
}

// ** ─── Host Info ────────────────────────────────────────────────────────
// Pubsub Returns Host Node MultiAddr
func (hn *SNRHost) Pubsub() *psub.PubSub {
	return hn.pubsub
}

// SendMessage writes a protobuf go data object to a network stream
func (hn *SNRHost) SendMessage(id peer.ID, p protocol.ID, data proto.Message) error {
	s, err := hn.NewStream(hn.ctxHost, id, p)
	if err != nil {
		return logger.Error("Failed to start stream", err)
	}
	defer s.Close()

	// marshall data to protobufs3 binary format
	bin, err := proto.Marshal(data)
	if err != nil {
		return logger.Error("Failed to marshal pb", err)
	}

	// Create Writer and write data to stream
	w := msgio.NewWriter(s)
	if err := w.WriteMsg(bin); err != nil {
		return logger.Error("Failed to write message to stream.", err)
	}
	return nil
}

// Stat returns the host stat info
func (hn *SNRHost) Stat() (*SNRHostStat, error) {
	// Get Public Key
	pubKey, err := device.KeyChain.GetPubKey(keychain.Account)
	if err != nil {
		return nil, logger.Error("Failed to get public key", err)
	}

	// Marshal Public Key
	buf, err := crypto.MarshalPublicKey(pubKey)
	if err != nil {
		return nil, logger.Error("Failed to marshal public key", err)
	}

	// Return Host Stat
	return &SNRHostStat{
		ID:        hn.ID(),
		PublicKey: string(buf),
		PeerID:    hn.ID().Pretty(),
		MultAddr:  hn.Addrs()[0].String(),
	}, nil
}
