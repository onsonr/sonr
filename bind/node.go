package bind

import (
	"context"

	sc "github.com/sonr-io/core/internal/client"
	net "github.com/sonr-io/core/internal/host"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

type Node struct {
	md.Callback

	// Properties
	call Callback
	ctx  context.Context

	// Client
	client sc.Client
	state  md.Lifecycle
	user   *md.User

	// Groups
	local  *net.TopicManager
	topics map[string]*net.TopicManager
}

// ^ Initializes New Node ^ //
func Initialize(reqBytes []byte, call Callback) *Node {
	// Unmarshal Request
	req := &md.InitializeRequest{}
	err := proto.Unmarshal(reqBytes, req)
	if err != nil {
		md.LogFatal(err)
		return nil
	}

	// Initialize Logger
	md.InitLogger(req)

	// Initialize Node
	mn := &Node{
		call:   call,
		ctx:    context.Background(),
		topics: make(map[string]*net.TopicManager, 10),
		state:  md.Lifecycle_ACTIVE,
	}

	// Create User
	if u, err := md.NewUser(req); err != nil {
		mn.handleError(err)
	} else {
		mn.user = u
	}

	// Create Client
	mn.client = sc.NewClient(mn.ctx, mn.user, mn.callback())
	return mn
}

// @ Starts Host and Connects
func (n *Node) Connect(data []byte) {
	// Unmarshal Request
	req := &md.ConnectionRequest{}
	err := proto.Unmarshal(data, req)
	if err != nil {
		md.LogFatal(err)
	}

	// Update User with Connection Request
	n.user.InitConnection(req)

	// Connect Host
	serr := n.client.Connect(req, n.user.KeyPair())
	if serr != nil {
		n.handleError(serr)
		n.setConnected(false)
	} else {
		// Update Status
		n.setConnected(true)
	}

	// Bootstrap Node
	n.local, serr = n.client.Bootstrap()
	if serr != nil {
		n.handleError(serr)
		n.setAvailable(false)
	} else {
		n.setAvailable(true)
	}
}

// ** ─── Node Binded Actions ────────────────────────────────────────────────────────
// @ Signing Request for Data
func (n *Node) Sign(data []byte) []byte {
	// Unmarshal Data to Request
	request := &md.AuthRequest{}
	err := proto.Unmarshal(data, request)
	if err != nil {
		// Handle Error
		n.handleError(md.NewUnmarshalError(err))

		// Initialize invalid Response
		invalidResp := md.AuthResponse{
			IsSigned: false,
		}

		// Send Invalid Response
		buf, err := proto.Marshal(&invalidResp)
		if err != nil {
			n.handleError(md.NewMarshalError(err))
			return nil
		}
		return buf
	}

	// Sign Buffer
	result := n.user.Sign(request)
	buf, err := proto.Marshal(result)
	if err != nil {
		n.handleError(md.NewMarshalError(err))
		return nil
	}
	return buf
}

// @ Verification Request for Signed Data
func (n *Node) Verify(data []byte) []byte {
	// Check Ready
	if n.isReady() {
		// Get Key Pair
		kp := n.user.KeyPair()

		// Unmarshal Data to Request
		request := &md.VerifyRequest{}
		if err := proto.Unmarshal(data, request); err != nil {
			// Handle Error
			n.handleError(md.NewUnmarshalError(err))

			// Send Invalid Response
			return md.NewInvalidVerifyResponseBuf()
		}

		// Check Request Type
		if request.GetType() == md.VerifyRequest_VERIFY {
			// Check type and Verify
			if request.IsBuffer() {
				// Verify Result
				result, err := kp.Verify(request.GetBufferValue(), request.GetSignedBuffer())
				if err != nil {
					return md.NewInvalidVerifyResponseBuf()
				}

				// Return Result
				return md.NewVerifyResponseBuf(result)
			} else if request.IsString() {
				// Verify Result
				result, err := kp.Verify([]byte(request.GetTextValue()), []byte(request.GetSignedText()))
				if err != nil {
					return md.NewInvalidVerifyResponseBuf()
				}

				// Return Result
				return md.NewVerifyResponseBuf(result)
			}
		} else {
			resp := n.user.VerifyRead()
			buf, err := proto.Marshal(resp)
			if err != nil {
				n.handleError(md.NewMarshalError(err))
			}
			return buf
		}

	}

	// Send Invalid Response
	return md.NewInvalidVerifyResponseBuf()
}

// @ Update proximity/direction and Notify Lobby
func (n *Node) Update(data []byte) {
	if n.isReady() {
		// Unmarshal Data to Request
		update := &md.UpdateRequest{}
		if err := proto.Unmarshal(data, update); err != nil {
			n.handleError(md.NewError(err, md.ErrorEvent_UNMARSHAL))
			return
		}

		// Check Update Request Type
		switch update.Data.(type) {
		// Update Position
		case *md.UpdateRequest_Position:
			n.user.UpdatePosition(update.GetPosition())

		// Update Contact
		case *md.UpdateRequest_Contact:
			n.user.UpdateContact(update.GetContact())

		// Update Peer Properties
		case *md.UpdateRequest_Properties:
			n.user.UpdateProperties(update.GetProperties())
		}

		// Notify Local Lobby
		err := n.client.Update(n.local)
		if err != nil {
			n.handleError(err)
			return
		}
	}
}

// @ Invite Processes Data and Sends Invite to Peer
func (n *Node) Invite(data []byte) {
	if n.isReady() {
		// Unmarshal Data to Request
		req := &md.InviteRequest{}
		if err := proto.Unmarshal(data, req); err != nil {
			n.handleError(md.NewError(err, md.ErrorEvent_UNMARSHAL))
			return
		}

		// Validate invite
		req = n.user.SignInvite(req)

		// Send Invite
		err := n.client.Invite(req, n.local)
		if err != nil {
			n.handleError(err)
			return
		}
	}
}

// @ Mail handles request for a message in Mailbox
func (n *Node) Mail(data []byte) []byte {
	// Check Ready
	if n.isReady() {
		// Unmarshal Data to Request
		req := &md.MailboxRequest{}
		if err := proto.Unmarshal(data, req); err != nil {
			n.handleError(md.NewUnmarshalError(err))
			return nil
		}

		// Handle Mail
		resp, serr := n.client.Mail(req)
		if serr != nil {
			n.handleError(serr)
			return nil
		}

		// Marshal Response
		buf, err := proto.Marshal(resp)
		if err != nil {
			n.handleError(md.NewMarshalError(err))
			return nil
		}

		// Return Response
		return buf
	}
	return nil
}

// @ Respond to an Invite with Decision
func (n *Node) Respond(data []byte) {
	if n.isReady() {
		// Unmarshal Data to Request
		resp := &md.InviteResponse{}
		if err := proto.Unmarshal(data, resp); err != nil {
			n.handleError(md.NewError(err, md.ErrorEvent_UNMARSHAL))
			return
		}

		// Send Response
		n.client.Respond(resp)

		// Update Status
		if resp.Decision {
			n.setStatus(md.Status_TRANSFER)
		} else {
			n.setStatus(md.Status_AVAILABLE)
		}
	}
}

// ** ─── Misc Methods ────────────────────────────────────────────────────────
// Action method handles misceallaneous actions for node
func (s *Node) Action(buf []byte) []byte {
	// Unmarshal Data to Request
	req := &md.ActionRequest{}
	if err := proto.Unmarshal(buf, req); err != nil {
		s.handleError(md.NewError(err, md.ErrorEvent_UNMARSHAL))
		return nil
	}

	// Check Action
	switch req.Action {
	case md.Action_PING:
		// Ping
		resp := &md.ActionResponse{
			Success: true,
			Action:  md.Action_PING,
		}

		// Marshal Response
		bytes, err := proto.Marshal(resp)
		if err != nil {
			s.handleError(md.NewMarshalError(err))
			return nil
		}
		return bytes
	case md.Action_LOCATION:
		// Location
		resp := &md.ActionResponse{
			Success: true,
			Action:  md.Action_LOCATION,
			Data: &md.ActionResponse_Location{
				Location: s.user.GetLocation(),
			},
		}

		// Marshal Response
		bytes, err := proto.Marshal(resp)
		if err != nil {
			s.handleError(md.NewMarshalError(err))
			return nil
		}
		return bytes
	case md.Action_URL_LINK:
		// URL Link
		resp := &md.ActionResponse{
			Success: true,
			Action:  md.Action_URL_LINK,
			Data: &md.ActionResponse_UrlLink{
				UrlLink: md.NewURLLink(req.GetData()),
			},
		}

		// Marshal Response
		bytes, err := proto.Marshal(resp)
		if err != nil {
			s.handleError(md.NewMarshalError(err))
			return nil
		}
		return bytes
	case md.Action_PAUSE:
		// Pause
		md.LogInfo("Lifecycle Pause Called")
		s.state = md.Lifecycle_PAUSED
		s.client.Lifecycle(s.state, s.local)
		md.GetState().Pause()
		resp := &md.ActionResponse{
			Success: true,
			Action:  md.Action_PAUSE,
			Data: &md.ActionResponse_Lifecycle{
				Lifecycle: s.state,
			},
		}

		// Marshal Response
		bytes, err := proto.Marshal(resp)
		if err != nil {
			s.handleError(md.NewMarshalError(err))
			return nil
		}
		return bytes
	case md.Action_RESUME:
		// Resume
		md.LogInfo("Lifecycle Resume Called")
		s.state = md.Lifecycle_ACTIVE
		s.client.Lifecycle(s.state, s.local)
		md.GetState().Resume()

		// Create Response
		resp := &md.ActionResponse{
			Success: true,
			Action:  md.Action_RESUME,
			Data: &md.ActionResponse_Lifecycle{
				Lifecycle: s.state,
			},
		}

		// Marshal Response
		bytes, err := proto.Marshal(resp)
		if err != nil {
			s.handleError(md.NewMarshalError(err))
			return nil
		}
		return bytes
	case md.Action_STOP:
		// Stop
		md.LogInfo("Lifecycle Stop Called")
		s.state = md.Lifecycle_STOPPED
		s.client.Lifecycle(s.state, s.local)

		// Create Response
		resp := &md.ActionResponse{
			Success: true,
			Action:  md.Action_STOP,
			Data: &md.ActionResponse_Lifecycle{
				Lifecycle: s.state,
			},
		}

		// Marshal Response
		bytes, err := proto.Marshal(resp)
		if err != nil {
			s.handleError(md.NewMarshalError(err))
			return nil
		}
		return bytes
	}
	return nil
}

// ** ─── Node Status Checks ────────────────────────────────────────────────────────
// # Checks if Node is Ready for Actions
func (n *Node) isReady() bool {
	return n.user.IsNotStatus(md.Status_STANDBY) || n.user.IsNotStatus(md.Status_FAILED)
}

// # Sets Node to be Connected Status
func (n *Node) setConnected(val bool) {
	// Update Status
	su := n.user.SetConnected(val)

	// Callback Status
	data, err := proto.Marshal(su)
	if err != nil {
		n.handleError(md.NewError(err, md.ErrorEvent_MARSHAL))
		return
	}
	n.call.OnStatus(data)
}

// # Sets Node to be Available Status
func (n *Node) setAvailable(val bool) {
	// Update Status
	su := n.user.SetAvailable(val)

	// Callback Status
	data, err := proto.Marshal(su)
	if err != nil {
		n.handleError(md.NewError(err, md.ErrorEvent_MARSHAL))
		return
	}
	n.call.OnStatus(data)
}

// # Sets Node to be (Provided) Status
func (n *Node) setStatus(newStatus md.Status) {
	// Set Status
	su := n.user.SetStatus(newStatus)

	// Callback Status
	data, err := proto.Marshal(su)
	if err != nil {
		n.handleError(md.NewError(err, md.ErrorEvent_MARSHAL))
		return
	}
	n.call.OnStatus(data)
}
