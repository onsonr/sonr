package bind

import (
	"context"
	"log"

	"github.com/getsentry/sentry-go"
	sc "github.com/sonr-io/core/pkg/client"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// @ Return URLLink
func URLLink(url string) []byte {
	// Create Link
	link := md.NewURLLink(url)

	// Marshal
	bytes, err := proto.Marshal(link)
	if err != nil {
		return nil
	}
	return bytes
}

// @ Gets User from Storj
func Storj(data []byte) []byte {
	// Unmarshal Request
	request := &md.StorjRequest{}
	proto.Unmarshal(data, request)

	switch request.Data.(type) {
	// @ Put USER
	case *md.StorjRequest_User:
		// Put User
		err := sc.PutUser(context.Background(), request.StorjApiKey, request.GetUser())
		if err != nil {
			sentry.CaptureException(err)

			// Create Response
			resp := &md.StorjResponse{
				Data: &md.StorjResponse_Success{
					Success: false,
				},
			}

			// Marshal
			bytes, err := proto.Marshal(resp)
			if err != nil {
				return nil
			}
			return bytes
		}
		// Create Response
		resp := &md.StorjResponse{
			Data: &md.StorjResponse_Success{
				Success: true,
			},
		}

		// Marshal
		bytes, err := proto.Marshal(resp)
		if err != nil {
			return nil
		}
		return bytes

	// @ Get USER
	case *md.StorjRequest_Prefix:
		// Get User from Uplink
		user, err := sc.GetUser(context.Background(), request.StorjApiKey, request.GetPrefix())
		if err != nil {
			sentry.CaptureException(err)
			return nil
		}

		// Create Response
		resp := &md.StorjResponse{
			Data: &md.StorjResponse_User{
				User: user,
			},
		}

		// Marshal
		bytes, err := proto.Marshal(resp)
		if err != nil {
			return nil
		}
		return bytes
	}
	return nil
}

// @ Create Remote Group
func (mn *Node) RemoteCreate(data []byte) []byte {
	if mn.isReady() {
		// Get Request
		request := &md.RemoteCreateRequest{}
		proto.Unmarshal(data, request)

		// Create Remote
		tm, resp, serr := mn.client.CreateRemote(request)
		if serr != nil {
			mn.handleError(serr)
			return nil
		}

		// Set Topic
		mn.topics[request.GetTopic()] = tm

		// Marshal
		buff, err := proto.Marshal(resp)
		if err != nil {
			mn.handleError(md.NewError(err, md.ErrorMessage_MARSHAL))
			return nil
		}
		return buff
	} else {
		log.Println("--- STATUS NOT READY: CANNOT (RemoteCreate) ---")
		return nil
	}
}

// @ Join Remote Group
func (mn *Node) RemoteJoin(data []byte) []byte {
	if mn.isReady() {
		// Get Request
		request := &md.RemoteJoinRequest{}
		err := proto.Unmarshal(data, request)
		if err != nil {
			return nil
		}

		// Create Remote
		tm, resp, serr := mn.client.JoinRemote(request)
		if serr != nil {
			mn.handleError(serr)
			return nil
		}

		// Set Topic
		mn.topics[request.GetTopic()] = tm

		// Marshal
		buff, err := proto.Marshal(resp)
		if err != nil {
			mn.handleError(md.NewError(err, md.ErrorMessage_MARSHAL))
			return nil
		}
		return buff
	} else {
		log.Println("--- STATUS NOT READY: CANNOT (RemoteJoin) ---")
		return nil
	}
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
