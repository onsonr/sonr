package node

import (
	"errors"
	"log"

	sentry "github.com/getsentry/sentry-go"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// ^ Interface: Callback is implemented from Plugin to receive updates ^
type Callback interface {
	OnConnected(data bool)     // Node Host has Bootstrapped
	OnReady(data bool)         // Node Host Connection Result
	OnRefreshed(data []byte)   // Lobby Updates
	OnEvent(data []byte)       // Lobby Event
	OnInvited(data []byte)     // User Invited
	OnRemoteStart(data string) // User started remote
	OnDirected(data []byte)    // User Direct-Invite from another Device
	OnResponded(data []byte)   // Peer has responded
	OnProgress(data float32)   // File Progress Updated
	OnReceived(data []byte)    // User Received File
	OnTransmitted(data []byte) // User Sent File
	OnError(data []byte)       // Internal Error
}

// ^ Passes node methods for FS/FQ ^
func (n *Node) FSCallback() md.FileCallback {
	return md.NewFileCallback(n.queued, n.multiQueued, n.error)
}

// ^ Passes node methods for Lobby ^
func (n *Node) LobbyCallback() md.LobbyCallback {
	return md.NewLobbyCallback(n.call.OnEvent, n.call.OnRefreshed, n.error, n.Peer)
}

// ^ Passes node methods for TransferController ^
func (n *Node) TransferCallback() md.TransferCallback {
	return md.NewTransferCallback(n.invited, n.call.OnResponded, n.call.OnProgress, n.received, n.transmitted, n.error)
}

// ^ queued Callback, Sends File Invite to Peer, and Notifies Client ^
func (sn *Node) queued(card *md.TransferCard, req *md.InviteRequest) {
	// Retreive Current File
	currFile := sn.fs.CurrentFile()
	if currFile != nil {
		card.Status = md.TransferCard_INVITE
		sn.transfer.NewOutgoing(currFile)

		// Create Invite Message
		invMsg := md.AuthInvite{
			From:    sn.peer,
			Payload: card.Payload,
			Card:    card,
		}

		// @ Check for Remote
		if req.IsRemote {
			// Retreive Current File
			currFile := sn.fs.CurrentFile()
			if currFile != nil {
				card.Status = md.TransferCard_INVITE
				sn.transfer.NewOutgoing(currFile)

				// Create Invite Message
				invMsg := md.AuthInvite{
					From:    sn.peer,
					Payload: card.Payload,
					Card:    card,
				}

				// Start Remote Point
				word, err := sn.transfer.StartRemotePoint(&invMsg)
				if err != nil {
					sn.error(err, "StartRemotePoint")
				}

				// Callback Point
				sn.call.OnRemoteStart(word)
			} else {
				sn.error(errors.New("No current file"), "internal:queued")
			}
		} else {
			// Get PeerID
			id, _, err := sn.lobby.Find(req.To.Id.Peer)

			// Check error
			if err != nil {
				sn.error(err, "Queued")
			}

			// Check if ID in PeerStore
			go func(inv *md.AuthInvite) {
				// Convert Protobuf to bytes
				msgBytes, err := proto.Marshal(inv)
				if err != nil {
					sn.error(err, "Marshal")
				}

				sn.transfer.RequestInvite(sn.host, id, msgBytes)
			}(&invMsg)
		}
	} else {
		sn.error(errors.New("No current file"), "internal:queued")
	}
}

// ^ multiQueued Callback, Sends File Invite to Peer, and Notifies Client ^
func (sn *Node) multiQueued(card *md.TransferCard, req *md.InviteRequest, count int) {
	// Get PeerID
	id, _, err := sn.lobby.Find(req.To.Id.Peer)

	// Check error
	if err != nil {
		sn.error(err, "Queued")
	}

	// Retreive Current File
	currFile := sn.fs.CurrentFile()
	if currFile != nil {
		card.Status = md.TransferCard_INVITE
		sn.transfer.NewOutgoing(currFile)

		// Create Invite Message
		invMsg := md.AuthInvite{
			From:    sn.peer,
			Payload: card.Payload,
			Card:    card,
		}

		// Check if ID in PeerStore
		go func(inv *md.AuthInvite) {
			// Convert Protobuf to bytes
			msgBytes, err := proto.Marshal(inv)
			if err != nil {
				sn.error(err, "Marshal")
			}

			sn.transfer.RequestInvite(sn.host, id, msgBytes)
		}(&invMsg)
	} else {
		sn.error(errors.New("No current file"), "internal:multiQueued")
	}
}

// ^ invite Callback with data for Lifecycle ^ //
func (sn *Node) invited(data []byte) {
	// Update Status
	sn.status = md.Status_INVITED
	// Callback with Data
	sn.call.OnInvited(data)
}

// ^ transmitted Callback middleware post transfer ^ //
func (sn *Node) transmitted(peer *md.Peer) {
	// Update Status
	sn.status = md.Status_AVAILABLE

	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(peer)
	if err != nil {
		sentry.CaptureException(err)
	}

	// Callback with Data
	sn.call.OnTransmitted(msgBytes)
}

// ^ received Callback middleware post transfer ^ //
func (sn *Node) received(card *md.TransferCard) {
	// Update Status
	sn.status = md.Status_AVAILABLE

	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(card)
	if err != nil {
		sentry.CaptureException(err)
	}

	// Callback with Data
	sn.call.OnReceived(msgBytes)
}

// ^ error Callback with error instance, and method ^
func (sn *Node) error(err error, method string) {
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
	sn.call.OnError(bytes)
}
