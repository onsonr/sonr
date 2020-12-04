package transfer

import (
	"context"
	"errors"
	"log"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	msgio "github.com/libp2p/go-msgio"
	sf "github.com/sonr-io/core/internal/file"
	md "github.com/sonr-io/core/internal/models"
)

// ChatRoomBufSize is the number of incoming messages to buffer for each topic.
const ChatRoomBufSize = 128

// Define Callback Function Types
type OnProtobuf func([]byte)
type OnError func(err error, method string)

// Package Error Callback
var onError OnError

// ^ Struct: Holds/Handles GRPC Calls and Handles Data Stream  ^ //
type PeerConnection struct {
	// Handlers
	auth     *Authorization
	safeFile *sf.SafeFile
	transfer *Transfer

	// Connection
	host   host.Host
	pubSub *pubsub.PubSub

	// Callbacks
	callback      OnProtobuf
	invitedCall   OnProtobuf
	respondedCall OnProtobuf
	progressCall  OnProgress
	completedCall OnProtobuf

	// Info
	olc         string
	dirs        *md.Directories
	currMessage *md.AuthMessage
	peerID      peer.ID
}

// ^ Initialize sets up new Peer Connection handler ^
func Initialize(h host.Host, ps *pubsub.PubSub, d *md.Directories, o string, ic OnProtobuf, rc OnProtobuf, pc OnProgress, cc OnProtobuf, ec OnError) (*PeerConnection, error) {
	// Set Package Level Callbacks
	onError = ec

	// Initialize Parameters into PeerConnection
	peerConn := &PeerConnection{
		host:          h,
		pubSub:        ps,
		olc:           o,
		dirs:          d,
		invitedCall:   ic,
		respondedCall: rc,
		progressCall:  pc,
		completedCall: cc,
	}

	// Set Handlers
	h.SetStreamHandler(protocol.ID("/sonr/transfer/data"), peerConn.HandleTransfer)

	// Create Auth Handler
	peerConn.auth = NewAuthRPC(peerConn)
	return peerConn, nil
}

// ^ Search for Peer in PubSub ^ //
func (pc *PeerConnection) Find(q string) peer.ID {
	// Iterate through PubSub in topic
	for _, id := range pc.pubSub.ListPeers(pc.olc) {
		// If Found Match
		if id.String() == q {
			return id
		}
	}
	return ""
}

// ^ Send Invite to a Peer ^ //
func (pc *PeerConnection) Invite(id string, info *md.Peer, sm *sf.SafeFile) {
	// Set SafeFile
	pc.safeFile = sm

	// Find Peer and Set
	pc.peerID = pc.Find(id)
	if pc.peerID == "" {
		onError(errors.New("Peer ID not Found"), "ID")
	}

	// Create Invite Message
	reqMsg := &md.AuthMessage{
		Event:    md.AuthMessage_REQUEST,
		From:     info,
		Metadata: sm.Metadata(),
	}

	// Send GRPC Call
	err := pc.auth.sendInvite(pc.peerID, reqMsg)
	if err != nil {
		onError(err, "Invite")
		log.Fatalln(err)
	}
}

// ^ Send Accept Message on Stream ^ //
func (pc *PeerConnection) Accept(selfInfo *md.Peer) error {
	// Find Save Path and Create Transfer
	savePath := "/" + pc.currMessage.Metadata.Name + "." + pc.currMessage.Metadata.Mime.Subtype
	pc.transfer = NewTransfer(savePath, pc.currMessage.Metadata, pc.currMessage.From, pc.progressCall, pc.completedCall)

	// Create Message
	respMsg := &md.AuthMessage{
		From:  selfInfo,
		Event: md.AuthMessage_ACCEPT,
	}

	// Send GRPC Call
	err := pc.auth.sendResponse(true, respMsg)
	if err != nil {
		onError(err, "Accept")
		log.Fatalln(err)
	}

	return nil
}

// ^ Send Decline Message on Stream ^ //
func (pc *PeerConnection) Decline(selfInfo *md.Peer) error {
	// Create Message
	respMsg := &md.AuthMessage{
		From:  selfInfo,
		Event: md.AuthMessage_DECLINE,
	}

	// Send GRPC Call
	err := pc.auth.sendResponse(false, respMsg)
	if err != nil {
		onError(err, "Decline")
		log.Fatalln(err)
	}

	// TODO: Reset Peer Info
	return nil
}

// ^ User has accepted ^ //
func (pc *PeerConnection) StartTransfer() {
	// Create New Auth Stream
	stream, err := pc.host.NewStream(context.Background(), pc.peerID, protocol.ID("/sonr/data/transfer"))
	if err != nil {
		onError(err, "Transfer")
		log.Fatalln(err)
	}

	// Initialize Writer
	writer := msgio.NewWriter(stream)
	meta := pc.safeFile.Metadata()

	// @ Check Type
	if pc.safeFile.Mime.Type == md.MIME_image {
		// Start Routine
		log.Println("Starting Base64 Write Routine")
		go writeBase64ToStream(writer, meta)
	} else {
		total := meta.Size

		// Start Routine
		log.Println("Starting Bytes Write Routine")
		go writeBytesToStream(writer, meta, total)
	}
}

// ^ Handle Incoming Stream ^ //
func (pc *PeerConnection) HandleTransfer(stream network.Stream) {
	// Set Stream
	log.Println("Stream Info: ", stream.Stat())

	// Route Data from Stream
	go func(reader msgio.ReadCloser, t *Transfer) {
		for {
			// @ Read Length Fixed Bytes
			buffer, err := reader.ReadMsg()
			if err != nil {
				onError(err, "ReadStream")
				log.Fatalln(err)
				break
			}

			// @ Unmarshal Bytes into Proto
			hasCompleted, err := t.AddBuffer(buffer)
			if err != nil {
				onError(err, "ReadStream")
				log.Fatalln(err)
				break
			}
			if hasCompleted {
				break
			}
		}
	}(msgio.NewReader(stream), pc.transfer)
}
