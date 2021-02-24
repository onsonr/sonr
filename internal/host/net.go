package host

import (
	"net"
	"os"
)

// ^ IPv4 returns the non loopback local IP of the host as IPv4 ^
func IPv4() string {
	// @1. Find IPv4 Address
	osHost, _ := os.Hostname()
	addrs, _ := net.LookupIP(osHost)
	var ipv4Ref string

	// Iterate through addresses
	for _, addr := range addrs {
		// @ Set IPv4
		if ipv4 := addr.To4(); ipv4 != nil {
			ipv4Ref = ipv4.String()
		} else {
			ipv4Ref = "0.0.0.0"
		}
	}
	// No IPv4 Found
	return ipv4Ref
}

// ^ IPv4 returns the non loopback local IP of the host as IPv4 ^
func IPv6() string {
	// @1. Find IPv4 Address
	osHost, _ := os.Hostname()
	addrs, _ := net.LookupIP(osHost)
	var ipv6Ref string

	// Iterate through addresses
	for _, addr := range addrs {
		// @ Set IPv4
		if ipv6 := addr.To16(); ipv6 != nil {
			ipv6Ref = ipv6.String()
		} else {
			ipv6Ref = "::"
		}
	}
	// No IPv4 Found
	return ipv6Ref
}
