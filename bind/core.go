package sonr

import (
	"context"
	"log"

	olc "github.com/google/open-location-code/go"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	sf "github.com/sonr-io/core/internal/file"
	sh "github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/internal/lobby"
	md "github.com/sonr-io/core/internal/models"
	tr "github.com/sonr-io/core/internal/transfer"
	"google.golang.org/protobuf/proto"
)

// @ Maximum Files in Node Cache
const maxFileBufferSize = 64

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
	ctx     context.Context
	olc     string
	device  *md.Device
	peer    *md.Peer
	contact *md.Contact
	status  md.Status

	// Networking Properties
	host   sh.SonrHost
	pubSub *pubsub.PubSub

	// Data Properties
	files       []*sf.ProcessedFile
	directories *md.Directories

	// References
	call     Callback
	lobby    *lobby.Lobby
	peerConn *tr.PeerConnection
}

// ^ NewNode Initializes Node with a host and default properties ^
func NewNode(reqBytes []byte, call Callback) *Node {
	// ** Create Context and Node - Begin Setup **
	node := new(Node)
	node.ctx = context.Background()
	node.call, node.files = call, make([]*sf.ProcessedFile, maxFileBufferSize)
	node.status = md.Status_NONE

	// ** Unmarshal Request **
	reqMsg := md.ConnectionRequest{}
	err := proto.Unmarshal(reqBytes, &reqMsg)
	if err != nil {
		node.error(err, "NewNode")
		return nil
	}

	// @1. Set OLC, Create Host, and Start Discovery
	node.olc = olc.Encode(float64(reqMsg.Latitude), float64(reqMsg.Longitude), 8)
	node.host, err = sh.NewHost(node.ctx, reqMsg.Directories, node.olc, reqMsg.Connectivity)
	if err != nil {
		node.error(err, "NewNode")
		return nil
	}

	// @3. Set Node User Information
	if err = node.setInfo(&reqMsg); err != nil {
		node.error(err, "NewNode")
		return nil
	}

	// @4. Setup Connection w/ Lobby and Set Stream Handlers
	if err = node.setConnection(node.ctx); err != nil {
		node.error(err, "NewNode")
		return nil
	}

	// ** Callback Node User Information ** //
	return node
}

// ^ queued Callback, Sends File Invite to Peer, and Notifies Client ^
func (sn *Node) queued(card *md.TransferCard, req *md.InviteRequest) {
	// Get PeerID
	id, _, err := sn.lobby.Find(req.To.Id)

	// Check error
	if err != nil {
		sn.error(err, "Queued")
	}

	// Retreive Current File
	currFile := sn.currentFile()
	card.Status = md.TransferCard_INVITE
	sn.peerConn.ProcessedFile = currFile

	// Create Invite Message
	invMsg := md.AuthInvite{
		From:    sn.peer,
		Payload: card.Payload,
		Card:    card,
	}

	// Check if ID in PeerStore
	go func(inv *md.AuthInvite) {
		// Convert Protobuf to bytes
		msgBytes, err := proto.Marshal(inv)
		if err != nil {
			sn.error(err, "Marshal")
		}

		sn.peerConn.Request(sn.host.Host, id, msgBytes)
	}(&invMsg)
}

// ^ multiQueued Callback, Sends File Invite to Peer, and Notifies Client ^
func (sn *Node) multiQueued(card *md.TransferCard, req *md.InviteRequest) {
	// Get PeerID
	id, _, err := sn.lobby.Find(req.To.Id)

	// Check error
	if err != nil {
		sn.error(err, "Queued")
	}

	// Retreive Current File
	currFile := sn.currentFile()
	card.Status = md.TransferCard_INVITE
	sn.peerConn.ProcessedFile = currFile

	// Create Invite Message
	invMsg := md.AuthInvite{
		From:    sn.peer,
		Payload: card.Payload,
		Card:    card,
	}

	// Check if ID in PeerStore
	go func(inv *md.AuthInvite) {
		// Convert Protobuf to bytes
		msgBytes, err := proto.Marshal(inv)
		if err != nil {
			sn.error(err, "Marshal")
		}

		sn.peerConn.Request(sn.host.Host, id, msgBytes)
	}(&invMsg)
}

// ^ invite Callback with data for Lifecycle ^ //
func (sn *Node) invited(invite *md.AuthInvite) {
	// Update Status
	sn.status = md.Status_INVITED

	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(invite)
	if err != nil {
		log.Println(err)
	}

	// Callback with Data
	sn.call.OnInvited(msgBytes)
}

// ^ transmitted Callback middleware post transfer ^ //
func (sn *Node) transmitted(peer *md.Peer) {
	// Update Status
	sn.status = md.Status_AVAILABLE

	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(peer)
	if err != nil {
		log.Println(err)
	}

	// Callback with Data
	sn.call.OnTransmitted(msgBytes)
}

// ^ received Callback middleware post transfer ^ //
func (sn *Node) received(card *md.TransferCard) {
	// Update Status
	sn.status = md.Status_AVAILABLE

	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(card)
	if err != nil {
		log.Println(err)
	}

	// Callback with Data
	sn.call.OnReceived(msgBytes)
}

// ^ error Callback with error instance, and method ^
func (sn *Node) error(err error, method string) {
	// Create Error ProtoBuf
	errorMsg := md.ErrorMessage{
		Message: err.Error(),
		Method:  method,
	}

	// Convert Message to bytes
	bytes, err := proto.Marshal(&errorMsg)
	if err != nil {
		log.Println("Cannot Marshal Error Protobuf: ", err)
	}
	// Send Callback
	sn.call.OnError(bytes)

	// Log In Core
	log.Fatalf("[Error] At Method %s : %s", err.Error(), method)
}
