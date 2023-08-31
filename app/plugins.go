package app

import (
	"github.com/sonrhq/core/config"
	"github.com/sonrhq/core/internal/highway"
)

// EnablePlugins enables the plugins.
func EnablePlugins() {
	if config.HighwayEnabled() {
		highway.StartAPI()
	}
}
