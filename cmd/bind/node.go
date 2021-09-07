package bind

import (
	"context"

	"github.com/sonr-io/core/tools/logger"
	"github.com/sonr-io/core/internal/room"
	ac "github.com/sonr-io/core/pkg/account"
	sc "github.com/sonr-io/core/pkg/client"
	"github.com/sonr-io/core/pkg/data"
	"go.uber.org/zap"

	"google.golang.org/protobuf/proto"
)

type Node struct {
	data.Callback
	// Properties
	call Callback
	ctx  context.Context

	// Client
	account ac.Account
	client  sc.Client
	state   data.Lifecycle

	// Rooms
	local   *room.RoomManager
	devices *room.RoomManager
	groups  map[string]*room.RoomManager
}

// Initializes New Node ^ //
func Initialize(reqBytes []byte, call Callback) *Node {
	// Unmarshal Request
	req := &data.InitializeRequest{}
	err := proto.Unmarshal(reqBytes, req)
	if err != nil {
		logger.Panic("Failed to unmarshal initialize request", zap.Error(err))
		return nil
	}

	// Initialize Logger
	logger.Init(req.Options.GetEnableLogging())

	// Create User
	u, serr := ac.OpenAccount(req, req.GetDevice())
	if serr != nil {
		logger.Panic("Failed to initialize user", zap.Error(serr.Error))
		return nil
	}
	// Initialize Node
	mn := &Node{
		call:    call,
		ctx:     context.Background(),
		groups:  make(map[string]*room.RoomManager, 10),
		state:   data.Lifecycle_ACTIVE,
		account: u,
	}

	// Create Client
	mn.client = sc.NewClient(mn.ctx, mn.account, mn.callback())
	return mn
}

// Starts Host and Connects
func (n *Node) Connect(buf []byte) {
	// Unmarshal Request
	req := &data.ConnectionRequest{}
	err := proto.Unmarshal(buf, req)
	if err != nil {
		logger.Panic("Failed to initialize user", zap.Error(err))
	}

	// Connect Host
	peer, serr := n.client.Connect(req)
	if serr != nil {
		n.handleError(serr)
		n.setConnected(false)
	} else {
		// Update Status
		n.setConnected(true)
	}

	// Bootstrap Node
	n.local, serr = n.client.Bootstrap(req)
	if serr != nil {
		n.handleError(serr)
		n.setAvailable(false)
	} else {
		n.setAvailable(true)
	}

	// Join Account Network
	if err := n.account.JoinNetwork(n.client.GetHost(), req, peer); err != nil {
		n.handleError(err)
		n.setAvailable(false)
	}
}

// ** ─── Node Status Checks ────────────────────────────────────────────────────────
// Sets Node to be Connected Status
func (n *Node) setConnected(val bool) {
	// Update Status
	su := n.account.SetConnected(val)

	// Callback Status
	res, err := proto.Marshal(su)
	if err != nil {
		n.handleError(data.NewError(err, data.ErrorEvent_MARSHAL))
		return
	}
	n.call.OnStatus(res)
}

// Sets Node to be Available Status
func (n *Node) setAvailable(val bool) {
	// Update Status
	su := n.account.SetAvailable(val)

	// Callback Status
	res, err := proto.Marshal(su)
	if err != nil {
		n.handleError(data.NewError(err, data.ErrorEvent_MARSHAL))
		return
	}
	n.call.OnStatus(res)
}
