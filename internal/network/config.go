package network

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
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

// @ Return MultiAddrs using Net Host
func MultiAddrs() []ma.Multiaddr {
	// Local IP lookup
	osHost, _ := os.Hostname()
	addrs, _ := net.LookupIP(osHost)
	allMultiAddrs := []ma.Multiaddr{}
	filteredMultiAddrs := []ma.Multiaddr{}

	// Iterate through Net Addrs
	for _, addr := range addrs {
		if addr.IsGlobalUnicast() {
			// Add ipv4
			if ipv4 := addr.To4(); ipv4 != nil {
				if IsNotPrivateIPv4(ipv4) {
					maddr, err := ma.NewMultiaddr(fmt.Sprintf("/ip4/%s/tcp/%d", ipv4, 0))
					if err == nil {
						allMultiAddrs = append(allMultiAddrs, maddr)
					}
				}
			}

			// Add ipv6
			if ipv6 := addr.To16(); ipv6 != nil {
				if IsNotPrivateIPv6(ipv6) {
					maddr, err := ma.NewMultiaddr(fmt.Sprintf("/ip6/%s/tcp/%d", ipv6, 0))
					if err == nil {
						allMultiAddrs = append(allMultiAddrs, maddr)
					}
				}

			}
		}
	}

	// Filter out Local Link Addrs
	for _, addr := range allMultiAddrs {
		// Skip link-local addrs, they're mostly useless.
		if !manet.IsIPUnspecified(addr) && !manet.IsIP6LinkLocal(addr) {
			filteredMultiAddrs = append(filteredMultiAddrs, addr)
		}
	}
	return filteredMultiAddrs
}

// @ Validates Not Private IPv4
func IsNotPrivateIPv4(ip net.IP) bool {
	for _, item := range manet.Private4 {
		if item.IP.Equal(ip) {
			return false
		}
	}
	return true
}

// @ Validates Not Private IPv6
func IsNotPrivateIPv6(ip net.IP) bool {
	for _, item := range manet.Private6 {
		if item.IP.Equal(ip) {
			return false
		}
	}
	return true
}
