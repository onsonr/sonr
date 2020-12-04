package transfer

import (
	"log"
	"time"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	gorpc "github.com/libp2p/go-libp2p-gorpc"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	msgio "github.com/libp2p/go-msgio"
	sf "github.com/sonr-io/core/internal/file"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
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
	// Connection
	host      host.Host
	pubSub    *pubsub.PubSub
	rpcClient *gorpc.Client
	rpcServer *gorpc.Server
	ascv      *Authorization

	// Data Handlers
	safeFile *sf.SafeMetadata
	transfer *Transfer

	// Callbacks
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
		// Connection
		host:   h,
		pubSub: ps,

		// Info
		olc:  o,
		dirs: d,

		// Callbacks
		invitedCall:   ic,
		respondedCall: rc,
		progressCall:  pc,
		completedCall: cc,
	}

	// Create Server/Client/Service
	peerConn.rpcServer = gorpc.NewServer(peerConn.host, protocol.ID("/sonr/rpc/auth"))
	peerConn.rpcClient = gorpc.NewClientWithServer(peerConn.host, protocol.ID("/sonr/rpc/auth"), peerConn.rpcServer)
	peerConn.ascv = &Authorization{
		peerConn: peerConn,
	}

	// Register Service
	err := peerConn.rpcServer.Register(peerConn.ascv)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("Created RPC AuthService")

	// Set Data Stream Handler
	peerConn.host.SetStreamHandler(protocol.ID("/sonr/data/transfer"), peerConn.HandleTransfer)
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
func (pc *PeerConnection) Invite(id peer.ID, info *md.Peer, sm *sf.SafeMetadata) {
	// Set SafeFile
	pc.safeFile = sm

	// Create Invite Message
	reqMsg := &md.AuthMessage{
		Event:    md.AuthMessage_REQUEST,
		From:     info,
		Metadata: sm.GetMetadata(),
	}

	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(reqMsg)
	if err != nil {
		onError(err, "sendInvite")
		log.Println(err)
	}

	// Send GRPC Call
	//go pc.auth.sendInvite(pc.peerID, reqMsg)
	go func(id peer.ID, data []byte) {
		// Initialize Vars
		var reply AuthReply
		var args AuthArgs
		args.Data = msgBytes

		// Set Data
		startTime := time.Now()

		// Call to Peer
		err = pc.rpcClient.Call(id, "Authorization", "Invite", args, &reply)
		if err != nil {
			onError(err, "sendInvite")
			log.Panicln(err)
		}

		// End Tracking
		endTime := time.Now()
		diff := endTime.Sub(startTime)
		log.Printf("Auth from %s: time=%s\n", id, diff)

		// Send Callback and Reset
		pc.respondedCall(reply.Data)

		// Handle Response
		if reply.Decision {
			// Begin Transfer
			pc.SendFile()
		}
	}(id, msgBytes)
}

// ^ Send Accept Message on Stream ^ //
func (pc *PeerConnection) SendResponse(decision bool, selfInfo *md.Peer) {
	// Initialize Message
	var respMsg *md.AuthMessage

	// Check Decision
	if decision {
		// Initialize Transfer
		savePath := "/" + pc.currMessage.Metadata.Name + "." + pc.currMessage.Metadata.Mime.Subtype
		pc.transfer = NewTransfer(savePath, pc.currMessage.Metadata, pc.currMessage.From, pc.progressCall, pc.completedCall)

		// Create Accept Response
		respMsg = &md.AuthMessage{
			From:  selfInfo,
			Event: md.AuthMessage_ACCEPT,
		}
	} else {
		// Reset Peer Info
		pc.peerID = ""
		pc.currMessage = nil

		// Create Decline Response
		respMsg = &md.AuthMessage{
			From:  selfInfo,
			Event: md.AuthMessage_DECLINE,
		}
	}

	// Send GRPC Call
	go pc.ascv.sendResponse(decision, respMsg)
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
