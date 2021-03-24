package user

import (
	"hash/fnv"
	"log"
	"math"

	mid "github.com/denisbrodbeck/machineid"
	"github.com/getsentry/sentry-go"
	md "github.com/sonr-io/core/internal/models"
)

// ^ Updates Current Contact Card ^
func (u *User) SetContact(newContact *md.Contact) error {
	// Set Node Contact
	u.contact = newContact

	// Update Peer Profile
	u.peer.Profile = &md.Profile{
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
	deviceID := u.device.GetId()

	// Get User ID
	userID := fnv.New32a()
	_, err := userID.Write([]byte(u.profile.GetUsername()))
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
	u.peer = &md.Peer{
		Id: &md.Peer_ID{
			Peer:   hID,
			Device: deviceID,
			User:   userID.Sum32(),
		},
		Profile:  u.profile,
		Platform: u.device.Platform,
		Model:    u.device.Model,
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
	u.peer.Position = &md.Position{
		Facing:           faceDir,
		FacingAntipodal:  faceAnpd,
		Heading:          headDir,
		HeadingAntipodal: headAnpd,
		Designation:      md.Position_Designation(desg % 32),
	}
}
