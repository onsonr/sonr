package main

import (
	md "github.com/sonr-io/core/pkg/models"
)

// ** ─── Node Status Checks ────────────────────────────────────────────────────────
// # Checks if Node is Ready for Actions
func (n *NodeServer) isReady() bool {
	return n.device.IsNotStatus(md.Status_STANDBY) || n.device.IsNotStatus(md.Status_FAILED)
}

// # Sets Node to be Connected Status
func (s *NodeServer) setConnected(val bool) {
	// Update Status
	su := s.device.SetConnected(val)

	// Callback Status
	s.statusEvents <- su
}

// # Sets Node to be Available Status
func (s *NodeServer) setAvailable(val bool) {
	// Update Status
	su := s.device.SetAvailable(val)

	// Callback Status
	s.statusEvents <- su
}

// # Sets Node to be (Provided) Status
func (s *NodeServer) setStatus(newStatus md.Status) {
	// Set Status
	su := s.device.SetStatus(newStatus)

	// Callback Status
	s.statusEvents <- su
}
