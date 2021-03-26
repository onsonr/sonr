package bind

import (
	"log"

	md "github.com/sonr-io/core/internal/models"
	net "github.com/sonr-io/core/internal/network"
	u "github.com/sonr-io/core/internal/user"
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

	// Create Host Options
	hostOpts, err := net.NewHostOpts(req)
	if err != nil {
		panic(err)
	}

	// Create Node
	mn.node = sn.NewNode(mn.config.contextNode(), hostOpts, mn.callbackNode())
	return mn
}

// **-----------------** //
// ** Network Actions ** //
// **-----------------** //
// @ Start Host
func (mn *MobileNode) Connect() {
	// ! Connect to Host
	// Get Key
	key, err := mn.user.PrivateKey()
	if err != nil {
		mn.config.setConnected(false)
		mn.call.OnConnected(false)
	}

	// Start Host
	err = mn.node.Connect(key)
	if err != nil {
		mn.config.setConnected(false)
		mn.call.OnConnected(false)
	}

	// Update Status
	mn.config.setConnected(true)
	mn.call.OnConnected(true)

	// ! Set User Peer
	err = mn.user.SetPeer(mn.node.ID())
	if err != nil {
		log.Println(err)
		mn.call.OnReady(false)
	}

	// ! Bootstrap to Peers
	err = mn.node.Bootstrap()
	if err != nil {
		log.Println("Failed to bootstrap node")
		mn.config.setBootstrapped(false)
		mn.call.OnReady(false)
	}

	// Update Status
	mn.config.setBootstrapped(true)

	// ! Join Local topic
	mn.local, err = mn.node.JoinLocal()
	if err != nil {
		log.Println("Failed to connect to local topic")
		mn.config.setJoinedLocal(false)
		mn.call.OnReady(false)
	}

	mn.config.setJoinedLocal(true)
	mn.call.OnReady(true)
}

// @ Return URL Metadata, Helper Method
func GetURLMetadata(url string) []byte {
	// Get Link Data
	data, err := md.GetPageInfoFromUrl(url)
	if err != nil {
		log.Println(err, " Failed to Parse URL")
	}

	// Marshal
	bytes, err := proto.Marshal(data)
	if err != nil {
		log.Println(err, " Failed to Parse URL")
	}
	return bytes
}

// **-------------------** //
// ** LifeCycle Actions ** //
// **-------------------** //
// @ Close Ends All Network Communication
func (mn *MobileNode) Pause() {
	mn.node.Pause()
}

// @ Close Ends All Network Communication
func (mn *MobileNode) Resume() {
	mn.node.Resume()
}

// @ Close Ends All Network Communication
func (mn *MobileNode) Stop() {
	// Check if Response Is Invited
	// mn.user.FS.Close()
	mn.node.Close()
}
