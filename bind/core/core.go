package core

import (
	"fmt"
	sonrHost "github.com/sonr-io/p2p/pkg/host"
)

// MobileHost is object which returns simple data back to dart
type MobileHost struct {
	id string
}

func startHost() MobileHost {
	h := sonrHost.start()
	return MobileHost{
		id: fmt.Printf("Hello World, my hosts ID is %s\n", h.ID()),
	}
}