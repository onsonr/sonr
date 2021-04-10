package network

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"
	md "github.com/sonr-io/core/pkg/models"
)

// ^ Return Bootstrap List Address Info ^ //
func GetBootstrapAddrInfo() ([]peer.AddrInfo, error) {
	// Create Bootstrapper List
	var bootstrappers []multiaddr.Multiaddr
	for _, s := range []string{
		// Libp2p Default
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmQCU2EcMqAqQPR2i9bChDtGNJchTbq5TbXJJ16u19uLTa",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmcZf59bWwK5XFi76CZX8cbJ4BhTzzA3gU1ZjYZcYW3dwt",
		"/ip4/104.131.131.82/tcp/4001/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ",
		"/ip4/104.131.131.82/udp/4001/quic/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ",
	} {
		ma, err := multiaddr.NewMultiaddr(s)
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

	// // Get iPv6 Addresses
	// ip6Addrs, err := iPv6Addrs()
	// if err != nil {
	// 	log.Println(err)
	// } else {
	// 	hasIpv6 = true
	// }

	// Add iPv4 Addresses
	if hasIpv4 {
		listenAddrs = append(listenAddrs, ip4Addrs...)
	}

	// // Add iPv6 Addresses
	// if hasIpv6 {
	// 	listenAddrs = append(listenAddrs, ip6Addrs...)
	// }

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
				fmt.Sprintf("/ip4/%s/udp/0/quic", ip4),
			}, nil

		}
	}
	return nil, errors.New("No IPV4 found")
}

// // @ Returns Node Public iPv6 Address
// func iPv6Addrs() ([]string, error) {
// 	osHost, _ := os.Hostname()
// 	addrs, _ := net.LookupIP(osHost)

// 	// Iterate through addresses
// 	for _, addr := range addrs {
// 		// @ Set IPv6
// 		if ipv6 := addr.To16(); ipv6 != nil {
// 			ip6 := ipv6.String()
// 			return []string{fmt.Sprintf("/ip6/%s/tcp/0", ip6)}, nil
// 		}
// 	}
// 	return nil, errors.New("No IPV6 Found")
// }

// ^ Returns Location from GeoIP ^ //
func Location(target *md.GeoIP) error {
	r, err := http.Get("https://api.ipgeolocationapi.com/geolocate")
	if err != nil {
		log.Fatalln(err)
	}
	return json.NewDecoder(r.Body).Decode(target)
}
