package transfer

import (
	"context"

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

// ^ Struct: Holds/Handles GRPC Calls and Handles Data Stream  ^ //
type TransferController struct {
	// Connection
	auth *AuthService

	// Data Handlers
	outgoing *OutgoingFile
	incoming *IncomingFile

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
	h.SetStreamHandler(protocol.ID("/sonr/transfer/data"), peerConn.HandleIncoming)
	rpcServer := gorpc.NewServer(h, protocol.ID("/sonr/transfer/auth"))

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

// ^  Prepare for Stream, Create Incoming Transfer ^ //
func (pc *TransferController) NewIncoming(inv *md.AuthInvite) {
	// Initialize Incoming
	pc.incoming = NewIncomingFile(inv, pc.dirs, pc.call)
}

// ^  Set Outgoing Transfer ^ //
func (pc *TransferController) NewOutgoing(pf *sf.ProcessedFile) {
	// Initialize Outgoing
	pc.outgoing = NewOutgoingFile(pf, pc.call)
}

// ^ User has accepted, Begin Sending Transfer ^ //
func (pc *TransferController) StartOutgoing(h host.Host, id peer.ID, peer *md.Peer) {
	// Create New Auth Stream
	stream, err := h.NewStream(context.Background(), id, protocol.ID("/sonr/transfer/data"))
	if err != nil {
		pc.call.Error(err)

	}

	// Initialize Writer
	writer := msgio.NewWriter(stream)

	// Start Routine
	go pc.outgoing.WriteBase64(writer, peer)
}

// ^ Handle Incoming Stream ^ //
func (pc *TransferController) HandleIncoming(stream network.Stream) {
	// Route Data from Stream
	go func(reader msgio.ReadCloser, t *IncomingFile) {
		for i := 0; ; i++ {
			// @ Read Length Fixed Bytes
			buffer, err := reader.ReadMsg()
			if err != nil {
				pc.call.Error(err)
				break
			}

			// @ Unmarshal Bytes into Proto
			hasCompleted, err := t.AddBuffer(i, buffer)
			if err != nil {
				pc.call.Error(err)
				break
			}

			// @ Check if All Buffer Received to Save
			if hasCompleted {
				// Sync file
				if err := pc.incoming.Save(); err != nil {
					pc.call.Error(err)
				}
				break
			}
			md.GetState().NeedsWait()
		}
	}(msgio.NewReader(stream), pc.incoming)
}
