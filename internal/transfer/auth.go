package transfer

import (
	"context"
	"errors"
	"log"

	"github.com/libp2p/go-libp2p-core/host"
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
	Data []byte
}

// Service Struct
type AuthService struct {
	// Current Data
	currArgs   AuthArgs
	currReply  *AuthReply
	inviteCall OnProtobuf
	peerConn   *PeerConnection
	authCh     chan *md.AuthMessage
}

// GRPC Callback
type OnGRPCall func(data []byte, from string) error

// ^ Calls Invite on Remote Peer ^ //
func (as *AuthService) InviteRequest(ctx context.Context, args AuthArgs, reply *AuthReply) error {
	log.Println("Received a Invite call: ", args.Data)
	// Set Current Message
	err := as.peerConn.SetCurrentMessage(args.Data)
	if err != nil {
		log.Println(err)
	}

	// Send Callback
	as.inviteCall(args.Data)

	select {
	// Received Auth Channel Message
	case m := <-as.authCh:
		log.Println("User has replied")
		// Convert Protobuf to bytes
		msgBytes, err := proto.Marshal(m)
		if err != nil {
			log.Println(err)
		}

		// Set Message data and call done
		reply.Data = msgBytes
		ctx.Done()
		return nil
		// Context is Done
	case <-ctx.Done():
		return ctx.Err()
	}
}

// ^ Send SendInvite to a Peer ^ //
func (pc *PeerConnection) SendInvite(h host.Host, id peer.ID, msgBytes []byte) {
	// Create Client
	rpcClient := gorpc.NewClient(h, protocol.ID("/sonr/rpc/auth"))

	// Set Data
	var reply AuthReply
	var args AuthArgs
	args.Data = msgBytes

	// Call to Peer
	done := make(chan *gorpc.Call, 1)
	err := rpcClient.Go(id, "AuthService", "InviteRequest", args, &reply, done)

	call := <-done
	if call.Error != nil {
		// Send Error
		onError(err, "sendInvite")
		log.Panicln(err)
	} else {
		//call.Reply := AuthReply

		// Send Callback and Reset
		pc.respondedCall(reply.Data)

		// Begin Transfer
		//pc.SendFile(h)
	}
}

// ^ Send Accept Message on Stream ^ //
func (pc *PeerConnection) Respond(decision bool, peer *md.Peer) {
	// @ Check Decision
	if decision {
		// Create Transfer File
		if pc.currMessage != nil {
			// Initialize Transfer
			log.Println("Preparing for Transfer")
			savePath := "/" + pc.currMessage.Metadata.Name + "." + pc.currMessage.Metadata.Mime.Subtype
			pc.transfer = NewTransfer(savePath, pc.currMessage.Metadata, pc.currMessage.From, pc.progressCall, pc.completedCall)
		} else {
			err := errors.New("AuthMessage wasnt cached")
			onError(err, "sendInvite")
		}

		// Create Accept Response
		respMsg := &md.AuthMessage{
			From:  peer,
			Event: md.AuthMessage_ACCEPT,
		}

		// Send to Channel
		pc.auth.authCh <- respMsg
	} else {
		// Reset Peer Info
		pc.peerID = ""
		pc.currMessage = nil

		// Create Decline Response
		respMsg := &md.AuthMessage{
			From:  peer,
			Event: md.AuthMessage_DECLINE,
		}

		// Send to Channel
		pc.auth.authCh <- respMsg
	}
}
