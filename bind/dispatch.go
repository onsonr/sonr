package sonr

import (
	"time"

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
	sn.peerConn.SafePreview = currFile

	// Create Invite Message
	invMsg := md.AuthInvite{
		From:    sn.peer,
		Payload: md.Payload_FILE,
		Preview: currFile.GetPreview(),
		Type:    md.AuthInvite_Device,
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
		From:    sn.peer,
		Payload: md.Payload_URL,
		Url:     url,
		Type:    md.AuthInvite_Device,
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
	sn.peerConn.SafePreview = currFile

	// Create Invite Message
	invMsg := md.AuthInvite{
		From:    sn.peer,
		Payload: md.Payload_FILE,
		Preview: currFile.GetPreview(),
		Type:    md.AuthInvite_Peer,
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

	// Create Invite Message with Payload
	invMsg := md.AuthInvite{
		From:    sn.peer,
		Payload: md.Payload_CONTACT,
		Contact: sn.contact,
		Type:    md.AuthInvite_Peer,
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
		From:    sn.peer,
		Payload: md.Payload_URL,
		Url:     url,
		Type:    md.AuthInvite_Peer,
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
