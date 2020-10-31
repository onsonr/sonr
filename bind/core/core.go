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
	h := sonrHost.Start()
	return MobileHost{
		id: fmt.Sprintf("%s\n", h.ID()),
	}
}
