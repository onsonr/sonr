package bind

import (
	"log"

	"github.com/getsentry/sentry-go"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/encoding/protojson"
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
		mn.status = md.Status_PENDING

		// @ 2. Check Transfer Type
		if req.Type == md.InviteRequest_Contact || req.Type == md.InviteRequest_URL {
			mn.node.Invite(req)

		} else {
			// File Transfer
			mn.fs.AddFromRequest(req)
		}
	}
}

// @ Respond to an Invite with Decision
func (mn *MobileNode) Respond(decs bool) {
	if mn.IsReady() {
		mn.node.Respond(decs)
		// Update Status
		if decs {
			mn.status = md.Status_INPROGRESS
		} else {
			mn.status = md.Status_AVAILABLE
		}
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
		// Convert String to Bytes
		request := md.LinkRequest{}

		// Convert to Peer Protobuf
		err := protojson.Unmarshal([]byte(json), &request)
		if err != nil {
			sentry.CaptureException(err)
		}

		// Link Device
		err = mn.fs.SaveDevice(request.Device)
		if err != nil {
			sentry.CaptureException(err)
		}
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
