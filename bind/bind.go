package bind

import (
	"log"

	md "github.com/sonr-io/core/internal/models"
	net "github.com/sonr-io/core/internal/network"
	sn "github.com/sonr-io/core/pkg/node"
	"google.golang.org/protobuf/proto"
)

// * Struct: Reference for Binded Proxy Node * //
type MobileNode struct {
	call            Callback
	ID              string
	DeviceID        string
	UserID          uint32
	node            *sn.Node
	hasStarted      bool
	hasBootstrapped bool
	hostOpts        *net.HostOptions
	status          md.Status
}

// @ Create New Mobile Node
func NewNode(reqBytes []byte, call Callback) *MobileNode {
	// Unmarshal Request
	req := md.ConnectionRequest{}
	err := proto.Unmarshal(reqBytes, &req)
	if err != nil {
		log.Fatalln(err)
	}

	hostOpts, err := net.NewHostOpts(&req)
	if err != nil {
		log.Println(err)
	}

	// Create New Sonr Client
	node := sn.NewNode(&req, call)

	// Return Mobile Node
	return &MobileNode{
		call:            call,
		node:            node,
		hasStarted:      false,
		hasBootstrapped: false,
		hostOpts:        hostOpts,
	}
}

// **-----------------** //
// ** Network Actions ** //
// **-----------------** //
// @ Start Host
func (mn *MobileNode) Connect() {
	// Start Node
	result := mn.node.Start(mn.hostOpts)
	if result {
		// Set Peer Info
		peer := mn.node.Peer()
		mn.ID = peer.Id.Peer
		mn.DeviceID = peer.Id.Device
		mn.UserID = peer.Id.User

		// Set Started
		mn.hasStarted = true

		// Bootstrap to Peers
		strapResult := mn.node.Bootstrap(mn.hostOpts)
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
