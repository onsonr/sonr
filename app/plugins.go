package app

import (
	"github.com/sonrhq/core/pkg/highway"
)

func EnablePlugins() {
	go highway.StartService()
}
