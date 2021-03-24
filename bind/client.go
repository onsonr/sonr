package bind

import (
	"log"

	"github.com/getsentry/sentry-go"
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
	// Initialize Node Logging
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://cbf88b01a5a5468fa77101f7dfc54f20@o549479.ingest.sentry.io/5672329",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	// Unmarshal Request
	req := md.ConnectionRequest{}
	err = proto.Unmarshal(reqBytes, &req)
	if err != nil {
		sentry.CaptureException(err)
	}

	// Create Mobile Node
	mn := &MobileNode{
		call:            call,
		hasStarted:      false,
		hasBootstrapped: false,
	}

	// Create New User
	mn.user, err = u.NewUser(&req, mn.nodeCallback())
	if err != nil {
		sentry.CaptureException(err)
	}

	// Create Host Options
	mn.hostOpts, err = net.NewHostOpts(&req, mn.user.FS, mn.user.PrivateKey())
	if err != nil {
		sentry.CaptureException(err)
		return nil
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
	// Connect to Host
	if ok := mn.node.Connect(mn.hostOpts); ok {
		// Set Started
		mn.hasStarted = true
		mn.user.SetPeer(mn.node.ID().String())

		// Bootstrap to Peers
		if ok := mn.node.Bootstrap(mn.hostOpts, mn.user.FS, mn.user.Peer, mn.user.PeerBuf); ok {
			mn.hasBootstrapped = true
		} else {
			log.Println("Failed to bootstrap node")
			return
		}
	} else {
		log.Println("Failed to start host")
		return
	}
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
