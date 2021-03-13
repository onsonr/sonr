package node

import (
	"log"
	"time"

	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/encoding/protojson"
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

// ^ Link with a QR Code ^ //
func (sn *Node) LinkDevice(json string) {
	// Convert String to Bytes
	request := md.LinkRequest{}

	// Convert to Peer Protobuf
	err := protojson.Unmarshal([]byte(json), &request)
	if err != nil {
		sn.error(err, "LinkDevice")
	}

	// Link Device
	err = sn.fs.AddDevice(request.Device, sn.directories.Documents)
	if err != nil {
		sn.error(err, "LinkDevice")
	}
}

// ^ Link with a QR Code ^ //
func (sn *Node) LinkRequest(name string) *md.LinkRequest {
	// Set Device
	device := sn.device
	device.Directories = sn.directories
	device.Name = name

	// Create Expiry - 1min 30s
	timein := time.Now().Local().Add(
		time.Minute*time.Duration(1) +
			time.Second*time.Duration(30))

	// Return Request
	return &md.LinkRequest{
		Device: device,
		Peer:   sn.Peer(),
		Expiry: int32(timein.Unix()),
	}
}

// ^ Peer returns Current Peer Info ^
func (sn *Node) Peer() *md.Peer {
	return sn.peer
}

// ^ Updates Current Contact Card ^
func (sn *Node) SetContact(newContact *md.Contact) {

	// Set Node Contact
	sn.contact = newContact

	// Update Peer Profile
	sn.peer.Profile = &md.Profile{
		FirstName: newContact.GetFirstName(),
		LastName:  newContact.GetLastName(),
		Picture:   newContact.GetPicture(),
	}

	// Set User Contact
	err := sn.fs.UpdateContact(newContact, sn.directories.Documents)
	if err != nil {
		sn.error(err, "SetContact")
	}
}
