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
	device  *md.Device
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

	// Initialize Device
	device := req.GetDevice()

	// Initialize Node
	mn := &Node{
		call:   call,
		ctx:    context.Background(),
		groups: make(map[string]*tp.RoomManager, 10),
		state:  md.Lifecycle_ACTIVE,
		device: device,
	}

	// Create User
	u, serr := ac.OpenAccount(req, device)
	if serr != nil {
		mn.handleError(serr)
		return nil
	}

	mn.account = u
	mn.device = device

	// Create Client
	mn.client = sc.NewClient(mn.ctx, mn.device, mn.callback())
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

	// Update User with Connection Request
	n.account.SetConnection(req)
	n.device.SetConnection(req)

	// Connect Host
	peer, isPrimary, serr := n.client.Connect(req, n.account)
	if serr != nil {
		n.handleError(serr)
		n.setConnected(false)
	} else {
		// Update Status
		n.setConnected(true)
		n.account.HandleSetPeer(peer, isPrimary)
	}

	// Bootstrap Node
	n.local, serr = n.client.Bootstrap(req)
	if serr != nil {
		n.handleError(serr)
		n.setAvailable(false)
	} else {
		n.setAvailable(true)
	}
}
