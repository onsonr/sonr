package network

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-netroute"
	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
	md "github.com/sonr-io/core/pkg/models"
)

type AddrsFactory func(addrs []ma.Multiaddr) []ma.Multiaddr

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

// @ Returns Address with Net Route
func MultiAddrs() ([]ma.Multiaddr, *md.SonrError) {
	// Initialize
	allMultiAddrs := []ma.Multiaddr{}
	filteredMultiAddrs := []ma.Multiaddr{}

	// Route IP Address
	if r, err := netroute.New(); err != nil {
		return nil, md.NewError(err, md.ErrorMessage_IP_LOCATE)
	} else {
		if _, _, localIPv4, err := r.Route(net.IPv4zero); err != nil {
			return nil, md.NewError(err, md.ErrorMessage_IP_LOCATE)
		} else if localIPv4.IsGlobalUnicast() {
			maddr, err := manet.FromIP(localIPv4)
			if err == nil {
				allMultiAddrs = append(allMultiAddrs, maddr)
			}
		}

		if _, _, localIPv6, err := r.Route(net.IPv6unspecified); err != nil {
			return nil, md.NewError(err, md.ErrorMessage_IP_LOCATE)
		} else if localIPv6.IsGlobalUnicast() {
			maddr, err := manet.FromIP(localIPv6)
			if err == nil {
				allMultiAddrs = append(allMultiAddrs, maddr)
			}
		}
	}

	// Resolve the interface addresses
	ifaceAddrs, err := manet.InterfaceMultiaddrs()
	if err != nil {

		// Add the loopback addresses to the filtered addrs and use them as the non-filtered addrs.
		// Then bail. There's nothing else we can do here.
		filteredMultiAddrs = append(filteredMultiAddrs, manet.IP4Loopback, manet.IP6Loopback)
		allMultiAddrs = filteredMultiAddrs

		// This usually shouldn't happen, but we could be in some kind
		// of funky restricted environment.
		return allMultiAddrs, md.NewError(err, md.ErrorMessage_IP_RESOLVE)
	}

	for _, addr := range ifaceAddrs {
		// Skip link-local addrs, they're mostly useless.
		if !manet.IsIP6LinkLocal(addr) {
			allMultiAddrs = append(allMultiAddrs, addr)
		}
	}

	// If netroute failed to get us any interface addresses, use all of
	// them.
	if len(filteredMultiAddrs) == 0 {
		// Add all addresses.
		filteredMultiAddrs = allMultiAddrs
		return filteredMultiAddrs, nil
	} else {
		// Only add loopback addresses. Filter these because we might
		// not _have_ an IPv6 loopback address.
		for _, addr := range allMultiAddrs {
			if manet.IsIPLoopback(addr) {
				filteredMultiAddrs = append(filteredMultiAddrs, addr)
			}
		}
	}

	return filteredMultiAddrs, nil
}

// @ Return MultiAddrs using Net Host
func MultiAddrsNet() []ma.Multiaddr {
	// Local IP lookup
	osHost, _ := os.Hostname()
	addrs, _ := net.LookupIP(osHost)
	allMultiAddrs := []ma.Multiaddr{}
	filteredMultiAddrs := []ma.Multiaddr{}

	// Iterate through Net Addrs
	for _, addr := range addrs {
		if addr.IsGlobalUnicast() {
			maddr, err := manet.FromIP(addr)
			if err == nil {
				allMultiAddrs = append(allMultiAddrs, maddr)
			}
		}
	}

	// Filter out Local Link Addrs
	for _, addr := range allMultiAddrs {
		// Skip link-local addrs, they're mostly useless.
		if !manet.IsIP6LinkLocal(addr) {
			filteredMultiAddrs = append(filteredMultiAddrs, addr)
		}
	}

	return filteredMultiAddrs
}

// ^ Returns Location from GeoIP ^ //
func Location(target *md.GeoIP) error {
	r, err := http.Get("https://api.ipgeolocationapi.com/geolocate")
	if err != nil {
		log.Fatalln(err)
	}
	return json.NewDecoder(r.Body).Decode(target)
}
