package net

import (
	"fmt"

	"net"
	"os"

	"github.com/pkg/errors"
)

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
