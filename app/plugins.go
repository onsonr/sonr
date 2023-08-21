package app

import (
	"github.com/sonr-io/sonr/internal/highway"
)

// EnablePlugins enables the plugins.
func EnablePlugins() {
	go highway.StartAPI()
}
