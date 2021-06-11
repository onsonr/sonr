package bind

import (
	"log"

	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// @ Start Host and Connect
func (mn *Node) Connect() []byte {
	// Connect Host
	err := mn.client.Connect(mn.user.KeyPrivate())
	if err != nil {
		mn.handleError(err)
		mn.setConnected(false)
		return nil
	} else {
		// Update Status
		mn.setConnected(true)
	}

	// Bootstrap Node
	mn.local, err = mn.client.Bootstrap()
	if err != nil {
		mn.handleError(err)
		mn.setAvailable(false)
		return nil
	} else {
		mn.setAvailable(true)
	}

	// Create ConnectResponse
	bytes, rerr := proto.Marshal(&md.ConnectionResponse{
		User: mn.user,
		Id:   mn.user.ID(),
	})

	// Handle Error
	if rerr != nil {
		mn.handleError(md.NewMarshalError(rerr))
		return nil
	}
	return bytes
}

// @ Update proximity/direction and Notify Lobby
func (mn *Node) Update(data []byte) {
	if mn.isReady() {
		// Initialize from Request
		update := &md.UpdateRequest{}
		if err := proto.Unmarshal(data, update); err != nil {
			mn.handleError(md.NewError(err, md.ErrorMessage_UNMARSHAL))
			return
		}

		// Update Peer
		mn.user.Update(update)

		// Notify Local Lobby
		err := mn.client.Update(mn.local)
		if err != nil {
			mn.handleError(err)
			return
		}
	}
}

// @ Invite Processes Data and Sends Invite to Peer
func (mn *Node) Invite(data []byte) {
	if mn.isReady() {
		// Update Status
		mn.setStatus(md.Status_PENDING)

		// Initialize from Request
		req := &md.InviteRequest{}
		if err := proto.Unmarshal(data, req); err != nil {
			mn.handleError(md.NewError(err, md.ErrorMessage_UNMARSHAL))
			return
		}

		// @ 1. Validate invite
		req = mn.user.ValidateInvite(req)

		// @ 2. Check Transfer Type
		if req.IsPayloadContact() {
			err := mn.client.InviteContact(req, mn.local, mn.user.Contact)
			if err != nil {
				mn.handleError(err)
				return
			}
		} else if req.IsPayloadUrl() {
			err := mn.client.InviteLink(req, mn.local)
			if err != nil {
				mn.handleError(err)
				return
			}
		} else {
			// Invite With file
			err := mn.client.InviteFile(req, mn.local)
			if err != nil {
				mn.handleError(err)
				return
			}
		}
	} else {
		log.Println("--- STATUS NOT READY: CANNOT (Invite) ---")
	}
}

// @ Respond to an Invite with Decision
func (mn *Node) Respond(data []byte) {
	if mn.isReady() {
		// Logging
		log.Println("--- Received Frontend Action for Response ---")

		// Initialize from Request
		req := &md.InviteResponse{}
		if err := proto.Unmarshal(data, req); err != nil {
			log.Println("--- FAILED: To Unmarshal Response ---")
			mn.handleError(md.NewError(err, md.ErrorMessage_UNMARSHAL))
			return
		}

		// Send Response
		mn.local.RespondToInvite(req)

		// Update Status
		if req.Decision {
			log.Println("--- Updated Status to Transfer ---")
			mn.setStatus(md.Status_TRANSFER)
		} else {
			log.Println("--- Updated Status to Available ---")
			mn.setStatus(md.Status_AVAILABLE)
		}
	} else {
		log.Println("--- STATUS NOT READY: CANNOT (Respond) ---")
	}
}

// @ Returns Node Location Protobuf as Bytes
func (mn *Node) Location() []byte {
	bytes, err := proto.Marshal(mn.user.Location)
	if err != nil {
		return nil
	}
	return bytes
}

// @ Returns Node User Protobuf as Bytes
func (mn *Node) User() []byte {
	bytes, err := proto.Marshal(mn.user)
	if err != nil {
		return nil
	}
	return bytes
}
