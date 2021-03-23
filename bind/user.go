package bind

import (
	"log"

	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// @ Update proximity/direction and Notify Lobby
func (mn *MobileNode) Update(facing float64, heading float64) {
	if mn.IsReady() {
		mn.node.Update(facing, heading)
	}
}

// @ Send Direct Message to Peer in Lobby
func (mn *MobileNode) Message(msg string, to string) {
	if mn.IsReady() {
		mn.node.Message(msg, to)
	}
}

// @ Invite Processes Data and Sends Invite to Peer
func (mn *MobileNode) Invite(reqBytes []byte) {
	if mn.IsReady() {
		// Initialize from Request
		req := &md.InviteRequest{}
		err := proto.Unmarshal(reqBytes, req)
		if err != nil {
			log.Println(err)
		}
		mn.node.Invite(req)
	}
}

// @ Respond to an Invite with Decision
func (mn *MobileNode) Respond(decs bool) {
	if mn.IsReady() {
		mn.node.Respond(decs)
	}
}

// @ Join Existing Group
func (mn *MobileNode) JoinRemote(data string) {
	if mn.IsReady() {
		mn.node.JoinRemote(data)
	}
}

// ** User Actions ** //
// @ Link with a QR Code
func (mn *MobileNode) LinkDevice(json string) {
	if mn.IsReady() {
		mn.node.LinkDevice(json)
	}
}

// @ Updates Current Contact Card
func (mn *MobileNode) SetContact(conBytes []byte) {
	if mn.IsReady() {
		// Unmarshal Data
		newContact := &md.Contact{}
		err := proto.Unmarshal(conBytes, newContact)
		if err != nil {
			log.Println(err)
		}
		mn.node.SetContact(newContact)
	}
}
