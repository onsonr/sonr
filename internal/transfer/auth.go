package transfer

import (
	"context"
	"log"
	"time"

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
	as.peerConn.SetCurrentMessage(args.Data)

	// Send Callback
	as.inviteCall(args.Data)

	select {
	// Received Auth Channel Message
	case m := <-as.authCh:
		log.Println("Auth Message Received")
		// Convert Protobuf to bytes
		msgBytes, err := proto.Marshal(m)
		if err != nil {
			log.Println(err)
		}

		reply.Data = msgBytes
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
	startTime := time.Now()

	// Call to Peer
	done := make(chan *gorpc.Call, 1)
	err := rpcClient.Go(id, "AuthService", "InviteRequest", args, &reply, done)

	call := <-done
	if call.Error != nil {
		// Track Execution
		endTime := time.Now()
		diff := endTime.Sub(startTime)
		log.Printf("Failed to call %s: time=%s\n", id, diff)

		// Send Error
		onError(err, "sendInvite")
		log.Panicln(err)
	} else {
		//call.Reply := AuthReply
		// End Tracking
		endTime := time.Now()
		diff := endTime.Sub(startTime)
		log.Printf("Response %s from %s: time=%s\n", id, reply.Data, diff)

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
		if pc.currMessage != nil {
			// Initialize Transfer
			savePath := "/" + pc.currMessage.Metadata.Name + "." + pc.currMessage.Metadata.Mime.Subtype
			pc.transfer = NewTransfer(savePath, pc.currMessage.Metadata, pc.currMessage.From, pc.progressCall, pc.completedCall)
		}
		// else {
		// 	err := errors.New("AuthMessage wasnt cached")
		// 	onError(err, "sendInvite")
		// 	log.Panicln(err)
		// }
	} else {
		// Reset Peer Info
		pc.peerID = ""
		pc.currMessage = nil
	}

	// @ Handle Decision
	if decision {
		// Create Accept Response
		respMsg := &md.AuthMessage{
			From:  peer,
			Event: md.AuthMessage_ACCEPT,
		}

		// Send to Channel
		pc.auth.authCh <- respMsg
	} else {
		// Create Decline Response
		respMsg := &md.AuthMessage{
			From:  peer,
			Event: md.AuthMessage_DECLINE,
		}

		// Send to Channel
		pc.auth.authCh <- respMsg
	}

}
