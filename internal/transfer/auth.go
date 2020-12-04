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

// *********************** //
// ** Callback Handling ** //
// *********************** //
// Define Callback Function Types
type OnAccepted func(meta *md.Metadata, peer *md.Peer)
type HandleAccepted func()

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
	onAccepted     OnAccepted
	handleAccepted HandleAccepted
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
	// Unmarshal Bytes into Proto
	protoMsg := &md.AuthMessage{}
	err := proto.Unmarshal(args.Data, protoMsg)
	if err != nil {
		onError(err, "read")
		log.Fatalln(err)
	}
}

// ********************* //
// ** Method Handling ** //
// ********************* //
// Handler Struct
type Authentication struct {
	// Rpc Properties
	rpcClient *gorpc.Client
	rpcServer *gorpc.Server
}

// ^ Set Sender as Server ^ //
func NewAuthentication(pc *PeerConnection) *Authentication {
	log.Println("Creating New Auth Handler")
	// Create Server Client
	rpcServer := gorpc.NewServer(pc.host, protocol.ID("/sonr/auth/handler"))
	rpcClient := gorpc.NewClientWithServer(pc.host, protocol.ID("/sonr/auth/caller"), rpcServer)

	// Create Service
	svc := AuthService{
		onAccepted:     pc.OnAccepted,
		handleAccepted: pc.HandleAccepted,
	}

	// Register Service
	err := rpcServer.Register(&svc)
	if err != nil {
		panic(err)
	}

	return &Authentication{
		// Rpc Properties
		rpcServer: rpcServer,
		rpcClient: rpcClient,
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
