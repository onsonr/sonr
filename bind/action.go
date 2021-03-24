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
		mn.user.SetPosition(facing, heading)
		mn.node.Update(mn.user.Peer())
	}
}

// @ Send Direct Message to Peer in Lobby
func (mn *MobileNode) Message(msg string, to string) {
	if mn.isReady() {
		mn.node.Message(msg, to, mn.user.Peer())
	}
}

// @ Invite Processes Data and Sends Invite to Peer
func (mn *MobileNode) Invite(reqBytes []byte) {
	if mn.isReady() {
		mn.status = md.Status_PENDING
		// Initialize from Request
		req := &md.InviteRequest{}
		err := proto.Unmarshal(reqBytes, req)
		if err != nil {
			log.Println(err)
		}

		// @ 2. Check Transfer Type
		if req.Type == md.InviteRequest_Contact {
			mn.node.InviteContact(req, mn.user.Peer(), mn.user.Contact())
		} else if req.Type == md.InviteRequest_URL {
			mn.node.InviteLink(req, mn.user.Peer())
		} else {
			err := mn.user.FS.AddFromRequest(req, mn.user.Peer())
			if err != nil {
				mn.error(err, "AddFromRequest")
			}
			mn.status = md.Status_AVAILABLE
		}
	}
}

// @ Respond to an Invite with Decision
func (mn *MobileNode) Respond(decs bool) {
	if mn.isReady() {
		mn.node.Respond(decs, mn.user.FS, mn.user.Peer(), mn.user.Contact())
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
		// mn.node.JoinRemote(data)
		log.Println(data)
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
		err = mn.user.SetContact(newContact)
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
	mn.user.FS.Close()
	mn.node.Close()
}
