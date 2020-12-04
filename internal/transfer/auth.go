package transfer

import (
	"context"
	"log"

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
	invited   OnProtobuf
	responded OnProtobuf
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

	// Set Peer ID
	as.peerConn.peerID = as.peerConn.Find(args.From)

	// Set Current Message
	err := proto.Unmarshal(args.Data, as.peerConn.currMessage)
	if err != nil {
		return err
	}

	// Send Callback
	as.invited(args.Data)
	return nil
}

// ^ Respond to Authorization Invite to Peer ^ //
func (as *AuthService) sendResponse(d bool, authMsg *md.AuthMessage) {
	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(authMsg)
	if err != nil {
		onError(err, "sendInvite")
		log.Println(err)
	}

	// Send Reply
	as.currReply.Data = msgBytes
	as.currReply.Decision = d
}
