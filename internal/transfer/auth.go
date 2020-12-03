package transfer

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	gorpc "github.com/libp2p/go-libp2p-gorpc"
)

// ****************** //
// ** GRPC Service ** //
// ****************** //
// Argument is AuthMessage protobuf
type AuthArgs struct {
	Message []byte
}

// Reply is also AuthMessage protobuf
type AuthReply struct {
	Message []byte
}

// Service Struct
type AuthService struct{}

// ^ Calls Invite on Remote Peer ^ //
func (t *AuthService) Invite(ctx context.Context, argType AuthArgs, replyType *AuthReply) error {
	log.Println("Received a Invite call: ", argType.Message)
	replyType.Message = argType.Message
	// Send Callback
	return nil
}

// ^ Calls Accept on Remote Peer ^ //
func (t *AuthService) Accept(ctx context.Context, argType AuthArgs, replyType *AuthReply) error {
	log.Println("Received a Accept call: ", argType.Message)
	replyType.Message = argType.Message
	// Send Callback
	return nil
}

// ********************* //
// ** Method Handling ** //
// ********************* //
// Handler Struct
type AuthHandler struct {
	rpcClient *gorpc.Client
	rpcServer *gorpc.Server
	peerConn  *PeerConnection
}

// ^ Set Sender as Server ^ //
func NewAuthHandler(pc *PeerConnection) *AuthHandler {
	log.Println("Setting Data Sender as RPC Server")
	rpcServer := gorpc.NewServer(pc.host, protocol.ID("/sonr/auth/handler"))
	rpcClient := gorpc.NewClientWithServer(pc.host, protocol.ID("/sonr/auth/caller"), rpcServer)

	svc := AuthService{}
	err := rpcServer.Register(&svc)
	if err != nil {
		panic(err)
	}
	return &AuthHandler{
		// Rpc Properties
		rpcServer: rpcServer,
		rpcClient: rpcClient,

		// References
		peerConn: pc,
	}
}

// ^ Send Authorization Invite to Data Sender ^ //
func (ah *AuthHandler) sendInvite(id peer.ID, msgBytes []byte) error {
	// Initialize Vars
	var reply AuthReply
	args := AuthArgs{
		Message: msgBytes,
	}

	// Set Message Data
	startTime := time.Now()

	// Call to Peer
	err := ah.rpcClient.Call(id, "AuthService", "Invite", args, &reply)
	if err != nil {
		panic(err)
	}

	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Printf("Auth from %s: time=%s\n", id, diff)
	return nil
}
