package net

import (
	"fmt"

	"net"
	"os"

	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/pkg/errors"
)

// BootstrapAddrInfo Returns Bootstrap List Address Info
func BootstrapAddrInfo() ([]peer.AddrInfo, error) {
	// Create Bootstrapper List
	var bootstrappers []ma.Multiaddr
	for _, s := range []string{
		// Libp2p Default
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmQCU2EcMqAqQPR2i9bChDtGNJchTbq5TbXJJ16u19uLTa",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmcZf59bWwK5XFi76CZX8cbJ4BhTzzA3gU1ZjYZcYW3dwt",
		"/ip4/104.131.131.82/tcp/4001/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ",
		"/ip4/104.131.131.82/udp/4001/quic/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ",
	} {
		ma, err := ma.NewMultiaddr(s)
		if err != nil {
			return nil, err
		}
		bootstrappers = append(bootstrappers, ma)
	}

	ds := make([]peer.AddrInfo, 0, len(bootstrappers))
	for i := range bootstrappers {
		info, err := peer.AddrInfoFromP2pAddr(bootstrappers[i])
		if err != nil {
			continue
		}
		ds = append(ds, *info)
	}
	return ds, nil
}

// FreePort asks the kernel for a free open port
func FreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

// PublicAddrStrs Returns Device Listening Addresses ^ //
func PublicAddrStrs() ([]string, error) {
	// Initialize
	listenAddrs := []string{}

	// // Set Initial Port
	port, err := FreePort()
	if err != nil {
		return nil, err
	}

	// 	// Get iPv4 Addresses
	ip4Addrs, err := iPv4Addrs(port)
	if err != nil {
		return nil, err
	}

	listenAddrs = append(listenAddrs, ip4Addrs...)
	// Return Listen Addr Strings
	return listenAddrs, nil
}

// iPv4Addrs Returns Node Public iPv4 Address
func iPv4Addrs(port int) ([]string, error) {
	// Find Hos
	osHost, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	// Find Public Address Strings
	addrs, err := net.LookupIP(osHost)
	if err != nil {
		return nil, err
	}

	// Iterate through addresses
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			return []string{
				fmt.Sprintf("/ip4/%s/tcp/%d", ipv4.String(), port),
			}, nil

		}
	}
	return nil, errors.New("No IPV4 found")
}
