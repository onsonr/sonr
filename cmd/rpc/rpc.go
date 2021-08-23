package main

import (
	"context"
	"fmt"
	"log"
	"net"

	tp "github.com/sonr-io/core/internal/topic"
	ac "github.com/sonr-io/core/pkg/account"
	sc "github.com/sonr-io/core/pkg/client"
	md "github.com/sonr-io/core/pkg/models"
	"github.com/sonr-io/core/pkg/util"
	"google.golang.org/grpc"
)

type NodeServer struct {
	md.NodeServiceServer
	ctx context.Context

	// Client
	account ac.Account
	client  sc.Client
	device  *md.Device
	state   md.Lifecycle

	// Groups
	local *tp.RoomManager
	Rooms map[string]*tp.RoomManager

	// Event Channels
	completeEvents  chan *md.CompleteEvent
	errorEvents     chan *md.ErrorEvent
	mailEvents      chan *md.MailEvent
	linkEvents      chan *md.LinkEvent
	progressEvents  chan *md.ProgressEvent
	statusEvents    chan *md.StatusEvent
	RoomEvents      chan *md.RoomEvent
	inviteRequests  chan *md.InviteRequest
	inviteResponses chan *md.InviteResponse

	// Callback Channels
	authResponses       chan *md.AuthResponse
	actionResponses     chan *md.ActionResponse
	connectionResponses chan *md.ConnectionResponse
	decisionResponses   chan *md.DecisionResponse
	linkResponses       chan *md.LinkResponse
	mailboxResponses    chan *md.MailboxResponse
	verifyResponses     chan *md.VerifyResponse
}

func main() {
	// Create a new gRPC server
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", util.RPC_SERVER_PORT))
	if err != nil {
		md.LogRPC("online", false)
		log.Fatal(err)
	}
	md.LogRPC("online", true)

	// Set GRPC Server
	chatServer := NodeServer{
		// Defaults
		ctx:   context.Background(),
		Rooms: make(map[string]*tp.RoomManager, 10),
		state: md.Lifecycle_ACTIVE,

		// Event Channels
		RoomEvents:      make(chan *md.RoomEvent, util.MAX_CHAN_DATA),
		mailEvents:      make(chan *md.MailEvent, util.MAX_CHAN_DATA),
		progressEvents:  make(chan *md.ProgressEvent, util.MAX_CHAN_DATA),
		completeEvents:  make(chan *md.CompleteEvent, util.MAX_CHAN_DATA),
		statusEvents:    make(chan *md.StatusEvent, util.MAX_CHAN_DATA),
		errorEvents:     make(chan *md.ErrorEvent, util.MAX_CHAN_DATA),
		inviteRequests:  make(chan *md.InviteRequest, util.MAX_CHAN_DATA),
		inviteResponses: make(chan *md.InviteResponse, util.MAX_CHAN_DATA),
		linkEvents:      make(chan *md.LinkEvent, util.MAX_CHAN_DATA),

		// Callback Channels
		authResponses:       make(chan *md.AuthResponse, util.MAX_CHAN_DATA),
		actionResponses:     make(chan *md.ActionResponse, util.MAX_CHAN_DATA),
		connectionResponses: make(chan *md.ConnectionResponse, util.MAX_CHAN_DATA),
		decisionResponses:   make(chan *md.DecisionResponse, util.MAX_CHAN_DATA),
		linkResponses:       make(chan *md.LinkResponse, util.MAX_CHAN_DATA),
		mailboxResponses:    make(chan *md.MailboxResponse, util.MAX_CHAN_DATA),
		verifyResponses:     make(chan *md.VerifyResponse, util.MAX_CHAN_DATA),
	}

	grpcServer := grpc.NewServer()

	// Register the gRPC service
	md.RegisterNodeServiceServer(grpcServer, &chatServer)
	if err := grpcServer.Serve(listener); err != nil {
		md.LogRPC("serve", false)
		log.Fatal(err)
	}
	md.LogRPC("serve", true)
}

// Initialize method is called when a new node is created
func (s *NodeServer) Initialize(ctx context.Context, req *md.InitializeRequest) (*md.NoResponse, error) {
	// Initialize Logger
	md.InitLogger(req)

	// Initialize Device
	device := req.GetDevice()

	// Create User
	u, serr := ac.OpenAccount(req, device)
	if serr != nil {
		s.handleError(serr)
		return &md.NoResponse{}, serr.Error
	}

	s.account = u
	s.device = device

	// Create Client
	s.client = sc.NewClient(s.ctx, s.device, s.callback())
	s.verifyResponses <- s.account.VerifyRead()
	// Return Blank Response
	return &md.NoResponse{}, nil
}

// Connect method starts this nodes host
func (s *NodeServer) Connect(ctx context.Context, req *md.ConnectionRequest) (*md.NoResponse, error) {
	// Update User with Connection Request
	s.device.SetConnection(req)

	// Connect Host
	peer, serr := s.client.Connect(req, s.account)
	if serr != nil {
		s.handleError(serr)
		s.setConnected(false)
	} else {
		// Update Status
		s.setConnected(true)
	}

	// Bootstrap Node
	s.local, serr = s.client.Bootstrap(req)
	if serr != nil {
		s.handleError(serr)
		s.setAvailable(false)
	} else {
		s.setAvailable(true)
	}

	// Join Account Network
	if err := s.account.JoinNetwork(s.client.GetHost(), req, peer); err != nil {
		s.handleError(err)
		s.setAvailable(false)
	}

	// Return Blank Response - Needs No Response Struc
	return &md.NoResponse{}, nil
}
