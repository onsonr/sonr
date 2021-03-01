package transfer

import (
	"context"
	"log"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	gorpc "github.com/libp2p/go-libp2p-gorpc"
	sf "github.com/sonr-io/core/internal/file"
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
	onInvite  md.OnInvite
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
	as.onInvite(receivedMessage)

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
func (pc *TransferController) Request(h host.Host, id peer.ID, msgBytes []byte) {
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
		onError(err, "Request")
	}

	// Send Callback and Reset
	pc.respondedCall(reply.Data)

	// Received Message
	responseMessage := md.AuthReply{}
	err = proto.Unmarshal(reply.Data, &responseMessage)
	if err != nil {
		// Send Error
		onError(err, "Unmarshal")
	}

	// Check Response for Accept
	if responseMessage.Decision && responseMessage.Type == md.AuthReply_Transfer {
		pc.StartTransfer(h, id, responseMessage.From)
	}
}

// ^ Send Authorize transfer on RPC ^ //
func (pc *TransferController) Authorize(decision bool, contact *md.Contact, peer *md.Peer) {
	// ** Get Current Message **
	offerMsg := pc.auth.inviteMsg

	// @ Check Reply Type for File
	switch offerMsg.Payload {
	case md.Payload_MEDIA:
		// @ Check Decision
		if decision {
			// Initialize Transfer
			pc.PrepareTransfer(offerMsg)

			// Create Accept Response
			respMsg := &md.AuthReply{
				From:     peer,
				Decision: true,
				Type:     md.AuthReply_Transfer,
			}

			// Send to Channel
			pc.auth.respCh <- respMsg

		} else {
			// Create Decline Response
			respMsg := &md.AuthReply{
				From:     peer,
				Decision: false,
				Type:     md.AuthReply_Transfer,
			}

			// Send to Channel
			pc.auth.respCh <- respMsg
		}
	case md.Payload_CONTACT:
		// @ Pass Contact Back
		// Create Accept Response
		card := sf.NewCardFromContact(peer, contact, md.TransferCard_REPLY)
		respMsg := &md.AuthReply{
			From: peer,
			Type: md.AuthReply_Contact,
			Card: &card,
		}

		// Send to Channel
		pc.auth.respCh <- respMsg
	default:
		break
	}
}

// ^ Send Authorize transfer on RPC ^ //
func (pc *TransferController) Cancel(peer *md.Peer) {
	// Create Decline Response
	respMsg := &md.AuthReply{
		From:     peer,
		Decision: false,
		Type:     md.AuthReply_Cancel,
	}

	// Send to Channel
	pc.auth.respCh <- respMsg
}
