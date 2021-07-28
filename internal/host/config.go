package host

import (
	"fmt"

	"net"
	"os"

	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/pkg/errors"
	md "github.com/sonr-io/core/pkg/models"
)

// ** ─── Address MANAGEMENT ────────────────────────────────────────────────────────
// # Return Bootstrap List Address Info
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

// # FreePort asks the kernel for a free open port
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

// # Return Device Listening Addresses ^ //
func PublicAddrStrs(cr *md.ConnectionRequest) ([]string, error) {
	// Initialize
	opts := cr.GetHostOptions()
	listenAddrs := []string{}
	hasIpv4 := false
	hasIpv6 := false

	// Set Initial Port
	port := 52006

	// Check for Port
	if opts.GetListenPort() != 0 {
		// Set Port from Options
		port = int(opts.GetListenPort())
	} else {
		// Set Port
		p, err := FreePort()
		if err != nil {
			md.LogError(err)
		} else {
			port = p
		}
	}

	// Check for Listen Addresses
	if len(opts.GetListenAddrs()) == 0 {
		// Get iPv4 Addresses
		ip4Addrs, err := iPv4Addrs(port)
		if err == nil {
			hasIpv4 = true
		}

		// Add iPv4 Addresses
		if hasIpv4 {
			listenAddrs = append(listenAddrs, ip4Addrs...)
		}

		// Neither iPv6 nor iPv4 found
		if !hasIpv4 && !hasIpv6 {
			return nil, errors.New("No IP Addresses found")
		}
	} else {
		for _, addr := range opts.GetListenAddrs() {
			// Get Address String
			addrStr, err := addr.MultiAddrStr(opts.GetIpv4Only(), port)
			if err != nil {
				md.LogError(err)
				continue // Skip this address
			}

			// Append Address List
			listenAddrs = append(listenAddrs, addrStr)
		}
	}

	// Return Listen Addr Strings
	return listenAddrs, nil
}

// # Returns Node Public iPv4 Address
func iPv4Addrs(port int) ([]string, error) {
	// Find Hos
	osHost, err := os.Hostname()
	if err != nil {
		md.LogError(err)
		return nil, err
	}

	// Find Public Address Strings
	addrs, err := net.LookupIP(osHost)
	if err != nil {
		md.LogError(err)
		return nil, err
	}

	// Iterate through addresses
	for _, addr := range addrs {
		// @ Set IPv4
		if ipv4 := addr.To4(); ipv4 != nil {
			ip4 := ipv4.String()
			return []string{
				fmt.Sprintf("/ip4/%s/tcp/%d", ip4, port),
			}, nil

		}
	}

	return nil, errors.New("No IPV4 found")
}

// # Handle Post-Connection result on hostNode
func handleConnectionResult(hh HostHandler, hostActive bool, textileActive bool, mdnsActive bool) {
	hh.OnConnected(&md.ConnectionResponse{
		HostActive:    hostActive,
		MdnsActive:    mdnsActive,
		TextileActive: textileActive,
	})
}
