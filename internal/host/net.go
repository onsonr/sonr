package host

import (
	"errors"
	"net"
)

// ^ localIP returns the non loopback local IP of the host ^
func localIP() (string, error) {
	// Find Interface Addresses
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", errors.New("Local IP Not Found")
	}

	// Loop and Find Local IP
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}

	// No IPv4 Found
	return "", errors.New("Local IP Not Found")
}
