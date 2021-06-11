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

// ** ─── Core Main Access Node ────────────────────────────────────────────────────────
type Node struct {
	md.Callback

	// Properties
	call    Callback
	ctx     context.Context
	request *md.ConnectionRequest

	// Client
	client *sc.Client
	user   *md.User

	// Groups
	local  *tpc.TopicManager
	topics map[string]*tpc.TopicManager

	// Miscellaneous
	store md.Store
}

// @ Initializes New Node
func NewNode(reqBytes []byte, call Callback) *Node {
	// Initialize Sentry
	sentry.Init(sentry.ClientOptions{
		Dsn: "https://cbf88b01a5a5468fa77101f7dfc54f20@o549479.ingest.sentry.io/5672329",
	})

	// Unmarshal Request
	req := &md.ConnectionRequest{}
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
	mn.initialize(req)
	return mn
}

// ** ─── Node Initializers ────────────────────────────────────────────────────────
// # Initializes Node with Client and Memory Store
func (mn *Node) initialize(req *md.ConnectionRequest) {
	// Set Type
	mn.request = req

	// Create Store - Start Auth Service
	if s, err := md.InitStore(req.GetDevice()); err == nil {
		mn.store = s
	}

	// Create User
	if u, err := md.NewUser(req, mn.store); err == nil {
		mn.user = u
	}

	// Create Client
	mn.client = sc.NewClient(mn.ctx, mn.user, mn.callback())
}

// @ Starts Host and Connects
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

// @ Returns Node Location Protobuf as Bytes
func (mn *Node) Location() []byte {
	bytes, err := proto.Marshal(mn.user.Location)
	if err != nil {
		return nil
	}
	return bytes
}

// ** ─── Node Binded Actions ────────────────────────────────────────────────────────
// @ Signing Request for Data
func (mn *Node) Sign(data []byte) []byte {
	if mn.isReady() {
		// Unmarshal Data to Request
		request := &md.SignRequest{}
		if err := proto.Unmarshal(data, request); err != nil {
			mn.handleError(md.NewUnmarshalError(err))
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
	} else {
		log.Println("--- STATUS NOT READY: CANNOT (Invite) ---")
	}
}

// @ Respond to an Invite with Decision
func (mn *Node) Respond(data []byte) {
	if mn.isReady() {
		// Logging
		log.Println("--- Received Frontend Action for Response ---")

		// Unmarshal Data to Request
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
