package main

import (
	md "github.com/sonr-io/core/pkg/models"
)

// ** ─── Node Status Checks ────────────────────────────────────────────────────────
// # Checks if Node is Ready for Actions
func (n *NodeServer) isReady() bool {
	return n.user.IsNotStatus(md.Status_STANDBY) || n.user.IsNotStatus(md.Status_FAILED)
}

// # Sets Node to be Connected Status
func (s *NodeServer) setConnected(val bool) {
	// Update Status
	su := s.user.SetConnected(val)

	// Callback Status
	s.statusEvents <- su
}

// # Sets Node to be Available Status
func (s *NodeServer) setAvailable(val bool) {
	// Update Status
	su := s.user.SetAvailable(val)

	// Callback Status
	s.statusEvents <- su
}

// # Sets Node to be (Provided) Status
func (s *NodeServer) setStatus(newStatus md.Status) {
	// Set Status
	su := s.user.SetStatus(newStatus)

	// Callback Status
	s.statusEvents <- su
}
