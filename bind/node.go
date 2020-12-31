package sonr

import (
	"log"
	"math"
	"time"

	sf "github.com/sonr-io/core/internal/file"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ^ Info returns ALL Peer Data as Bytes^
func (sn *Node) Info() []byte {
	// Convert to bytes to view in plugin
	data, err := proto.Marshal(sn.peer)
	if err != nil {
		log.Println("Error Marshaling Lobby Data ", err)
		return nil
	}
	return data
}

// ^ Updates Current Contact Card ^
func (sn *Node) SetContact(conBytes []byte) {
	newContact := &md.Contact{}
	err := proto.Unmarshal(conBytes, newContact)
	if err != nil {
		log.Println(err)
	}
	sn.contact = newContact
}

// ^ Update proximity/direction and Notify Lobby ^ //
func (sn *Node) Update(direction float64) {
	// ** Initialize ** //
	// Update User Values
	sn.peer.Direction = math.Round(direction*100) / 100

	// Inform Lobby
	err := sn.lobby.Update()
	if err != nil {
		sn.error(err, "Update")
	}
}

// ^ AddFile adds generates Metadata and Thumbnail ^ //
func (sn *Node) AddFile(path string) {
	//@2. Initialize SafeFile
	safeMeta := sf.NewMetadata(path, sn.call.OnQueued, sn.error)
	sn.files = append(sn.files, safeMeta)
}

// ^ Send Invite with a File ^ //
func (sn *Node) InviteWithFile(peerId string) {
	// Get PeerID
	id, _, err := sn.lobby.Find(peerId)

	// Check error
	if err != nil {
		sn.error(err, "Invite")
	}

	// Create Invite Message with Payload
	time.Sleep(time.Millisecond * 100)

	// Retreive Current File
	currFile := sn.currentFile()
	sn.peerConn.SafeMeta = currFile

	// Create Invite Message
	invMsg := md.AuthInvite{
		From:    sn.peer,
		Payload: md.Payload_FILE,
		File:    currFile.GetMetadata(),
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
		sn.error(err, "Invite")
	}

	// Create Invite Message with Payload
	invMsg := md.AuthInvite{
		From:    sn.peer,
		Payload: md.Payload_CONTACT,
		Contact: sn.contact,
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
		sn.error(err, "Invite")
	}

	// Create Invite Message with Payload
	invMsg := md.AuthInvite{
		From:    sn.peer,
		Payload: md.Payload_URL,
		Url:     url,
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

// ^ Respond to an Invitation ^ //
func (sn *Node) Respond(decision bool) {
	// @ Check Decision

	// Send Response on PeerConnection
	sn.peerConn.Authorize(decision, sn.contact, sn.peer)
}

// ^ Reset Current Queued File Metadata ^ //
func (sn *Node) ResetFile() {
	// Reset Files Slice
	sn.files = nil
	sn.files = make([]*sf.SafeMetadata, maxFileBufferSize)
}
