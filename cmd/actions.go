package bind

import (
	"context"

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
	}
	return nil
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
	}
	return nil
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
		req := &md.AuthInvite{}
		if err := proto.Unmarshal(data, req); err != nil {
			mn.handleError(md.NewError(err, md.ErrorMessage_UNMARSHAL))
			return
		}

		// @ 2. Check Transfer Type
		if req.Payload == md.Payload_CONTACT || req.Payload == md.Payload_FLAT_CONTACT {
			err := mn.client.InviteContact(req, mn.local, mn.user.Contact)
			if err != nil {
				mn.handleError(err)
				return
			}
		} else if req.Payload == md.Payload_URL {
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
	}
}

// @ Respond to an Invite with Decision
func (mn *Node) Respond(data []byte) {
	if mn.isReady() {
		// Initialize from Request
		req := &md.AuthReply{}
		if err := proto.Unmarshal(data, req); err != nil {
			mn.handleError(md.NewError(err, md.ErrorMessage_UNMARSHAL))
			return
		}

		mn.client.Respond(req, mn.local)
		// Update Status
		if req.Decision {
			mn.setStatus(md.Status_INPROGRESS)
		} else {
			mn.setStatus(md.Status_AVAILABLE)
		}
	}
}
