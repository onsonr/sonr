package bind

import (
	"errors"
	"log"

	"github.com/getsentry/sentry-go"
	"google.golang.org/protobuf/proto"
	md "github.com/sonr-io/core/internal/models"
)

// * Interface: Callback is implemented from Plugin to receive updates * //
type Callback interface {
	OnConnected(data bool)     // Node Host has Bootstrapped
	OnReady(data bool)         // Node Host has Bootstrapped
	OnRefreshed(data []byte)   // Lobby Updates
	OnEvent(data []byte)       // Lobby Event
	OnRemoteStart(data []byte) // User started remote
	OnInvited(data []byte)     // User Invited
	OnDirected(data []byte)    // User Direct-Invite from another Device
	OnResponded(data []byte)   // Peer has responded
	OnProgress(data float32)   // File Progress Updated
	OnReceived(data []byte)    // User Received File
	OnTransmitted(data []byte) // User Sent File
	OnError(data []byte)       // Internal Error
}

// ^ queued Callback, Sends File Invite to Peer, and Notifies Client ^
func (mn *MobileNode) queued(card *md.TransferCard, req *md.InviteRequest) {
	// Retreive Current File
	currFile := n.fs.CurrentFile()
	if currFile != nil {
		card.Status = md.TransferCard_INVITE
		n.transfer.NewOutgoing(currFile)

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
				n.error(err, "StartRemotePoint")
			}
		} else {
			// Get PeerID
			id, _, err := n.GetPeer(n.local, req.To.Id.Peer)
			if err != nil {
				n.error(err, "Queued")
			}

			// Check if ID in PeerStore
			go func(inv *md.AuthInvite) {
				// Convert Protobuf to bytes
				msgBytes, err := proto.Marshal(inv)
				if err != nil {
					n.error(err, "Marshal")
				}

				n.transfer.RequestInvite(n.host, id, msgBytes)
			}(&invMsg)
		}
	} else {
		n.error(errors.New("No current file"), "internal:queued")
	}
}

// ^ multiQueued Callback, Sends File Invite to Peer, and Notifies Client ^
func (mn *MobileNode) multiQueued(card *md.TransferCard, req *md.InviteRequest, count int) {
	// Get PeerID
	id, _, err := n.GetPeer(n.local, req.To.Id.Peer)
	// Check error
	if err != nil {
		n.error(err, "Queued")
	}

	// Retreive Current File
	currFile := n.fs.CurrentFile()
	if currFile != nil {
		card.Status = md.TransferCard_INVITE
		n.transfer.NewOutgoing(currFile)

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
				n.error(err, "StartRemotePoint")
			}
		} else {
			// Check if ID in PeerStore
			go func(inv *md.AuthInvite) {
				// Convert Protobuf to bytes
				msgBytes, err := proto.Marshal(inv)
				if err != nil {
					n.error(err, "Marshal")
				}

				mn.node.transfer.RequestInvite(mn.host, id, msgBytes)
			}(&invMsg)
		}
	} else {
		mn.error(errors.New("No current file"), "internal:multiQueued")
	}
}

// ^ invite Callback with data for Lifecycle ^ //
func (mn *MobileNode) invited(data []byte) {
	// Update Status
	mn.status = md.Status_INVITED
	// Callback with Data
	mn.call.OnInvited(data)
}

// ^ transmitted Callback middleware post transfer ^ //
func (mn *MobileNode) transmitted(peer *md.Peer) {
	// Update Status
	mn.status = md.Status_AVAILABLE

	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(peer)
	if err != nil {
		sentry.CaptureException(err)
	}

	// Callback with Data
	mn.call.OnTransmitted(msgBytes)
}

// ^ received Callback middleware post transfer ^ //
func (mn *MobileNode) received(card *md.TransferCard) {
	// Update Status
	mn.status = md.Status_AVAILABLE

	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(card)
	if err != nil {
		sentry.CaptureException(err)
	}

	// Callback with Data
	mn.call.OnReceived(msgBytes)
}

// ^ error Callback with error instance, and method ^
func (mn *MobileNode) error(err error, method string) {
	// Log Error
	sentry.CaptureException(err)

	// Create Error ProtoBuf
	errorMsg := md.ErrorMessage{
		Message: err.Error(),
		Method:  method,
	}

	// Convert Message to bytes
	bytes, err := proto.Marshal(&errorMsg)
	if err != nil {
		log.Println("Cannot Marshal Error Protobuf: ", err)
	}
	// Send Callback
	mn.call.OnError(bytes)
}
