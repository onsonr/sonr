package transfer

import (
	"bytes"
	"log"
	"strings"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	gorpc "github.com/libp2p/go-libp2p-gorpc"
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
	// Connection
	host      host.Host
	pubSub    *pubsub.PubSub
	rpcClient *gorpc.Client
	rpcServer *gorpc.Server
	ascv      Authorization

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
func Initialize(h host.Host, ps *pubsub.PubSub, d *md.Directories, o string, ic OnProtobuf, rc OnProtobuf, pc OnProgress, cc OnProtobuf, ec OnError) (PeerConnection, error) {
	// Set Package Level Callbacks
	onError = ec

	// Initialize Parameters into PeerConnection
	peerConn := PeerConnection{
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
	peerConn.ascv = Authorization{
		peerConn: &peerConn,
	}

	// Register Service
	err := peerConn.rpcServer.Register(&peerConn.ascv)
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

// ^ Create new SonrFile struct with meta and documents directory ^ //
func NewTransfer(savePath string, meta *md.Metadata, own *md.Peer, op OnProgress, oc OnProtobuf) *Transfer {
	return &Transfer{
		// Inherited Properties
		metadata:   meta,
		path:       savePath,
		owner:      own,
		onProgress: op,
		onComplete: oc,

		// Builders
		stringsBuilder: new(strings.Builder),
		bytesBuilder:   new(bytes.Buffer),

		// Tracking
		count: 0,
		size:  0,
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
