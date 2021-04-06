package bind

import (
	"log"

	"github.com/sonr-io/core/internal/network"
	tpc "github.com/sonr-io/core/internal/topic"
	sc "github.com/sonr-io/core/pkg/client"
	md "github.com/sonr-io/core/pkg/models"
	u "github.com/sonr-io/core/pkg/user"
	"google.golang.org/protobuf/proto"
)

// * Struct: Reference for Binded Proxy Node * //
type Node struct {
	// Properties
	call    Callback
	config  nodeConfig
	connreq *md.ConnectionRequest

	// Client
	client *sc.Client
	user   *u.User

	// Groups
	local  *tpc.TopicManager
	topics map[string]*tpc.TopicManager
}

// @ Create New Mobile Node
func NewMobileNode(reqBytes []byte, call Callback) *Node {
	// Unmarshal Request
	req := &md.ConnectionRequest{}
	err := proto.Unmarshal(reqBytes, req)
	if err != nil {
		panic(err)
	}

	// Create Mobile Node
	mn := &Node{
		call:    call,
		config:  newNodeConfig(),
		connreq: req,
		topics:  make(map[string]*tpc.TopicManager, 10),
	}

	// Create New User
	mn.user, err = u.NewUser(req, mn.callbackNode())
	if err != nil {
		panic(err)
	}

	// Create Client
	mn.client = sc.NewClient(mn.contextNode(), req, mn.callbackNode())
	return mn
}

// @ Create New Desktop Node
func NewDesktopNode(req *md.ConnectionRequest, call Callback) *Node {
	// Get Location by IP
	geoIP := md.GeoIP{}
	err := network.Location(&geoIP)
	if err != nil {
		return nil
	}

	// Modify Request
	req.AttachGeoToRequest(&geoIP)

	// Create Mobile Node
	mn := &Node{
		call:    call,
		config:  newNodeConfig(),
		connreq: req,
		topics:  make(map[string]*tpc.TopicManager, 10),
	}

	// Create New User
	mn.user, err = u.NewUser(req, mn.callbackNode())
	if err != nil {
		panic(err)
	}

	// Create Client
	mn.client = sc.NewClient(mn.contextNode(), req, mn.callbackNode())
	return mn
}

// **-----------------** //
// ** Network Actions ** //
// **-----------------** //
// @ Start Host and Connect
func (mn *Node) Connect() {
	// Connect Host
	err := mn.client.Connect(mn.user.PrivateKey())
	if err != nil {
		log.Println("Failed to start host")
		mn.setConnected(false)
		return
	} else {
		// Update Status
		mn.setConnected(true)
	}

	// Bootstrap Node
	err = mn.client.Bootstrap()
	if err != nil {
		log.Println("Failed to bootstrap node")
		mn.setBootstrapped(false)
		return
	} else {
		mn.setBootstrapped(true)
	}

	mn.local, err = mn.client.JoinLocal()
	if err != nil {
		log.Println("Failed to join local pubsub")
		mn.setJoinedLocal(false)
		return
	} else {
		mn.setJoinedLocal(true)
	}
}

// **-------------------** //
// ** LifeCycle Actions ** //
// **-------------------** //
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
	mn.config.Ctx.Done()
}
