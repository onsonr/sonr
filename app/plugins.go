package app

import (
	"github.com/sonrhq/core/internal/highway"
)

// EnablePlugins enables the plugins.
func EnablePlugins() {
	go highway.StartAPI()
}
