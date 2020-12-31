package transfer

import (
	"bytes"
	"context"
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
	"github.com/sonr-io/core/internal/lifecycle"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// Define Callback Function Types
type OnResponded func(isReceiver bool, data []byte)
type OnCompleted func(isReceiver bool, data []byte)
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
	respondedCall OnResponded
	progressCall  OnProgress
	completedCall OnCompleted

	// Info
	olc  string
	dirs *md.Directories
}

// ^ Initialize sets up new Peer Connection handler ^
func Initialize(h host.Host, ps *pubsub.PubSub, d *md.Directories, o string, ic OnProtobuf, rc OnResponded, pc OnProgress, compCall OnCompleted, ec OnError) (*PeerConnection, error) {
	// Set Package Level Callbacks
	onError = ec

	// Initialize Parameters into PeerConnection
	peerConn := &PeerConnection{
		olc:           o,
		dirs:          d,
		invitedCall:   ic,
		respondedCall: rc,
		progressCall:  pc,
		completedCall: compCall,
	}

	// Create GRPC Client/Server and Set Data Stream Handler
	h.SetStreamHandler(protocol.ID("/sonr/data/transfer"), peerConn.HandleTransfer)
	rpcServer := gorpc.NewServer(h, protocol.ID("/sonr/rpc/auth"))

	// Create AuthService
	ath := AuthService{
		onInvite: ic,
		respCh:   make(chan *md.AuthReply, 1),
	}

	// Register Service
	err := rpcServer.Register(&ath)
	if err != nil {
		return nil, err
	}

	// Set RPC Services
	peerConn.auth = &ath
	return peerConn, nil
}

// ^  Prepare for Stream, Create new Transfer ^ //
func (pc *PeerConnection) PrepareTransfer(meta *md.Metadata, own *md.Peer) *Transfer {
	// Create Transfer
	return &Transfer{
		// Inherited Properties
		meta:       meta,
		owner:      own,
		path:       pc.dirs.Temporary + "/" + meta.Name + "." + meta.Mime.Subtype,
		onProgress: pc.progressCall,
		onComplete: pc.completedCall,

		// Builders
		stringsBuilder: new(strings.Builder),
		bytesBuilder:   new(bytes.Buffer),
	}
}

// ^ User has accepted, Begin Sending Transfer ^ //
func (pc *PeerConnection) StartTransfer(h host.Host, id peer.ID, peer *md.Peer) {
	// Create New Auth Stream
	stream, err := h.NewStream(context.Background(), id, protocol.ID("/sonr/data/transfer"))
	if err != nil {
		onError(err, "Transfer")
		log.Fatalln(err)
	}

	// Marshal Peer to bytes
	peerBytes, err := proto.Marshal(peer)
	if err != nil {
		onError(err, "Transfer")
		log.Fatalln(err)
	}

	// Initialize Writer
	writer := msgio.NewWriter(stream)
	meta := pc.SafeMeta.GetMetadata()

	// @ Check Type
	if pc.SafeMeta.Mime.Type == md.MIME_image {
		// Start Routine
		log.Println("Starting Base64 Write Routine")
		go writeBase64ToStream(writer, pc.completedCall, meta, peerBytes)
	} else {
		// Start Routine
		log.Println("Starting Bytes Write Routine")
		go writeBytesToStream(writer, pc.completedCall, meta, peerBytes)
	}
}

// ^ Handle Incoming Stream ^ //
func (pc *PeerConnection) HandleTransfer(stream network.Stream) {
	// Route Data from Stream
	go func(reader msgio.ReadCloser, t *Transfer) {
		for i := 0; ; i++ {
			// @ Read Length Fixed Bytes
			buffer, err := reader.ReadMsg()
			if err != nil {
				onError(err, "ReadStream")
				log.Fatalln(err)
				break
			}

			// @ Unmarshal Bytes into Proto
			hasCompleted, err := t.addBuffer(i, buffer)
			if err != nil {
				onError(err, "ReadStream")
				log.Fatalln(err)
				break
			}

			// @ Check if All Buffer Received to Save
			if hasCompleted {
				// Sync file
				if err := pc.transfer.save(); err != nil {
					onError(err, "SaveFile")
					log.Fatalln(err)
				}
				break
			}
			lifecycle.GetState().NeedsWait()
		}
	}(msgio.NewReader(stream), pc.transfer)
}
