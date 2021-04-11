package network

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	exip "github.com/glendc/go-external-ip"
	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
	"github.com/pkg/errors"
	md "github.com/sonr-io/core/pkg/models"
)

type AddrsFactory func(addrs []ma.Multiaddr) []ma.Multiaddr

// ^ Returns Location from GeoIP ^ //
func Location(target *md.GeoIP) error {
	r, err := http.Get("https://api.ipgeolocationapi.com/geolocate")
	if err != nil {
		log.Fatalln(err)
	}
	return json.NewDecoder(r.Body).Decode(target)
}

// ^ Return Bootstrap List Address Info ^ //
func GetBootstrapAddrInfo() ([]peer.AddrInfo, error) {
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

// ^ Failsafe to Return External IP ^ //
func GetExternalIPMultiAddr(port int) ma.Multiaddr {
	// Default consensus
	consensus := exip.DefaultConsensus(nil, nil)

	// Return IP
	externalIP, err := consensus.ExternalIP()
	if err != nil {
		return nil
	}

	// Check for ipv4
	if ipv4 := externalIP.To4(); ipv4 != nil {
		maddr, err := ma.NewMultiaddr(fmt.Sprintf("/ipv4/%s/tcp/%d", ipv4.String(), port))
		if err != nil {
			return nil
		}
		return maddr
	}

	// Check for ipv6
	if ipv6 := externalIP.To16(); ipv6 != nil {
		maddr, err := ma.NewMultiaddr(fmt.Sprintf("/ipv6/%s/tcp/%d", ipv6.String(), port))
		if err != nil {
			return nil
		}
		return maddr
	}
	return nil
}

// ^ GetFreePort asks the kernel for a free open port that is ready to use. ^
func GetFreePort() (int, error) {
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

// ^ GetFreePort asks the kernel for free open ports that are ready to use. ^
func GetFreePorts(count int) ([]int, error) {
	var ports []int
	for i := 0; i < count; i++ {
		addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
		if err != nil {
			return nil, err
		}

		l, err := net.ListenTCP("tcp", addr)
		if err != nil {
			return nil, err
		}
		defer l.Close()
		ports = append(ports, l.Addr().(*net.TCPAddr).Port)
	}
	return ports, nil
}

// ^ Return Device Listening Addresses ^ //
func GetListenAddrStrings() ([]string, error) {
	// Initialize
	listenAddrs := []string{}
	hasIpv4 := false
	hasIpv6 := false

	// Get iPv4 Addresses
	ip4Addrs, err := iPv4Addrs()
	if err != nil {
		log.Println(err)
	} else {
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

	// Return Listen Addr Strings
	return listenAddrs, nil
}

// @ Returns Node Public iPv4 Address
func iPv4Addrs() ([]string, error) {
	osHost, _ := os.Hostname()
	addrs, _ := net.LookupIP(osHost)
	// Iterate through addresses
	for _, addr := range addrs {
		// @ Set IPv4
		if ipv4 := addr.To4(); ipv4 != nil {
			ip4 := ipv4.String()
			return []string{
				fmt.Sprintf("/ip4/%s/tcp/0", ip4),
			}, nil

		}
	}
	return nil, errors.New("No IPV4 found")
}

// @ Return MultiAddrs using Net Host
func MultiAddrs() []ma.Multiaddr {
	// Local IP lookup
	osHost, _ := os.Hostname()
	addrs, _ := net.LookupIP(osHost)
	multiAddrs := []ma.Multiaddr{}

	// Get Free Port
	port, err := GetFreePort()
	if err != nil {
		port = 60214
	}

	// // Add External IP
	// if extMultiAddr := GetExternalIPMultiAddr(port); extMultiAddr != nil {
	// 	multiAddrs = append(multiAddrs, extMultiAddr)
	// }

	// Iterate through Net Addrs
	for _, addr := range addrs {
		// Add ipv4
		if ipv4 := addr.To4(); ipv4 != nil {
			if IsValidIPv4(ipv4) {
				maddr, err := ma.NewMultiaddr(fmt.Sprintf("/ip4/%s/tcp/%d", ipv4.String(), port))
				if err == nil {
					multiAddrs = append(multiAddrs, maddr)
				}
			}
		}
	}
	return multiAddrs
}

// @ Validates Not Private IPv4
func IsValidIPv4(ip net.IP) bool {
	// Local Link
	if !ip.IsGlobalUnicast() {
		return false
	}

	// Known Private iPv4
	for _, item := range manet.Private4 {
		if item.IP.Equal(ip) {
			return false
		}
	}
	return true
}

// @ Validates Not Private IPv6
func IsValidIPv6(ip net.IP) bool {
	// Local Link
	if !ip.IsGlobalUnicast() {
		return false
	}

	// Remove Unspecified IPv6
	if ip.IsUnspecified() {
		return false
	}

	// Known Private iPv6
	for _, item := range manet.Private6 {
		if item.IP.Equal(ip) {
			return false
		}
	}
	return true
}
