package context

import (
	"fmt"
	"log"
	"net"
	"time"
)

// CheckReachable checks if the given host and port is reachable.
func CheckReachable(host string, port int) bool {
	timeout := 1 * time.Second
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeout)
	if err != nil {
		log.Println("Site unreachable, error: ", err)
        return false
	}
	conn.Close()
	return true
}
