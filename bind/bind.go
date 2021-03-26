package bind

import (
	"log"

	md "github.com/sonr-io/core/internal/models"
	u "github.com/sonr-io/core/internal/user"
	dt "github.com/sonr-io/core/pkg/data"
	sn "github.com/sonr-io/core/pkg/node"
	tpc "github.com/sonr-io/core/pkg/topic"
	"google.golang.org/protobuf/proto"
)

// * Struct: Reference for Binded Proxy Node * //
type MobileNode struct {
	call   Callback
	config mobileConfig
	node   *sn.Node
	user   *u.User
	local  *tpc.TopicManager
	topics []*tpc.TopicManager
}

// @ Create New Mobile Node
func NewNode(reqBytes []byte, call Callback) *MobileNode {
	// Unmarshal Request
	req := &md.ConnectionRequest{}
	err := proto.Unmarshal(reqBytes, req)
	if err != nil {
		panic(err)
	}

	// Create Mobile Node
	mn := &MobileNode{
		call:   call,
		config: newMobileConfig(),
		topics: make([]*tpc.TopicManager, 10),
	}

	// Create New User
	mn.user, err = u.NewUser(req, mn.callbackNode())
	if err != nil {
		panic(err)
	}

	// Create Node
	mn.node = sn.NewNode(mn.contextNode(), req, mn.callbackNode())
	return mn
}

// **-----------------** //
// ** Network Actions ** //
// **-----------------** //
// @ Start Host and Connect
func (mn *MobileNode) Connect() {
	// Get Private Key and Connect Host
	key, err := mn.user.PrivateKey()
	if err != nil {
		log.Println("Failed to retreive private key")
		mn.setConnected(false)
		return
	}

	// Connect Host
	err = mn.node.Start(key)
	if err != nil {
		log.Println("Failed to start host")
		mn.setConnected(false)
		return
	} else {
		// Update Status
		mn.setConnected(true)
	}

	// Set User Peer
	err = mn.user.SetPeer(mn.node.Host.ID())
	if err != nil {
		log.Println("Failed to set peer")
		return
	}

	// Bootstrap Node
	err = mn.node.Bootstrap()
	if err != nil {
		log.Println("Failed to bootstrap node")
		mn.setBootstrapped(false)
		return
	} else {
		mn.setBootstrapped(true)
	}

	mn.local, err = mn.node.JoinLocal()
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
func (mn *MobileNode) Pause() {
	dt.GetState().Pause()
}

// @ Close Ends All Network Communication
func (mn *MobileNode) Resume() {
	dt.GetState().Resume()
}

// @ Close Ends All Network Communication
func (mn *MobileNode) Stop() {
	mn.node.Close()
}
