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

// Define Callback Function Types
type OnProtobuf func([]byte)
type OnError func(err error, method string)

// Package Error Callback
var onError OnError

// ^ Struct: Holds/Handles GRPC Calls and Handles Data Stream  ^ //
type PeerConnection struct {
	// Connection
	auth *AuthService

	// Data Handlers
	SafeMeta *sf.SafeMetadata
	transfer *Transfer

	// Callbacks
	invitedCall   OnProtobuf
	respondedCall OnProtobuf
	progressCall  OnProgress
	completedCall OnProtobuf

	// Info
	olc    string
	dirs   *md.Directories
	peerID peer.ID
}

// ^ Initialize sets up new Peer Connection handler ^
func Initialize(h host.Host, ps *pubsub.PubSub, d *md.Directories, o string, ic OnProtobuf, rc OnProtobuf, pc OnProgress, cc OnProtobuf, ec OnError) (*PeerConnection, error) {
	// Set Package Level Callbacks
	onError = ec

	// Initialize Parameters into PeerConnection
	peerConn := &PeerConnection{}
	peerConn.olc = o
	peerConn.dirs = d
	peerConn.invitedCall = ic
	peerConn.respondedCall = rc
	peerConn.progressCall = pc
	peerConn.completedCall = cc

	// Create GRPC Client/Server and Set Data Stream Handler
	h.SetStreamHandler(protocol.ID("/sonr/data/transfer"), peerConn.HandleTransfer)
	rpcServer := gorpc.NewServer(h, protocol.ID("/sonr/rpc/auth"))

	// Create AuthService
	ath := AuthService{
		inviteCall: ic,
		authCh:     make(chan *md.AuthMessage, 100),
	}

	// Register Service
	err := rpcServer.Register(&ath)
	if err != nil {
		return nil, err
	}
	log.Println("Created RPC AuthService")

	// Set RPC Services
	peerConn.auth = &ath
	return peerConn, nil
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

			// @ Check if All Buffer Received
			if hasCompleted {
				pc.transfer.Save()
				break
			}
		}
	}(msgio.NewReader(stream), pc.transfer)
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
		count:       0,
		currentSize: 0,
	}
}
