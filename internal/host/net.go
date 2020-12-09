//nolint
package host

import (
	"fmt"
	"net"
	"os"

	//tor "berty.tech/go-libp2p-tor-transport"
	//torConfig "berty.tech/go-libp2p-tor-transport/config"
	"github.com/libp2p/go-libp2p"
	//"github.com/libp2p/go-libp2p-core/transport"
	quic "github.com/libp2p/go-libp2p-quic-transport"
	"github.com/libp2p/go-libp2p/config"
)

// ^ Find IPv4 Address ^
func GetIPv4() string {
	osHost, _ := os.Hostname()
	addrs, _ := net.LookupIP(osHost)

	// Iterate through addresses
	for _, addr := range addrs {
		// @ Set IPv4
		if ipv4 := addr.To4(); ipv4 != nil {
			return ipv4.String()
		}
	}
	return ""
}

// ^ Helper: Method Returns Wifi Options ^ //
// TODO:
// func getTorTransport() (transport.Transport, error) {
// 	torTransport, err := tor.NewBuilder( // Create a builder
// 		torConfig.EnableEmbeded,
// 		torConfig.DoSlowStart,
// 		torConfig.AllowTcpDial, // Some Configurator are already ready to use.
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return torTransport, nil
// }

// ^ Helper: Method Returns Wifi Options ^ //
func getWifiOptions() []config.Option {
	// @1. Find IPv4 Address
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
	var opts []config.Option
	opts = append(opts,
		// Add listening Addresses
		libp2p.ListenAddrStrings(
			fmt.Sprintf("/ip4/%s/tcp/0", ipv4Ref),
			"/ip6/::/tcp/0",

			fmt.Sprintf("/ip4/%s/udp/0/quic", ipv4Ref),
			"/ip6/::/udp/0/quic"),

		// support QUIC
		libp2p.Transport(quic.NewTransport))
	return opts
}
