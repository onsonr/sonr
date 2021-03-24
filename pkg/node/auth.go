package node

import (
	"context"
	"errors"

	sentry "github.com/getsentry/sentry-go"
	"github.com/libp2p/go-libp2p-core/peer"
	msgio "github.com/libp2p/go-msgio"
	dt "github.com/sonr-io/core/internal/data"
	sf "github.com/sonr-io/core/internal/file"

	md "github.com/sonr-io/core/internal/models"
	tr "github.com/sonr-io/core/pkg/transfer"
	fs "github.com/sonr-io/core/pkg/user"
	"google.golang.org/protobuf/proto"
)

// ^ Invite Processes Data and Sends Invite to Peer ^ //
func (n *Node) InviteLink(req *md.InviteRequest) {
	// @ 3. Send Invite to Peer
	if n.HasPeer(n.local, req.To.Id.Peer) {
		// Get PeerID and Check error
		id, _, err := n.FindPeerInTopic(n.local, req.To.Id.Peer)
		if err != nil {
			sentry.CaptureException(err)
		}

		// Set Contact
		card := dt.NewCardFromUrl(n.peer, req.Url, md.TransferCard_DIRECT)
		invMsg := md.AuthInvite{
			IsRemote: req.IsRemote,
			From:     n.peer,
			Payload:  md.Payload_URL,
			Card:     &card,
		}

		// Run Routine
		go n.handleAuthInviteResponse(id, &invMsg, nil)
	} else {
		n.call.Error(errors.New("Invalid Peer"), "Invite")
	}
}

// ^ Invite Processes Data and Sends Invite to Peer ^ //
func (n *Node) InviteContact(req *md.InviteRequest) {
	// @ 3. Send Invite to Peer
	if n.HasPeer(n.local, req.To.Id.Peer) {
		// Get PeerID and Check error
		id, _, err := n.FindPeerInTopic(n.local, req.To.Id.Peer)
		if err != nil {
			sentry.CaptureException(err)
		}

		// Set Contact
		req.Contact = n.contact

		// Get Card
		card := dt.NewCardFromContact(n.peer, req.Contact, md.TransferCard_DIRECT)
		invMsg := md.AuthInvite{
			IsRemote: req.IsRemote,
			From:     n.peer,
			Payload:  md.Payload_CONTACT,
			Card:     &card,
		}

		// Run Routine
		go n.handleAuthInviteResponse(id, &invMsg, nil)
	} else {
		n.call.Error(errors.New("Invalid Peer"), "Invite")
	}
}

// ^ Invite Processes Data and Sends Invite to Peer ^ //
func (n *Node) InviteFile(card *md.TransferCard, req *md.InviteRequest, cf *sf.ProcessedFile) {
	card.Status = md.TransferCard_INVITE

	// Create Invite Message
	invMsg := md.AuthInvite{
		From:    n.peer,
		Payload: card.Payload,
		Card:    card,
	}

	// @ Check for Remote
	if req.IsRemote {

	} else {
		// Get PeerID
		id, _, err := n.FindPeerInTopic(n.local, req.To.Id.Peer)
		if err != nil {
			n.call.Error(err, "Queued")
		}

		// Run Routine
		go n.handleAuthInviteResponse(id, &invMsg, cf)
	}
}

// ^ Respond to an Invitation ^ //
func (n *Node) Respond(decision bool, fs *fs.SonrFS) {
	// Check Decision
	if decision {
		n.host.SetStreamHandler(n.router.Transfer(), n.handleTransferIncoming)
		n.incoming = tr.CreateIncomingFile(n.auth.invite, fs, n.call)
	}

	// @ Pass Contact Back
	if n.auth.invite.Payload == md.Payload_CONTACT {
		// Create Accept Response
		card := dt.NewCardFromContact(n.peer, n.contact, md.TransferCard_REPLY)
		resp := &md.AuthReply{
			IsRemote: n.auth.invite.IsRemote,
			From:     n.peer,
			Type:     md.AuthReply_Contact,
			Card:     &card,
		}
		// Send to Channel
		n.auth.respCh <- resp
	} else {
		// Create Accept Response
		resp := &md.AuthReply{
			IsRemote: n.auth.invite.IsRemote,
			From:     n.peer,
			Type:     md.AuthReply_Transfer,
			Decision: decision,
		}
		// Send to Channel
		n.auth.respCh <- resp
	}
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
	call   dt.NodeCallback
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

// ^ User has accepted, Begin Sending Transfer ^ //
func (n *Node) NewOutgoingTransfer(id peer.ID, peer *md.Peer, pf *sf.ProcessedFile) {
	// Create New Auth Stream
	stream, err := n.host.NewStream(n.ctx, id, n.router.Transfer())
	if err != nil {
		n.call.Error(err, "StartOutgoing")
	}

	outFile := tr.CreateOutgoingFile(pf, n.call)

	// Initialize Writer
	writer := msgio.NewWriter(stream)

	// Start Routine
	go outFile.WriteBase64(writer, peer)
}
