package client

import (
	md "github.com/sonr-io/core/pkg/models"
	sn "github.com/sonr-io/core/pkg/node"
)

// @ Interface: Callback is implemented from Plugin to receive updates
type Callback interface {
	OnRefreshed(data []byte)   // Lobby Updates
	OnEvent(data []byte)       // Lobby Event
	OnInvited(data []byte)     // User Invited
	OnDirected(data []byte)    // User Direct-Invite from another Device
	OnResponded(data []byte)   // Peer has responded
	OnProgress(data float32)   // File Progress Updated
	OnReceived(data []byte)    // User Received File
	OnTransmitted(data []byte) // User Sent File
	OnError(data []byte)       // Internal Error
}

// @ Struct: Node for Alternative Client
type ClientNode struct {
	ID       string
	DeviceID string
	UserID   uint32
	node     *sn.Node
	Status   md.Status
}

// Generate QRCode for Linking
func (sn *ClientNode) LinkRequest(name string) *md.LinkRequest {
	lreq := sn.node.LinkRequest(name)
	return lreq
}
