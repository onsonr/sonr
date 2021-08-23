package bind

import (
	"context"

	tp "github.com/sonr-io/core/internal/topic"
	ac "github.com/sonr-io/core/pkg/account"
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
	account ac.Account
	client  sc.Client
	state   md.Lifecycle

	// Rooms
	local   *tp.RoomManager
	devices *tp.RoomManager
	groups  map[string]*tp.RoomManager
}

// Initializes New Node ^ //
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

	// Create User
	u, serr := ac.OpenAccount(req, req.GetDevice())
	if serr != nil {
		md.LogError(serr.Error)
		return nil
	}
	// Initialize Node
	mn := &Node{
		call:    call,
		ctx:     context.Background(),
		groups:  make(map[string]*tp.RoomManager, 10),
		state:   md.Lifecycle_ACTIVE,
		account: u,
	}

	// Create Client
	mn.client = sc.NewClient(mn.ctx, mn.account, mn.callback())
	return mn
}

// Starts Host and Connects
func (n *Node) Connect(data []byte) {
	// Unmarshal Request
	req := &md.ConnectionRequest{}
	err := proto.Unmarshal(data, req)
	if err != nil {
		md.LogFatal(err)
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
	data, err := proto.Marshal(su)
	if err != nil {
		n.handleError(md.NewError(err, md.ErrorEvent_MARSHAL))
		return
	}
	n.call.OnStatus(data)
}

// Sets Node to be Available Status
func (n *Node) setAvailable(val bool) {
	// Update Status
	su := n.account.SetAvailable(val)

	// Callback Status
	data, err := proto.Marshal(su)
	if err != nil {
		n.handleError(md.NewError(err, md.ErrorEvent_MARSHAL))
		return
	}
	n.call.OnStatus(data)
}

// Sets Node to be (Provided) Status
func (n *Node) setStatus(newStatus md.Status) {
	// Set Status
	su := n.account.SetStatus(newStatus)

	// Callback Status
	data, err := proto.Marshal(su)
	if err != nil {
		n.handleError(md.NewError(err, md.ErrorEvent_MARSHAL))
		return
	}
	n.call.OnStatus(data)
}
