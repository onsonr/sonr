package host

import (
	"github.com/sonr-io/core/tools/logger"
	"github.com/sonr-io/core/tools/net"
)

func getListenAddrStrings(listenAddrs ...HostListenAddr) []string {
	var listenAddresses []string
	// Add Listen Addresses
	if len(listenAddrs) > 0 {
		// Set Initial Port
		port, err := net.FreePort()
		if err != nil {
			logger.Warn("Failed to get free port", err)
			port = 60214
		}

		// Build MultAddr Address Strings
		for _, addr := range listenAddrs {
			listenAddresses = append(listenAddresses, addr.MultiAddrStr(port))
		}
	} else {
		addrs, err := net.PublicAddrStrs()
		if err != nil {
			logger.Warn("Failed to get public addresses", err)
			listenAddresses = []string{}
		} else {
			listenAddresses = addrs
		}
	}
	return listenAddresses
}
