package bind

import (
	"log"

	"github.com/libp2p/go-libp2p-core/crypto"
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
// @ Start Host
func (mn *MobileNode) Connect() {
	// Initialize
	startChan := make(chan bool)
	bootstrapChan := make(chan bool)
	localChan := make(chan bool)

	// Get Private Key and Connect Host
	key, err := mn.user.PrivateKey()
	if err != nil {
		mn.setConnected(false)
	} else {
		// Connect Host
		go mn.start(key, startChan)

		// Await Routine Responses
		for i := 0; i < 3; i++ {
			select {
			// @ On Connection
			case status := <-startChan:
				// Update Status
				mn.setConnected(status)

				// Set User Peer
				if status {
					err = mn.user.SetPeer(mn.node.Host.ID())
					if err != nil {
						log.Println(err)
						break
					} else {
						// Begin Bootstrap
						go mn.bootstrap(bootstrapChan)
					}
				}

				// @ On Bootstrap
			case status := <-bootstrapChan:
				// Update Status and Join Local
				mn.setBootstrapped(status)
				if status {
					go mn.joinLocal(localChan)
				} else {
					break
				}

				// @ On Local Join
			case status := <-localChan:
				// Update Status
				mn.setJoinedLocal(status)
				break
			}
		}
	}

	// Close Channels
	close(startChan)
	close(bootstrapChan)
	close(localChan)
}

// @ Start Nodes Host
func (mn *MobileNode) start(key crypto.PrivKey, done chan bool) {
	err := mn.node.Start(key)
	if err != nil {
		log.Println("Failed to start host")
		done <- false
	} else {
		done <- true
	}
}

// @ Bootstrap Host to DHT
func (mn *MobileNode) bootstrap(done chan bool) {
	err := mn.node.Bootstrap()
	if err != nil {
		log.Println("Failed to bootstrap node")
		done <- false
	} else {
		done <- true
	}
}

// @ Join Local Pubsub
func (mn *MobileNode) joinLocal(done chan bool) {
	var err error
	mn.local, err = mn.node.JoinLocal()
	if err != nil {
		log.Println("Failed to join local pubsub")
		done <- false
	} else {
		done <- true
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
