package main

import (
	"context"
	"fmt"
	"net"

	"github.com/sonr-io/core/tools/logger"
	"github.com/sonr-io/core/internal/room"
	"github.com/sonr-io/core/pkg/account"
	"github.com/sonr-io/core/pkg/client"
	"github.com/sonr-io/core/pkg/data"
	"github.com/sonr-io/core/pkg/util"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type NodeServer struct {
	data.NodeServiceServer
	ctx context.Context

	// Client
	account account.Account
	client  client.Client
	state   data.Lifecycle

	// Groups
	local *room.RoomManager
	Rooms map[string]*room.RoomManager

	// Event Channels
	completeEvents  chan *data.CompleteEvent
	errorEvents     chan *data.ErrorEvent
	mailEvents      chan *data.MailEvent
	linkEvents      chan *data.LinkEvent
	progressEvents  chan *data.ProgressEvent
	statusEvents    chan *data.StatusEvent
	RoomEvents      chan *data.RoomEvent
	inviteRequests  chan *data.InviteRequest
	inviteResponses chan *data.InviteResponse

	// Callback Channels
	authResponses       chan *data.AuthResponse
	actionResponses     chan *data.ActionResponse
	connectionResponses chan *data.ConnectionResponse
	decisionResponses   chan *data.DecisionResponse
	linkResponses       chan *data.LinkResponse
	mailboxResponses    chan *data.MailboxResponse
	verifyResponses     chan *data.VerifyResponse
}

func main() {
	// Create a new gRPC server
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", util.RPC_SERVER_PORT))
	if err != nil {
		logger.Panic("Failed to bind to port", zap.Error(err))
	}

	// Set GRPC Server
	chatServer := NodeServer{
		// Defaults
		ctx:   context.Background(),
		Rooms: make(map[string]*room.RoomManager, 10),
		state: data.Lifecycle_ACTIVE,

		// Event Channels
		RoomEvents:      make(chan *data.RoomEvent, util.MAX_CHAN_DATA),
		mailEvents:      make(chan *data.MailEvent, util.MAX_CHAN_DATA),
		progressEvents:  make(chan *data.ProgressEvent, util.MAX_CHAN_DATA),
		completeEvents:  make(chan *data.CompleteEvent, util.MAX_CHAN_DATA),
		statusEvents:    make(chan *data.StatusEvent, util.MAX_CHAN_DATA),
		errorEvents:     make(chan *data.ErrorEvent, util.MAX_CHAN_DATA),
		inviteRequests:  make(chan *data.InviteRequest, util.MAX_CHAN_DATA),
		inviteResponses: make(chan *data.InviteResponse, util.MAX_CHAN_DATA),
		linkEvents:      make(chan *data.LinkEvent, util.MAX_CHAN_DATA),

		// Callback Channels
		authResponses:       make(chan *data.AuthResponse, util.MAX_CHAN_DATA),
		actionResponses:     make(chan *data.ActionResponse, util.MAX_CHAN_DATA),
		connectionResponses: make(chan *data.ConnectionResponse, util.MAX_CHAN_DATA),
		decisionResponses:   make(chan *data.DecisionResponse, util.MAX_CHAN_DATA),
		linkResponses:       make(chan *data.LinkResponse, util.MAX_CHAN_DATA),
		mailboxResponses:    make(chan *data.MailboxResponse, util.MAX_CHAN_DATA),
		verifyResponses:     make(chan *data.VerifyResponse, util.MAX_CHAN_DATA),
	}

	grpcServer := grpc.NewServer()

	// Register the gRPC service
	data.RegisterNodeServiceServer(grpcServer, &chatServer)
	if err := grpcServer.Serve(listener); err != nil {
		logger.Panic("Failed to Register node service server", zap.Error(err))
	}
}

// Initialize method is called when a new node is created
func (s *NodeServer) Initialize(ctx context.Context, req *data.InitializeRequest) (*data.NoResponse, error) {
	// Initialize Logger
	var serr *data.SonrError
	logger.Init(req.Options.GetEnableLogging())

	// Create User
	s.account, serr = account.OpenAccount(req, req.GetDevice())
	if serr != nil {
		s.handleError(serr)
		return &data.NoResponse{}, serr.Error
	}

	// Create Client
	s.client = client.NewClient(s.ctx, s.account, s.callback())

	// Return Blank Response
	return &data.NoResponse{}, nil
}

// Connect method starts this nodes host
func (s *NodeServer) Connect(ctx context.Context, req *data.ConnectionRequest) (*data.NoResponse, error) {
	// Connect Host
	peer, serr := s.client.Connect(req)
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
	return &data.NoResponse{}, nil
}

// ** ─── Node Status Checks ────────────────────────────────────────────────────────
// Sets Node to be Connected Status
func (s *NodeServer) setConnected(val bool) {
	// Update Status
	su := s.account.SetConnected(val)

	// Callback Status
	s.statusEvents <- su
}

// Sets Node to be Available Status
func (s *NodeServer) setAvailable(val bool) {
	// Update Status
	su := s.account.SetAvailable(val)

	// Callback Status
	s.statusEvents <- su
}
