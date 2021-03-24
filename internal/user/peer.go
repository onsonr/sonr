package user

import (
	"hash/fnv"
	"log"
	"math"

	mid "github.com/denisbrodbeck/machineid"
	"github.com/getsentry/sentry-go"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ^ Get Peer returns Current Peer ^ //
func (u *User) GetPeer() *md.Peer {
	return u.Peer
}

// ^ Peer returns Current Peer Info as Buffer
func (u *User) GetPeerBuf() []byte {
	// Convert to bytes
	buf, err := proto.Marshal(u.Peer)
	if err != nil {
		sentry.CaptureException(err)
		return nil
	}
	return buf
}

// ^ Updates Current Contact Card ^
func (u *User) SetContact(newContact *md.Contact) error {
	// Set Node Contact
	u.Contact = newContact

	// Update Peer Profile
	u.Peer.Profile = &md.Profile{
		FirstName: newContact.GetFirstName(),
		LastName:  newContact.GetLastName(),
		Picture:   newContact.GetPicture(),
	}

	// Load User
	user, err := u.LoadUser()
	if err != nil {
		return err
	}

	// Save User
	if err := u.SaveUser(user); err != nil {
		return err
	}
	return nil
}

// ^ SetPeer configures Peer Ref from Host ID Reference ^ //
func (u *User) SetPeer(hID string) {
	// Initialize
	deviceID := u.Device.GetId()

	// Get User ID
	userID := fnv.New32a()
	_, err := userID.Write([]byte(u.Profile.GetUsername()))
	if err != nil {
		log.Println(err)
	}

	// Check if ID not provided
	if deviceID == "" {
		// Generate ID
		if id, err := mid.ProtectedID("Sonr"); err != nil {
			sentry.CaptureException(err)
			deviceID = ""
		} else {
			deviceID = id
		}
	}

	// Set Peer
	u.Peer = &md.Peer{
		Id: &md.Peer_ID{
			Peer:   hID,
			Device: deviceID,
			User:   userID.Sum32(),
		},
		Profile:  u.Profile,
		Platform: u.Device.Platform,
		Model:    u.Device.Model,
	}
}

// ^ Update proximity/direction and Notify Lobby ^ //
func (u *User) SetPosition(facing float64, heading float64) {
	// Update User Values
	var faceDir float64
	var faceAnpd float64
	var headDir float64
	var headAnpd float64
	faceDir = math.Round(facing*100) / 100
	headDir = math.Round(heading*100) / 100
	desg := int((facing / 11.25) + 0.25)

	// Find Antipodal
	if facing > 180 {
		faceAnpd = math.Round((facing-180)*100) / 100
	} else {
		faceAnpd = math.Round((facing+180)*100) / 100
	}

	// Find Antipodal
	if heading > 180 {
		headAnpd = math.Round((heading-180)*100) / 100
	} else {
		headAnpd = math.Round((heading+180)*100) / 100
	}

	// Set Position
	u.Peer.Position = &md.Position{
		Facing:           faceDir,
		FacingAntipodal:  faceAnpd,
		Heading:          headDir,
		HeadingAntipodal: headAnpd,
		Designation:      md.Position_Designation(desg % 32),
	}
}
