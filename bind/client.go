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
	mn.user = u.NewUser(&req, mn.nodeCallback())
	key, err := mn.user.GetPrivateKey()
	if err != nil {
		sentry.CaptureException(err)
	}

	// Create Host Options
	mn.hostOpts, err = net.NewHostOpts(&req, mn.user.FS, key)
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
	// Start Node
	mn.user.SetPeer(mn.node.ID().String())
	result := mn.node.Init(mn.hostOpts)

	// Check Result
	if result {
		// Set Started
		mn.hasStarted = true

		// Bootstrap to Peers
		strapResult := mn.node.Bootstrap(mn.hostOpts, mn.user.FS, mn.user.GetPeer, mn.user.GetPeerBuf)
		if strapResult {
			mn.hasBootstrapped = true
		} else {
			log.Println("Failed to bootstrap node")
		}
	} else {
		log.Println("Failed to start host")
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
