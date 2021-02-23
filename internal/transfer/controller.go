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
	md "github.com/sonr-io/core/internal/models"
)

// ^ Struct: Holds/Handles GRPC Calls and Handles Data Stream  ^ //
type TransferController struct {
	// Connection
	auth *AuthService

	// Data Handlers
	ProcessedFile *sf.ProcessedFile
	transfer      *IncomingFile

	// Callbacks
	call md.TransferCallback

	// Info
	olc  string
	dirs *md.Directories
}

// ^ Initialize sets up new Peer Connection handler ^
func Initialize(h host.Host, ps *pubsub.PubSub, d *md.Directories, o string, tc md.TransferCallback) (*TransferController, error) {
	// Initialize Parameters into PeerConnection
	peerConn := &TransferController{
		olc:  o,
		dirs: d,
		call: tc,
	}

	// Create GRPC Client/Server and Set Data Stream Handler
	h.SetStreamHandler(protocol.ID("/sonr/data/transfer"), peerConn.HandleTransfer)
	rpcServer := gorpc.NewServer(h, protocol.ID("/sonr/rpc/auth"))

	// Create AuthService
	ath := AuthService{
		call:   tc,
		respCh: make(chan *md.AuthReply, 1),
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
func (tc *TransferController) PrepareIncoming(inv *md.AuthInvite) {
	// Initialize Transfer
	tc.transfer = NewIncomingFile(inv, tc.dirs, tc.call)
}

// ^ User has accepted, Begin Sending Transfer ^ //
func (tc *TransferController) StartTransfer(h host.Host, id peer.ID, peer *md.Peer) {
	// Create New Auth Stream
	stream, err := h.NewStream(context.Background(), id, protocol.ID("/sonr/data/transfer"))
	if err != nil {
		tc.call.OnError(err, "Transfer")
		log.Fatalln(err)
	}

	// Initialize Writer
	writer := msgio.NewWriter(stream)

	// Start Routine
	go writeBase64ToStream(writer, tc.call, tc.ProcessedFile, peer)
}

// ^ Handle Incoming Stream ^ //
func (tc *TransferController) HandleTransfer(stream network.Stream) {
	// Route Data from Stream
	go func(reader msgio.ReadCloser, t *IncomingFile) {
		for i := 0; ; i++ {
			// @ Read Length Fixed Bytes
			buffer, err := reader.ReadMsg()
			if err != nil {
				tc.call.OnError(err, "ReadStream")
				log.Fatalln(err)
				break
			}

			// @ Unmarshal Bytes into Proto
			hasCompleted, err := t.AddBuffer(i, buffer)
			if err != nil {
				tc.call.OnError(err, "ReadStream")
				log.Fatalln(err)
				break
			}

			// @ Check if All Buffer Received to Save
			if hasCompleted {
				// Sync file
				if err := tc.transfer.Save(); err != nil {
					tc.call.OnError(err, "SaveFile")
					log.Fatalln(err)
				}
				break
			}
			lifecycle.GetState().NeedsWait()
		}
	}(msgio.NewReader(stream), tc.transfer)
}
