package transfer

import (
	"context"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	gorpc "github.com/libp2p/go-libp2p-gorpc"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	msgio "github.com/libp2p/go-msgio"
	sf "github.com/sonr-io/core/internal/file"
	md "github.com/sonr-io/core/internal/models"
	fs "github.com/sonr-io/core/pkg/data"
	net "github.com/sonr-io/core/pkg/net"
)

// ^ Struct: Holds/Handles GRPC Calls and Handles Data Stream  ^ //
type TransferController struct {
	// Connection
	ctx    context.Context
	auth   *AuthService
	remote *RemotePoint

	// Networking
	host   host.Host
	pubsub *pubsub.PubSub
	router *net.ProtocolRouter

	// Data Handlers
	outgoing *OutgoingFile
	incoming *IncomingFile

	// Callbacks
	call md.TransferCallback

	// Info
	fs *fs.SonrFS
}

// ^ Initialize sets up new Peer Connection handler ^
func Initialize(ctx context.Context, h host.Host, ps *pubsub.PubSub, fs *fs.SonrFS, pr *net.ProtocolRouter, tc md.TransferCallback) (*TransferController, error) {
	// Initialize Parameters into PeerConnection
	peerConn := &TransferController{
		router: pr,
		fs:     fs,
		call:   tc,
		host:   h,
		ctx:    ctx,
		pubsub: ps,
	}

	// Create GRPC Client/Server and Set Data Stream Handler
	h.SetStreamHandler(peerConn.router.Transfer(), peerConn.HandleIncoming)
	rpcServer := gorpc.NewServer(h, peerConn.router.Auth())

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
	pc.incoming = NewIncomingFile(inv, pc.fs, pc.call)
}

// ^  Set Outgoing Transfer ^ //
func (pc *TransferController) NewOutgoing(pf *sf.ProcessedFile) {
	// Initialize Outgoing
	pc.outgoing = NewOutgoingFile(pf, pc.call)
}

// ^ User has accepted, Begin Sending Transfer ^ //
func (pc *TransferController) StartOutgoing(h host.Host, id peer.ID, peer *md.Peer) {
	// Create New Auth Stream
	stream, err := h.NewStream(context.Background(), id, pc.router.Transfer())
	if err != nil {
		pc.call.Error(err, "StartOutgoing")
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
				pc.call.Error(err, "HandleIncoming:ReadMsg")
				break
			}

			// @ Unmarshal Bytes into Proto
			hasCompleted, err := t.AddBuffer(i, buffer)
			if err != nil {
				pc.call.Error(err, "HandleIncoming:AddBuffer")
				break
			}

			// @ Check if All Buffer Received to Save
			if hasCompleted {
				// Sync file
				if err := pc.incoming.Save(); err != nil {
					pc.call.Error(err, "HandleIncoming:Save")
				}
				break
			}
			md.GetState().NeedsWait()
		}
	}(msgio.NewReader(stream), pc.incoming)
}
