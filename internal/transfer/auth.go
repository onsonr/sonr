package transfer

import (
	"context"
	"log"
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	gorpc "github.com/libp2p/go-libp2p-gorpc"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ****************** //
// ** GRPC Service ** //
// ****************** //
// Argument is AuthMessage protobuf
type AuthArgs struct {
	Data []byte
	From string
}

// Reply is also AuthMessage protobuf
type AuthReply struct {
	Data     []byte
	Decision bool
}

// Service Struct
type AuthService struct {
	// Current Data
	currArgs  AuthArgs
	currReply *AuthReply
	peerConn  *PeerConnection
}

// ^ Calls Invite on Remote Peer ^ //
func (as *AuthService) Invite(ctx context.Context, args AuthArgs, reply *AuthReply) error {
	log.Println("Received a Invite call: ", args.Data)
	// Process Message
	err := as.processInvite(args, reply)
	if err != nil {
		onError(err, "process")
		panic(err)
	}
	return nil
}

// ^ Processes Accept Event ^ //
func (as *AuthService) processInvite(args AuthArgs, reply *AuthReply) error {
	// Set Current Data
	as.currArgs = args
	as.currReply = reply

	// Set Current Message
	err := proto.Unmarshal(args.Data, as.peerConn.currMessage)
	if err != nil {
		return err
	}

	// Set Peer ID
	as.peerConn.peerID = as.peerConn.Find(args.From)

	// Send Callback
	as.peerConn.invitedCall(args.Data)
	return nil
}

// ********************* //
// ** Method Handling ** //
// ********************* //
// Handler Struct
type Authorization struct {
	rpcClient   *gorpc.Client
	rpcServer   *gorpc.Server
	authService AuthService
	peerConn    *PeerConnection
}

// ^ Set Sender as Server ^ //
func NewAuthRPC(pc *PeerConnection) *Authorization {
	log.Println("Creating New Auth Handler")
	// Create Server Client
	rpcServer := gorpc.NewServer(pc.host, protocol.ID("/sonr/rpc/auth"))
	rpcClient := gorpc.NewClient(pc.host, protocol.ID("/sonr/rpc/auth"))

	// Create Service
	svc := AuthService{
		peerConn: pc,
	}

	// Register Service
	err := rpcServer.Register(&svc)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("Setup RPC Auth")

	return &Authorization{
		// Rpc Properties
		rpcServer:   rpcServer,
		rpcClient:   rpcClient,
		authService: svc,
		peerConn:    pc,
	}
}

// ^ Send Authorization Invite to Peer ^ //
func (ah *Authorization) sendInvite(id peer.ID, authMsg *md.AuthMessage) error {
	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(authMsg)
	if err != nil {
		return err
	}

	// Initialize Vars
	var reply AuthReply
	args := AuthArgs{
		Data: msgBytes,
	}

	// Set Data
	startTime := time.Now()

	// Call to Peer
	err = ah.rpcClient.Call(id, "AuthService", "Invite", args, &reply)
	if err != nil {
		onError(err, "sendInvite")
		log.Fatalln(err)
	}

	// End Tracking
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	log.Printf("Auth from %s: time=%s\n", id, diff)

	// Send Callback and Reset
	ah.peerConn.respondedCall(reply.Data)

	// Handle Response
	if reply.Decision {
		// Begin Transfer
		ah.peerConn.StartTransfer()
	} else {
		// TODO: Reset RPC Data
	}

	return nil
}

// ^ Respond to Authorization Invite to Peer ^ //
func (ah *Authorization) sendResponse(d bool, authMsg *md.AuthMessage) error {
	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(authMsg)
	if err != nil {
		return err
	}

	// Send Reply
	ah.authService.currReply.Data = msgBytes
	ah.authService.currReply.Decision = d
	return nil
}
