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
	done      chan *gorpc.Call
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

	// Set Peer ID
	as.peerConn.peerID = as.peerConn.Find(args.From)

	// Set Current Message
	err := proto.Unmarshal(args.Data, as.peerConn.currMessage)
	if err != nil {
		return err
	}

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
	// Create Server
	rpcServer := gorpc.NewServer(pc.host, protocol.ID("/sonr/rpc/auth"))
	svc := AuthService{
		peerConn: pc,
		done:     make(chan *gorpc.Call, 1),
	}

	// Register Service
	err := rpcServer.Register(&svc)
	if err != nil {
		log.Panicln(err)
	}

	// Create Client
	rpcClient := gorpc.NewClientWithServer(pc.host, protocol.ID("/sonr/rpc/auth"), rpcServer)
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
func (ah *Authorization) sendInvite(id *peer.ID, authMsg *md.AuthMessage) {
	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(authMsg)
	if err != nil {
		onError(err, "sendInvite")
		log.Println(err)
	}

	// Initialize Vars
	var reply AuthReply
	var args AuthArgs
	args.Data = msgBytes

	// Set Data
	startTime := time.Now()

	// Call to Peer
	err = ah.rpcClient.Call(*id, "AuthService", "Invite", args, &reply)
	if err != nil {
		onError(err, "sendInvite")
		log.Panicln(err)
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
		ah.peerConn.SendFile()
	} else {
		// TODO: Reset RPC Data
	}
}

// ^ Respond to Authorization Invite to Peer ^ //
func (ah *Authorization) sendResponse(d bool, authMsg *md.AuthMessage) {
	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(authMsg)
	if err != nil {
		onError(err, "sendInvite")
		log.Println(err)
	}

	// Send Reply
	ah.authService.currReply.Data = msgBytes
	ah.authService.currReply.Decision = d
}
