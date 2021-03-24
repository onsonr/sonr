package bind

import (
	"log"

	"github.com/getsentry/sentry-go"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// @ Update proximity/direction and Notify Lobby
func (mn *MobileNode) Update(facing float64, heading float64) {
	if mn.isReady() {
		mn.node.Update(facing, heading)
	}
}

// @ Send Direct Message to Peer in Lobby
func (mn *MobileNode) Message(msg string, to string) {
	if mn.isReady() {
		mn.node.Message(msg, to)
	}
}

// @ Invite Processes Data and Sends Invite to Peer
func (mn *MobileNode) Invite(reqBytes []byte) {
	if mn.isReady() {
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
	if mn.isReady() {
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
	if mn.isReady() {
		mn.node.JoinRemote(data)
	}
}

// ** User Actions ** //
// @ Updates Current Contact Card
func (mn *MobileNode) SetContact(conBytes []byte) {
	if mn.isReady() {
		// Unmarshal Data
		newContact := &md.Contact{}
		err := proto.Unmarshal(conBytes, newContact)
		if err != nil {
			log.Println(err)
		}

		// Update Node Profile
		mn.node.SetContact(newContact)

		// Set User Contact
		err = mn.fs.SaveContact(newContact)
		if err != nil {
			sentry.CaptureException(err)
		}
	}
}

// **-------------------** //
// ** LifeCycle Actions ** //
// **-------------------** //
// @ Checks for is Ready
func (mn *MobileNode) isReady() bool {
	return mn.hasBootstrapped && mn.hasStarted
}

// @ Close Ends All Network Communication
func (mn *MobileNode) Pause() {
	mn.node.Pause()
}

// @ Close Ends All Network Communication
func (mn *MobileNode) Resume() {
	mn.node.Resume()
}

// @ Close Ends All Network Communication
func (mn *MobileNode) Stop() {
	// Check if Response Is Invited
	mn.node.Close()
}
