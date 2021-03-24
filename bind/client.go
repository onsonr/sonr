package bind

import (
	"log"

	"github.com/getsentry/sentry-go"
	dt "github.com/sonr-io/core/internal/data"
	md "github.com/sonr-io/core/internal/models"
	net "github.com/sonr-io/core/internal/network"
	sn "github.com/sonr-io/core/pkg/node"
	dq "github.com/sonr-io/core/pkg/user"
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
	fs              *dq.SonrFS
	profile         *md.Profile
	contact         *md.Contact
	device          *md.Device
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
		log.Fatalln(err)
	}

	// Create Mobile Node
	mn := &MobileNode{
		call:            call,
		hasStarted:      false,
		hasBootstrapped: false,
		profile:         req.GetProfile(),
		contact:         req.Contact,
		device:          req.Device,
	}

	// Create Host Options
	mn.fs = dq.InitFS(&req, mn.profile, mn.nodeCallback())
	mn.hostOpts, err = net.NewHostOpts(&req, mn.fs)
	if err != nil {
		log.Println(err)
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
	peerId := dq.GetPeerID(mn.device, mn.profile, mn.node.ID().String())
	result := mn.node.Init(mn.hostOpts, peerId)

	// Check Result
	if result {
		// Set Started
		mn.hasStarted = true

		// Bootstrap to Peers
		strapResult := mn.node.Bootstrap(mn.hostOpts, mn.fs)
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
