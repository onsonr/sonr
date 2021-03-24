package bind

import (
	"log"

	md "github.com/sonr-io/core/internal/models"
	net "github.com/sonr-io/core/internal/network"
	u "github.com/sonr-io/core/internal/user"
	dt "github.com/sonr-io/core/pkg/data"
	sn "github.com/sonr-io/core/pkg/node"
	"google.golang.org/protobuf/proto"
)

// * Struct: Reference for Binded Proxy Node * //
type MobileNode struct {
	call            Callback
	node            *sn.Node
	hasStarted      bool
	hasBootstrapped bool
	hostOpts        *net.HostOptions
	status          md.Status
	user            *u.User
}

// @ Create New Mobile Node
func NewNode(reqBytes []byte, call Callback) *MobileNode {
	// // Initialize Node Logging
	// err := sentry.Init(sentry.ClientOptions{
	// 	Dsn: "https://cbf88b01a5a5468fa77101f7dfc54f20@o549479.ingest.sentry.io/5672329",
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// Unmarshal Request
	req := &md.ConnectionRequest{}
	err := proto.Unmarshal(reqBytes, req)
	if err != nil {
		// sentry.CaptureException(err)
		panic(err)
	}

	// Create Mobile Node
	mn := &MobileNode{
		call:            call,
		hasStarted:      false,
		hasBootstrapped: false,
	}

	// Create New User
	mn.user, err = u.NewUser(req, mn.nodeCallback())
	if err != nil {
		// sentry.CaptureException(err)
		panic(err)
	}

	// Create Host Options
	mn.hostOpts, err = net.NewHostOpts(req, mn.user.PrivateKey())
	if err != nil {
		// sentry.CaptureException(err)
		panic(err)
	}

	// Create Node
	mn.node = sn.NewNode(mn.hostOpts, mn.nodeCallback())
	return mn
}

// **-----------------** //
// ** Network Actions ** //
// **-----------------** //
// @ Start Host
func (mn *MobileNode) Connect() {
	// ! Connect to Host
	err := mn.node.Connect(mn.hostOpts)
	if err != nil {
		log.Println("Failed to start host")
		// sentry.CaptureException(err)
		mn.call.OnConnected(false)
		return
	}

	// Set Started
	mn.hasStarted = true
	err = mn.user.SetPeer(mn.node.ID())
	if err != nil {
		log.Println(err)
		return
	}
	mn.call.OnConnected(true)

	// ! Bootstrap to Peers
	err = mn.node.Bootstrap(mn.hostOpts, mn.user.FS)
	if err != nil {
		log.Println("Failed to bootstrap node")
		// sentry.CaptureException(err)
		mn.call.OnReady(false)
		return
	}

	// Update Status
	mn.hasBootstrapped = true
	mn.call.OnReady(true)
}

// @ Return URL Metadata, Helper Method
func GetURLMetadata(url string) []byte {
	// Get Link Data
	data, err := dt.GetPageInfoFromUrl(url)
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
// @ Checks for is Ready
func (mn *MobileNode) isReady() bool {
	return mn.hasBootstrapped && mn.hasStarted
}

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
