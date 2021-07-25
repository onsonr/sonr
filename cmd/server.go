package main

import (
	"context"
	"fmt"
	"log"
	"net"

	sc "github.com/sonr-io/core/internal/client"
	sh "github.com/sonr-io/core/internal/host"
	md "github.com/sonr-io/core/pkg/models"
	"github.com/sonr-io/core/pkg/util"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

type NodeServer struct {
	md.NodeServiceServer
	ctx context.Context

	// Client
	client sc.Client
	state  md.Lifecycle
	user   *md.User

	// Groups
	local  *sh.TopicManager
	topics map[string]*sh.TopicManager

	// Callback Channels
	connectionResponses chan *md.ConnectionResponse
	completeEvents      chan *md.CompleteEvent
	inviteRequests      chan *md.InviteRequest
	inviteResponses     chan *md.InviteResponse
	errorMessages       chan *md.ErrorMessage
	mailEvents          chan *md.MailEvent
	progressEvents      chan *md.ProgressEvent
	statusEvents        chan *md.StatusEvent
	topicEvents         chan *md.TopicEvent
}

func main() {
	// Find Open Port
	port, err := sh.FreePort()
	if err != nil {
		port = 9000
		md.LogFatal(err)
	}
	md.LogInfo(fmt.Sprintf("(SONR_RPC)-PORT=%d", port))

	// Create a new gRPC server
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		md.LogInfo("(SONR_RPC)-ONLINE=false")
		log.Fatal(err)
	}
	md.LogInfo("(SONR_RPC)-ONLINE=true")

	// Set GRPC Server
	chatServer := NodeServer{
		ctx:                 context.Background(),
		topics:              make(map[string]*sh.TopicManager, 10),
		state:               md.Lifecycle_Active,
		topicEvents:         make(chan *md.TopicEvent, util.MAX_CHAN_DATA),
		mailEvents:          make(chan *md.MailEvent, util.MAX_CHAN_DATA),
		progressEvents:      make(chan *md.ProgressEvent, util.MAX_CHAN_DATA),
		completeEvents:      make(chan *md.CompleteEvent, util.MAX_CHAN_DATA),
		statusEvents:        make(chan *md.StatusEvent, util.MAX_CHAN_DATA),
		errorMessages:       make(chan *md.ErrorMessage, util.MAX_CHAN_DATA),
		inviteRequests:      make(chan *md.InviteRequest, util.MAX_CHAN_DATA),
		inviteResponses:     make(chan *md.InviteResponse, util.MAX_CHAN_DATA),
		connectionResponses: make(chan *md.ConnectionResponse, util.MAX_CHAN_DATA),
	}

	grpcServer := grpc.NewServer()

	// Register the gRPC service
	md.RegisterNodeServiceServer(grpcServer, &chatServer)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

// Action method handles misceallaneous actions for node
func (s *NodeServer) Action(ctx context.Context, req *md.ActionRequest) (*md.ActionResponse, error) {
	// Check Action
	switch req.Action {
	case md.Action_PING:
		// Ping
		md.LogInfo("Action: Ping Called")
		return &md.ActionResponse{
			Success: true,
			Action:  md.Action_PING,
		}, nil
	case md.Action_LOCATION:
		// Location
		md.LogInfo("Action: Location Called")
		return &md.ActionResponse{
			Success: true,
			Action:  md.Action_LOCATION,
			Data: &md.ActionResponse_Location{
				Location: s.user.GetLocation(),
			},
		}, nil
	case md.Action_URL_LINK:
		// URL Link
		md.LogInfo("Action: URL Link Called")
		return &md.ActionResponse{
			Success: true,
			Action:  md.Action_URL_LINK,
			Data: &md.ActionResponse_UrlLink{
				UrlLink: md.NewURLLink(req.GetData()),
			},
		}, nil
	case md.Action_PAUSE:
		// Pause
		md.LogInfo("Action: Lifecycle Pause Called")
		s.state = md.Lifecycle_Paused
		s.client.Lifecycle(s.state, s.local)
		md.GetState().Pause()
		return &md.ActionResponse{
			Success: true,
			Action:  md.Action_PAUSE,
			Data: &md.ActionResponse_Lifecycle{
				Lifecycle: s.state,
			},
		}, nil
	case md.Action_RESUME:
		// Resume
		md.LogInfo("Action: Lifecycle Resume Called")
		s.state = md.Lifecycle_Active
		s.client.Lifecycle(s.state, s.local)
		md.GetState().Resume()

		return &md.ActionResponse{
			Success: true,
			Action:  md.Action_RESUME,
			Data: &md.ActionResponse_Lifecycle{
				Lifecycle: s.state,
			},
		}, nil
	case md.Action_STOP:
		// Stop
		md.LogInfo("Lifecycle Stop Called")
		s.state = md.Lifecycle_Stopped
		s.client.Lifecycle(s.state, s.local)
		return &md.ActionResponse{
			Success: true,
			Action:  md.Action_STOP,
			Data: &md.ActionResponse_Lifecycle{
				Lifecycle: s.state,
			},
		}, nil
	}
	return nil, nil
}

// Initialize method is called when a new node is created
func (s *NodeServer) Initialize(ctx context.Context, req *md.InitializeRequest) (*md.GenericResponse, error) {
	// Initialize Logger
	md.InitLogger(req)

	// Create User
	if u, err := md.NewUser(req); err != nil {
		s.handleError(err)
		return nil, err.Error
	} else {
		s.user = u
	}

	// Create Client
	s.client = sc.NewClient(s.ctx, s.user, s.callback())

	// Return Blank Response
	return nil, nil
}

// Connect method starts this nodes host
func (s *NodeServer) Connect(ctx context.Context, req *md.ConnectionRequest) (*md.GenericResponse, error) {
	// Update User with Connection Request
	s.user.InitConnection(req)

	// Connect Host
	serr := s.client.Connect(req, s.user.KeyPair())
	if serr != nil {
		s.handleError(serr)
		s.setConnected(false)
	} else {
		// Update Status
		s.setConnected(true)
	}

	// Bootstrap Node
	s.local, serr = s.client.Bootstrap()
	if serr != nil {
		s.handleError(serr)
		s.setAvailable(false)
	} else {
		s.setAvailable(true)
	}

	// Return Blank Response
	return nil, nil
}

// Sign method signs data with user's private key
func (s *NodeServer) Sign(ctx context.Context, req *md.AuthRequest) (*md.AuthResponse, error) {
	md.LogInfo("Sign Called")
	return s.user.Sign(req), nil
}

// Verify validates user Keys
func (s *NodeServer) Verify(ctx context.Context, req *md.VerifyRequest) (*md.VerifyResponse, error) {
	// Verify Node is Ready
	if s.isReady() {
		// Get Key Pair
		kp := s.user.KeyPair()

		// Check Request Type
		if req.GetType() == md.VerifyRequest_VERIFY {
			// Check type and Verify
			if req.IsBuffer() {
				// Verify Result
				result, err := kp.Verify(req.GetBufferValue(), req.GetSignedBuffer())
				if err != nil {
					return &md.VerifyResponse{IsVerified: false}, err
				}

				// Return Result
				return &md.VerifyResponse{IsVerified: result}, nil
			} else if req.IsString() {
				// Verify Result
				result, err := kp.Verify([]byte(req.GetTextValue()), []byte(req.GetSignedText()))
				if err != nil {
					return &md.VerifyResponse{IsVerified: false}, err
				}

				// Return Result
				return &md.VerifyResponse{IsVerified: result}, nil
			}
		} else {
			return s.user.VerifyRead(), nil
		}
	}
	return &md.VerifyResponse{IsVerified: false}, fmt.Errorf("Node is not ready")
}

// Update proximity/direction/contact/properties and notify Lobby
func (s *NodeServer) Update(ctx context.Context, req *md.UpdateRequest) (*md.GenericResponse, error) {
	// Verify Node is Ready
	if s.isReady() {
		// Check Update Request Type
		switch req.Data.(type) {
		// Update Position
		case *md.UpdateRequest_Position:
			s.user.UpdatePosition(req.GetPosition())

		// Update Contact
		case *md.UpdateRequest_Contact:
			s.user.UpdateContact(req.GetContact())

		// Update Peer Properties
		case *md.UpdateRequest_Properties:
			s.user.UpdateProperties(req.GetProperties())

		// Restart Connection
		case *md.UpdateRequest_Connectivity:
			local, err := s.client.Restart(req, s.user.KeyPair())
			if err != nil {
				s.handleError(err)
			}
			s.local = local
		}

		// Notify Local Lobby
		err := s.client.Update(s.local)
		if err != nil {
			s.handleError(err)
			return nil, err.Error
		}
	}

	// Return Blank Response
	return nil, nil
}

// Invite pushes Invite request to Peer
func (s *NodeServer) Invite(ctx context.Context, req *md.InviteRequest) (*md.GenericResponse, error) {
	// Verify Node is Ready
	if s.isReady() {
		// Validate invite
		req = s.user.SignInvite(req)

		// Send Invite
		err := s.client.Invite(req, s.local)
		if err != nil {
			s.handleError(err)
			return nil, err.Error
		}
	}
	// Return Blank Response
	return nil, nil
}

// Respond handles a respond request
func (s *NodeServer) Respond(ctx context.Context, req *md.InviteResponse) (*md.GenericResponse, error) {
	// Verify Node is Ready
	if s.isReady() {
		// Send Response
		s.client.Respond(req)

		// Update Status
		if req.Decision {
			s.setStatus(md.Status_TRANSFER)
		} else {
			s.setStatus(md.Status_AVAILABLE)
		}

		// Return Blank Response
		return nil, nil
	}
	return nil, fmt.Errorf("Node is not ready")
}

// Mail method handles a mail request
func (s *NodeServer) Mail(ctx context.Context, req *md.MailboxRequest) (*md.MailboxResponse, error) {
	// Verify Node is Ready
	if s.isReady() {
		// Handle Mail
		resp, serr := s.client.Mail(req)
		if serr != nil {
			s.handleError(serr)
			return nil, serr.Error
		}

		// Return Response
		return resp, nil
	}
	return nil, fmt.Errorf("Node is not ready")
}

// OnComplete is called when a complete event is received
func (s *NodeServer) OnComplete(req *md.GenericRequest, stream md.NodeService_OnCompleteServer) error {
	for {
		select {
		case m := <-s.completeEvents:
			stream.Send(m)
		case <-s.ctx.Done():
			return nil
		}
		md.GetState().NeedsWait()
	}
}

// OnInvite is called when user is invited by a Peer
func (s *NodeServer) OnInvite(req *md.GenericRequest, stream md.NodeService_OnInviteServer) error {
	for {
		select {
		case m := <-s.inviteRequests:
			stream.Send(m)
		case <-s.ctx.Done():
			return nil
		}
		md.GetState().NeedsWait()
	}
}

// OnReply is called when a peer responds to invite
func (s *NodeServer) OnReply(req *md.GenericRequest, stream md.NodeService_OnReplyServer) error {
	for {
		select {
		case m := <-s.inviteResponses:
			stream.Send(m)
		case <-s.ctx.Done():
			return nil
		}
		md.GetState().NeedsWait()
	}
}

// OnMail is called when a new mail is received
func (s *NodeServer) OnMail(req *md.GenericRequest, stream md.NodeService_OnMailServer) error {
	for {
		select {
		case m := <-s.mailEvents:
			stream.Send(m)
		case <-s.ctx.Done():
			return nil
		}
		md.GetState().NeedsWait()
	}
}

// OnProgress is called when a file is being transferred
func (s *NodeServer) OnProgress(req *md.GenericRequest, stream md.NodeService_OnProgressServer) error {
	for {
		select {
		case m := <-s.progressEvents:
			stream.Send(m)
		case <-s.ctx.Done():
			return nil
		}
		md.GetState().NeedsWait()
	}
}

// OnStatus is called when the node receives a status event
func (s *NodeServer) OnStatus(req *md.GenericRequest, stream md.NodeService_OnStatusServer) error {
	for {
		select {
		case m := <-s.statusEvents:
			stream.Send(m)
		case <-s.ctx.Done():
			return nil
		}
		md.GetState().NeedsWait()
	}
}

// OnTopic is called when Topic Event is received
func (s *NodeServer) OnTopic(req *md.GenericRequest, stream md.NodeService_OnTopicServer) error {
	for {
		select {
		case m := <-s.topicEvents:
			stream.Send(m)
		case <-s.ctx.Done():
			return nil
		}
		md.GetState().NeedsWait()
	}
}

// # Passes binded Methods to Node
func (s *NodeServer) callback() md.Callback {
	return md.Callback{
		OnEvent:    s.handleEvent,
		OnRequest:  s.handleRequest,
		OnResponse: s.handleResponse,
		OnError:    s.handleError,
		SetStatus:  s.setStatus,
	}
}

// ** ─── Node Status Checks ────────────────────────────────────────────────────────
// # Checks if Node is Ready for Actions
func (n *NodeServer) isReady() bool {
	return n.user.IsNotStatus(md.Status_STANDBY) || n.user.IsNotStatus(md.Status_FAILED)
}

// # Sets Node to be Connected Status
func (s *NodeServer) setConnected(val bool) {
	// Update Status
	su := s.user.SetConnected(val)

	// Callback Status
	s.statusEvents <- su
}

// # Sets Node to be Available Status
func (s *NodeServer) setAvailable(val bool) {
	// Update Status
	su := s.user.SetAvailable(val)

	// Callback Status
	s.statusEvents <- su
}

// # Sets Node to be (Provided) Status
func (s *NodeServer) setStatus(newStatus md.Status) {
	// Set Status
	su := s.user.SetStatus(newStatus)

	// Callback Status
	s.statusEvents <- su
}

// Handle Event and Send to Channel after unmarshal
func (s *NodeServer) handleEvent(buf []byte) {
	// Unmarshal Generic Event
	event := &md.GenericEvent{}
	err := proto.Unmarshal(buf, event)
	if err != nil {
		md.LogFatal(err)
		return
	}

	// Switch case event type
	switch event.GetType() {
	case md.GenericEvent_COMPLETE:
		// Unmarshal Complete Event
		ce := &md.CompleteEvent{}
		err = proto.Unmarshal(event.GetData(), ce)
		if err != nil {
			md.LogFatal(err)
			return
		}

		// Send Event to Channel
		s.completeEvents <- ce
	case md.GenericEvent_PROGRESS:
		// Unmarshal Progress Event
		pe := &md.ProgressEvent{}
		err = proto.Unmarshal(event.GetData(), pe)
		if err != nil {
			md.LogFatal(err)
			return
		}

		// Send Event to Channel
		s.progressEvents <- pe
	case md.GenericEvent_TOPIC:
		// Unmarshal Topic Event
		te := &md.TopicEvent{}
		err = proto.Unmarshal(event.GetData(), te)
		if err != nil {
			md.LogFatal(err)
			return
		}

		// Send Event to Channel
		s.topicEvents <- te

	case md.GenericEvent_MAIL:
		// Unmarshal Mail Event
		me := &md.MailEvent{}
		err = proto.Unmarshal(event.GetData(), me)
		if err != nil {
			md.LogFatal(err)
			return
		}

		// Send Event to Channel
		s.mailEvents <- me
	}
}

// Handle Request and Send to Channel after unmarshal
func (s *NodeServer) handleRequest(buf []byte) {
	// Unmarshal Generic Request
	request := &md.GenericRequest{}
	err := proto.Unmarshal(buf, request)
	if err != nil {
		md.LogFatal(err)
		return
	}
	// Switch case request type
	switch request.GetType() {
	case md.GenericRequest_INVITE:
		// Unmarshal Invite Request
		ir := &md.InviteRequest{}
		err = proto.Unmarshal(request.GetData(), ir)
		if err != nil {
			md.LogFatal(err)
			return
		}
		// Send Request to Channel
		s.inviteRequests <- ir
	}
}

// Handle Request and Send to Channel after unmarshal
func (s *NodeServer) handleResponse(buf []byte) {
	// Unmarshal Generic Response
	response := &md.GenericResponse{}
	err := proto.Unmarshal(buf, response)
	if err != nil {
		md.LogFatal(err)
		return
	}
	// Switch case response type
	switch response.GetType() {
	case md.GenericResponse_CONNECTION:
		// Unmarshal Connection Response
		cr := &md.ConnectionResponse{}
		err = proto.Unmarshal(response.GetData(), cr)
		if err != nil {
			md.LogFatal(err)
			return
		}

		// Send Response to Channel
		s.connectionResponses <- cr
	case md.GenericResponse_REPLY:
		// Unmarshal Reply Response
		rr := &md.InviteResponse{}
		err = proto.Unmarshal(response.GetData(), rr)
		if err != nil {
			md.LogFatal(err)
			return
		}
		// Send Response to Channel
		s.inviteResponses <- rr
	}
}

// # handleError Callback with handleError instance, and method
func (s *NodeServer) handleError(errMsg *md.SonrError) {
	// Check for Error
	if errMsg.HasError {
		// Send Callback
		s.errorMessages <- errMsg.Message()
	}
}
