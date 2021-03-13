package node

import (
	"context"
	"log"

	olc "github.com/google/open-location-code/go"
	"github.com/libp2p/go-libp2p-core/host"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	sh "github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/internal/lobby"
	tr "github.com/sonr-io/core/internal/transfer"
	dq "github.com/sonr-io/core/pkg/data"
	md "github.com/sonr-io/core/pkg/models"

	sentry "github.com/getsentry/sentry-go"
)

// ^ Interface: Callback is implemented from Plugin to receive updates ^
type Callback interface {
	OnRefreshed(data []byte)   // Lobby Updates
	OnEvent(data []byte)       // Lobby Event
	OnInvited(data []byte)     // User Invited
	OnDirected(data []byte)    // User Direct-Invite from another Device
	OnResponded(data []byte)   // Peer has responded
	OnProgress(data float32)   // File Progress Updated
	OnReceived(data []byte)    // User Received File
	OnTransmitted(data []byte) // User Sent File
	OnError(data []byte)       // Internal Error
}

// ^ Struct: Main Node handles Networking/Identity/Streams ^
type Node struct {
	// Properties
	ctx         context.Context
	olc         string
	fs          *dq.SonrFS
	directories *md.Directories
	device      *md.Device
	peer        *md.Peer
	contact     *md.Contact
	Status      md.Status

	// Networking Properties
	host   host.Host
	pubSub *pubsub.PubSub

	// References
	call     Callback
	lobby    *lobby.Lobby
	peerConn *tr.TransferController
	queue    *dq.FileQueue
}

// ^ NewNode Initializes Node with a host and default properties ^
func NewNode(req *md.ConnectionRequest, call Callback) *Node {
	// ** Create Context and Node - Begin Setup **
	node := new(Node)
	node.ctx = context.Background()
	node.call = call

	// ** Initialize Node Logging ** //
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://cbf88b01a5a5468fa77101f7dfc54f20@o549479.ingest.sentry.io/5672329",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	// Create New Profile from Request
	profile := &md.Profile{
		Username:  req.GetUsername(),
		FirstName: req.Contact.GetFirstName(),
		LastName:  req.Contact.GetLastName(),
		Picture:   req.Contact.GetPicture(),
		Platform:  req.Device.GetPlatform(),
	}

	// Set File System
	node.fs = dq.InitFS(req, profile)

	// @1. Set OLC, Create Host, and Start Discovery
	node.queue = dq.InitQueue(req.Directories, profile, node.queued, node.multiQueued, node.error)
	node.Status = md.Status_NONE
	node.olc = olc.Encode(float64(req.Latitude), float64(req.Longitude), 8)
	node.host, err = sh.NewHost(node.ctx, req.Directories, node.olc)
	if err != nil {
		node.error(err, "NewNode")
		return nil
	}

	// @3. Set Node User Information
	if err = node.setInfo(req, profile); err != nil {
		node.error(err, "NewNode")
		return nil
	}

	// @4. Setup Connection w/ Lobby and Set Stream Handlers
	if err = node.setConnection(node.ctx); err != nil {
		node.error(err, "NewNode")
		return nil
	}

	// ** Callback Node User Information ** //
	sentry.CaptureMessage("Node Started Successfully.")
	return node
}
