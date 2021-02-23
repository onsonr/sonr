package transfer

import (
	"context"
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
}

// Reply is also AuthMessage protobuf
type AuthResponse struct {
	Data []byte
}

// Service Struct
type AuthService struct {
	// Current Data
	call      md.TransferCallback
	respCh    chan *md.AuthReply
	inviteMsg *md.AuthInvite
}

// ^ Calls Invite on Remote Peer ^ //
func (as *AuthService) Invited(ctx context.Context, args AuthArgs, reply *AuthResponse) error {
	// Received Message
	receivedMessage := &md.AuthInvite{}
	err := proto.Unmarshal(args.Data, receivedMessage)
	if err != nil {
		return err
	}

	// Set Current Message
	as.inviteMsg = receivedMessage

	// Send Callback
	as.call.OnInvited(receivedMessage)

	// Hold Select for Invite Type
	if !as.inviteMsg.IsDirect {
		select {
		// Received Auth Channel Message
		case m := <-as.respCh:

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
			return nil
		}
	}

	// Begin Direct Transfer
	log.Println("Direct Transfer")
	return nil
}

// ^ Send Request to a Peer ^ //
func (tc *TransferController) Request(h host.Host, id peer.ID, msgBytes []byte) {
	// Initialize Data
	rpcClient := gorpc.NewClient(h, protocol.ID("/sonr/rpc/auth"))
	var reply AuthResponse
	var args AuthArgs
	args.Data = msgBytes

	// Call to Peer
	done := make(chan *gorpc.Call, 1)
	err := rpcClient.Go(id, "AuthService", "Invited", args, &reply, done)

	// Initiate Call on transfer
	call := <-done
	if call.Error != nil {
		// Send Error
		tc.call.OnError(err, "Request")
	}

	// Send Callback and Reset
	tc.call.OnResponded(reply.Data)

	// Received Message
	responseMessage := md.AuthReply{}
	err = proto.Unmarshal(reply.Data, &responseMessage)
	if err != nil {
		// Send Error
		tc.call.OnError(err, "Unmarshal")
	}

	// Check Response for Accept
	if responseMessage.Decision && responseMessage.Payload == md.Payload_UNDEFINED {
		// Begin Transfer
		tc.StartTransfer(h, id, responseMessage.From)
	}
}

// ^ Send Authorize transfer on RPC ^ //
func (tc *TransferController) Authorize(decision bool, contact *md.Contact, peer *md.Peer) {
	// ** Get Current Message **
	offerMsg := tc.auth.inviteMsg

	// @ Check Reply Type for File
	switch offerMsg.Payload {
	case md.Payload_MEDIA:
		// @ Check Decision
		if decision {
			// Initialize Transfer
			tc.PrepareIncoming(offerMsg)

			// Create Accept Response
			respMsg := &md.AuthReply{
				From:     peer,
				Decision: true,
				Payload:  md.Payload_UNDEFINED,
			}

			// Send to Channel
			tc.auth.respCh <- respMsg

		} else {
			// Create Decline Response
			respMsg := &md.AuthReply{
				From:     peer,
				Decision: false,
				Payload:  md.Payload_UNDEFINED,
			}

			// Send to Channel
			tc.auth.respCh <- respMsg
		}
	case md.Payload_CONTACT:
		// @ Pass Contact Back
		// Create Accept Response
		card := md.NewCardFromContact(peer, contact, md.TransferCard_REPLY)
		respMsg := &md.AuthReply{
			From:    peer,
			Payload: md.Payload_CONTACT,
			Card:    &card,
		}

		// Send to Channel
		tc.auth.respCh <- respMsg
	default:
		break
	}
}
