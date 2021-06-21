package bind

import (
	"context"
	"log"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	tpc "github.com/sonr-io/core/internal/topic"
	sc "github.com/sonr-io/core/pkg/client"
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
	user   *md.User

	// Groups
	local  *tpc.TopicManager
	topics map[string]*tpc.TopicManager

	// Miscellaneous
	store md.Store
}

// ^ Initializes New Node ^ //
func NewNode(reqBytes []byte, call Callback) *Node {
	// Initialize Sentry
	sentry.Init(sentry.ClientOptions{
		Dsn: "https://cbf88b01a5a5468fa77101f7dfc54f20@o549479.ingest.sentry.io/5672329",
	})

	// Unmarshal Request
	req := &md.InitializeRequest{}
	err := proto.Unmarshal(reqBytes, req)
	if err != nil {
		sentry.CaptureException(errors.Wrap(err, "Unmarshalling Connection Request"))
		return nil
	}
	// Initialize Node
	mn := &Node{
		call:   call,
		ctx:    context.Background(),
		topics: make(map[string]*tpc.TopicManager, 10),
	}

	// Create Store - Start Auth Service
	if s, err := md.InitStore(req.GetDevice()); err != nil {
		mn.handleError(err)
	} else {
		mn.store = s
	}

	// Create User
	if u, err := md.NewUser(req, mn.store); err != nil {
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
		sentry.CaptureException(errors.Wrap(err, "Unmarshalling Connection Request"))
	}

	// Update User with Connection Request
	n.user.InitConnection(req)
	req.ApiKeys = n.user.APIKeys()
	req.Point = n.user.GetRouter().Rendevouz

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
	if err != nil {
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
		log.Println("Failed to Unmarshal Sign Request")

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
	result := n.user.Sign(request, n.store)
	buf, err := proto.Marshal(result)
	if err != nil {
		n.handleError(md.NewMarshalError(err))
		return nil
	}
	return buf
}

// @ Store Request for Payload into MemoryStore
func (n *Node) Store(data []byte) []byte {
	// Unmarshal Data to Request
	request := &md.StoreRequest{}
	if err := proto.Unmarshal(data, request); err != nil {
		// Handle Error
		n.handleError(md.NewUnmarshalError(err))

		// Create Error Response
		resp := &md.StoreResponse{
			Error: md.NewUnmarshalError(err).Message(),
		}

		// Marshal Data
		bytes, err := proto.Marshal(resp)
		if err != nil {
			return nil
		}

		// Send Invalid Response
		return bytes
	}

	// Handle Request with Store
	resp := n.store.Handle(request)

	// Marshal Data
	bytes, err := proto.Marshal(resp)
	if err != nil {
		return nil
	}
	return bytes
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
			n.handleError(md.NewError(err, md.ErrorMessage_UNMARSHAL))
			return
		}

		// Update Peer
		n.user.Update(update)

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
		// Update Status
		n.setStatus(md.Status_PENDING)

		// Unmarshal Data to Request
		req := &md.InviteRequest{}
		if err := proto.Unmarshal(data, req); err != nil {
			n.handleError(md.NewError(err, md.ErrorMessage_UNMARSHAL))
			return
		}

		// @ 1. Validate invite
		req = n.user.ValidateInvite(req)

		// @ 2. Check Transfer Type
		err := n.client.Invite(req, n.local)
		if err != nil {
			n.handleError(err)
			return
		}
	}
}

// @ Respond to an Invite with Decision
func (n *Node) Respond(data []byte) {
	if n.isReady() {
		// Unmarshal Data to Request
		req := &md.InviteResponse{}
		if err := proto.Unmarshal(data, req); err != nil {

			n.handleError(md.NewError(err, md.ErrorMessage_UNMARSHAL))
			return
		}

		// Send Response
		n.local.RespondToInvite(req)

		// Update Status
		if req.Decision {

			n.setStatus(md.Status_TRANSFER)
		} else {

			n.setStatus(md.Status_AVAILABLE)
		}
	}
}

// ** ─── Misc Methods ────────────────────────────────────────────────────────
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

// @ Returns Node Location Protobuf as Bytes
func (n *Node) Location() []byte {
	bytes, err := proto.Marshal(n.user.Router.GetLocation())
	if err != nil {
		return nil
	}
	return bytes
}

// ** ─── LifeCycle Actions ────────────────────────────────────────────────────────
// @ Close Ends All Network Communication
func (n *Node) Pause() {
	md.GetState().Pause()
}

// @ Close Ends All Network Communication
func (n *Node) Resume() {
	md.GetState().Resume()
}

// @ Close Ends All Network Communication
func (n *Node) Stop() {
	n.client.Close(n.local)
	n.ctx.Done()
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
		n.handleError(md.NewError(err, md.ErrorMessage_MARSHAL))
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
		n.handleError(md.NewError(err, md.ErrorMessage_MARSHAL))
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
		n.handleError(md.NewError(err, md.ErrorMessage_MARSHAL))
		return
	}
	n.call.OnStatus(data)
}
