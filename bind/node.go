package bind

import (
	"context"

	sc "github.com/sonr-io/core/pkg/client"
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
	client    sc.Client
	state     md.Lifecycle
	user      *md.User

	// Groups
	local  *net.TopicManager
	topics map[string]*net.TopicManager
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

// Starts Host and Connects
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
	n.local, serr = n.client.Bootstrap(req)
	if serr != nil {
		n.handleError(serr)
		n.setAvailable(false)
	} else {
		n.setAvailable(true)
	}
}
