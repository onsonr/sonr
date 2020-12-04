package transfer

import (
	"context"
	"log"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	msgio "github.com/libp2p/go-msgio"
	sf "github.com/sonr-io/core/internal/file"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ChatRoomBufSize is the number of incoming messages to buffer for each topic.
const ChatRoomBufSize = 128

// Define Callback Function Types
type OnProtobuf func(data []byte)
type OnError func(err error, method string)

// Package Error Callback
var onError OnError

// ^ Struct: Holds/Handles GRPC Calls and Handles Data Stream  ^ //
type PeerConnection struct {
	// Handlers
	auth     *Authentication
	safeFile *sf.SafeFile
	transfer Transfer

	// Connection
	host host.Host

	// Callbacks
	invitedCall   OnProtobuf
	respondedCall OnProtobuf
	progressCall  OnProgress
	completedCall OnProtobuf

	// Info
	dirs   *md.Directories
	peerID peer.ID
}

// ^ Initialize sets up new Peer Connection handler ^
func Initialize(h host.Host, d *md.Directories, ic OnProtobuf, rc OnProtobuf, pc OnProgress, cc OnProtobuf, ec OnError) (*PeerConnection, error) {
	// Set Package Level Callbacks
	onError = ec

	// Initialize Parameters into PeerConnection
	peerConn := &PeerConnection{
		host:          h,
		dirs:          d,
		invitedCall:   ic,
		respondedCall: rc,
		progressCall:  pc,
		completedCall: cc,
	}

	// Set Handlers
	h.SetStreamHandler(protocol.ID("/sonr/transfer/data"), peerConn.HandleTransfer)

	// Create Auth Handler
	peerConn.auth = NewAuthentication(peerConn)
	return peerConn, nil
}

// ^ Send Invite to a Peer ^ //
func (pc *PeerConnection) Invite(id peer.ID, info *md.Peer, sm *sf.SafeFile) {
	// @1. Set PeerConnection Details
	pc.peerID = id
	pc.safeFile = sm

	// @2. Create Invite Message
	reqMsg := &md.AuthMessage{
		Event:    md.AuthMessage_REQUEST,
		From:     info,
		Metadata: sm.Metadata(),
	}

	// Convert to bytes
	bytes, err := proto.Marshal(reqMsg)
	if err != nil {
		onError(err, "Invite")
		log.Fatalln(err)
	}

	// Send GRPC Call
	err = pc.auth.sendInvite(id, bytes)
	if err != nil {
		onError(err, "Invite")
		log.Fatalln(err)
	}
}

// ^ User has accepted ^ //
func (pc *PeerConnection) OnAccepted(meta *md.Metadata, peer *md.Peer) {
	// Create Save Path
	savePath := "/" + meta.Name + "." + meta.Mime.Subtype

	// Set Transfer
	pc.transfer = NewTransfer(savePath, meta, peer, pc.progressCall, pc.completedCall)
}

// ^ User has accepted ^ //
func (pc *PeerConnection) HandleAccepted() {
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

	// Initialize Routine
	go pc.ReadStream(msgio.NewReader(stream))
}
