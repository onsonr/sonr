package bind

import (
	"context"

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
	client *sc.Client
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
func (mn *Node) Connect(data []byte) {
	// Unmarshal Request
	req := &md.ConnectionRequest{}
	err := proto.Unmarshal(data, req)
	if err != nil {
		sentry.CaptureException(errors.Wrap(err, "Unmarshalling Connection Request"))
	}

	// Update User with Connection Request
	mn.user.InitConnection(req)

	// Connect Host
	serr := mn.client.Connect(mn.user.APIKeys(), mn.user.KeyPair())
	if serr != nil {
		mn.handleError(serr)
		mn.setConnected(false)
	} else {
		// Update Status
		mn.setConnected(true)
	}

	// Bootstrap Node
	mn.local, serr = mn.client.Bootstrap()
	if err != nil {
		mn.handleError(serr)
		mn.setAvailable(false)
	} else {
		mn.setAvailable(true)
	}
}

// ** ─── Node Binded Actions ────────────────────────────────────────────────────────
// @ Signing Request for Data
func (mn *Node) Sign(data []byte) []byte {
	if mn.isReady() {
		// Unmarshal Data to Request
		request := &md.SignRequest{}
		if err := proto.Unmarshal(data, request); err != nil {
			// Handle Error
			mn.handleError(md.NewUnmarshalError(err))

			// Send Invalid Response
			if buf, err := proto.Marshal(md.NewInvalidSignResponse()); err != nil {
				mn.handleError(md.NewMarshalError(err))
				return nil
			} else {
				return buf
			}
		}

		// Initialize Result List
		signedList := make([][]byte, request.Count())

		// Get Key Pair
		kp := mn.user.KeyPair()
		if kp != nil {
			// Check Data Type
			if request.IsBuffers() {
				// Iterate Buffer Values
				for i, v := range request.BuffersList() {
					// Sign Buffer
					r, err := kp.Sign(v)
					if err != nil {
						break
					}

					// Set Value
					signedList[i] = r
				}
			} else if request.IsStrings() {
				// Iterate String Values
				for i, v := range request.StringsList() {
					// Sign String
					r, err := kp.Sign([]byte(v))
					if err != nil {
						break
					}

					// Set Value
					signedList[i] = r
				}
			}

			// Check if Validated
			if len(signedList) == request.Count() {
				// Send Valid Response
				if buf, err := proto.Marshal(md.NewValidSignResponse(signedList, request.IsStrings())); err != nil {
					mn.handleError(md.NewMarshalError(err))
					return nil
				} else {
					return buf
				}
			}
		}
	}

	// Send Invalid Response
	if buf, err := proto.Marshal(md.NewInvalidSignResponse()); err != nil {
		mn.handleError(md.NewMarshalError(err))
		return nil
	} else {
		return buf
	}
}

// @ Store Request for Payload into MemoryStore
func (mn *Node) Store(data []byte) []byte {
	// Unmarshal Data to Request
	request := &md.StoreRequest{}
	if err := proto.Unmarshal(data, request); err != nil {
		// Handle Error
		mn.handleError(md.NewUnmarshalError(err))

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
	resp := mn.store.Handle(request)

	// Marshal Data
	bytes, err := proto.Marshal(resp)
	if err != nil {
		return nil
	}
	return bytes
}

// @ Verification Request for Signed Data
func (mn *Node) Verify(data []byte) []byte {
	// Check Ready
	if mn.isReady() {
		// Get Key Pair
		kp := mn.user.KeyPair()

		// Unmarshal Data to Request
		request := &md.VerifyRequest{}
		if err := proto.Unmarshal(data, request); err != nil {
			// Handle Error
			mn.handleError(md.NewUnmarshalError(err))

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
func (mn *Node) Update(data []byte) {
	if mn.isReady() {
		// Unmarshal Data to Request
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

		// Unmarshal Data to Request
		req := &md.InviteRequest{}
		if err := proto.Unmarshal(data, req); err != nil {
			mn.handleError(md.NewError(err, md.ErrorMessage_UNMARSHAL))
			return
		}

		// @ 1. Validate invite
		req = mn.user.ValidateInvite(req)

		// @ 2. Check Transfer Type
		if req.IsPayloadContact() {
			err := mn.client.InviteContact(req, mn.local, req.GetContact())
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
	}
}

// @ Respond to an Invite with Decision
func (mn *Node) Respond(data []byte) {
	if mn.isReady() {
		// Unmarshal Data to Request
		req := &md.InviteResponse{}
		if err := proto.Unmarshal(data, req); err != nil {

			mn.handleError(md.NewError(err, md.ErrorMessage_UNMARSHAL))
			return
		}

		// Send Response
		mn.local.RespondToInvite(req)

		// Update Status
		if req.Decision {

			mn.setStatus(md.Status_TRANSFER)
		} else {

			mn.setStatus(md.Status_AVAILABLE)
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
func (mn *Node) Location() []byte {
	bytes, err := proto.Marshal(mn.user.Location)
	if err != nil {
		return nil
	}
	return bytes
}

// ** ─── LifeCycle Actions ────────────────────────────────────────────────────────
// @ Close Ends All Network Communication
func (mn *Node) Pause() {
	md.GetState().Pause()
}

// @ Close Ends All Network Communication
func (mn *Node) Resume() {
	md.GetState().Resume()
}

// @ Close Ends All Network Communication
func (mn *Node) Stop() {
	mn.client.Close()
	mn.ctx.Done()
}

// ** ─── Node Status Checks ────────────────────────────────────────────────────────
// # Checks if Node is Ready for Actions
func (mn *Node) isReady() bool {
	return mn.user.IsNotStatus(md.Status_STANDBY) || mn.user.IsNotStatus(md.Status_FAILED)
}

// # Sets Node to be Connected Status
func (mn *Node) setConnected(val bool) {
	// Update Status
	su := mn.user.SetConnected(val)

	// Callback Status
	data, err := proto.Marshal(su)
	if err != nil {
		mn.handleError(md.NewError(err, md.ErrorMessage_MARSHAL))
		return
	}
	mn.call.OnStatus(data)
}

// # Sets Node to be Available Status
func (mn *Node) setAvailable(val bool) {
	// Update Status
	su := mn.user.SetAvailable(val)

	// Callback Status
	data, err := proto.Marshal(su)
	if err != nil {
		mn.handleError(md.NewError(err, md.ErrorMessage_MARSHAL))
		return
	}
	mn.call.OnStatus(data)
}

// # Sets Node to be (Provided) Status
func (mn *Node) setStatus(newStatus md.Status) {
	// Set Status
	su := mn.user.SetStatus(newStatus)

	// Callback Status
	data, err := proto.Marshal(su)
	if err != nil {
		mn.handleError(md.NewError(err, md.ErrorMessage_MARSHAL))
		return
	}
	mn.call.OnStatus(data)
}
