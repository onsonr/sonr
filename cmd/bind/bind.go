package bind

import (
	"log"

	md "github.com/sonr-io/core/pkg/models"
	sn "github.com/sonr-io/core/pkg/node"
	"google.golang.org/protobuf/proto"
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

// @ Struct: Reference for Binded Proxy Node
type MobileNode struct {
	ID       string
	DeviceID string
	UserID   uint32
	node     *sn.Node
}

// Create New Mobile Node
func NewNode(reqBytes []byte, call Callback) *MobileNode {
	// ** Unmarshal Request **
	req := md.ConnectionRequest{}
	err := proto.Unmarshal(reqBytes, &req)
	if err != nil {
		panic(err)
	}

	// Create New Sonr Client
	node := sn.NewNode(&req, call)
	peer := node.Peer()

	// Return Mobile Node
	return &MobileNode{
		node:     node,
		ID:       peer.Id.Peer,
		DeviceID: peer.Id.Device,
		UserID:   peer.Id.User,
	}
}

// **-----------------** //
// ** Network Actions ** //
// **-----------------** //
// Start Host
func (sn *MobileNode) Start() {
	err := sn.node.Start()
	if err != nil {
		log.Fatalln(err)
	}
}

// Initiate Bootstrapping
func (sn *MobileNode) Bootstrap() {
	err := sn.node.Bootstrap()
	if err != nil {
		log.Fatalln(err)
	}
}

// **--------------** //
// ** Node Actions ** //
// **--------------** //
// Update proximity/direction and Notify Lobby
func (sn *MobileNode) Update(facing float64, heading float64) {
	sn.node.Update(facing, heading)
}

// Invite Processes Data and Sends Invite to Peer
func (sn *MobileNode) Invite(reqBytes []byte) {
	// @ 1. Initialize from Request
	req := &md.InviteRequest{}
	err := proto.Unmarshal(reqBytes, req)
	if err != nil {
		log.Println(err)
	}

	sn.node.Invite(req)
}

// Respond to an Invitation
func (sn *MobileNode) Respond(decision bool) {
	sn.node.Respond(decision)
}

// ** User Actions ** //
// Info returns ALL Peer Data as Bytes
func (sn *MobileNode) Info() []byte {
	info := sn.node.Info()
	return info
}

//  Link with a QR Code
func (sn *MobileNode) LinkDevice(json string) {
	sn.node.LinkDevice(json)
}

// Updates Current Contact Card
func (sn *MobileNode) SetContact(conBytes []byte) {
	// Unmarshal Data
	newContact := &md.Contact{}
	err := proto.Unmarshal(conBytes, newContact)
	if err != nil {
		log.Println(err)
	}
	sn.node.SetContact(newContact)
}

// **-------------------** //
// ** LifeCycle Actions ** //
// **-------------------** //
// Close Ends All Network Communication
func (sn *MobileNode) Pause() {
	sn.node.Pause()
}

// Close Ends All Network Communication
func (sn *MobileNode) Resume() {
	sn.node.Resume()
}

// Close Ends All Network Communication
func (sn *MobileNode) Stop() {
	sn.node.Stop()
}
