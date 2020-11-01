package core

import (
	"fmt"

	"github.com/libp2p/go-libp2p-core/host"
	sonrHost "github.com/sonr-io/p2p/pkg/host"
)

// MobileHost is object which returns simple data back to dart
type MobileHost struct {
	id   string
	host *host.Host
}

// Start begins the mobile host
func Start() string {
	h := sonrHost.NewHost()
	return fmt.Sprintf("%s\n", h.ID())
}
