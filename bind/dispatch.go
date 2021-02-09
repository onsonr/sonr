package sonr

import (
	"time"

	sf "github.com/sonr-io/core/internal/file"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ^ Send Invite with a File ^ //
func (sn *Node) DirectWithFile(peerId string) {
	// Get PeerID
	id, _, err := sn.lobby.Find(peerId)

	// Check error
	if err != nil {
		sn.error(err, "DirectWithFile")
	}

	// Create Invite Message with Payload
	time.Sleep(time.Millisecond * 100)

	// Retreive Current File
	currFile := sn.currentFile()
	card := currFile.GetTransferCard()
	card.Status = md.TransferCard_DIRECT
	sn.peerConn.SafePreview = currFile

	// Create Invite Message
	invMsg := md.AuthInvite{
		From:     sn.peer,
		Payload:  card.Payload,
		Card:     card,
		IsFile:   true,
		IsDirect: true,
	}

	// Check if ID in PeerStore
	go func(inv *md.AuthInvite) {
		// Convert Protobuf to bytes
		msgBytes, err := proto.Marshal(inv)
		if err != nil {
			sn.error(err, "Marshal")
		}

		sn.peerConn.Request(sn.host, id, msgBytes)
	}(&invMsg)
}

// ^ Send Invite with URL Link ^ //
func (sn *Node) DirectWithURL(peerId string, url string) {
	// Get PeerID
	id, _, err := sn.lobby.Find(peerId)

	// Check error
	if err != nil {
		sn.error(err, "DirectWithURL")
	}

	// Create Invite Message with Payload
	invMsg := md.AuthInvite{
		From:     sn.peer,
		Payload:  md.Payload_URL,
		Card:     sf.NewCardFromUrl(sn.peer.Profile, url, md.TransferCard_DIRECT),
		IsFile:   false,
		IsDirect: true,
	}

	// Check if ID in PeerStore
	go func(inv *md.AuthInvite) {
		// Convert Protobuf to bytes
		msgBytes, err := proto.Marshal(inv)
		if err != nil {
			sn.error(err, "Marshal")
		}

		sn.peerConn.Request(sn.host, id, msgBytes)
	}(&invMsg)
}

// ^ Send Invite with a File ^ //
func (sn *Node) InviteWithFile(peerId string) {
	// Get PeerID
	id, _, err := sn.lobby.Find(peerId)

	// Check error
	if err != nil {
		sn.error(err, "InviteWithFile")
	}

	// Create Invite Message with Payload
	time.Sleep(time.Millisecond * 100)

	// Retreive Current File
	currFile := sn.currentFile()
	card := currFile.GetTransferCard()
	card.Status = md.TransferCard_INVITE
	sn.peerConn.SafePreview = currFile

	// Create Invite Message
	invMsg := md.AuthInvite{
		From:     sn.peer,
		Payload:  card.Payload,
		Card:     card,
		IsDirect: false,
		IsFile:   true,
	}

	// Check if ID in PeerStore
	go func(inv *md.AuthInvite) {
		// Convert Protobuf to bytes
		msgBytes, err := proto.Marshal(inv)
		if err != nil {
			sn.error(err, "Marshal")
		}

		sn.peerConn.Request(sn.host, id, msgBytes)
	}(&invMsg)
}

// ^ Send Invite with User Contact Card ^ //
func (sn *Node) InviteWithContact(peerId string) {
	// Get PeerID
	id, _, err := sn.lobby.Find(peerId)

	// Check error
	if err != nil {
		sn.error(err, "InviteWithContact")
	}

	// Create Invite Message with Payload and Card
	invMsg := md.AuthInvite{
		From:     sn.peer,
		Payload:  md.Payload_CONTACT,
		Card:     sf.NewCardFromContact(sn.peer.Profile, sn.contact, md.TransferCard_INVITE),
		IsFile:   false,
		IsDirect: false,
	}

	// Check if ID in PeerStore
	go func(inv *md.AuthInvite) {
		// Convert Protobuf to bytes
		msgBytes, err := proto.Marshal(inv)
		if err != nil {
			sn.error(err, "Marshal")
		}

		sn.peerConn.Request(sn.host, id, msgBytes)
	}(&invMsg)
}

// ^ Send Invite with URL Link ^ //
func (sn *Node) InviteWithURL(peerId string, url string) {
	// Get PeerID
	id, _, err := sn.lobby.Find(peerId)

	// Check error
	if err != nil {
		sn.error(err, "InviteWithURL")
	}

	// Create Invite Message with Payload
	invMsg := md.AuthInvite{
		From:     sn.peer,
		Payload:  md.Payload_URL,
		Card:     sf.NewCardFromUrl(sn.peer.Profile, url, md.TransferCard_INVITE),
		IsFile:   false,
		IsDirect: false,
	}

	// Check if ID in PeerStore
	go func(inv *md.AuthInvite) {
		// Convert Protobuf to bytes
		msgBytes, err := proto.Marshal(inv)
		if err != nil {
			sn.error(err, "Marshal")
		}

		sn.peerConn.Request(sn.host, id, msgBytes)
	}(&invMsg)
}
