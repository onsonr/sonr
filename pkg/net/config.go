package net

import (
	"net"
	"os"
)

// ^ Router Protocol ID Option ^ //
type protocolRouterOption struct {
	local     bool
	group     bool
	groupName string
}

// ! Option to Set Protocol ID for Local
func SetIDForLocal() *protocolRouterOption {
	return &protocolRouterOption{
		local:     true,
		group:     false,
		groupName: "",
	}
}

// ! Option to Set Protocol ID for a Group, TODO
func SetIDForGroup(name string) *protocolRouterOption {
	return &protocolRouterOption{
		local:     false,
		group:     true,
		groupName: name,
	}
}

// @ Returns Node Public IPv4 Address
func IPv4() string {
	osHost, _ := os.Hostname()
	addrs, _ := net.LookupIP(osHost)
	ipv4Ref := "0.0.0.0"
	// Iterate through addresses
	for _, addr := range addrs {
		// @ Set IPv4
		if ipv4 := addr.To4(); ipv4 != nil {
			ipv4Ref = ipv4.String()
		}
	}
	return ipv4Ref
}

// @ Returns Node Public IPv6 Address
func IPv6() string {
	osHost, _ := os.Hostname()
	addrs, _ := net.LookupIP(osHost)
	ipv6Ref := "::"

	// Iterate through addresses
	for _, addr := range addrs {
		// @ Set IPv4
		if ipv6 := addr.To16(); ipv6 != nil {
			ipv6Ref = ipv6.String()
		}
	}
	return ipv6Ref
}
