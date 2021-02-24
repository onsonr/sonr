package transfer

import (
	"context"
	"log"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	gorpc "github.com/libp2p/go-libp2p-gorpc"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	msgio "github.com/libp2p/go-msgio"
	sf "github.com/sonr-io/core/internal/file"
	"github.com/sonr-io/core/internal/lifecycle"
	lf "github.com/sonr-io/core/internal/lifecycle"
	md "github.com/sonr-io/core/internal/models"
)

// Package Error Callback
var onError lf.OnError

// ^ Struct: Holds/Handles GRPC Calls and Handles Data Stream  ^ //
type PeerConnection struct {
	// Connection
	auth *AuthService

	// Data Handlers
	ProcessedFile *sf.ProcessedFile
	transfer      *sf.TransferFile

	// Callbacks
	invitedCall     lf.OnInvite
	respondedCall   lf.OnProtobuf
	progressCall    lf.OnProgress
	receivedCall    lf.OnReceived
	transmittedCall lf.OnTransmitted

	// Info
	olc  string
	dirs *md.Directories
}

// ^ Initialize sets up new Peer Connection handler ^
func Initialize(h host.Host, ps *pubsub.PubSub, d *md.Directories, o string, tc lf.TransferCallbacks) (*PeerConnection, error) {
	// Set Package Level Callbacks
	onError = tc.CallError

	// Initialize Parameters into PeerConnection
	peerConn := &PeerConnection{
		olc:             o,
		dirs:            d,
		invitedCall:     tc.CallInvited,
		respondedCall:   tc.CallResponded,
		progressCall:    tc.CallProgress,
		receivedCall:    tc.CallReceived,
		transmittedCall: tc.CallTransmitted,
	}

	// Create GRPC Client/Server and Set Data Stream Handler
	h.SetStreamHandler(protocol.ID("/sonr/data/transfer"), peerConn.HandleTransfer)
	rpcServer := gorpc.NewServer(h, protocol.ID("/sonr/rpc/auth"))

	// Create AuthService
	ath := AuthService{
		onInvite: tc.CallInvited,
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
func (pc *PeerConnection) PrepareTransfer(inv *md.AuthInvite) {
	// Initialize Transfer
	pc.transfer = sf.NewTransfer(inv, pc.dirs, pc.progressCall, pc.receivedCall)
}

// ^ User has accepted, Begin Sending Transfer ^ //
func (pc *PeerConnection) StartTransfer(h host.Host, id peer.ID, peer *md.Peer) {
	// Create New Auth Stream
	stream, err := h.NewStream(context.Background(), id, protocol.ID("/sonr/data/transfer"))
	if err != nil {
		onError(err, "Transfer")
		log.Fatalln(err)
	}

	// Initialize Writer
	writer := msgio.NewWriter(stream)

	// Start Routine
	go writeBase64ToStream(writer, pc.transmittedCall, pc.ProcessedFile, peer)
}

// ^ Handle Incoming Stream ^ //
func (pc *PeerConnection) HandleTransfer(stream network.Stream) {
	// Route Data from Stream
	go func(reader msgio.ReadCloser, t *sf.TransferFile) {
		for i := 0; ; i++ {
			// @ Read Length Fixed Bytes
			buffer, err := reader.ReadMsg()
			if err != nil {
				onError(err, "ReadStream")
				log.Fatalln(err)
				break
			}

			// @ Unmarshal Bytes into Proto
			hasCompleted, err := t.AddBuffer(i, buffer)
			if err != nil {
				onError(err, "ReadStream")
				log.Fatalln(err)
				break
			}

			// @ Check if All Buffer Received to Save
			if hasCompleted {
				// Sync file
				if err := pc.transfer.Save(); err != nil {
					onError(err, "SaveFile")
					log.Fatalln(err)
				}
				break
			}
			lifecycle.GetState().NeedsWait()
		}
	}(msgio.NewReader(stream), pc.transfer)
}
