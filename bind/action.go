package bind

import (
	"log"

	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// @ Update proximity/direction and Notify Lobby
func (mn *MobileNode) Update(facing float64, heading float64) {
	if mn.config.isReady() {
		mn.user.SetPosition(facing, heading)
		err := mn.node.Update(mn.local, mn.user.Peer())
		if err != nil {
			log.Println(err)
			return
		}
	}
}

// @ Send Direct Message to Peer in Lobby
func (mn *MobileNode) Message(msg string, to string) {
	if mn.config.isReady() {
		err := mn.node.Message(mn.local, msg, to, mn.user.Peer())
		if err != nil {
			log.Println(err)
			return
		}
	}
}

// @ Invite Processes Data and Sends Invite to Peer
func (mn *MobileNode) Invite(reqBytes []byte) {
	if mn.config.isReady() {
		// Update Status
		mn.config.Status = md.Status_PENDING

		// Initialize from Request
		req := &md.InviteRequest{}
		if err := proto.Unmarshal(reqBytes, req); err != nil {
			log.Println(err)
			return
		}

		// @ 2. Check Transfer Type
		if req.Type == md.InviteRequest_Contact {
			err := mn.node.InviteContact(req, mn.local, mn.user.Peer(), mn.user.Contact())
			if err != nil {
				log.Println(err)
				return
			}
		} else if req.Type == md.InviteRequest_URL {
			err := mn.node.InviteLink(req, mn.local, mn.user.Peer())
			if err != nil {
				log.Println(err)
				return
			}
		} else {
			// @ Add File to Queue
			hasCompleted := make(chan bool)
			go func() {
				if err := mn.user.FS.EnqueueFromRequest(req, mn.user.Peer(), hasCompleted); err != nil {
					log.Println(err)
					return
				}
			}()

			// @ Wait For Done
			done := <-hasCompleted
			if done {
				// Retreive Current File
				currFile, err := mn.user.FS.DequeueOut()
				if err != nil {
					log.Println(err)
					return
				}

				// Send Invite
				err = mn.node.InviteFile(req, mn.local, mn.user.Peer(), currFile)
				if err != nil {
					log.Println(err)
					return
				}
			}
		}
	}
}

// @ Respond to an Invite with Decision
func (mn *MobileNode) Respond(decs bool) {
	if mn.config.isReady() {
		mn.node.Respond(decs, mn.user.FS, mn.user.Peer(), mn.local, mn.user.Contact())
		// Update Status
		if decs {
			mn.config.Status = md.Status_INPROGRESS
		} else {
			mn.config.Status = md.Status_AVAILABLE
		}
	}
}

// @ Join Existing Group
func (mn *MobileNode) JoinRemote(data string) {
	if mn.config.isReady() {
		// mn.node.JoinRemote(data)
		log.Println(data)
		return
	}
}

// ** User Actions ** //
// @ Updates Current Contact Card
func (mn *MobileNode) SetContact(conBytes []byte) {
	if mn.config.isReady() {
		// Unmarshal Data
		newContact := &md.Contact{}
		err := proto.Unmarshal(conBytes, newContact)
		if err != nil {
			log.Println(err)
			return
		}

		// Update Node Profile
		err = mn.user.SetContact(newContact)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
