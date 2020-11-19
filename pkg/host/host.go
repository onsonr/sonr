package host

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	libp2pquic "github.com/libp2p/go-libp2p-quic-transport"
	secio "github.com/libp2p/go-libp2p-secio"
	libp2ptls "github.com/libp2p/go-libp2p-tls"
)

// NewBasicHost creates a host without any options
func NewHost(ctx *context.Context) (host.Host, error) {
	// Find IPv4 Address
	osHost, _ := os.Hostname()
	addrs, _ := net.LookupIP(osHost)
	var ipv4Ref string

	// Iterate through addresses
	for _, addr := range addrs {
		// @ Set IPv4
		if ipv4 := addr.To4(); ipv4 != nil {
			ipv4Ref = ipv4.String()
		}
	}

	// Create Libp2p Host
	h, err := libp2p.New(*ctx,
		// Add listening Addresses
		libp2p.ListenAddrStrings(
			fmt.Sprintf("/ip4/%s/tcp/0", ipv4Ref),
			fmt.Sprintf("/ip4/%s/udp/0/quic", ipv4Ref)),

		// support TLS connections
		libp2p.Security(libp2ptls.ID, libp2ptls.New),

		// support secio connections
		libp2p.Security(secio.ID, secio.New),

		// support QUIC
		libp2p.Transport(libp2pquic.NewTransport),

		// support any other default transports (TCP)
		libp2p.DefaultTransports,
	)
	return h, err
}
