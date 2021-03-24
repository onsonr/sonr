package node

import (
	"context"
	"errors"

	sentry "github.com/getsentry/sentry-go"
	"github.com/sonr-io/core/internal/file"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ^ Invite Processes Data and Sends Invite to Peer ^ //
func (n *Node) Invite(req *md.InviteRequest) {
	// @ 3. Send Invite to Peer
	// Set Contact
	req.Contact = n.contact
	invMsg := md.NewInviteFromRequest(req, n.peer)

	if req.IsRemote {
		// Start Remote Point
		err := n.transfer.StartRemote(&invMsg)
		if err != nil {
			n.call.Error(err, "StartRemotePoint")
		}
	} else {
		if n.HasPeer(n.local, req.To.Id.Peer) {
			// Get PeerID and Check error
			id, _, err := n.GetPeer(n.local, req.To.Id.Peer)
			if err != nil {
				sentry.CaptureException(err)
			}

			// Run Routine
			go func(inv *md.AuthInvite) {
				// Convert Protobuf to bytes
				msgBytes, err := proto.Marshal(inv)
				if err != nil {
					sentry.CaptureException(err)
				}

				n.transfer.RequestInvite(n.host, id, msgBytes)
			}(&invMsg)
		} else {
			n.call.Error(errors.New("Invalid Peer"), "Invite")
		}
	}
}

// ^ Invite Processes Data and Sends Invite to Peer ^ //
func (n *Node) InviteFile(card *md.TransferCard, req *md.InviteRequest, cf *file.ProcessedFile) {
	card.Status = md.TransferCard_INVITE
	n.transfer.NewOutgoing(cf)

	// Create Invite Message
	invMsg := md.AuthInvite{
		From:    n.peer,
		Payload: card.Payload,
		Card:    card,
	}

	// @ Check for Remote
	if req.IsRemote {
		// Start Remote Point
		err := n.transfer.StartRemote(&invMsg)
		if err != nil {
			sentry.CaptureException(err)
			n.call.Error(err, "StartRemotePoint")
		}
	} else {
		// Get PeerID
		id, _, err := n.GetPeer(n.local, req.To.Id.Peer)
		if err != nil {
			n.call.Error(err, "Queued")
		}

		// Check if ID in PeerStore
		go func(inv *md.AuthInvite) {
			// Convert Protobuf to bytes
			msgBytes, err := proto.Marshal(inv)
			if err != nil {
				n.call.Error(err, "Marshal")
			}
			n.transfer.RequestInvite(n.host, id, msgBytes)
		}(&invMsg)
	}
}

// ^ Respond to an Invitation ^ //
func (n *Node) Respond(decision bool) {
	// Send Response on PeerConnection
	n.transfer.Authorize(decision, n.contact, n.peer)

}

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
	call   md.TransferCallback
	respCh chan *md.AuthReply
	invite *md.AuthInvite
}

// ^ Calls Invite on Remote Peer ^ //
func (as *AuthService) Invited(ctx context.Context, args AuthArgs, reply *AuthResponse) error {
	// Received Message
	receivedMessage := md.AuthInvite{}
	err := proto.Unmarshal(args.Data, &receivedMessage)
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	// Set Current Message
	as.invite = &receivedMessage

	// Send Callback
	as.call.Invited(args.Data)

	// Hold Select for Invite Type
	select {
	// Received Auth Channel Message
	case m := <-as.respCh:

		// Convert Protobuf to bytes
		msgBytes, err := proto.Marshal(m)
		if err != nil {
			sentry.CaptureException(err)
			return err
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
