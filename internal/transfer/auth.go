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
)

// *********************** //
// ** Callback Handling ** //
// *********************** //
// Define Callback Function Types
type OnInvited func(meta *md.Metadata, peer *md.Peer)
type OnAccepted func()

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
type AuthService struct {
}

// ^ Calls Invite on Remote Peer ^ //
func (as *AuthService) Invite(ctx context.Context, args AuthArgs, reply *AuthReply) error {
	log.Println("Received a Invite call: ", args.Data)
	as.process(args)
	reply.Data = args.Data
	// Send Callback
	return nil
}

// ^ Calls Accept on Remote Peer ^ //
func (as *AuthService) Accept(ctx context.Context, args AuthArgs, reply *AuthReply) error {
	log.Println("Received a Accept call: ", args.Data)
	as.process(args)
	reply.Data = args.Data
	// Send Callback
	return nil
}

// ^ Processes Message Content ^ //
func (as *AuthService) process(args AuthArgs) {

}

// ********************* //
// ** Method Handling ** //
// ********************* //
// Handler Struct
type Authentication struct {
	// Rpc Properties
	rpcClient *gorpc.Client
	rpcServer *gorpc.Server

	// Callbacks

}

// ^ Set Sender as Server ^ //
func NewAuthentication(pc *PeerConnection) *Authentication {
	log.Println("Creating New Auth Handler")
	// Create Server Client
	rpcServer := gorpc.NewServer(pc.host, protocol.ID("/sonr/auth/handler"))
	rpcClient := gorpc.NewClientWithServer(pc.host, protocol.ID("/sonr/auth/caller"), rpcServer)

	// Register Functions
	svc := AuthService{}
	err := rpcServer.Register(&svc)
	if err != nil {
		panic(err)
	}

	return &Authentication{
		// Rpc Properties
		rpcServer: rpcServer,
		rpcClient: rpcClient,

		// Callback Properties
		
	}
}

// ^ Send Authorization Invite to Data Sender ^ //
func (ah *Authentication) sendInvite(id peer.ID, msgBytes []byte) error {
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
