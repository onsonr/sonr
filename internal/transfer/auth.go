package transfer

import (
	"context"
	"fmt"
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
}

// Reply is also AuthMessage protobuf
type AuthReply struct {
	Data []byte
}

// Service Struct
type Authorization struct {
	// Reference
	peerConn *PeerConnection

	// Current Data
	currArgs  AuthArgs
	currReply *AuthReply
}

// ^ Calls Invite on Remote Peer ^ //
func (as *Authorization) Invite(ctx context.Context, args AuthArgs, reply *AuthReply) error {
	log.Println("Received a Invite call: ", args.Data)
	// Process Message
	err := as.processInvite(args, reply)
	if err != nil {
		onError(err, "process")
		panic(err)
	}

	//reply.Data = args.Data
	return nil
}

// ^ Calls Accept on Remote Peer ^ //
func (as *Authorization) Accept(ctx context.Context, args AuthArgs, reply *AuthReply) error {
	log.Println("Received a Accept call: ", args.Data)
	// Process Message
	err := as.processAccept(args, reply)
	if err != nil {
		onError(err, "process")
		panic(err)
	}

	//reply.Data = args.Data
	// Send Callback
	return nil
}

// ^ Processes Accept Event ^ //
func (as *Authorization) processInvite(args AuthArgs, reply *AuthReply) error {
	// Set Current Data
	as.currArgs = args
	as.currReply = reply

	// Send Callback
	as.peerConn.invitedCall(args.Data)
	return nil
}

// ^ Processes Accept Event ^ //
func (as *Authorization) processAccept(args AuthArgs, reply *AuthReply) error {
	// Set Current Data
	as.currArgs = args
	as.currReply = reply

	// Send Callback
	as.peerConn.respondedCall(args.Data)

	// Initiate Transfer
	as.peerConn.HandleAccepted()
	return nil
}

// ^ Processes Accept Event ^ //
func (as *Authorization) processDecline(args AuthArgs, reply *AuthReply) error {
	// Set Current Data
	as.currArgs = args
	as.currReply = reply

	// Unmarshal Bytes into Proto
	protoMsg := &md.AuthMessage{}
	err := proto.Unmarshal(args.Data, protoMsg)
	if err != nil {
		return err
	}

	// Send Callback by Event Type
	return nil
}

// ********************* //
// ** Method Handling ** //
// ********************* //
// Handler Struct
type AuthRPC struct {
	// Connection
	rpcClient *gorpc.Client
	rpcServer *gorpc.Server
}

// ^ Set Sender as Server ^ //
func NewAuthentication(pc *PeerConnection) *AuthRPC {
	log.Println("Creating New Auth Handler")
	// Create Server Client
	rpcServer := gorpc.NewServer(pc.host, protocol.ID("/sonr/auth/handler"))
	rpcClient := gorpc.NewClientWithServer(pc.host, protocol.ID("/sonr/auth/caller"), rpcServer)

	// Create Service
	svc := Authorization{
		peerConn: pc,
	}

	// Register Service
	err := rpcServer.Register(&svc)
	if err != nil {
		panic(err)
	}

	return &AuthRPC{
		// Rpc Properties
		rpcServer: rpcServer,
		rpcClient: rpcClient,
	}
}

// ^ Send Authorization Invite to Peer ^ //
func (ah *AuthRPC) sendInvite(id peer.ID, msgBytes []byte) error {
	// Initialize Vars
	var reply AuthReply
	args := AuthArgs{
		Data: msgBytes,
	}

	// Set Message Data
	startTime := time.Now()

	// Call to Peer
	err := ah.rpcClient.Call(id, "AuthService", "Invite", args, &reply)
	if err != nil {
		onError(err, "sendInvite")
		panic(err)
	}

	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Printf("Auth from %s: time=%s\n", id, diff)
	return nil
}

// ^ Send Response as Accept to Invite ^ //
func (ah *AuthRPC) sendAccept(id peer.ID, msgBytes []byte) error {
	// Initialize Vars
	var reply AuthReply
	args := AuthArgs{
		Data: msgBytes,
	}

	// Set Message Data
	startTime := time.Now()

	// Call to Peer
	err := ah.rpcClient.Call(id, "AuthService", "Invite", args, &reply)
	if err != nil {
		onError(err, "sendInvite")
		panic(err)
	}

	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Printf("Auth from %s: time=%s\n", id, diff)
	return nil
}
